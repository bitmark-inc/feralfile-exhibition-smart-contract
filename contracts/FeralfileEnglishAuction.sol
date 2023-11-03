// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./IFeralfileSaleData.sol";
import "./ECDSASigner.sol";
import "./IFeralfileVault.sol";

contract IFeralfileExhibitionV4 is IFeralfileSaleData {
    function buyArtworks(
        bytes32 r_,
        bytes32 s_,
        uint8 v_,
        SaleData calldata saleData_
    ) external payable {}
}

contract FeralfileEnglishAuction is Ownable, IFeralfileSaleData, ECDSASigner {
    struct Auction {
        uint256 id;
        uint256 startAt;
        uint256 endAt;
        uint256 extendDuration;
        uint256 extendThreshold;
        uint256 minIncreaseFactor;
        uint256 minIncreaseAmount;
        uint256 minPrice;
        bool isSettled;
    }

    struct Bid {
        address bidder;
        uint256 amount;
        bool fromFeralFile;
    }

    struct LatestAuctionStatus {
        Bid latestBid;
        uint256 endAt;
        bool isSettled;
    }

    mapping(uint256 => Auction) public auctions;
    mapping(uint256 => Bid) public highestBids;

    constructor(address signer_) ECDSASigner(signer_) {}

    function ongoingAuction(uint256 id_) external view returns (bool) {
        require(
            auctions[id_].id > 0,
            "FeralfileEnglishAuction: auction not found"
        );

        return
            auctions[id_].startAt <= block.timestamp &&
            auctions[id_].endAt >= block.timestamp;
    }

    function registerAuctions(Auction[] calldata auctions_) external onlyOwner {
        for (uint256 i = 0; i < auctions_.length; i++) {
            Auction memory auction_ = auctions[auctions_[i].id];

            require(
                auction_.id == 0,
                "FeralfileEnglishAuction: auction already exist"
            );

            require(
                auctions_[i].id > 0,
                "FeralfileEnglishAuction: invalid auction id"
            );

            require(
                auctions_[i].startAt >= block.timestamp &&
                    auctions_[i].endAt > block.timestamp,
                "FeralfileEnglishAuction: auction start time and end time should be in the future"
            );

            require(
                auctions_[i].endAt > auctions_[i].startAt,
                "FeralfileEnglishAuction: auction end time should be after start time"
            );

            require(
                auctions_[i].minIncreaseFactor > 0,
                "FeralfileEnglishAuction: auction min increase factor should be greater than 0"
            );

            require(
                auctions_[i].minIncreaseAmount > 0,
                "FeralfileEnglishAuction: auction min increase amount should be greater than 0"
            );

            require(
                auctions_[i].minPrice > 0,
                "FeralfileEnglishAuction: auction min price should be greater than 0"
            );

            auctions[auctions_[i].id] = auctions_[i];
        }
    }

    function listHighestBids(
        uint256[] memory aucIds_
    ) external view returns (LatestAuctionStatus[] memory) {
        LatestAuctionStatus[] memory results = new LatestAuctionStatus[](
            aucIds_.length
        );
        for (uint i = 0; i < aucIds_.length; i++) {
            Auction memory auction_ = auctions[aucIds_[i]];
            results[i].latestBid = highestBids[aucIds_[i]];
            results[i].endAt = auction_.endAt;
            results[i].isSettled = auction_.isSettled;
        }
        return results;
    }

    function isValidNewBid(
        uint256 auctionID_,
        uint256 amount_
    ) external view returns (bool) {
        Auction memory auction_ = auctions[auctionID_];
        Bid memory highestBid_ = highestBids[auctionID_];
        _validateBidRequest(auction_, highestBid_, amount_, block.timestamp);
        return true;
    }

    function _validateBidRequest(
        Auction memory auction_,
        Bid memory highestBid_,
        uint256 amount_,
        uint256 timestamp_
    ) internal pure {
        // make sure the auction exist
        require(auction_.id > 0, "FeralfileEnglishAuction: auction not found");
        // check if auction is on going
        require(
            auction_.startAt <= timestamp_ && timestamp_ < auction_.endAt,
            "FeralfileEnglishAuction: auction not started or ended"
        );
        // check if auction is not settled
        require(
            !auction_.isSettled,
            "FeralfileEnglishAuction: auction already settled"
        );

        if (highestBid_.amount > 0) {
            // check bidding amount follow the minimum increment
            uint256 minIncreaseAmount_ = (highestBid_.amount *
                auction_.minIncreaseFactor) / 100;
            if (minIncreaseAmount_ < auction_.minIncreaseAmount) {
                minIncreaseAmount_ = auction_.minIncreaseAmount;
            }
            require(
                amount_ >= highestBid_.amount + minIncreaseAmount_,
                "FeralfileEnglishAuction: bid amount should follow the minimum increment"
            );
        } else {
            // check bidding amount follow the minimum price
            require(
                amount_ >= auction_.minPrice,
                "FeralfileEnglishAuction: bid amount should be greater than minimum price"
            );
        }
    }

    function _placeBid(
        uint256 auctionID_,
        address bidder_,
        uint256 amount_,
        uint256 timestamp_,
        bool fromFeralFile_
    ) private {
        Auction memory auction_ = auctions[auctionID_];
        Bid memory highestBid_ = highestBids[auctionID_];
        // verify the bid is valid
        _validateBidRequest(auction_, highestBid_, amount_, timestamp_);
        // transfer last winning bid amount to last bidder if fromFeralFile is false
        if (!highestBid_.fromFeralFile && highestBid_.amount > 0) {
            payable(highestBid_.bidder).transfer(highestBid_.amount);
        }
        // replace the new bid to bid map
        highestBids[auctionID_] = Bid({
            bidder: bidder_,
            amount: amount_,
            fromFeralFile: fromFeralFile_
        });
        // if the gap to end time is lower than threshold, update the auction end time to block.timestamp + extend duration
        if (auction_.endAt - block.timestamp <= auction_.extendThreshold) {
            auctions[auctionID_].endAt =
                block.timestamp +
                auction_.extendDuration;
        }
        // emit new bid event
        emit NewBid(auctionID_, msg.sender, msg.value, fromFeralFile_);
    }

    function placeBid(uint256 auctionID_) external payable {
        _placeBid(auctionID_, msg.sender, msg.value, block.timestamp, false);
    }

    function placeSignedBid(
        uint256 auctionID_,
        address bidder_,
        uint256 amount_,
        uint256 timestamp_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) external {
        bytes32 message_ = keccak256(
            abi.encode(block.chainid, auctionID_, bidder_, amount_, timestamp_)
        );
        require(
            isValidSignature(message_, r_, s_, v_),
            "FeralfileEnglishAuction: invalid signature"
        );

        _placeBid(auctionID_, bidder_, amount_, timestamp_, true);
    }

    function settleAuction(
        uint256 auctionID_,
        address tokenAddr_,
        address vaultAddr_,
        SaleData memory saleData_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) external onlyOwner {
        Auction memory auction_ = auctions[auctionID_];
        require(auction_.id != 0, "FeralfileEnglishAuction: auction not found");

        require(
            auction_.endAt <= block.timestamp,
            "FeralfileEnglishAuction: auction not ended"
        );

        Bid memory highestBid_ = highestBids[auctionID_];
        require(
            highestBid_.bidder != address(0),
            "FeralfileEnglishAuction: no bid"
        );

        require(
            !auction_.isSettled,
            "FeralfileEnglishAuction: auction already settled"
        );

        auctions[auctionID_].isSettled = true;

        saleData_.destination = highestBid_.bidder;
        saleData_.payByVaultContract = true;

        // transfer winning bid amount to Feral File Vault contract if winning bid is crypto bid
        if (!highestBid_.fromFeralFile && highestBid_.amount > 0) {
            payable(vaultAddr_).transfer(highestBid_.amount);
        }
        IFeralfileExhibitionV4(tokenAddr_).buyArtworks(r_, s_, v_, saleData_);

        emit AuctionSettled(
            auctionID_,
            tokenAddr_,
            highestBid_.bidder,
            highestBid_.amount
        );
    }

    event NewBid(
        uint256 indexed auctionId,
        address indexed bidder,
        uint256 indexed amount,
        bool fromFeralFile
    );

    event AuctionSettled(
        uint256 indexed auctionId,
        address indexed contractAddress,
        address indexed winner,
        uint256 amount
    );
}
