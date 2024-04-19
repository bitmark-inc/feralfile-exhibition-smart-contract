// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./FeralfileArtworkV4.sol";

contract FeralfileExhibitionV4_1 is FeralfileExhibitionV4 {
    mapping(address => uint256) public advances;

    error InvalidAdvanceAddressesAndAmounts();
    error InvalidAdvanceAddress();
    error InvalidAdvanceAmount();
    error InvalidSignature();

    constructor(
        string memory name_,
        string memory symbol_,
        bool burnable_,
        bool bridgeable_,
        address signer_,
        address vault_,
        address costReceiver_,
        string memory contractURI_,
        uint256[] memory seriesIds_,
        uint256[] memory seriesMaxSupplies_
    )
        FeralfileExhibitionV4(
            name_,
            symbol_,
            burnable_,
            bridgeable_,
            signer_,
            vault_,
            costReceiver_,
            contractURI_,
            seriesIds_,
            seriesMaxSupplies_
        )
    {}

    /// @notice set advances setting
    /// @param addresses_ - the addresses to set advances
    /// @param amounts_ - the amounts to set advances
    function setAdvanceSetting(
        address[] calldata addresses_,
        uint256[] calldata amounts_
    ) public onlyAuthorized {
        if (addresses_.length != amounts_.length) {
            revert InvalidAdvanceAddressesAndAmounts();
        }
        for (uint256 i = 0; i < addresses_.length; i++) {
            if (addresses_[i] == address(0)) {
                revert InvalidAdvanceAddress();
            }
            if (amounts_[i] == 0 && advances[addresses_[i]] == 0) {
                revert InvalidAdvanceAmount();
            }
            if (amounts_[i] > 0) {
                advances[addresses_[i]] = amounts_[i];
            }
        }
    }

    /// @notice pay to get artworks to a destination address. The pricing, costs and other details is included in the saleData
    /// @param r_ - part of signature for validating parameters integrity
    /// @param s_ - part of signature for validating parameters integrity
    /// @param v_ - part of signature for validating parameters integrity
    /// @param saleData_ - the sale data
    function buyArtworks(
        bytes32 r_,
        bytes32 s_,
        uint8 v_,
        SaleData calldata saleData_
    ) external payable override {
        require(_selling, "FeralfileExhibitionV4: sale is not started");
        super._checkContractOwnedToken();
        validateSaleData(saleData_);

        saleData_.payByVaultContract
            ? vault.payForSale(r_, s_, v_, saleData_)
            : require(
                saleData_.price == msg.value,
                "FeralfileExhibitionV4: invalid payment amount"
            );

        bytes32 message = keccak256(
            abi.encode(block.chainid, address(this), saleData_)
        );

        if (!isValidSignature(message, r_, s_, v_)) {
            revert InvalidSignature();
        }

        uint256 itemRevenue;
        if (saleData_.price > saleData_.cost) {
            itemRevenue =
                (saleData_.price - saleData_.cost) /
                saleData_.tokenIds.length;
        }

        uint256 distributedRevenue;
        uint256 platformRevenue;
        for (uint256 i = 0; i < saleData_.tokenIds.length; i++) {
            // send NFT
            _safeTransfer(
                address(this),
                saleData_.destination,
                saleData_.tokenIds[i],
                ""
            );
            // distribute royalty
            RevenueShare[] memory revenueShares = saleData_.revenueShares[i];
            uint256 remainingRev = itemRevenue;

            // deduct advances payment from revenue
            for (uint256 j = 0; j < revenueShares.length; j++) {
                if (
                    advances[revenueShares[j].recipient] > 0 && remainingRev > 0
                ) {
                    if (advances[revenueShares[j].recipient] >= remainingRev) {
                        platformRevenue += remainingRev;
                        advances[revenueShares[j].recipient] -= remainingRev;
                        remainingRev = 0;
                    } else {
                        platformRevenue += advances[revenueShares[j].recipient];
                        remainingRev -= advances[revenueShares[j].recipient];
                        advances[revenueShares[j].recipient] = 0;
                    }
                }
            }

            // distribute revenue
            if (remainingRev > 0) {
                for (uint256 j = 0; j < revenueShares.length; j++) {
                    address recipient = revenueShares[j].recipient;
                    uint256 rev = (remainingRev * revenueShares[j].bps) / 10000;
                    if (recipient == costReceiver) {
                        platformRevenue += rev;
                        continue;
                    }
                    distributedRevenue += rev;
                    payable(recipient).transfer(rev);
                }
            }

            emit BuyArtwork(saleData_.destination, saleData_.tokenIds[i]);
        }

        require(
            saleData_.price - saleData_.cost >=
                distributedRevenue + platformRevenue,
            "FeralfileExhibitionV4: total bps over 10,000"
        );

        // Transfer cost, platform revenue and remaining funds
        uint256 leftOver = saleData_.price - distributedRevenue;
        if (leftOver > 0) {
            payable(costReceiver).transfer(leftOver);
        }
    }
}
