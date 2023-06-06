// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract FeralfileVault is Ownable {
    mapping(address => bool) public exhibitionContract;

    modifier onlyExhibitionContract() {
        require(
            exhibitionContract[msg.sender],
            "Only exhibition contract can call this function."
        );
        _;
    }

    function setExhibitionContract(address ec_) external onlyOwner {
        exhibitionContract[ec_] = true;
    }

    function unsetExhibitionContract(address ec_) external onlyOwner {
        exhibitionContract[ec_] = false;
    }

    function pay(uint256 amount_) external onlyExhibitionContract {
        require(address(this).balance >= amount_, "insufficient balance");
        payable(msg.sender).transfer(amount_);
    }

    fallback() external payable {}

    receive() external payable {}
}
