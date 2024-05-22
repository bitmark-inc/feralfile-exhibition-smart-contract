// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import { FeralfileExhibitionV4_1 } from "./FeralfileArtworkV4_1.sol";

contract FeralfileExhibitionV4_2 is FeralfileExhibitionV4_1 {

    error ArtworkEditionNotFound();

    mapping(uint256 => string) public artworkParameters; // artworkEditionID => parameters as string

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

    /// @notice Update the parameters of an edition to a new value
    function updateArtworkParameters(uint256 tokenId, string calldata parameters)
        external
        onlyAuthorized
    {
        if (!_exists(tokenId)) {
            revert ArtworkEditionNotFound();
        }

        artworkParameters[tokenId] = parameters;
    }

    /// @notice Get the parameters of an edition
    function getArtworkParameters(
        uint256 tokenId
    ) public view virtual returns (string memory) {
        if (!_exists(tokenId)) {
            revert ArtworkEditionNotFound();
        }

        return artworkParameters[tokenId];
    }
}
