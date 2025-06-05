// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {SSTORE2} from "@0xsequence/sstore2/contracts/SSTORE2.sol";

library LibBytes {
    //----------------------------------------------------------
    // Errors
    //----------------------------------------------------------
    error ErrEmptySSTORE2Pointer();
    error ErrOutOfBounds();

    /// @notice Copy bytes from src to dest
    /// @param dest The destination bytes
    /// @param destOffset The offset to copy to
    /// @param src The source bytes
    /// @param len The length to copy
    function memcpy(
        bytes memory dest,
        uint256 destOffset,
        bytes memory src,
        uint256 len
    ) public view {
        if (len == 0) return;
        if (destOffset + len > dest.length) {
            revert ErrOutOfBounds();
        }
        assembly {
            if iszero(
                staticcall(
                    gas(),
                    4,
                    add(src, 32),
                    len,
                    add(add(dest, 32), destOffset),
                    len
                )
            ) {
                revert(0, 0)
            }
        }
    }

    /// @notice Join SSTORE2 chunks into a single bytes array
    /// @param pointers The SSTORE2 pointers to join
    /// @return The joined bytes array
    function sstore2Join(
        address[] memory pointers
    ) public view returns (bytes memory) {
        if (pointers.length == 0) {
            return new bytes(0);
        }

        uint256 total;
        for (uint256 i = 0; i < pointers.length; ) {
            uint256 length = pointers[i].code.length;
            if (length <= 1) {
                revert ErrEmptySSTORE2Pointer();
            }
            total += length - 1; // skip STOP byte
            unchecked {
                ++i;
            }
        }

        bytes memory result = new bytes(total);
        uint256 offset;
        for (uint256 i = 0; i < pointers.length; ) {
            bytes memory chunk = SSTORE2.read(pointers[i]);
            uint256 length = chunk.length;
            memcpy(result, offset, chunk, length);
            offset += length;
            unchecked {
                ++i;
            }
        }

        return result;
    }
}
