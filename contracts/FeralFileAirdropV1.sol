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
    string public constant version = "v1";

    // Airdropped addresses
    mapping(address => bool) public airdroppedAddresses;

    // already ended flag
    bool private _ended;

    // Token type
    Type public tokenType;

    constructor(Type tokenType_, string memory tokenURI_) ERC1155(tokenURI_) {
        tokenType = tokenType_;
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
    /// @param tokenID token ID
    /// @param amount_ amount of tokens to mint
    function mint(uint256 tokenID, uint256 amount_) external onlyAuthorized {
        require(!_ended, "FeralFileAirdropV1: already ended");
        require(
            (tokenType == Type.Fungible && amount_ > 0) ||
                (tokenType == Type.NonFungible && amount_ == 1),
            "FeralFileAirdropV1: amount mismatch"
        );
        _mint(address(this), tokenID, amount_, "");
    }

    /// @notice airdrop tokens to a list of addresses
    /// @param tokenID token ID
    /// @param to_ list of addresses to airdrop to
    function airdrop(
        uint256 tokenID,
        address[] calldata to_
    ) external onlyAuthorized {
        require(!_ended, "FeralFileAirdropV1: already ended");
        require(
            (tokenType == Type.Fungible && to_.length > 0) ||
                (tokenType == Type.NonFungible && to_.length == 1),
            "FeralFileAirdropV1: amount mismatch"
        );

        uint256[] memory _tokenIDs = new uint256[](1);
        _tokenIDs[0] = tokenID;

        for (uint256 i = 0; i < to_.length; i++) {
            require(
                !airdroppedAddresses[to_[i]],
                "FeralFileAirdropV1: already airdropped"
            );

            uint256[] memory _amounts = new uint256[](1);
            _amounts[0] = 1; // 1 token for each address

            _safeBatchTransferFrom(
                address(this),
                to_[i],
                _tokenIDs,
                _amounts,
                ""
            );

            airdroppedAddresses[to_[i]] = true;
        }
    }

    /// @notice burn remaining tokens
    /// @param tokenID token ID
    function burnRemaining(uint256 tokenID) external onlyOwner {
        _burn(address(this), tokenID, balanceOf(address(this), tokenID));
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
