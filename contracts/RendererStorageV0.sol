// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {Base64} from "@openzeppelin/contracts/utils/Base64.sol";

library RendererStorageV0 {
    //----------------------------------------------------------
    // Errors
    //----------------------------------------------------------
    error ErrEmptyField();

    //----------------------------------------------------------
    // Constants
    //----------------------------------------------------------
    bytes private constant PLACEHOLDER = bytes("{{TEXTURE_JSON}}");

    /// @notice Generate the token URI for a renderer
    /// @param renderer The renderer bytes
    /// @param textureURI The texture URI
    /// @param imageURI The image URI
    /// @param tokenName The token name
    /// @return The token URI
    function tokenURI(
        bytes memory renderer,
        string memory textureURI,
        string memory imageURI,
        string memory tokenName
    ) public pure returns (string memory) {
        if (
            renderer.length == 0 ||
            bytes(textureURI).length == 0 ||
            bytes(imageURI).length == 0 ||
            bytes(tokenName).length == 0
        ) {
            revert ErrEmptyField();
        }

        // inject the texture URI into the renderer
        bytes memory patchedRenderer = _injectTexture(renderer, textureURI);
        string memory rendererB64 = Base64.encode(patchedRenderer);

        // prettier-ignore
        bytes memory json = abi.encodePacked(
            "{\"animation_url\":\"data:text/html;base64,",
            rendererB64,
            "\",\"image\":\"",
            imageURI,
            "\",\"name\":\"",
            tokenName,
            "\"}"
        );

        return
            string(
                abi.encodePacked(
                    "data:application/json;base64,",
                    Base64.encode(json)
                )
            );
    }

    /// @notice Replace the first `{{TEXTURE_JSON}}` occurrence with the provided url.
    /// @param renderer The renderer bytes to modify
    /// @param url The URL to inject
    /// @return The modified renderer bytes
    function _injectTexture(
        bytes memory renderer,
        string memory url
    ) internal pure returns (bytes memory) {
        bytes memory urlBytes = bytes(url);
        uint256 rendererLen = renderer.length;
        uint256 placeholderLen = PLACEHOLDER.length;

        // Early return for edge cases
        if (rendererLen < placeholderLen || urlBytes.length == 0) {
            return renderer;
        }

        // Find placeholder with optimized search
        uint256 pos = type(uint256).max;
        uint256 searchLimit = rendererLen - placeholderLen + 1;

        for (uint256 i = 0; i < searchLimit; i++) {
            // Quick first-byte check
            if (renderer[i] == PLACEHOLDER[0]) {
                // Full comparison only if first byte matches
                bool found = true;
                for (uint256 j = 1; j < placeholderLen; j++) {
                    if (renderer[i + j] != PLACEHOLDER[j]) {
                        found = false;
                        break;
                    }
                }
                if (found) {
                    pos = i;
                    break;
                }
            }
        }

        // No placeholder found
        if (pos == type(uint256).max) {
            return renderer;
        }

        // Efficient concatenation using bytes.concat
        bytes memory prefix = new bytes(pos);
        bytes memory suffix = new bytes(rendererLen - pos - placeholderLen);

        // Copy prefix and suffix
        for (uint256 i = 0; i < pos; i++) {
            prefix[i] = renderer[i];
        }

        uint256 suffixStart = pos + placeholderLen;
        for (uint256 i = 0; i < suffix.length; i++) {
            suffix[i] = renderer[suffixStart + i];
        }

        return bytes.concat(prefix, urlBytes, suffix);
    }
}
