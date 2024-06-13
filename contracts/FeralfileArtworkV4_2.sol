// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {Nonces} from "./Nonces.sol";
import {FeralfileExhibitionV4_1} from "./FeralfileArtworkV4_1.sol";
import {IFeralfileVaultV2} from "./IFeralfileVaultV2.sol";
import {FeralfileSaleDataV2} from "./FeralfileSaleDataV2.sol";

contract FeralfileExhibitionV4_2 is
    FeralfileExhibitionV4_1,
    FeralfileSaleDataV2,
    Nonces
{
    error TokenIDNotFound();
    error FunctionNotSupported();
    error SaleNotStarted();
    error InvalidPaymentAmount();
    error TotalBpsOver();
    error InvalidAddress();

    // vault contract instance
    IFeralfileVaultV2 public vaultV2;

    mapping(uint256 => uint256) private seriesNextPurchasableTokenIds; // seriesID -> tokenID

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
        uint256[] memory seriesNextPurchasableTokenIds_
    )
        FeralfileExhibitionV4_1(
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
        vaultV2 = IFeralfileVaultV2(payable(vault_));
        for (uint256 i = 0; i < seriesIds_.length; i++) {
            seriesNextPurchasableTokenIds[seriesIds_[i]] = seriesNextPurchasableTokenIds_[i];
        }
    }

    /// @notice Set vaultV2 contract
    /// @dev don't allow to set vaultV2 as zero address
    function setVaultV2(address vault_) external onlyOwner {
        if (vault_ == address(0)) {
            revert InvalidAddress();
        }

        vaultV2 = IFeralfileVaultV2(payable(vault_));
    }

    /// @notice override revert setVault
    function setVault(address) external pure override {
        revert FunctionNotSupported();
    }

    /// @notice pay to get artworks to a destination address. The pricing, costs and other details is included in the saleData
    /// @param r_ - part of signature for validating parameters integrity
    /// @param s_ - part of signature for validating parameters integrity
    /// @param v_ - part of signature for validating parameters integrity
    /// @param saleData_ - the sale data
    function buyBulkArtworks(
        bytes32 r_,
        bytes32 s_,
        uint8 v_,
        SaleDataV2 calldata saleData_
    ) external payable {
        if (!_selling) {
            revert SaleNotStarted();
        }

        super._checkContractOwnedToken();
        validateSaleDataV2(saleData_);

        if (saleData_.payByVaultContract) {
            vaultV2.payForSaleV2(r_, s_, v_, saleData_);
        } else {
            if (saleData_.price != msg.value) {
                revert InvalidPaymentAmount();
            }
        }

        bytes32 message = keccak256(
            abi.encode(block.chainid, address(this), saleData_)
        );

        if (!isValidSignature(message, r_, s_, v_)) {
            revert InvalidSignature();
        }

        //check nonce
        _useCheckedNonce(saleData_.destination, saleData_.nonce);

        uint256 itemRevenue;
        if (saleData_.price > saleData_.cost) {
            itemRevenue =
                (saleData_.price - saleData_.cost) /
                saleData_.quantity;
        }

        uint256 nextPurchasableTokenId = seriesNextPurchasableTokenIds[saleData_.seriesID];
        uint256 i = 0;
        while (i < saleData_.quantity) {
            uint256 tokenIdForSale = nextPurchasableTokenId;

            if (!_exists(tokenIdForSale)) {
                revert TokenIDNotFound();
            }

            nextPurchasableTokenId++;
            if (ownerOf(tokenIdForSale) != address(this)) {
                continue;
            }

            // send NFT
            _safeTransfer(
                address(this),
                saleData_.destination,
                tokenIdForSale,
                ""
            );

            emit BuyArtworkV2(
                saleData_.destination,
                tokenIdForSale,
                saleData_.nonce
            );
            i++;
        }

        // save next sale token id for seriesID
        seriesNextPurchasableTokenIds[saleData_.seriesID] = nextPurchasableTokenId;

        // distribute royalty
        uint256 distributedRevenue;
        uint256 platformRevenue;

        RevenueShare[] memory revenueShares = saleData_.revenueShares;
        uint256 remainingRev = itemRevenue;

        // deduct advances payment from revenue
        for (uint256 j = 0; j < revenueShares.length && remainingRev > 0; j++) {
            uint256 remainingAdvanceAmount = advances[
                revenueShares[j].recipient
            ];
            uint256 rev = remainingAdvanceAmount >= remainingRev
                ? remainingRev
                : remainingAdvanceAmount;
            uint256 totalRev = saleData_.quantity * rev;
            platformRevenue += totalRev;
            advances[revenueShares[j].recipient] -= totalRev;
            remainingRev -= totalRev;
        }

        // distribute revenue
        if (remainingRev > 0) {
            for (uint256 j = 0; j < revenueShares.length; j++) {
                address recipient = revenueShares[j].recipient;
                uint256 rev = (remainingRev * revenueShares[j].bps) / 10000;
                uint256 totalRev = saleData_.quantity * rev;
                if (recipient == costReceiver) {
                    platformRevenue += totalRev;
                    continue;
                }
                distributedRevenue += totalRev;
                payable(recipient).transfer(totalRev);
            }
        }

        if (
            saleData_.price - saleData_.cost <
            distributedRevenue + platformRevenue
        ) {
            revert TotalBpsOver();
        }

        // Transfer cost, platform revenue and remaining funds
        uint256 leftOver = saleData_.price - distributedRevenue;
        if (leftOver > 0) {
            payable(costReceiver).transfer(leftOver);
        }
    }

    /// @notice override revert buyArtworks
    function buyArtworks(
        bytes32,
        bytes32,
        uint8,
        SaleData calldata
    ) external payable override {
        revert FunctionNotSupported();
    }

    /// @notice Event emitted when Artwork has been sold with the additional nonce
    event BuyArtworkV2(
        address indexed buyer,
        uint256 indexed tokenId,
        uint256 nonce
    );
}
