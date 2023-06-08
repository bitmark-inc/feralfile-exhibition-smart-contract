// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

import "./FeralfileSaleStruct.sol";
import "./ECDSASign.sol";

contract FeralfileVault is Ownable, FeralfileSaleStruct, ECDSASign {
    mapping(bytes32 => bool) private _paidSale;

    constructor(address signer_) ECDSASign(signer_) {}

    /// @notice pay for buyArtwork to a FFV4 contract destination.
    /// @param from - a FFV4 exhibition contract address
    /// @param r_ - part of signature for validating parameters integrity
    /// @param s_ - part of signature for validating parameters integrity
    /// @param v_ - part of signature for validating parameters integrity
    /// @param saleData_ - the sale data
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
        require(!_paidSale[requestHash], "FeralfileVault: paid sale");
        require(
            verifySignature(requestHash, r_, s_, v_),
            "FeralfileVault: invalid signature"
        );
        require(
            address(this).balance >= saleData_.price,
            "FeralfileVault: insufficient balance"
        );
        _paidSale[requestHash] = true;
        payable(from).transfer(saleData_.price);
    }

    function withdrawFunds() external onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }

    receive() external payable {}
}
