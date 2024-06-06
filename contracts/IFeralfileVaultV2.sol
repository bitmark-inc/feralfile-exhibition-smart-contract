// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

import "./IFeralfileSaleDataV2.sol";
import "./ECDSASigner.sol";

interface IFeralfileVaultV2 is IFeralfileSaleDataV2 {
    function payForSaleV2(
        bytes32 r_,
        bytes32 s_,
        uint8 v_,
        SaleDataV2 calldata saleData_
    ) external;

    function withdrawFund(uint256 weiAmount) external;

    receive() external payable;
}
