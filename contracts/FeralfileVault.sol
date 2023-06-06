// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

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

    function setExhibitionContract(address ec) external onlyOwner {
        exhibitionContract[ec] = true;
    }

    function unsetExhibitionContract(address ec) external onlyOwner {
        exhibitionContract[ec] = false;
    }

    function pay(uint256 amount) external onlyExhibitionContract {
        require(address(this).balance >= amount, "insufficient balance");
        payable(msg.sender).transfer(amount);
    }

    function withdrawFunds() external onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }

    // @notice able to recieve funds
    receive() external payable {}
}
