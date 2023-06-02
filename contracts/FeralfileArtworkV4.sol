// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "./Authorizable.sol";
import "./UpdateableOperatorFilterer.sol";

contract FeralfileExhibitionV4 is
    IERC165,
    ERC721Enumerable,
    Authorizable,
    UpdateableOperatorFilterer
{
    using Strings for uint256;

    // version code of contract
    string public constant codeVersion = "FeralfileExhibitionV4";

    // burnable
    bool public isBurnable;

    // bridgeable
    bool public isBridgeable;

    // token base URI
    string internal _tokenBaseURI;

    // contract URI
    string private _contractURI;

    constructor(
        string memory name_,
        string memory symbol_,
        bool isBurnable_,
        bool isBridgeable_
    ) ERC721(name_, symbol_) {
        isBurnable = isBurnable_;
        isBridgeable = isBridgeable_;
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view virtual override(ERC721Enumerable, IERC165) returns (bool) {
        return
            interfaceId == type(IERC721Enumerable).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /// @notice Return a count of series registered in this exhibition
    function totalSeries() public view virtual returns (uint256) {}

    /// @notice Update the IPFS cid of an artwork to a new value
    function updateArtworkIPFSCid(
        uint256 tokenId,
        string memory ipfsCID
    ) external onlyAuthorized {}

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    function tokenURI(
        uint256 tokenId
    ) public view virtual override returns (string memory) {
        return super.tokenURI(tokenId);
    }

    /// @notice Update the base URI for all tokens
    function setTokenBaseURI(string memory baseURI_) external onlyAuthorized {
        _tokenBaseURI = baseURI_;
    }

    /// @notice A URL for the opensea storefront-level metadata
    function contractURI() public view returns (string memory) {
        return _contractURI;
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

    /// @notice burn artworks
    /// @param artworkIDs_ - the list of artwork id will be burned
    function burnArtowks(uint256[] memory artworkIDs_) public {}

    event NewArtwork(
        address indexed owner,
        uint256 indexed seriesID,
        uint256 indexed tokenIndex
    );
    event BurnArtwork(uint256 indexed artworkID);
}
