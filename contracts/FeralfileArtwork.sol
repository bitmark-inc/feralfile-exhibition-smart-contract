// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";

import "./Authorizable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract FeralfileExhibition is ERC721Enumerable, Authorizable, IERC2981 {
    using Strings for uint256;

    // Exihibition title
    string public exhibitionTitle;

    // royalty payout address
    address public royaltyPayoutAddress;

    // The maximum limit of edition size for each exhibitions
    uint256 public maxEditionPerArtwork;

    // the biase points of royalty payments for each secondary sales
    uint256 public secondarySaleRoyaltyBPS = 0;

    // the maximum biase points of royalty payments
    uint256 public constant maxRoyaltyBPS = 100_00;

    // token base URI
    string private _tokenBaseURI;

    // contract URI
    string private _contractURI;

    /// @notice A structure for Feral File artwork
    struct Artwork {
        string title;
        string artistName;
        string fingerprint;
        uint256 editionSize;
    }

    struct ArtworkEdition {
        uint256 editionID;
        string ipfsCID;
    }

    uint256[] private _allArtworks;
    mapping(uint256 => Artwork) public artworks; // artworkID => Artwork
    mapping(uint256 => ArtworkEdition) public artworkEditions; // artworkEditionID => ArtworkEdition
    mapping(uint256 => uint256[]) internal allArtworkEditions; // artworkID => []ArtworkEditionID
    mapping(uint256 => bool) internal registeredBitmarks; // bitmarkID => bool
    mapping(string => bool) internal registeredIPFSCIDs; // ipfsCID => bool

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _exhibitionTitle,
        uint256 _maxEditionPerArtwork,
        uint256 _secondarySaleRoyaltyBPS,
        address _royaltyPayoutAddress,
        string memory contractURI_,
        string memory tokenBaseURI_
    ) ERC721(_name, _symbol) {
        require(
            _maxEditionPerArtwork > 0,
            "maxEdition of each artwork in an exhibition needs to be greater than zero"
        );
        require(
            _secondarySaleRoyaltyBPS <= MAX_ROYALITY_BPS,
            "royalty BPS for secondary sales can not be greater than the maximum royalty BPS"
        );
        require(
            _royaltyPayoutAddress != address(0),
            "invalid royalty payout address"
        );

        exhibitionTitle = _exhibitionTitle;
        maxEditionPerArtwork = _maxEditionPerArtwork;
        secondarySaleRoyaltyBPS = _secondarySaleRoyaltyBPS;
        royaltyPayoutAddress = _royaltyPayoutAddress;
        _contractURI = contractURI_;
        _tokenBaseURI = tokenBaseURI_;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC721Enumerable, IERC165)
        returns (bool)
    {
        return
            interfaceId == type(IERC721Enumerable).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /// @notice Call to create an artwork in the exhibition
    /// @param _fingerprint - the fingerprint of an artwork
    /// @param _title - the title of an artwork
    /// @param _artistName - the artist of an artwork
    /// @param _editionSize - the maximum edition size of an artwork
    function createArtwork(
        string memory _fingerprint,
        string memory _title,
        string memory _artistName,
        uint256 _editionSize
    ) external onlyAuthorized {
        require(bytes(_title).length != 0, "title can not be empty");
        require(bytes(_artistName).length != 0, "artist can not be empty");
        require(
            bytes(_fingerprint).length != 0,
            "fingerprint can not be empty"
        );
        require(_editionSize > 0, "edition size needs to be at least 1");
        require(
            _editionSize <= maxEditionPerArtwork,
            "artwork edition size exceeds the maximum edition size of the exhibition"
        );

        uint256 _artworkID = uint256(keccak256(abi.encode(_fingerprint)));

        /// @notice make sure the artwork have not been registered
        require(
            bytes(artworks[_artworkID].fingerprint).length == 0,
            "an artwork with the same fingerprint has already registered"
        );

        Artwork memory _artwork = Artwork(
            _fingerprint,
            _title,
            _artistName,
            _editionSize
        );

        _allArtworks.push(_artworkID);
        artworks[_artworkID] = _artwork;

        emit NewArtwork(_artworkID);
    }

    /// @notice Return a count of artworks registered in this exhibition
    function totalArtworks() public view virtual returns (uint256) {
        return _allArtworks.length;
    }

    /// @notice Return the token identifier for the `index`th artwork
    function getArtworkByIndex(uint256 index)
        public
        view
        virtual
        returns (uint256)
    {
        require(
            index < totalArtworks(),
            "artworks: global index out of bounds"
        );
        return _allArtworks[index];
    }

    /// @notice Swap an existent artwork from bitmark to ERC721
    /// @param _artworkID - the artwork id where the new edition is referenced to
    /// @param _bitmarkID - the bitmark id of artwork edition before swapped
    /// @param _editionNumber - the edition number of the artwork edition
    /// @param _owner - the owner address of the new minted token
    /// @param _ipfsCID - the IPFS cid for the new token
    function swapArtworkFromBitmark(
        uint256 _artworkID,
        uint256 _bitmarkID,
        uint256 _editionNumber,
        address _owner,
        string memory _ipfsCID
    ) external onlyAuthorized {
        /// @notice the edition size is not set implies the artwork is not created
        require(artworks[_artworkID].editionSize > 0, "artwork is not found");
        /// @notice The range of _editionNumber should be between 0 (AP) ~ artwork.editionSize
        require(
            _editionNumber <= artworks[_artworkID].editionSize,
            "edition number exceed the edition size of the artwork"
        );
        require(_owner != address(0), "invalid owner address");
        require(!registeredBitmarks[_bitmarkID], "bitmark id has registered");
        require(!registeredIPFSCIDs[_ipfsCID], "ipfs id has registered");

        uint256 editionID = _artworkID + _editionNumber;
        require(
            artworkEditions[editionID].editionID == 0,
            "the edition is existent"
        );

        ArtworkEdition memory edition = ArtworkEdition(editionID, _ipfsCID);

        artworkEditions[editionID] = edition;
        allArtworkEditions[_artworkID].push(editionID);

        registeredBitmarks[_bitmarkID] = true;
        registeredIPFSCIDs[_ipfsCID] = true;

        _safeMint(_owner, editionID);
        emit NewArtworkEdition(_owner, _artworkID, editionID);
    }

    /// @notice Update the IPFS cid of an edition to a new value
    function updateArtworkEditionIPFSCid(
        uint256 _tokenId,
        string memory _ipfsCID
    ) external onlyAuthorized {
        ArtworkEdition storage edition = artworkEditions[_tokenId];
        require(edition.editionID != 0, "artwork edition is not found");
        require(!registeredIPFSCIDs[_ipfsCID], "ipfs id has registered");

        delete registeredIPFSCIDs[edition.ipfsCID];
        registeredIPFSCIDs[_ipfsCID] = true;
        edition.ipfsCID = _ipfsCID;
    }

    /// @notice setRoyaltyPayoutAddress assigns a payout address so
    //          that we can split the royalty.
    function setRoyaltyPayoutAddress(address payoutAddress)
        external
        onlyAuthorized
    {
        require(payoutAddress != address(0), "invalid royalty payout address");
        royaltyPayoutAddress = payoutAddress;
    }

    /// @notice Return the edition counts for an artwork
    function totalEditionOfArtwork(uint256 artworkID)
        public
        view
        returns (uint256)
    {
        return allArtworkEditions[artworkID].length;
    }

    /// @notice Return the edition id of an artwork by index
    function getArtworkEditionByIndex(uint256 artworkID, uint256 index)
        public
        view
        returns (uint256)
    {
        require(index < totalEditionOfArtwork(artworkID));
        return allArtworkEditions[artworkID][index];
    }

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    function tokenURI(uint256 tokenId)
        public
        view
        virtual
        override
        returns (string memory)
    {
        require(
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );

        string memory baseURI = _tokenBaseURI;
        if (bytes(baseURI).length == 0) {
            baseURI = "ipfs://";
        }

        return
            string(
                abi.encodePacked(
                    baseURI,
                    artworkEditions[tokenId].ipfsCID,
                    "/metadata.json"
                )
            );
    }

    /// @notice Update the base URI for all tokens
    function setTokenBaseURI(string memory baseURI_) external onlyAuthorized {
        _tokenBaseURI = baseURI_;
    }

    /// @notice A URL for the opensea storefront-level metadata
    function contractURI() public view returns (string memory) {
        return _contractURI;
    }

    /// @notice Called with the sale price to determine how much royalty
    //          is owed and to whom.
    /// @param _tokenId - the NFT asset queried for royalty information
    /// @param _salePrice - the sale price of the NFT asset specified by _tokenId
    /// @return receiver - address of who should be sent the royalty payment
    /// @return royaltyAmount - the royalty payment amount for _salePrice
    function royaltyInfo(uint256 _tokenId, uint256 _salePrice)
        external
        view
        override
        returns (address receiver, uint256 royaltyAmount)
    {
        ArtworkEdition memory edition = artworkEditions[_tokenId];
        require(edition.editionID != 0, "artwork edition is not found");

        receiver = royaltyPayoutAddress;

        royaltyAmount = (_salePrice / maxRoyaltyBPS) * secondarySaleRoyaltyBPS;
    }

    event NewArtwork(uint256 indexed artworkID);
    event NewArtworkEdition(
        address indexed owner,
        uint256 indexed artworkID,
        uint256 indexed editionID
    );
}
