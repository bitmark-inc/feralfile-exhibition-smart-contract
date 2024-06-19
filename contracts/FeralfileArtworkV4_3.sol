// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {FeralfileExhibitionV4_1} from "./FeralfileArtworkV4_1.sol";

contract FeralfileExhibitionV4_3 is FeralfileExhibitionV4_1 {
    error InvalidOwner();
    error TokenIsNonMergeable();
    error InvalidLength();
    error InvalidMergingStatus();

    struct MergeArtworkInfo {
        uint256 singleSeriesId;
        uint256 mergedSeriesId;
        uint256 nextTokenId;
    }
    MergeArtworkInfo private mergeArtworkInfo;

    // merging
    bool internal _merging;

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
        MergeArtworkInfo memory mergeArtworkInfo_
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
        mergeArtworkInfo = mergeArtworkInfo_;
    }

    /// @notice burns multiples mergeable artworks and mint a new artworks
    /// @param tokenIds - list of tokenIds to be burned
    function mergeArtworks(uint256[] calldata tokenIds) external {
        if (!_merging) {
            revert InvalidMergingStatus();
        }

        if (tokenIds.length < 2) {
            revert InvalidLength();
        }

        // Burn artworks
        for (uint256 i = 0; i < tokenIds.length; i++) {
            uint256 tokenId = tokenIds[i];
            Artwork memory artwork = _allArtworks[tokenId];

            if (
                artwork.seriesId != mergeArtworkInfo.singleSeriesId &&
                artwork.seriesId != mergeArtworkInfo.mergedSeriesId
            ) {
                revert TokenIsNonMergeable();
            }

            if (ownerOf(tokenId) != _msgSender()) {
                revert InvalidOwner();
            }

            _burnArtwork(tokenId);
        }

        // Mint new artwork
        _mintArtwork(
            mergeArtworkInfo.mergedSeriesId,
            mergeArtworkInfo.nextTokenId,
            _msgSender()
        );
        mergeArtworkInfo.nextTokenId++;
    }

    /// @notice Start token merging
    function startMerging() external onlyOwner {
        if (_merging) {
            revert InvalidMergingStatus();
        }

        _merging = true;
    }

    /// @notice Pause token merging
    function pauseMerging() public onlyOwner {
        if (!_merging) {
            revert InvalidMergingStatus();
        }

        _merging = false;
    }
}
