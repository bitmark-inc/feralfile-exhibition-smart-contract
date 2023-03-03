// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./FeralfileArtworkV3_3.sol";

contract MoMABurn is Ownable {
    using SafeMath for uint256;
    address constant DEAD_ADDRESS = 0x000000000000000000000000000000000000dEaD;

    ERC721 public burnToken; // Address of token need to burn
    FeralfileExhibitionV3_3 public mintToken; // Address of new token will be mint after burned
    uint256 private tier1ArtworkID;
    uint256 private tier2ArtworkID;
    uint256 private tier3ArtworkID;
    uint256 private tier1Counter;
    uint256 private tier2Counter;
    uint256 private tier3Counter;
    string private tier1IpfsCID;
    string private tier2IpfsCID;
    string private tier3IpfsCID;

    constructor(
        address _burnToken,
        address _mintToken,
        uint256 _tier1ArtworkID,
        uint256 _tier2ArtworkID,
        uint256 _tier3ArtworkID,
        string memory _tier1IpfsCID,
        string memory _tier2IpfsCID,
        string memory _tier3IpfsCID
    ) {
        burnToken = ERC721(_burnToken);
        mintToken = FeralfileExhibitionV3_3(_mintToken);
        tier1ArtworkID = _tier1ArtworkID;
        tier2ArtworkID = _tier2ArtworkID;
        tier3ArtworkID = _tier3ArtworkID;
        tier1Counter = 1;
        tier2Counter = 1;
        tier3Counter = 1;
        tier1IpfsCID = _tier1IpfsCID;
        tier2IpfsCID = _tier2IpfsCID;
        tier3IpfsCID = _tier3IpfsCID;
    }

    // Events
    event NewBurnAddress(address indexed token);
    event NewMintAddress(address indexed token);

    function updateBurnToken(address _token) public onlyOwner {
        require(_token != address(0), "Invalid contract address");

        burnToken = ERC721(_token);

        emit NewBurnAddress(_token);
    }

    function updateMintToken(address _token) public onlyOwner {
        require(_token != address(0), "Invalid contract address");

        mintToken = FeralfileExhibitionV3_3(_token);

        emit NewMintAddress(_token);
    }

    function updateMintArtworks(
        uint256 _tier1ArtworkID,
        uint256 _tier2ArtworkID,
        uint256 _tier3ArtworkID,
        string memory _tier1IpfsCID,
        string memory _tier2IpfsCID,
        string memory _tier3IpfsCID
    ) public onlyOwner {
        tier1ArtworkID = _tier1ArtworkID;
        tier2ArtworkID = _tier2ArtworkID;
        tier3ArtworkID = _tier3ArtworkID;
        tier1IpfsCID = _tier1IpfsCID;
        tier2IpfsCID = _tier2IpfsCID;
        tier3IpfsCID = _tier3IpfsCID;
    }

    function burnAndMint(uint256[] memory _burnedEditions) public {
        require(
            _burnedEditions.length == 3 ||
                _burnedEditions.length == 6 ||
                _burnedEditions.length == 9,
            "Invalid number of burning editions"
        );

        bool isApprovalForAll = burnToken.isApprovedForAll(
            _msgSender(),
            address(this)
        );

        for (uint256 i = 0; i < _burnedEditions.length; i++) {
            require(
                burnToken.ownerOf(_burnedEditions[i]) == _msgSender(),
                "Invalid owner of burned token"
            );

            if (!isApprovalForAll) {
                require(
                    burnToken.getApproved(_burnedEditions[i]) == address(this),
                    "Token need to approve for this contract"
                );
            }

            burnToken.safeTransferFrom(
                _msgSender(),
                DEAD_ADDRESS,
                _burnedEditions[i]
            );
        }

        uint256 artworkID = tier1ArtworkID;
        uint256 editionNumber = tier1Counter;
        string memory ipfsCID = tier1IpfsCID;

        if (_burnedEditions.length == 3) {
            tier1Counter++;
        }

        if (_burnedEditions.length == 6) {
            artworkID = tier2ArtworkID;
            editionNumber = tier2Counter;
            ipfsCID = tier2IpfsCID;
            tier2Counter++;
        }

        if (_burnedEditions.length == 9) {
            artworkID = tier3ArtworkID;
            editionNumber = tier3Counter;
            ipfsCID = tier3IpfsCID;
            tier3Counter++;
        }

        FeralfileExhibitionV3_3.MintArtworkParam[]
            memory _mintParams = new FeralfileExhibitionV3_3.MintArtworkParam[](
                1
            );

        _mintParams[0] = FeralfileExhibitionV3_3.MintArtworkParam(
            artworkID,
            editionNumber,
            _msgSender(),
            _msgSender(),
            ipfsCID
        );

        mintToken.batchMint(_mintParams);
    }
}
