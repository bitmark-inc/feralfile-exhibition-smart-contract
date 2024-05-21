// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { Authorizable } from "./Authorizable.sol";

contract FeralfileToken is ERC20, Authorizable {

    // Constructor to initialize the token with a name and symbol
    constructor(string memory name_, string memory symbol_) ERC20(name_, symbol_) {}

    /// @notice Mint new tokens to an owner
    /// @param amount The amount of tokens to mint
    /// @param owner The address of the owner to receive the minted tokens
    function mintToken(uint256 amount, address owner) external onlyAuthorized {
        _mint(owner, amount);
    }
}