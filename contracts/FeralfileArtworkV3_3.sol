// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

import "./FeralfileArtworkV3_2.sol";

contract FeralfileExhibitionV3_3 is FeralfileExhibitionV3_2 {
    using Strings for uint256;
    using Counters for Counters.Counter;

    address public minter;
    Counters.Counter private tokenCounter;
    mapping(uint256 => string) private artworkCIDs; // artworkID => artwork cids

    modifier onlyMinter() {
        require(_msgSender() == minter, "Only minter can call this function");
        _;
    }

    constructor(
        string memory name_,
        string memory symbol_,
        uint256 secondarySaleRoyaltyBPS_,
        address royaltyPayoutAddress_,
        string memory contractURI_,
        string memory tokenBaseURI_,
        bool isBurnable_,
        bool isBridgeable_,
        address minter_
    )
        FeralfileExhibitionV3_2(
            name_,
            symbol_,
            secondarySaleRoyaltyBPS_,
            royaltyPayoutAddress_,
            contractURI_,
            tokenBaseURI_,
            isBurnable_,
            isBridgeable_
        )
    {
        minter = minter_;
    }

    /// @notice tokenURI A distinct Uniform Resource Identifier (URI) for a given asset.
    /// @param tokenId - the id of token
    function tokenURI(
        uint256 tokenId
    ) public view virtual override returns (string memory) {
        require(
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );

        ArtworkEditionIndex memory aei = allArtworkEditionsIndex[tokenId];

        require(
            bytes(artworkCIDs[aei.artworkID]).length != 0,
            "Artwork CID does not exist"
        );

        string memory baseURI = _tokenBaseURI;
        if (bytes(baseURI).length == 0) {
            baseURI = "ipfs://";
        }

        return
            string(
                abi.encodePacked(
                    baseURI,
                    artworkCIDs[aei.artworkID],
                    "/",
                    tokenId.toString()
                )
            );
    }

    /// @notice getArtworkIDByTokenID use for getting artwork ID from token ID
    /// @param _tokenId - the id of edition
    function getArtworkIDByTokenID(
        uint256 _tokenId
    ) public view returns (uint256) {
        return allArtworkEditionsIndex[_tokenId].artworkID;
    }

    /// @notice updateArtworkCIDs use for update CIDs for artworks
    /// @param _artworkIDs - array of artwork IDs
    /// @param _artworkCIDs - array of ipfs CIDs
    function updateArtworkCIDs(
        uint256[] memory _artworkIDs,
        string[] memory _artworkCIDs
    ) public onlyAuthorized {
        require(_artworkIDs.length == _artworkCIDs.length, "Invalid input");
        for (uint256 i = 0; i < _artworkIDs.length; i++) {
            require(
                bytes(artworks[_artworkIDs[i]].title).length != 0,
                "Artwork does not exist"
            );
            artworkCIDs[_artworkIDs[i]] = _artworkCIDs[i];
        }
    }

    /// @notice batchMint use for override parent method and only mintArtworkEdition can create new edition
    function batchMint(
        MintArtworkParam[] memory /* mintParams_ */
    ) external pure override {
        revert("Not support");
    }

    /// @notice mintArtworkEdition use for mint new token to owner
    /// @param _artworkID - the id of artwork
    /// @param _owner - the owner will be received minted token
    function mintArtworkEdition(
        uint256 _artworkID,
        address _owner
    ) public onlyMinter {
        require(
            bytes(artworks[_artworkID].title).length != 0,
            "Artwork does not exist"
        );

        require(
            bytes(artworkCIDs[_artworkID]).length != 0,
            "Artwork CID does not exist"
        );

        uint256 editionNumber = tokenCounter.current();

        _mintArtwork(
            _artworkID,
            editionNumber,
            _owner,
            _owner,
            (_artworkID + editionNumber).toString()
        );

        tokenCounter.increment();
    }
}
