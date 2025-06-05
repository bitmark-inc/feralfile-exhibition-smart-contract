// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {FeralfileExhibitionV4_1} from "./FeralfileArtworkV4_1.sol";
import {SSTORE2} from "@0xsequence/sstore2/contracts/SSTORE2.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {RendererStorageV0} from "./RendererStorageV0.sol";

import {LibBytes} from "./LibBytes.sol";

contract FeralfileExhibitionV4_4 is FeralfileExhibitionV4_1 {
    using Strings for uint256;
    using RendererStorageV0 for bytes;
    using LibBytes for address[];

    uint256 public constant CHUNK_SIZE = 24000; // Single SSTORE2 write up to 24,576 bytes
    uint256 public constant RENDERER_BLOB_MAX_SIZE = 72000; // 3x CHUNK_SIZE (24,000) costs around 15.5M gas as target block gas limit

    //----------------------------------------------------------
    // Structs
    //----------------------------------------------------------
    struct RendererTokenData {
        string imageURI;
        string textureURI;
        uint256 index;
    }

    //----------------------------------------------------------
    // Errors
    //----------------------------------------------------------
    error ErrSeriesDoesNotExist(uint256 seriesId);
    error ErrTokenDoesNotExist(uint256 tokenId);
    error ErrLengthMismatch();
    error ErrEmptyArray();
    error ErrEmptyString();
    error ErrEmptyBytes();
    error ErrSeriesHasRenderer();
    error ErrSeriesHasNoRenderer();
    error ErrRendererBlobTooLarge();
    error ErrUnsupportedCharacters();
    error ErrEmptyRenderer();
    error ErrEmptySeriesName();

    //----------------------------------------------------------
    // Events
    //----------------------------------------------------------
    event NewSeriesName(uint256 indexed seriesId, string name);
    event NewSeriesRenderer(uint256 indexed seriesId, address[] pointers);
    event DeleteSeriesRenderer(uint256 indexed seriesId);
    event NewRendererTokenData(uint256 indexed tokenId, RendererTokenData data);

    //----------------------------------------------------------
    // State Variables
    //----------------------------------------------------------

    /**
     * @dev A mapping of series IDs to a list of addresses that can be used to render the series
     */
    mapping(uint256 => address[]) private _seriesRendererPointers;

    /**
     * @dev A mapping of series IDs to the name of the series
     */
    mapping(uint256 => string) public seriesNames;

    /**
     * @dev A mapping of token IDs to the data used to render the token
     */
    mapping(uint256 => RendererTokenData) public rendererTokenData;

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

    //----------------------------------------------------------
    // Modifiers
    //----------------------------------------------------------
    modifier seriesExists(uint256 seriesId) {
        if (!_checkSeriesExists(seriesId)) {
            revert ErrSeriesDoesNotExist(seriesId);
        }
        _;
    }

    //----------------------------------------------------------
    // External Functions
    //----------------------------------------------------------

    /// @notice Set the series renderer for a series
    /// @param seriesId The series ID to set
    /// @param blob The renderer bytes
    function setSeriesRenderer(
        uint256 seriesId,
        bytes calldata blob
    ) external onlyAuthorized seriesExists(seriesId) {
        if (blob.length == 0) {
            revert ErrEmptyBytes();
        }
        if (blob.length > RENDERER_BLOB_MAX_SIZE) {
            revert ErrRendererBlobTooLarge();
        }

        // cleanup the old renderer
        if (_seriesRendererPointers[seriesId].length > 0) {
            delete _seriesRendererPointers[seriesId];
            emit DeleteSeriesRenderer(seriesId);
        }

        // store the new renderer
        uint256 offset = 0;
        while (offset < blob.length) {
            uint256 len = blob.length - offset;
            if (len > CHUNK_SIZE) len = CHUNK_SIZE;

            // Copy slice to new bytes
            address ptr = SSTORE2.write(blob[offset:offset + len]);
            _seriesRendererPointers[seriesId].push(ptr);
            offset += len;
        }

        emit NewSeriesRenderer(seriesId, _seriesRendererPointers[seriesId]);
    }

    /// @notice Get the series renderer from the SSTORE2 pointers
    /// @param seriesId The series ID to read
    /// @return renderer The renderer bytes
    function getSeriesRenderer(
        uint256 seriesId
    ) external view seriesExists(seriesId) returns (string memory) {
        if (!_hasRenderer(seriesId)) {
            revert ErrSeriesHasNoRenderer();
        }

        return string(_readSeriesRenderer(seriesId));
    }

    /// @notice Get the token URI
    /// @param tokenId The token ID to read
    /// @return uri The token URI
    function tokenURI(
        uint256 tokenId
    ) public view override returns (string memory) {
        // Validate token exists
        if (!_exists(tokenId)) {
            revert ErrTokenDoesNotExist(tokenId);
        }

        // Validate token is a renderer token
        uint256 seriesId = _allArtworks[tokenId].seriesId;
        if (!_hasRenderer(seriesId)) {
            return super.tokenURI(tokenId);
        }

        // Validate renderer token data is set
        RendererTokenData memory data = rendererTokenData[tokenId];
        if (bytes(data.imageURI).length == 0) {
            revert ErrEmptyRenderer();
        }

        // Validate series name is set
        string memory seriesName = seriesNames[seriesId];
        if (bytes(seriesName).length == 0) {
            revert ErrEmptySeriesName();
        }

        // Read renderer
        bytes memory renderer = _readSeriesRenderer(seriesId);
        string memory name = string(
            abi.encodePacked(seriesName, " #", Strings.toString(data.index))
        );

        return renderer.tokenURI(data.textureURI, data.imageURI, name);
    }

    /// @notice Set the names for a list of series IDs
    /// @param seriesIds The series IDs to set
    /// @param names The names to set
    function setSeriesNames(
        uint256[] calldata seriesIds,
        string[] calldata names
    ) external onlyAuthorized {
        // Validation
        if (seriesIds.length != names.length) {
            revert ErrLengthMismatch();
        }

        for (uint256 i = 0; i < seriesIds.length; ) {
            if (bytes(names[i]).length == 0) {
                revert ErrEmptyString();
            }

            _checkForUnsupportedCharacters(names[i]);

            unchecked {
                ++i;
            }
        }

        _validateRendererSeries(seriesIds);

        // Set
        for (uint256 i = 0; i < seriesIds.length; ) {
            seriesNames[seriesIds[i]] = names[i];

            // Emit event
            emit NewSeriesName(seriesIds[i], names[i]);

            unchecked {
                ++i;
            }
        }
    }

    /// @notice Set the renderer token data for a list of token IDs
    /// @param tokenIds The token IDs to set
    /// @param data The renderer token data to set
    function setRendererTokenData(
        uint256[] calldata tokenIds,
        RendererTokenData[] calldata data
    ) external onlyAuthorized {
        // Validation
        if (tokenIds.length != data.length) {
            revert ErrLengthMismatch();
        }

        for (uint256 i = 0; i < data.length; ) {
            _validateRendererTokenData(data[i]);

            unchecked {
                ++i;
            }
        }

        _validateRendererTokens(tokenIds);

        // Set
        for (uint256 i = 0; i < tokenIds.length; ) {
            rendererTokenData[tokenIds[i]] = data[i];

            // Emit event
            emit NewRendererTokenData(tokenIds[i], data[i]);

            unchecked {
                ++i;
            }
        }
    }

    //----------------------------------------------------------
    // Internal Functions
    //----------------------------------------------------------

    /// @notice Read the series renderer from the SSTORE2 pointers
    /// @param seriesId The series ID to read
    /// @return renderer The renderer bytes
    function _readSeriesRenderer(
        uint256 seriesId
    ) private view returns (bytes memory renderer) {
        return _seriesRendererPointers[seriesId].sstore2Join();
    }

    /// @notice Check if the series has a renderer
    /// @param seriesId The series ID to check
    /// @return hasRenderer True if the series has a renderer, false otherwise
    function _hasRenderer(uint256 seriesId) private view returns (bool) {
        return _seriesRendererPointers[seriesId].length > 0;
    }

    /// @notice Check if the series exists
    /// @param seriesId The series ID to check
    /// @return exists True if the series exists, false otherwise
    function _checkSeriesExists(uint256 seriesId) private view returns (bool) {
        return _seriesMaxSupplies[seriesId] > 0;
    }

    /// @notice Validate the tokens are renderer tokens
    /// @param tokenIds The token IDs to validate
    function _validateRendererTokens(uint256[] calldata tokenIds) private view {
        if (tokenIds.length == 0) {
            revert ErrEmptyArray();
        }

        for (uint256 i = 0; i < tokenIds.length; ) {
            if (!_exists(tokenIds[i])) {
                revert ErrTokenDoesNotExist(tokenIds[i]);
            }
            if (!_hasRenderer(_allArtworks[tokenIds[i]].seriesId)) {
                revert ErrSeriesHasNoRenderer();
            }

            unchecked {
                ++i;
            }
        }
    }

    /// @notice Validate the series are renderer series
    /// @param seriesIds The series IDs to validate
    function _validateRendererSeries(
        uint256[] calldata seriesIds
    ) private view {
        if (seriesIds.length == 0) {
            revert ErrEmptyArray();
        }

        for (uint256 i = 0; i < seriesIds.length; ) {
            if (!_checkSeriesExists(seriesIds[i])) {
                revert ErrSeriesDoesNotExist(seriesIds[i]);
            }
            if (!_hasRenderer(seriesIds[i])) {
                revert ErrSeriesHasNoRenderer();
            }

            unchecked {
                ++i;
            }
        }
    }

    /// @notice Validate the renderer token data
    /// @param data The renderer token data to validate
    function _validateRendererTokenData(
        RendererTokenData memory data
    ) private pure {
        if (bytes(data.imageURI).length == 0) {
            revert ErrEmptyString();
        }

        if (bytes(data.textureURI).length == 0) {
            revert ErrEmptyString();
        }

        // Check for unsupported characters in URIs
        _checkForUnsupportedCharacters(data.imageURI);
        _checkForUnsupportedCharacters(data.textureURI);
    }

    /// @notice Check if a string contains unsupported characters (quotes and backslashes)
    /// @param str The string to check
    function _checkForUnsupportedCharacters(string memory str) private pure {
        bytes memory strBytes = bytes(str);
        for (uint256 i = 0; i < strBytes.length; ) {
            // Check for double quote ("), backslash (\), and control characters
            bytes1 b = strBytes[i];
            if (b == 0x22 || b == 0x5C || b < 0x20) {
                revert ErrUnsupportedCharacters();
            }

            unchecked {
                ++i;
            }
        }
    }
}
