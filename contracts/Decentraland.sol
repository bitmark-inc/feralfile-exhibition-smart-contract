// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Decentraland interface
contract Decentraland {
    // function decodeTokenId(uint256 value)
    //     external
    //     pure
    //     returns (int256, int256)
    // {}

    function ownerOf(uint256 tokenId) public view returns (address) {}
}

contract MockDecentraland {
    function ownerOf(uint256 tokenId) public view returns (address) {
        return 0x0177Ba2aD05cb12D6Dd637963e040092676EE8DE;
    }
}
