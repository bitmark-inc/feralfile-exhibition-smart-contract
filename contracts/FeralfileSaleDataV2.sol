// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./IFeralfileSaleDataV2.sol";

contract FeralfileSaleDataV2 is IFeralfileSaleDataV2 {
    function validateSaleDataV2(SaleDataV2 calldata saleData_) internal view {
        require(
            saleData_.expiryTime > block.timestamp,
            "FeralfileSaleData: sale is expired"
        );
    }
}
