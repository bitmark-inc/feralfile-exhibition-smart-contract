// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "./Authorizable.sol";

contract FeralFileAirdropV1 is ERC1155, Authorizable {
    enum Type {
        Fungible,
        NonFungible
    }

    // Version
    string public constant codeVersion = "FeralFileAirdropV1";

    // Airdropped addresses
    mapping(address => bool) public airdroppedAddresses;

    // Token type
    Type public tokenType;

    // Burnable flag
    bool public burnable;

    // Bridgeable flag
    bool public bridgeable;

    // Contract URI
    string public contractURI;

    // Ended flag
    bool private _ended;

    constructor(
        Type tokenType_,
        string memory tokenURI_,
        string memory contractURI_,
        bool burnable_,
        bool bridgeable_
    ) ERC1155(tokenURI_) {
        contractURI = contractURI_;
        tokenType = tokenType_;
        burnable = burnable_;
        bridgeable = bridgeable_;
    }

    // @notice check if the airdrop is ended
    function isEnded() external view returns (bool) {
        return _ended;
    }

    // @notice end the airdrop
    function end() external onlyOwner {
        require(!_ended, "FeralFileAirdropV1: already ended");
        _ended = true;
    }

    /// @notice Mint tokens to the contract
    /// @param tokenID_ token ID
    /// @param amount_ amount of tokens to mint
    function mint(uint256 tokenID_, uint256 amount_) external onlyAuthorized {
        require(!_ended, "FeralFileAirdropV1: already ended");
        _checkProperTokenAmount(amount_);
        _mint(address(this), tokenID_, amount_, "");
    }

    /// @notice airdrop tokens to a list of addresses
    /// @param tokenID_ token ID
    /// @param to_ list of addresses to airdrop to
    function airdrop(
        uint256 tokenID_,
        address[] calldata to_
    ) external onlyAuthorized {
        require(!_ended, "FeralFileAirdropV1: already ended");
        _checkProperTokenAmount(to_.length); // to_ length is the amount of tokens to airdrop

        for (uint256 i = 0; i < to_.length; i++) {
            require(
                !airdroppedAddresses[to_[i]],
                "FeralFileAirdropV1: already airdropped"
            );

            _safeTransferFrom(address(this), to_[i], tokenID_, 1, "");

            airdroppedAddresses[to_[i]] = true;
        }
    }

    /// @notice burn remaining tokens
    /// @param tokenID_ token ID
    function burnRemaining(uint256 tokenID_) external onlyOwner {
        _burn(address(this), tokenID_, balanceOf(address(this), tokenID_));
    }

    /// @notice burn tokens
    /// @param tokenID_ token ID
    /// @param amount_ amount of tokens to burn
    function burn(uint256 tokenID_, uint256 amount_) external {
        require(burnable, "FeralFileAirdropV1: not burnable");
        _checkProperTokenAmount(amount_);
        _burn(_msgSender(), tokenID_, amount_);
    }

    /// @notice check proper amount of tokens
    function _checkProperTokenAmount(uint256 amount_) internal view {
        require(
            (tokenType == Type.Fungible && amount_ > 0) ||
                (tokenType == Type.NonFungible && amount_ == 1),
            "FeralFileAirdropV1: amount mismatch"
        );
    }

    function onERC1155Received(
        address,
        address from_,
        uint256,
        uint256,
        bytes memory
    ) public pure returns (bytes4) {
        require(
            from_ == address(0),
            "FeralFileAirdropV1: not allowed to send token back"
        );
        return this.onERC1155Received.selector;
    }
}
