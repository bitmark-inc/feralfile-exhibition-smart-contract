// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract FFV2 {
    struct FF2Artwork {
        string title;
        string artistName;
        string fingerprint;
        uint256 editionSize;
    }

    mapping(uint256 => FF2Artwork) public artworks;

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public {}
}

contract FFV33 {
    struct FF3Artwork {
        string title;
        string artistName;
        string fingerprint;
        uint256 editionSize;
        uint256 AEAmount;
        uint256 PPAmount;
    }
    mapping(uint256 => FF3Artwork) public artworks;

    function mintArtworkEdition(uint256 _artworkID, address _owner) public {}
}

contract MoMABurn is Ownable {
    using SafeMath for uint256;
    address constant DEAD_ADDRESS = 0x000000000000000000000000000000000000dEaD;

    address public burnContractAddress; // Address of token need to burn
    address public tier1ContractAddress; // Address of new token for Tier 1
    address public tier2ContractAddress; // Address of new token for Tier 2
    address public tier3ContractAddress; // Address of new token for Tier 3

    uint256 private burnArtworkID; // Artwork ID will be burned
    uint256 private tier1ArtworkID; // Artwork ID of Tier 1
    uint256 private tier2ArtworkID; // Artwork ID of Tier 2
    uint256 private tier3ArtworkID; // Artwork ID of Tier 3

    bool public isBurnEnabled;

    struct FFArtworkInfo {
        string fingerprint;
        uint256 editionSize;
    }

    constructor() {
        isBurnEnabled = false;
    }

    /// @notice setBurnAndMintParams use for update burn and mint contract addresses and artwork IDs
    /// @param _burnContractAddress - new burn contract address
    /// @param _tier1Address - new burn contract address for tier 1
    /// @param _tier2Address - new burn contract address for tier 2
    /// @param _tier3Address - new burn contract address for tier 3
    /// @param _burnArtworkID - artwork ID will be burned
    /// @param _tier1ArtworkID - new artwork ID for tier 1
    /// @param _tier2ArtworkID - new artwork ID for tier 2
    /// @param _tier3ArtworkID - new artwork ID for tier 3
    function setBurnAndMintParams(
        address _burnContractAddress,
        address _tier1Address,
        address _tier2Address,
        address _tier3Address,
        uint256 _burnArtworkID,
        uint256 _tier1ArtworkID,
        uint256 _tier2ArtworkID,
        uint256 _tier3ArtworkID
    ) public onlyOwner {
        require(
            _burnContractAddress != address(0) &&
                _tier1Address != address(0) &&
                _tier2Address != address(0) &&
                _tier3Address != address(0),
            "Invalid contract address"
        );
        FFArtworkInfo memory burnArtwork = getV2ContractArtwork(
            _burnContractAddress,
            _burnArtworkID
        );
        FFArtworkInfo memory tier1Artwork = getV3ContractArtwork(
            _tier1Address,
            _tier1ArtworkID
        );
        FFArtworkInfo memory tier2Artwork = getV3ContractArtwork(
            _tier2Address,
            _tier2ArtworkID
        );
        FFArtworkInfo memory tier3Artwork = getV3ContractArtwork(
            _tier3Address,
            _tier3ArtworkID
        );

        require(
            burnArtwork.editionSize > 0 &&
                bytes(burnArtwork.fingerprint).length != 0 &&
                tier1Artwork.editionSize > 0 &&
                bytes(tier1Artwork.fingerprint).length != 0 &&
                tier2Artwork.editionSize > 0 &&
                bytes(tier2Artwork.fingerprint).length != 0 &&
                tier3Artwork.editionSize > 0 &&
                bytes(tier3Artwork.fingerprint).length != 0,
            "Invalid minting parameters"
        );

        burnContractAddress = _burnContractAddress;
        tier1ContractAddress = _tier1Address;
        tier2ContractAddress = _tier2Address;
        tier3ContractAddress = _tier3Address;
        burnArtworkID = _burnArtworkID;
        tier1ArtworkID = _tier1ArtworkID;
        tier2ArtworkID = _tier2ArtworkID;
        tier3ArtworkID = _tier3ArtworkID;
    }

    /// @notice setBurnEnabled use for update burn enable status
    /// @param _isBurnEnabled - the condition is true or false
    function setBurnEnabled(bool _isBurnEnabled) public onlyOwner {
        isBurnEnabled = _isBurnEnabled;
    }

    /// @notice burnAndMint use for burn old editions and mint new edition
    /// @param _burnedEditions - array of edition IDs to burn
    function burnAndMint(uint256[] memory _burnedEditions) external {
        require(isBurnEnabled, "Burn and mint status is not enabled");

        require(
            _burnedEditions.length == 3 ||
                _burnedEditions.length == 6 ||
                _burnedEditions.length == 9,
            "Invalid number of burning editions"
        );

        FFArtworkInfo memory burnArtwork = getV2ContractArtwork(
            burnContractAddress,
            burnArtworkID
        );

        for (uint256 i = 0; i < _burnedEditions.length; i++) {
            // Check duplicate with others
            for (uint256 j = i + 1; j < _burnedEditions.length; j++) {
                if (_burnedEditions[i] == _burnedEditions[j]) {
                    revert("Invalid burning editions");
                }
            }

            require(
                bytes(burnArtwork.fingerprint).length != 0 &&
                    burnArtwork.editionSize > 0 &&
                    _burnedEditions[i] >= burnArtworkID &&
                    _burnedEditions[i] <=
                    burnArtworkID + burnArtwork.editionSize,
                "Edition ID is not support for burning"
            );

            FFV2(burnContractAddress).safeTransferFrom(
                _msgSender(),
                DEAD_ADDRESS,
                _burnedEditions[i]
            );
        }

        address mintContractAddr = tier1ContractAddress;
        uint256 artworkID = tier1ArtworkID;

        if (_burnedEditions.length == 6) {
            mintContractAddr = tier2ContractAddress;
            artworkID = tier2ArtworkID;
        }

        if (_burnedEditions.length == 9) {
            mintContractAddr = tier3ContractAddress;
            artworkID = tier3ArtworkID;
        }

        FFV33(mintContractAddr).mintArtworkEdition(artworkID, _msgSender());
    }

    /// @notice getV2ContractArtwork use for get artwork info from V2 contract
    /// @param _contractAddress - contract address
    /// @param _artworkID - artwork ID
    function getV2ContractArtwork(
        address _contractAddress,
        uint256 _artworkID
    ) private view returns (FFArtworkInfo memory) {
        (, , string memory fingerprint, uint256 editionSize) = FFV2(
            _contractAddress
        ).artworks(_artworkID);
        return FFArtworkInfo(fingerprint, editionSize);
    }

    /// @notice getV3ContractArtwork use for get artwork info from V3 contract
    /// @param _contractAddress - contract address
    /// @param _artworkID - artwork ID
    function getV3ContractArtwork(
        address _contractAddress,
        uint256 _artworkID
    ) private view returns (FFArtworkInfo memory) {
        (, , string memory fingerprint, uint256 editionSize, , ) = FFV33(
            _contractAddress
        ).artworks(_artworkID);
        return FFArtworkInfo(fingerprint, editionSize);
    }
}
