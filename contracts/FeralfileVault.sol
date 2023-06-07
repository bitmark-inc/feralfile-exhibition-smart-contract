// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

import "./FeralfileSaleStruct.sol";
import "./ECDSASign.sol";

contract FeralfileVault is Ownable, FeralfileSaleStruct, ECDSASign {
    mapping(bytes32 => bool) private paidSale;

    constructor(address signer_) ECDSASign(signer_) {}

    function payForSale(
        address from,
        bytes32 r_,
        bytes32 s_,
        uint8 v_,
        SaleData calldata saleData_
    ) external {
        require(
            saleData_.payByVaultContract,
            "FeralfileVault: not pay by vault"
        );
        isValidSaleData(saleData_);
        bytes32 requestHash = keccak256(
            abi.encode(block.chainid, from, saleData_)
        );
        require(!paidSale[requestHash], "FeralfileVault: paid sale");
        require(
            isValidSignature(requestHash, r_, s_, v_),
            "FeralfileVault: invalid signature"
        );
        require(
            address(this).balance >= saleData_.price,
            "FeralfileVault: insufficient balance"
        );
        paidSale[requestHash] = true;
        payable(from).transfer(saleData_.price);
    }

    function withdrawFunds() external onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }

    // @notice able to recieve funds
    receive() external payable {}
}
