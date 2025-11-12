// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./FeralfileArtworkV4_1.sol";

contract FeralfileExhibitionV4_5 is FeralfileExhibitionV4_1 {
    // Struct to represent token ID to index mapping
    struct TokenIndex {
        uint256 tokenId;
        uint256 index;
    }

    // Token ID prefix shard (exhibition ID without dashes)
    string public tokenIdPrefixShard;

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
        uint256[] memory seriesMaxSupplies_,
        string memory tokenIdPrefixShard_
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
    {
        require(
            bytes(tokenIdPrefixShard_).length == TOKEN_ID_PREFIX_LENGTH,
            "FeralfileExhibitionV4_5: invalid prefix shard length"
        );
        tokenIdPrefixShard = tokenIdPrefixShard_;
    }

    /// @notice Calculate artwork indexes from token IDs and return mapping
    /// @param tokenIds_ - array of token IDs to calculate indexes for
    /// @return tokenIndexes - array of TokenIndex structs containing tokenId to index mappings
    function tokenIndexesByOwner(uint256[] calldata tokenIds_)
        external
        view
        returns (TokenIndex[] memory tokenIndexes)
    {
        tokenIndexes = new TokenIndex[](tokenIds_.length);
        
        for (uint256 i = 0; i < tokenIds_.length; ++i) {
            uint256 tokenId = tokenIds_[i];
            
            // Get the artwork to retrieve seriesId
            Artwork memory artwork = _allArtworks[tokenId];
            require(
                artwork.tokenId != 0,
                "FeralfileExhibitionV4_5: token does not exist"
            );
            
            // Calculate the artwork index from tokenId
            uint256 artworkIndex = _calculateArtworkIndex(tokenId, artwork.seriesId);
            
            // Create TokenIndex mapping
            tokenIndexes[i] = TokenIndex({
                tokenId: tokenId,
                index: artworkIndex
            });
        }
        
        return tokenIndexes;
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

