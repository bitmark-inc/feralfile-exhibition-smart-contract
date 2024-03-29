// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "./Authorizable.sol";

contract FeralfileExhibitionV3 is ERC721Enumerable, Authorizable, IERC2981 {
    using Strings for uint256;

    // royalty payout address
    address public royaltyPayoutAddress;

    // the basis points of royalty payments for each secondary sales
    uint256 public immutable secondarySaleRoyaltyBPS;

    // the maximum basis points of royalty payments
    uint256 public constant MAX_ROYALITY_BPS = 100_00;

    // version code of contract
    string public constant codeVersion = "FeralfileExhibitionV3";

    // burnable
    bool public isBurnable;

    // bridgeable
    bool public isBridgeable;

    // token base URI
    string internal _tokenBaseURI;

    // contract URI
    string private _contractURI;

    /// @notice A structure for Feral File artwork
    struct Artwork {
        string title;
        string artistName;
        string fingerprint;
        uint256 editionSize;
        uint256 AEAmount;
        uint256 PPAmount;
    }

    struct ArtworkEdition {
        uint256 editionID;
        string ipfsCID;
    }

    struct TransferArtworkParam {
        address from;
        address to;
        uint256 tokenID;
        uint256 expireTime;
        bytes32 r_;
        bytes32 s_;
        uint8 v_;
    }

    struct MintArtworkParam {
        uint256 artworkID;
        uint256 edition;
        address artist;
        address owner;
        string ipfsCID;
    }

    struct ArtworkEditionIndex {
        uint256 artworkID;
        uint256 index;
    }

    uint256[] private _allArtworks;
    mapping(uint256 => Artwork) public artworks; // artworkID => Artwork
    mapping(uint256 => ArtworkEdition) public artworkEditions; // artworkEditionID => ArtworkEdition
    mapping(uint256 => uint256[]) internal allArtworkEditions; // artworkID => []ArtworkEditionID
    mapping(string => bool) internal registeredIPFSCIDs; // ipfsCID => bool
    mapping(uint256 => ArtworkEditionIndex) internal allArtworkEditionsIndex; // editionID => ArtworkEditionIndex

    constructor(
        string memory name_,
        string memory symbol_,
        uint256 secondarySaleRoyaltyBPS_,
        address royaltyPayoutAddress_,
        string memory contractURI_,
        string memory tokenBaseURI_,
        bool isBurnable_,
        bool isBridgeable_
    ) ERC721(name_, symbol_) {
        require(
            secondarySaleRoyaltyBPS_ <= MAX_ROYALITY_BPS,
            "royalty BPS for secondary sales can not be greater than the maximum royalty BPS"
        );
        require(
            royaltyPayoutAddress_ != address(0),
            "invalid royalty payout address"
        );

        secondarySaleRoyaltyBPS = secondarySaleRoyaltyBPS_;
        royaltyPayoutAddress = royaltyPayoutAddress_;
        _contractURI = contractURI_;
        _tokenBaseURI = tokenBaseURI_;
        isBurnable = isBurnable_;
        isBridgeable = isBridgeable_;
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
    function _createArtwork(
        string memory fingerprint,
        string memory title,
        string memory artistName,
        uint256 editionSize,
        uint256 aeAmount,
        uint256 ppAmount
    ) internal returns (uint256) {
        require(bytes(title).length != 0, "title can not be empty");
        require(bytes(artistName).length != 0, "artist can not be empty");
        require(bytes(fingerprint).length != 0, "fingerprint can not be empty");
        require(editionSize > 0, "edition size needs to be at least 1");

        uint256 artworkID = uint256(keccak256(abi.encode(fingerprint)));

        /// @notice make sure the artwork have not been registered
        require(
            bytes(artworks[artworkID].fingerprint).length == 0,
            "an artwork with the same fingerprint has already registered"
        );

        Artwork memory artwork = Artwork(
            title = title,
            artistName = artistName,
            fingerprint = fingerprint,
            editionSize = editionSize,
            aeAmount = aeAmount,
            ppAmount = ppAmount
        );

        _allArtworks.push(artworkID);
        artworks[artworkID] = artwork;

        emit NewArtwork(artworkID);

        return artworkID;
    }

    /// @notice createArtworks use for create list of artworks in a transaction
    /// @param artworks_ - the array of artwork
    function createArtworks(Artwork[] memory artworks_)
        external
        onlyAuthorized
    {
        for (uint256 i = 0; i < artworks_.length; i++) {
            _createArtwork(
                artworks_[i].fingerprint,
                artworks_[i].title,
                artworks_[i].artistName,
                artworks_[i].editionSize,
                artworks_[i].AEAmount,
                artworks_[i].PPAmount
            );
        }
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
            string(abi.encodePacked(baseURI, artworkEditions[tokenId].ipfsCID));
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

    /// @notice isValidRequest validates a message by ecrecover to ensure
    //          it is signed by owner of token.
    /// @param message_ - the raw message for signing
    /// @param owner_ - owner address of token
    /// @param r_ - part of signature for validating parameters integrity
    /// @param s_ - part of signature for validating parameters integrity
    /// @param v_ - part of signature for validating parameters integrity
    function isValidRequest(
        bytes32 message_,
        address owner_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) internal pure returns (bool) {
        address signer = ECDSA.recover(
            ECDSA.toEthSignedMessageHash(message_),
            v_,
            r_,
            s_
        );
        return signer == owner_;
    }

    /// @notice authorizedTransfer use for transfer list of items in a transaction
    /// @param transferParams_ - the array of transfer parameters
    function authorizedTransfer(TransferArtworkParam[] memory transferParams_)
        external
        onlyAuthorized
    {
        for (uint256 i = 0; i < transferParams_.length; i++) {
            _authorizedTransfer(transferParams_[i]);
        }
    }

    function _authorizedTransfer(TransferArtworkParam memory transferParam_)
        private
    {
        require(
            _exists(transferParam_.tokenID),
            "ERC721: artwork edition is not found"
        );

        require(
            _isApprovedOrOwner(transferParam_.from, transferParam_.tokenID),
            "ERC721: caller is not token owner nor approved"
        );

        require(
            block.timestamp <= transferParam_.expireTime,
            "FeralfileExhibitionV3: the transfer request is expired"
        );

        bytes32 requestHash = keccak256(
            abi.encode(
                transferParam_.from,
                transferParam_.to,
                transferParam_.tokenID,
                transferParam_.expireTime
            )
        );

        require(
            isValidRequest(
                requestHash,
                transferParam_.from,
                transferParam_.r_,
                transferParam_.s_,
                transferParam_.v_
            ),
            "FeralfileExhibitionV3: the transfer request is not authorized"
        );

        _safeTransfer(
            transferParam_.from,
            transferParam_.to,
            transferParam_.tokenID,
            ""
        );
    }

    /// @notice batchMint is function mint array of tokens
    /// @param mintParams_ - the array of transfer parameters
    function batchMint(MintArtworkParam[] memory mintParams_)
        external
        virtual
        onlyAuthorized
    {
        for (uint256 i = 0; i < mintParams_.length; i++) {
            _mintArtwork(
                mintParams_[i].artworkID,
                mintParams_[i].edition,
                mintParams_[i].artist,
                mintParams_[i].owner,
                mintParams_[i].ipfsCID
            );
        }
    }

    /// @notice mint artwork to ERC721
    /// @param artworkID_ - the artwork id where the new edition is referenced to
    /// @param editionNumber_ - the edition number of the artwork edition
    /// @param artist_ - the artist address of the new minted token
    /// @param owner_ - the owner address of the new minted token
    /// @param ipfsCID_ - the IPFS cid for the new token
    function _mintArtwork(
        uint256 artworkID_,
        uint256 editionNumber_,
        address artist_,
        address owner_,
        string memory ipfsCID_
    ) internal {
        /// @notice the edition size is not set implies the artwork is not created
        require(
            artworks[artworkID_].editionSize > 0,
            "FeralfileExhibitionV3: artwork is not found"
        );
        /// @notice The range of editionNumber should be between 0 to artwork.editionSize + artwork.AEAmount + artwork.PPAmount - 1
        require(
            editionNumber_ <
                artworks[artworkID_].editionSize +
                    artworks[artworkID_].AEAmount +
                    artworks[artworkID_].PPAmount,
            "FeralfileExhibitionV3: edition number exceed the edition size of the artwork"
        );
        require(artist_ != address(0), "invalid artist address");
        require(owner_ != address(0), "invalid owner address");
        require(!registeredIPFSCIDs[ipfsCID_], "ipfs id has registered");

        uint256 editionID = artworkID_ + editionNumber_;
        require(
            artworkEditions[editionID].editionID == 0,
            "FeralfileExhibitionV3: the edition is existent"
        );

        ArtworkEdition memory edition = ArtworkEdition(editionID, ipfsCID_);

        artworkEditions[editionID] = edition;
        allArtworkEditions[artworkID_].push(editionID);
        allArtworkEditionsIndex[editionID] = ArtworkEditionIndex(
            artworkID_,
            allArtworkEditions[artworkID_].length - 1
        );

        registeredIPFSCIDs[ipfsCID_] = true;

        _safeMint(artist_, editionID);

        if (artist_ != owner_) {
            _safeTransfer(artist_, owner_, editionID, "");
        }

        emit NewArtworkEdition(owner_, artworkID_, editionID);
    }

    /// @notice remove an edition from allArtworkEditions
    /// @param editionID - the edition id where we are going to remove from allArtworkEditions
    function _removeEditionFromAllArtworkEditions(uint256 editionID) private {
        ArtworkEditionIndex
            memory artworkEditionIndex = allArtworkEditionsIndex[editionID];

        require(
            artworkEditionIndex.artworkID > 0,
            "FeralfileExhibitionV3: artworkID is no found for the artworkEditionIndex"
        );

        uint256[] storage artworkEditions_ = allArtworkEditions[
            artworkEditionIndex.artworkID
        ];

        require(
            artworkEditions_.length > 0,
            "FeralfileExhibitionV3: no editions in this artwork of allArtworkEditions"
        );

        uint256 lastEditionID = artworkEditions_[artworkEditions_.length - 1];

        // Swap between the last token and the to-delete token and pop up the last token
        artworkEditions_[artworkEditionIndex.index] = lastEditionID;
        artworkEditions_.pop();
        allArtworkEditionsIndex[lastEditionID].index = artworkEditionIndex
            .index;
        delete allArtworkEditionsIndex[editionID];
    }

    /// @notice burn editions
    /// @param editionIDs_ - the list of edition id will be burned
    function burnEditions(uint256[] memory editionIDs_) public {
        require(isBurnable, "FeralfileExhibitionV3: not allow burn edition");

        for (uint256 i = 0; i < editionIDs_.length; i++) {
            require(
                _exists(editionIDs_[i]),
                "ERC721: artwork edition is not found"
            );
            require(
                _isApprovedOrOwner(_msgSender(), editionIDs_[i]),
                "ERC721: caller is not token owner nor approved"
            );
            ArtworkEdition memory edition = artworkEditions[editionIDs_[i]];

            delete registeredIPFSCIDs[edition.ipfsCID];
            delete artworkEditions[editionIDs_[i]];

            _removeEditionFromAllArtworkEditions(editionIDs_[i]);

            _burn(editionIDs_[i]);

            emit BurnArtworkEdition(editionIDs_[i]);
        }
    }

    event NewArtwork(uint256 indexed artworkID);
    event NewArtworkEdition(
        address indexed owner,
        uint256 indexed artworkID,
        uint256 indexed editionID
    );
    event BurnArtworkEdition(uint256 indexed editionID);
}
