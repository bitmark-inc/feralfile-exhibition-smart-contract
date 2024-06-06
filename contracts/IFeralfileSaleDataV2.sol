// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {IFeralfileSaleData} from "./IFeralfileSaleData.sol";

interface IFeralfileSaleDataV2 is IFeralfileSaleData {
    struct SaleDataV2 {
        uint256 price; // in wei
        uint256 cost; // in wei
        uint256 expiryTime;
        address destination;
        uint256 nonce;
        uint256 seriesID;
        uint16 quantity;
        RevenueShare[] revenueShares; // address and royalty bps (500 means 5%)
        bool payByVaultContract; // get eth from vault contract, used by credit card pay that proxy by ITX
    }
}
