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
        string fingerprint;
        string title;
        string description;
        address artist;
        string medium;
        uint256 editionSize;
        bool minted; // the field is designed to limit the minting process
    }

    struct ArtworkEdition {
        uint256 editionID;
        uint256 editionNumber;
        uint256 artworkID;
        uint256 bitmarkID;
        string prevProvenance;
        string ipfsCID;
    }

    // Exihibition information
    string public title;
    address public curator;
    string public curator_notes;
    uint256 public maxEdition;
    uint256 public basePrice;

    string private _tokenBaseURI;

    uint256[] private _allArtworks;
    mapping(uint256 => Artwork) public artworks; // artworkID => Artwork
    mapping(uint256 => ArtworkEdition) public artworkEditions; // artworkEditionID => ArtworkEdition
    mapping(uint256 => uint256[]) internal allArtworkEditions; // artworkID => []ArtworkEditionID
    mapping(uint256 => bool) internal registeredBitmarks; // bitmarkID => bool
    mapping(string => bool) internal registeredIPFSCIDs; // ipfsCID => bool

    constructor(
        string memory _title,
        string memory _symbol,
        address _curator,
        uint256 _maxEdition,
        uint256 _basePrice,
        string memory tokenBaseURI_
    ) ERC721(_title, _symbol) {
        require(_curator != address(0), "invalid curator address");
        require(_maxEdition > 0, "maxEdition needs to be greater than zero");
        require(_basePrice > 0, "basePrice needs to be greater than zero");

        title = _title;
        curator = _curator;
        maxEdition = _maxEdition;
        basePrice = _basePrice;
        _tokenBaseURI = tokenBaseURI_;
    }

    // Create an artwork for an exhibition
    function createArtwork(
        string memory _fingerprint,
        string memory _title,
        string memory _description,
        address _artist,
        string memory _medium,
        uint256 _editionSize
    ) public onlyAuthorized {
        require(
            bytes(_fingerprint).length != 0,
            "fingerprint can not be empty"
        );
        require(bytes(_title).length != 0, "title can not be empty");
        require(_artist != address(0), "invalid artist address");
        require(bytes(_medium).length != 0, "medium can not be empty");
        require(_editionSize > 0, "edition size needs to be at least 1");
        require(
            _editionSize <= maxEdition,
            "edition size exceeds the maximum edition size of the exhibition"
        );

        uint256 _artworkID = uint256(keccak256(abi.encode(_fingerprint)));

        // make sure an artwork have not been registered
        require(
            artworks[_artworkID].artist == address(0),
            "an artwork with the same fingerprint has already registered"
        );

        Artwork memory _artwork = Artwork(
            _fingerprint,
            _title,
            _description,
            _artist,
            _medium,
            _editionSize,
            false
        );

        _allArtworks.push(_artworkID);
        artworks[_artworkID] = _artwork;

        emit NewArtwork(_artist, _artworkID);
    }

    // Return a count of artworks registered in this exhibition
    function totalArtworks() public view virtual returns (uint256) {
        return _allArtworks.length;
    }

    // Return the token identifier for the `_index`th artwork
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

    // Swap an existent artwork from bitmark to ERC721
    function swapArtworkFromBitmarks(
        uint256 _artworkID,
        uint256 _bitmarkID,
        uint256 _editionNumber,
        address _owner,
        string memory _prevProvenance,
        string memory _ipfsCID
    ) public onlyAuthorized {
        Artwork memory artwork = artworks[_artworkID];
        require(artwork.artist != address(0), "artwork is not found");

        // The range of _editionNumber should be between 0 (AP) ~ artwork.editionSize
        require(
            _editionNumber <= artwork.editionSize,
            "edition number exceed the edition size of the artwork"
        );

        require(_owner != address(0), "invalid artist address");

        uint256 editionID = _artworkID + _editionNumber;
        require(
            artworkEditions[editionID].editionID == 0,
            "the edition is existent"
        );
        require(!registeredBitmarks[_bitmarkID], "bitmark id has registered");
        require(!registeredIPFSCIDs[_ipfsCID], "ipfs id has registered");

        ArtworkEdition memory edition = ArtworkEdition(
            editionID,
            _editionNumber,
            _artworkID,
            _bitmarkID,
            _prevProvenance,
            _ipfsCID
        );

        artworkEditions[editionID] = edition;
        allArtworkEditions[_artworkID].push(editionID);

        registeredBitmarks[_bitmarkID] = true;
        registeredIPFSCIDs[_ipfsCID] = true;

        _mint(_owner, editionID);
        emit NewArtworkEdition(_owner, _artworkID, editionID);
    }

    // Return the edition counts for an artwork
    function totalEditionOfArtwork(uint256 artworkID)
        public
        view
        returns (uint256)
    {
        return allArtworkEditions[artworkID].length;
    }

    // Return the edition id of an artwork by index
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

    event NewArtwork(address creator, uint256 artworkID);
    event NewArtworkEdition(
        address owner,
        uint256 artworkID,
        uint256 editionID
    );
}
