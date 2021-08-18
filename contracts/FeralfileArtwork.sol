// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract Authorizable is Ownable {
    address public trustee;

    constructor() {
        trustee = address(0x0);
    }

    modifier onlyAuthorized() {
        require(msg.sender == trustee || msg.sender == owner());
        _;
    }

    function setTrustee(address _newTrustee) public onlyOwner {
        trustee = _newTrustee;
    }
}

contract FeralfileExhibition is ERC721Enumerable, Authorizable {
    using Strings for uint256;

    struct Artwork {
        bytes32 fingerprint;
        string title;
        address artist;
        string medium;
        string data;
        uint256 editionSize;
        uint256 initialPrice;
        bool minted; // the field is designed to limit the minting process
    }

    struct ArtworkEdition {
        uint256 editionID;
        uint256 artworkID;
        uint256 bitmarkID;
        string ipfsCID;
    }

    // Exihibition information
    string public title;
    address public curator;
    uint256 public maxEdition;
    uint256 public basePrice;

    string private _tokenBaseURI;

    mapping(uint256 => Artwork) public artworks; // artworkID => Artwork
    mapping(uint256 => ArtworkEdition) public artworkEditions; // artworkEditionID => ArtworkEdition
    mapping(uint256 => uint256[]) internal allArtworkEditions; // artworkID => []ArtworkEditionID
    mapping(string => bool) internal registeredIPFSCIDs; // ipfsCID => bool

    constructor(
        string memory _title,
        string memory _symbol,
        address _curator,
        uint256 _maxEdition,
        uint256 _basePrice,
        string memory tokenBaseURI_
    ) ERC721(_title, _symbol) {
        title = _title;
        curator = _curator;
        maxEdition = _maxEdition;
        basePrice = _basePrice;
        _tokenBaseURI = tokenBaseURI_;
    }

    // Create an artwork for an exhibition
    function createArtwork(
        bytes32 _fingerprint,
        string memory _title,
        address _artist,
        string memory _medium,
        string memory _data,
        uint256 _editionSize,
        uint256 _initialPrice
    ) public onlyAuthorized {
        require(_fingerprint.length != 0, "fingerprint can not be empty");
        require(
            _editionSize <= maxEdition,
            "edition size exceeds the maximum edition size of the exhibition"
        );

        uint256 _artworkID = uint256(keccak256(abi.encode(_fingerprint)));

        // make sure an artwork have not been registered
        require(
            artworks[_artworkID].fingerprint == 0,
            "duplicated fingerprint"
        );

        Artwork memory _artwork = Artwork(
            _fingerprint,
            _title,
            _artist,
            _medium,
            _data,
            _editionSize,
            _initialPrice,
            false
        );

        artworks[_artworkID] = _artwork;

        emit NewArtwork(_artist, _artworkID);
    }

    // Mint editions for an artwork. For each artwork, it can only mint once.
    function mintArtwork(uint256 _artworkID, string[] memory _ipfsCIDs)
        public
        onlyAuthorized
    {
        Artwork memory _artwork = artworks[_artworkID];
        require(_artwork.fingerprint != 0, "artwork is not found");
        require(!_artwork.minted);

        for (uint256 i = 0; i < _ipfsCIDs.length; i++) {
            uint256 editionID = uint256(
                keccak256(
                    abi.encode(
                        _artworkID,
                        allArtworkEditions[_artworkID].length
                    )
                )
            );

            // ensure the IPFS id is not used by other artwork
            require(
                artworkEditions[editionID].editionID == 0,
                "duplicated edition id"
            );

            // ensure the IPFS id is not used by other artwork
            require(!registeredIPFSCIDs[_ipfsCIDs[i]], "ipfs id registered");

            ArtworkEdition memory edition = ArtworkEdition(
                editionID,
                _artworkID,
                0,
                _ipfsCIDs[i]
            );

            registeredIPFSCIDs[_ipfsCIDs[i]] = true;
            artworkEditions[editionID] = edition;
            allArtworkEditions[_artworkID].push(editionID);

            _mint(_artwork.artist, editionID);
            emit NewArtworkEdition(_artwork.artist, _artworkID, editionID);
        }
    }

    // Swap an existent artwork from bitmark to ERC721
    function swapArtworkFromBitmarks(
        uint256 _artworkID,
        uint256 _bitmarkIDs,
        address _newOwner,
        string memory _ipfsCID
    ) public onlyAuthorized {
        Artwork memory artwork = artworks[_artworkID];
        require(artwork.fingerprint != 0, "artwork is not found");
        require(!artwork.minted);

        uint256 editionID = uint256(
            keccak256(
                abi.encode(_artworkID, allArtworkEditions[_artworkID].length)
            )
        );

        string memory ipfsCID = artworkEditions[editionID].ipfsCID;

        require(!registeredIPFSCIDs[ipfsCID], "ipfs id registered");

        ArtworkEdition memory edition = ArtworkEdition(
            editionID,
            _artworkID,
            _bitmarkIDs,
            _ipfsCID
        );

        artworkEditions[editionID] = edition;
        allArtworkEditions[_artworkID].push(editionID);

        _mint(_newOwner, editionID);
        emit NewArtworkEdition(_newOwner, _artworkID, editionID);
    }

    function totalEditionOfArtwork(uint256 artworkID)
        public
        view
        returns (uint256)
    {
        return allArtworkEditions[artworkID].length;
    }

    function getArtworkEditionByIndex(uint256 artworkID, uint256 index)
        public
        view
        returns (uint256)
    {
        require(index < totalEditionOfArtwork(artworkID));
        return allArtworkEditions[artworkID][index];
    }

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

        string memory baseURI = _baseURI();
        if (bytes(baseURI).length == 0) {
            return "";
        }

        string memory ipfsID = artworkEditions[tokenId].ipfsCID;

        return string(abi.encodePacked(baseURI, ipfsID, "/metadata.json"));
    }

    function setArtworkBaseURI(string memory baseURI_) public onlyAuthorized {
        _tokenBaseURI = baseURI_;
    }

    function _baseURI() internal view virtual override returns (string memory) {
        return _tokenBaseURI;
    }

    event NewArtwork(address _creator, uint256 _artworkID);
    event NewArtworkEdition(
        address _owner,
        uint256 _artworkID,
        uint256 _editionID
    );
}
