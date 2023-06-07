// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract FeralfileSaleStruct {
    struct RevenueShare {
        address recipient;
        uint256 bps;
    }

    struct SaleData {
        uint256 price; // in wei
        uint256 cost; // in wei
        uint256 expiryTime;
        address destination;
        uint256[] tokenIds;
        RevenueShare[][] revenueShares; // address and royalty bps (500 means 5%)
        bool payByVaultContract; // get eth from vault contract, used by credit card pay that proxy by ITX
    }

    function isValidSaleData(SaleData calldata saleData_) internal view {
        require(saleData_.tokenIds.length > 0, "SaleData: tokenIds is empty");
        require(
            saleData_.tokenIds.length == saleData_.revenueShares.length,
            "SaleData: tokenIds and revenueShares length mismatch"
        );
        require(
            saleData_.expiryTime > block.timestamp,
            "SaleData: sale is expired"
        );
    }
}
