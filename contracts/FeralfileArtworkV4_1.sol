// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./FeralfileArtworkV4.sol";

contract FeralfileExhibitionV4_1 is FeralfileExhibitionV4 {
    struct AdvancedData {
        address artist;
        uint256 amount;
    }

    mapping(address => uint256) private artistAdvancedAmounts;

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
        uint256[] memory seriesMaxSupplies_,
        AdvancedData[] memory data_
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
    {
        for (uint256 i = 0; i < data_.length; i++) {
            artistAdvancedAmounts[data_[i].artist] = data_[i].amount;
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

        require(
            isValidSignature(message, r_, s_, v_),
            "FeralfileExhibitionV4: invalid signature"
        );

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
            if (itemRevenue > 0) {
                // distribute royalty
                RevenueShare[] memory revenueShares = saleData_.revenueShares[
                    i
                ];
                for (uint256 j = 0; j < revenueShares.length; j++) {
                    address recipient = revenueShares[j].recipient;
                    uint256 rev = (itemRevenue * revenueShares[j].bps) / 10000;
                    if (recipient == costReceiver) {
                        platformRevenue += rev;
                        continue;
                    }
                    // check if artist has advanced amount
                    if (artistAdvancedAmounts[recipient] > 0) {
                        if (artistAdvancedAmounts[recipient] >= rev) {
                            artistAdvancedAmounts[recipient] -= rev;
                            platformRevenue += rev;
                            rev = 0;
                        } else {
                            rev -= artistAdvancedAmounts[recipient];
                            artistAdvancedAmounts[recipient] = 0;
                        }
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
