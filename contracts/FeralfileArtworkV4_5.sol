// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./FeralfileArtworkV4_1.sol";

contract FeralfileExhibitionV4_5 is FeralfileExhibitionV4_1 {
    // Struct to represent token ID to index mapping
    struct TokenIndex {
        uint256 tokenId;
        uint256 index;
    }

    // Constants for token ID calculation
    uint256 private constant TOKEN_ID_PREFIX_LENGTH = 32; // 128 bits in hex
    uint256 private constant SERIES_MULTIPLIER = 1000000;

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

    /// @notice Get the artwork indexes for all tokens in the same series owned by the same owner
    /// @param tokenId_ - the token ID to use as reference
    /// @return indexes - array of artwork indexes for tokens in the same series owned by the same owner
    function seriesArtworkIndexesOfOwner(uint256 tokenId_)
        external
        view
        returns (uint256[] memory indexes)
    {
        // Get the artwork to retrieve seriesId
        Artwork memory referenceArtwork = _allArtworks[tokenId_];
        require(
            referenceArtwork.tokenId != 0,
            "FeralfileExhibitionV4_5: token does not exist"
        );
        // Get the owner of the reference token
        address owner = ownerOf(tokenId_);
        
        uint256 targetSeriesId = referenceArtwork.seriesId;
        
        // Get all token IDs owned by the owner
        uint256[] memory allTokenIds = this.tokensOfOwner(owner);
        
        // First pass: count how many tokens belong to the same series
        uint256 matchCount = 0;
        for (uint256 i = 0; i < allTokenIds.length; ++i) {
            Artwork memory artwork = _allArtworks[allTokenIds[i]];
            if (artwork.seriesId == targetSeriesId) {
                ++matchCount;
            }
        }
        
        // Initialize the result array with the correct size
        indexes = new uint256[](matchCount);
        
        // Second pass: populate the result array with indexes
        uint256 resultIndex = 0;
        for (uint256 i = 0; i < allTokenIds.length; ++i) {
            uint256 tokenId = allTokenIds[i];
            Artwork memory artwork = _allArtworks[tokenId];
            
            if (artwork.seriesId == targetSeriesId) {
                // Calculate the artwork index from tokenId and store it directly
                indexes[resultIndex] = _calculateArtworkIndex(tokenId, artwork.seriesId);
                ++resultIndex;
            }
        }
        
        return indexes;
    }

    /// @notice Internal function to calculate artwork index from token ID
    /// @param tokenId_ - the token ID
    /// @param seriesId_ - the series ID of the artwork
    /// @return artworkIndex - the calculated artwork index
    function _calculateArtworkIndex(uint256 tokenId_, uint256 seriesId_)
        internal
        pure
        returns (uint256 artworkIndex)
    {
        // Token ID structure (in hex):
        // [prefix_shard (32 hex chars / 128 bits)][part2 (remaining bits)]
        // where part2 = seriesId * 1000000 + artworkIndex
        
        // Extract part2 by removing the prefix (shift right by 128 bits)
        // First, we need to get the lower part of the tokenId
        // Since tokenId is uint256 (256 bits), and prefix is 128 bits
        // part2 occupies the lower 128 bits
        
        // Create a mask for lower 128 bits: 2^128 - 1
        uint256 mask128 = type(uint128).max;
        uint256 part2 = tokenId_ & mask128;
        
        // Calculate: artworkIndex = part2 - (seriesId * SERIES_MULTIPLIER)
        uint256 seriesOffset = seriesId_ * SERIES_MULTIPLIER;
        require(
            part2 >= seriesOffset,
            "FeralfileExhibitionV4_5: invalid token ID structure"
        );
        
        artworkIndex = part2 - seriesOffset;
        
        return artworkIndex;
    }

    /// @notice Get artwork index for a single token ID
    /// @param tokenId_ - the token ID to get index for
    /// @return artworkIndex - the artwork index
    function getArtworkIndex(uint256 tokenId_)
        external
        view
        returns (uint256 artworkIndex)
    {
        Artwork memory artwork = _allArtworks[tokenId_];
        require(
            artwork.tokenId != 0,
            "FeralfileExhibitionV4_5: token does not exist"
        );
        
        return _calculateArtworkIndex(tokenId_, artwork.seriesId);
    }
}

