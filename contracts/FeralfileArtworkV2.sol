// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";

import "./Authorizable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract FeralfileExhibitionV2 is ERC721Enumerable, Authorizable, IERC2981 {
    using Strings for uint256;

    // royalty payout address
    address public royaltyPayoutAddress;

    // The maximum limit of edition size for each exhibitions
    uint256 public immutable maxEditionPerArtwork;

    // the basis points of royalty payments for each secondary sales
    uint256 public immutable secondarySaleRoyaltyBPS;

    // the maximum basis points of royalty payments
    uint256 public constant MAX_ROYALITY_BPS = 100_00;

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
        string memory name_,
        string memory symbol_,
        uint256 maxEditionPerArtwork_,
        uint256 secondarySaleRoyaltyBPS_,
        address royaltyPayoutAddress_,
        string memory contractURI_,
        string memory tokenBaseURI_
    ) ERC721(name_, symbol_) {
        require(
            maxEditionPerArtwork_ > 0,
            "maxEdition of each artwork in an exhibition needs to be greater than zero"
        );
        require(
            secondarySaleRoyaltyBPS_ <= MAX_ROYALITY_BPS,
            "royalty BPS for secondary sales can not be greater than the maximum royalty BPS"
        );
        require(
            royaltyPayoutAddress_ != address(0),
            "invalid royalty payout address"
        );

        maxEditionPerArtwork = maxEditionPerArtwork_;
        secondarySaleRoyaltyBPS = secondarySaleRoyaltyBPS_;
        royaltyPayoutAddress = royaltyPayoutAddress_;
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
    /// @param fingerprint - the fingerprint of an artwork
    /// @param title - the title of an artwork
    /// @param artistName - the artist of an artwork
    /// @param editionSize - the maximum edition size of an artwork
    function createArtwork(
        string memory fingerprint,
        string memory title,
        string memory artistName,
        uint256 editionSize
    ) external onlyAuthorized {
        require(bytes(title).length != 0, "title can not be empty");
        require(bytes(artistName).length != 0, "artist can not be empty");
        require(bytes(fingerprint).length != 0, "fingerprint can not be empty");
        require(editionSize > 0, "edition size needs to be at least 1");
        require(
            editionSize <= maxEditionPerArtwork,
            "artwork edition size exceeds the maximum edition size of the exhibition"
        );

        uint256 artworkID = uint256(keccak256(abi.encode(fingerprint)));

        /// @notice make sure the artwork have not been registered
        require(
            bytes(artworks[artworkID].fingerprint).length == 0,
            "an artwork with the same fingerprint has already registered"
        );

        Artwork memory artwork = Artwork(
            fingerprint,
            title,
            artistName,
            editionSize
        );

        _allArtworks.push(artworkID);
        artworks[artworkID] = artwork;

        emit NewArtwork(artworkID);
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
    /// @param artworkID - the artwork id where the new edition is referenced to
    /// @param bitmarkID - the bitmark id of artwork edition before swapped
    /// @param editionNumber - the edition number of the artwork edition
    /// @param owner - the owner address of the new minted token
    /// @param ipfsCID - the IPFS cid for the new token
    function swapArtworkFromBitmark(
        uint256 artworkID,
        uint256 bitmarkID,
        uint256 editionNumber,
        address owner,
        string memory ipfsCID
    ) external onlyAuthorized {
        /// @notice the edition size is not set implies the artwork is not created
        require(artworks[artworkID].editionSize > 0, "artwork is not found");
        /// @notice The range of editionNumber should be between 0 (AP) ~ artwork.editionSize
        require(
            editionNumber <= artworks[artworkID].editionSize,
            "edition number exceed the edition size of the artwork"
        );
        require(owner != address(0), "invalid owner address");
        require(!registeredBitmarks[bitmarkID], "bitmark id has registered");
        require(!registeredIPFSCIDs[ipfsCID], "ipfs id has registered");

        uint256 editionID = artworkID + editionNumber;
        require(
            artworkEditions[editionID].editionID == 0,
            "the edition is existent"
        );

        ArtworkEdition memory edition = ArtworkEdition(editionID, ipfsCID);

        artworkEditions[editionID] = edition;
        allArtworkEditions[artworkID].push(editionID);

        registeredBitmarks[bitmarkID] = true;
        registeredIPFSCIDs[ipfsCID] = true;

        _safeMint(owner, editionID);
        emit NewArtworkEdition(owner, artworkID, editionID);
    }

    /// @notice Update the IPFS cid of an edition to a new value
    function updateArtworkEditionIPFSCid(uint256 tokenId, string memory ipfsCID)
        external
        onlyAuthorized
    {
        require(_exists(tokenId), "artwork edition is not found");
        require(!registeredIPFSCIDs[ipfsCID], "ipfs id has registered");

        ArtworkEdition storage edition = artworkEditions[tokenId];
        delete registeredIPFSCIDs[edition.ipfsCID];
        registeredIPFSCIDs[ipfsCID] = true;
        edition.ipfsCID = ipfsCID;
    }

    /// @notice setRoyaltyPayoutAddress assigns a payout address so
    //          that we can split the royalty.
    /// @param royaltyPayoutAddress_ - the new royalty payout address
    function setRoyaltyPayoutAddress(address royaltyPayoutAddress_)
        external
        onlyAuthorized
    {
        require(
            royaltyPayoutAddress_ != address(0),
            "invalid royalty payout address"
        );
        royaltyPayoutAddress = royaltyPayoutAddress_;
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
    /// @param tokenId - the NFT asset queried for royalty information
    /// @param salePrice - the sale price of the NFT asset specified by tokenId
    /// @return receiver - address of who should be sent the royalty payment
    /// @return royaltyAmount - the royalty payment amount for salePrice
    function royaltyInfo(uint256 tokenId, uint256 salePrice)
        external
        view
        override
        returns (address receiver, uint256 royaltyAmount)
    {
        require(
            _exists(tokenId),
            "ERC2981: query royalty info for nonexistent token"
        );

        receiver = royaltyPayoutAddress;

        royaltyAmount =
            (salePrice * secondarySaleRoyaltyBPS) /
            MAX_ROYALITY_BPS;
    }

    event NewArtwork(uint256 indexed artworkID);
    event NewArtworkEdition(
        address indexed owner,
        uint256 indexed artworkID,
        uint256 indexed editionID
    );
}
