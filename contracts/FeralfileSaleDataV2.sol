// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface IFeralfileSaleDataV2 {
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
        uint256 biddingUnix; // unix timestamp for bidding
    }
}

contract FeralfileSaleDataV2 is IFeralfileSaleDataV2 {
    function validateSaleData(SaleData calldata saleData_) internal view {
        require(
            saleData_.tokenIds.length > 0,
            "FeralfileSaleData: tokenIds is empty"
        );
        require(
            saleData_.tokenIds.length == saleData_.revenueShares.length,
            "FeralfileSaleData: tokenIds and revenueShares length mismatch"
        );
        require(
            saleData_.expiryTime > block.timestamp,
            "FeralfileSaleData: sale is expired"
        );
    }
}