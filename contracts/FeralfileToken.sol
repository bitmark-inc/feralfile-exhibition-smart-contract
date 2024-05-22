// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { Authorizable } from "./Authorizable.sol";

contract FeralfileToken is ERC20, Authorizable {

    // Constructor to initialize the token with a name and symbol
    constructor(string memory name_, string memory symbol_) ERC20(name_, symbol_) {}

    /// @notice Mint new tokens to an owner
    /// @param owner The address of the owner to receive the minted tokens
    /// @param amount The amount of tokens to mint
    function mint(address owner, uint256 amount) external onlyAuthorized {
        _mint(owner, amount);
    }

    /// @notice Mint tokens for multiple owners
    /// @param owners An array of addresses to receive the minted tokens
    /// @param amounts An array of amounts of tokens to mint for each respective owner
    function batchMint(address[] calldata owners, uint256[] calldata amounts) external virtual onlyAuthorized {
        require(owners.length == amounts.length, "FeralfileToken: owners and amounts length mismatch");
        for (uint256 i = 0; i < owners.length; i++) {
            _mint(owners[i], amounts[i]);
        }
    }
}