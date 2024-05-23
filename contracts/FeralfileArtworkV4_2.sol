// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import { FeralfileExhibitionV4_1 } from "./FeralfileArtworkV4_1.sol";
import { Base64 } from "@openzeppelin/contracts/utils/Base64.sol";

contract FeralfileExhibitionV4_2 is FeralfileExhibitionV4_1 {

    error TokenIDNotFound();
    error MintNotEnabled();

    struct TokenInfo {
        string thumbnail;
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

    string public tokenPlaceholderURL;
    string public tokenPlaceholderThubmnail;

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
    /// @dev the function iterates over the array of MintDataWithIndex to call the internal function mintArtworksWithIndex
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

    /// @notice Update the thumbnail & parameters of an edition to a new value
    function updateTokenInformation(uint256 tokenId, string calldata thumbnail, bytes calldata parameters)
        external
        onlyAuthorized
    {
        if (!_exists(tokenId)) {
            revert TokenIDNotFound();
        }

        _tokenInfos[tokenId] = TokenInfo(thumbnail, parameters);
    }

    /// @notice _buildArtworkData returns an object of artwork which would push to the actually artwork
    /// @param parameters - the parameters for building artwork data
    function _buildArtworkData(bytes memory parameters)
        private
        pure
        returns (string memory)
    {
        return Base64.encode(parameters);
    }

    /// @notice _buildPreviewURL returns the preview url
    /// @param parameters - the token paramters
    function _buildPreviewURL(bytes memory parameters)
        private
        view
        returns (string memory)
    {
        if (parameters.length == 0) {
            return tokenPlaceholderURL;
        }

        return _buildIframe(_buildArtworkData(parameters), tokenBaseURI);
    }

    /// @notice _buildThumbnailURL returns the preview url
    /// @param thumbnail - the token thumbnail
    function _buildThumbnailURL(string memory thumbnail)
        private
        view
        returns (string memory)
    {
        if (bytes(thumbnail).length == 0) {
            return tokenPlaceholderThubmnail;
        }

        return thumbnail;
    }

    /// @notice _buildIframe returns a base64 encoded data for ff-frame
    /// @param artworkData - the artwork data which would bring into the artwork
    /// @param iframeURI - the artwork URL to be loaded into the iframe
    function _buildIframe(string memory artworkData, string memory iframeURI)
        private
        pure
        returns (string memory)
    {
        return 
            Base64.encode(
                abi.encodePacked(
                    "<!DOCTYPE html><html lang=\"en\"><head><script> var defaultArtworkData= ",
                    artworkData,
                    "</script><script>",
                    "let allowOrigins={\"https://feralfile.com\":!0};function resizeIframe(t){let e=document.getElementById(\"mainframe\");t&&(e.style.minHeight=t+\"px\")}",
                    "function initData(){pushArtworkDataToIframe(defaultArtworkData)}function pushArtworkDataToIframe(t){t&&document.getElementById(\"mainframe\").contentWindow.postMessage(t,\"*\")}",
                    "function updateArtworkData(t){document.getElementById(\"mainframe\").contentWindow.postMessage(t,\"*\")}",
                    "window.addEventListener(\"message\",function(t){allowOrigins[t.origin]?\"update-artwork-data\"===t.data.type&&updateArtworkData(t.data.artworkData):\"object\"==typeof t.data&&\"resize-iframe\"===t.data.type&&resizeIframe(t.data.newHeight)});</script>",
                    "</head><body style=\"overflow-x:hidden;padding:0;margin:0;width: 100%;\" onload=\"initData()\">",
                    "<iframe id=\"mainframe\" style=\"display:block;padding:0;margin:0;border:none;width:100%;height:100vh;\" src=\"",
                    iframeURI,
                    "\"></iframe> </body></html>"
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
                "\", \"external_url\":\"https://feralfile.com\", \"image\":\"", _buildThumbnailURL(tokenInfo.thumbnail),
                "\", \"animation_url\":\"data:text/html;base64,", _buildPreviewURL(tokenInfo.parameters),
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
