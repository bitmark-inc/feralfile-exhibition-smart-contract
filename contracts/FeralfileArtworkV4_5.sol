// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {FeralfileExhibitionV4_1} from "./FeralfileArtworkV4_1.sol";

contract FeralfileExhibitionV4_5 is FeralfileExhibitionV4_1 {

    // Constants for token ID calculation
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
    /// @return tokenIds - array of token Id in the same series owned by the same owner
    function seriesArtworksOfOwner(uint256 tokenId_)
        external
        view
        returns (uint256[] memory tokenIds)
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
        tokenIds = new uint256[](matchCount);
        
        // Second pass: populate the result array with token IDs
        uint256 resultIndex = 0;
        for (uint256 i = 0; i < allTokenIds.length; ++i) {
            uint256 tokenId = allTokenIds[i];
            Artwork memory artwork = _allArtworks[tokenId];
            if (artwork.seriesId == targetSeriesId) {
                tokenIds[resultIndex] = tokenId;
                ++resultIndex;
            }
        }
        
        return tokenIds;
    }
}

