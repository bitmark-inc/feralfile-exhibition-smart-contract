// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import { FeralfileExhibitionV4_1 } from "./FeralfileArtworkV4_1.sol";
import { Base64 } from "@openzeppelin/contracts/utils/Base64.sol";

contract FeralfileExhibitionV4_2 is FeralfileExhibitionV4_1 {

    error TokenIDNotFound();
    error MintNotEnabled();
    error FunctionNotSupported();
    error EmptyURI();

    struct TokenInfo {
        string imageURI;
        bytes parameters;
    }

    struct MintDataWithIndex {
        uint256 seriesId;
        uint256 tokenId;
        address owner;
        uint256 tokenIndex;
    }

    mapping(uint256 => TokenInfo) public _tokenInfos; // tokenID => tokenInfo
    mapping(uint256 => uint256) public _tokenIndexes; // tokenID => tokenIndex

    string public tokenPlaceholderAnimationURI;
    string public tokenPlaceholderImageURI;

    string public artworkFileURI;

    constructor(
        string memory name_,
        string memory symbol_,
        bool burnable_,
        bool bridgeable_,
        address signer_,
        address vault_,
        address costReceiver_,
        string memory contractURI_,
        uint256[] memory seriesIds_,
        uint256[] memory seriesMaxSupplies_
    )
        FeralfileExhibitionV4_1(
            name_,
            symbol_,
            burnable_,
            bridgeable_,
            signer_,
            vault_,
            costReceiver_,
            contractURI_,
            seriesIds_,
            seriesMaxSupplies_
        )
    {}
    
    /// @notice Mint new collection of Artwork
    /// @dev the function iterates over the array of MintDataWithIndex to call the internal function _mintArtwork
    /// @param data an array of MintDataWithIndex
    function mintArtworksWithIndex(
        MintDataWithIndex[] calldata data
    ) external onlyAuthorized {
        if (!mintable) {
            revert MintNotEnabled();
        }

        for (uint256 i = 0; i < data.length; i++) {
            _tokenIndexes[data[i].tokenId] = data[i].tokenIndex;
            _mintArtwork(data[i].seriesId, data[i].tokenId, data[i].owner);
        }
    }

    /// @notice overrdie revert mint new collection of Artwork
    function mintArtworks(MintData[] calldata) external override view onlyAuthorized {
        revert FunctionNotSupported();
    }

    /// @notice Update the imageURI & parameters of an edition to a new value
    function updateTokenInformation(uint256 tokenId, string calldata imageURI, bytes calldata parameters)
        external
        onlyAuthorized
    {
        if (!_exists(tokenId)) {
            revert TokenIDNotFound();
        }

        _tokenInfos[tokenId] = TokenInfo(imageURI, parameters);
    }


    /// @notice Update the base URI for all tokens
    function setArtworkFileURI(string calldata uri) external virtual onlyOwner {
        if (bytes(uri).length == 0) {
            revert EmptyURI();
        }

        artworkFileURI = uri;
    }

    /// @notice Update tokenPlaceholderImageURI
    function setTokenPlaceholderImageURI(string calldata uri) external virtual onlyOwner {
        if (bytes(uri).length == 0) {
            revert EmptyURI();
        }

        tokenPlaceholderImageURI = uri;
    }

    /// @notice Update tokenPlaceholderAnimationURI
    function setTokenPlaceholderAnimationURI(string calldata uri) external virtual onlyOwner {
        if (bytes(uri).length == 0) {
            revert EmptyURI();
        }

        tokenPlaceholderAnimationURI = uri;
    }

    /// @notice _buildAnimationURI returns the animation url
    /// @param parameters - the token paramters
    function _buildAnimationURI(bytes memory parameters)
        private
        view
        returns (string memory)
    {
        if (parameters.length == 0) {
            return tokenPlaceholderAnimationURI;
        }

        return _buildIframe(Base64.encode(parameters), artworkFileURI);
    }

    /// @notice _buildImageURI returns the image url
    /// @param imageURI - the token imageURI
    function _buildImageURI(string memory imageURI)
        private
        view
        returns (string memory)
    {
        if (bytes(imageURI).length == 0) {
            return tokenPlaceholderImageURI;
        }

        return imageURI;
    }

    /// @notice _buildIframe returns a base64 encoded data for ff-frame
    /// @param tokenParams - the paramteres which would bring into the artwork
    /// @param artworkURI - the artwork file URI to be loaded into the iframe
    function _buildIframe(string memory tokenParams, string memory artworkURI)
        private
        pure
        returns (string memory)
    {
        return 
            Base64.encode(
                abi.encodePacked(
                    "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\">",
                    "<script>var defaultArtworkData=",
                    tokenParams,
                    ";let allowOrigins = {\"https://feralfile.com\": !0};function initData(){document.getElementById(\"mainframe\").contentWindow.postMessage(defaultArtworkData, \"*\");}</script>",
                    "</head><body style=\"overflow-x:hidden;padding:0;margin:0;width: 100%;\" onload=\"initData()\"><iframe id=\"mainframe\" style=\"display:block;padding:0;margin:0;border:none;width:100%;height:100vh;\" src=\"",
                    artworkURI,
                    "\"></iframe></body></html>"
                )
            );
    }

    /// @notice _buildDataURL returns a base64 encoded data for ff-frame
    /// @param tokenID - the token ID for building artwork data
    function _buildDataURL(uint256 tokenID) private view returns (string memory) {
        TokenInfo memory tokenInfo = _tokenInfos[tokenID];
        uint256 tokenIndex = _tokenIndexes[tokenID];

        string memory tokenName = string(
            abi.encodePacked(name(), " #", tokenIndex+1)
        );

        string memory json = string(
            abi.encodePacked(
                "{\"name\":\"", tokenName,
                "\", \"external_url\":\"https://feralfile.com\", \"image\":\"", _buildImageURI(tokenInfo.imageURI),
                "\", \"animation_url\":\"data:text/html;base64,", _buildAnimationURI(tokenInfo.parameters),
                "\"}"
            )
        );

        return string(
            abi.encodePacked(
                "data:application/json;base64,", Base64.encode(bytes(json))
            )
        );
    }

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    function tokenURI(uint256 tokenId)
        public
        view
        virtual
        override
        returns (string memory)
    {
        if (!_exists(tokenId)) {
            revert TokenIDNotFound();
        }

        return _buildDataURL(tokenId);
    }
}
