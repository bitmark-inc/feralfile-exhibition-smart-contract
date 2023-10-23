// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./FeralfileSaleData.sol";
import "./ECDSASigner.sol";
import "./IFeralfileVault.sol";

contract FFV4 is FeralfileSaleData {
    function buyArtworks(
        bytes32 r_,
        bytes32 s_,
        uint8 v_,
        SaleData calldata saleData_
    ) external payable {}
}

contract FeralfileEnglishAuction is Ownable, FeralfileSaleData, ECDSASigner {
    struct EnglishAuction {
        bytes32 id;
        uint256 startAt;
        uint256 endAt;
        uint256 extendDuration;
        uint256 extendThreshold;
        uint256 minIncreaseFactor;
        uint256 minIncreaseAmount;
        uint256 minPrice;
    }

    struct Bid {
        address bidder;
        uint256 amount;
        bool fromFeralFile;
    }

    mapping(bytes32 => EnglishAuction) public auctions;
    mapping(bytes32 => Bid) public auctionLatestBid;

    constructor(address signer_) ECDSASigner(signer_) {}

    function getAuctionID(
        address tokenContractAddr_,
        uint256[] memory tokenIDs_
    ) public view returns (bytes32) {
        // order the tokenIDs based on certain rule
        for (uint i = 0; i < tokenIDs_.length - 1; i++) {
            uint minIndex = i;
            for (uint j = i + 1; j < tokenIDs_.length; j++) {
                if (tokenIDs_[j] < tokenIDs_[minIndex]) {
                    minIndex = j;
                }
            }
            (tokenIDs_[i], tokenIDs_[minIndex]) = (
                tokenIDs_[minIndex],
                tokenIDs_[i]
            );
        }

        // hash the (contractAddr + ordered tokenIDs)
        bytes32 message = keccak256(
            abi.encode(block.chainid, tokenContractAddr_, tokenIDs_)
        );

        // return hash as identifier
        return message;
    }

    function auctionOngoing(bytes32 aucId_) public view returns (uint256) {
        // check the auction status and return
        require(
            auctions[aucId_].id != 0,
            "FeralfileEnglishAuction: auction not found"
        );

        return block.timestamp;

        // return
        //     auctions[aucId_].startAt <= block.timestamp &&
        //     auctions[aucId_].endAt >= block.timestamp;
    }

    function setEnglishAuctions(
        EnglishAuction[] calldata auctions_
    ) external onlyOwner {
        // for-loop to set the auction
        // able to reset an auction if there's no winner and highest bid = 0
        for (uint256 i = 0; i < auctions_.length; i++) {
            EnglishAuction memory auc_ = auctions[auctions_[i].id];

            require(
                auc_.id == 0,
                "FeralfileEnglishAuction: auction already exist"
            );

            require(
                auctions_[i].startAt > block.timestamp,
                "FeralfileEnglishAuction: auction start time should be in the future"
            );

            require(
                auctions_[i].endAt > auctions_[i].startAt,
                "FeralfileEnglishAuction: auction end time should be after start time"
            );

            require(
                auctions_[i].extendDuration > 0,
                "FeralfileEnglishAuction: auction extend duration should be greater than 0"
            );

            require(
                auctions_[i].extendThreshold > 0,
                "FeralfileEnglishAuction: auction extend threshold should be greater than 0"
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

    function isValidBidRequest(
        EnglishAuction memory auction_,
        Bid memory lastBid_,
        address bidder_,
        uint256 amount_,
        uint256 timestamp_
    ) internal pure returns (bool) {
        // make sure the auction exist
        require(
            auction_.id.length > 0,
            "FeralfileEnglishAuction: auction not found"
        );
        // check if auction started
        require(
            auction_.startAt < timestamp_,
            "FeralfileEnglishAuction: auction not started"
        );
        // check if now() under auction end time
        require(
            auction_.endAt > timestamp_,
            "FeralfileEnglishAuction: auction already ended"
        );
        // check if current winner not equal to current bidder (optional)
        require(
            lastBid_.bidder != bidder_,
            "FeralfileEnglishAuction: bidder is the current winner"
        );
        // check bidding amount > highest bid amount
        require(
            amount_ > lastBid_.amount,
            "FeralfileEnglishAuction: bid amount should be greater than highest bid amount"
        );
        if (lastBid_.amount > 0) {
            // check bidding amount follow the minimum increment
            uint256 minIncreaseAmount_ = (lastBid_.amount *
                auction_.minIncreaseFactor) / 100;
            if (minIncreaseAmount_ < auction_.minIncreaseAmount) {
                minIncreaseAmount_ = auction_.minIncreaseAmount;
            }
            require(
                amount_ >= lastBid_.amount + minIncreaseAmount_,
                "FeralfileEnglishAuction: bid amount should follow the minimum increment"
            );
        } else {
            // check bidding amount follow the minimum price
            require(
                amount_ >= auction_.minPrice,
                "FeralfileEnglishAuction: bid amount should be greater than minimum price"
            );
        }
        return true;
    }

    function bidOnAuction(bytes32 aucId_) external payable {
        // get the auction object
        EnglishAuction memory auction_ = auctions[aucId_];
        // get the last auction bid object
        Bid memory lastBid_ = auctionLatestBid[aucId_];
        // verify the bid is valid
        require(
            isValidBidRequest(
                auction_,
                lastBid_,
                msg.sender,
                msg.value,
                block.timestamp
            ),
            "FeralfileEnglishAuction: invalid bid"
        );
        // transfer last winning bid amount to last bidder if fromFeralFile is false
        if (!lastBid_.fromFeralFile) {
            payable(lastBid_.bidder).transfer(lastBid_.amount);
        }
        // replace the new bid to bid map
        auctionLatestBid[aucId_] = Bid({
            bidder: msg.sender,
            amount: msg.value,
            fromFeralFile: false
        });
        // if the gap to end time is lower than threshold, update the auction end time to now() + extend duration
        if (auction_.endAt - block.timestamp < auction_.extendThreshold) {
            auctions[aucId_].endAt = block.timestamp + auction_.extendDuration;
        }
        // emit new bid event
        emit NewBid(aucId_, msg.sender, msg.value);
    }

    function bidOnAuctionByFeralFile(
        bytes32 aucId_,
        address bidder_,
        uint256 amount_,
        uint256 timestamp_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) external {
        // get the auction object
        EnglishAuction memory auction_ = auctions[aucId_];
        // get the last auction bid object
        Bid memory lastBid_ = auctionLatestBid[aucId_];
        // verify the bid is valid
        require(
            isValidBidRequest(auction_, lastBid_, bidder_, amount_, timestamp_),
            "FeralfileEnglishAuction: invalid bid"
        );
        // make sure the auction exist
        // get the auction object
        // get the last auction bid object
        // check if current winner not equal to current bidder (optional)
        // check if now() under auction end time
        // check bidding amount > highest bid amount
        // check bidding amount follow the minimum increment
        // check timestamp should in 10 mins

        // check the signature of (aucId + bidder + amount + timestamp) to make sure the request is from Feral File
        bytes32 message_ = keccak256(
            abi.encode(block.chainid, aucId_, bidder_, amount_, timestamp_)
        );
        require(
            isValidSignature(message_, r_, s_, v_),
            "FeralfileEnglishAuction: invalid signature"
        );
        // transfer last winning bid amount to last bidder if fromFeralFile is false
        if (!lastBid_.fromFeralFile) {
            payable(lastBid_.bidder).transfer(lastBid_.amount);
        }
        // replace the new bid to bid map
        auctionLatestBid[aucId_] = Bid({
            bidder: bidder_,
            amount: amount_,
            fromFeralFile: true
        });
        // if the gap to end time is lower than threshold, update the auction end time to now() + extend duration
        if (auction_.endAt - block.timestamp < auction_.extendThreshold) {
            auctions[aucId_].endAt = block.timestamp + auction_.extendDuration;
        }
        // emit new bid event
        emit NewBid(aucId_, bidder_, amount_);
    }

    function settleAuction(
        bytes32 r_,
        bytes32 s_,
        uint8 v_,
        address tokenContractAddr_,
        SaleData memory saleData_
    ) external onlyOwner {
        // transfer winning bid amount to Feral File Vault contract if winning bid is crypto bid
        bytes32 aucId_ = keccak256(
            abi.encodePacked(
                getAuctionID(tokenContractAddr_, saleData_.tokenIds)
            )
        );
        EnglishAuction memory auction_ = auctions[aucId_];
        require(
            auction_.id.length > 0,
            "FeralfileEnglishAuction: auction not found"
        );
        Bid memory lastBid_ = auctionLatestBid[aucId_];
        require(
            lastBid_.bidder != address(0),
            "FeralfileEnglishAuction: no bid"
        );

        saleData_.destination = lastBid_.bidder;
        FFV4(tokenContractAddr_).buyArtworks(r_, s_, v_, saleData_);
    }

    event NewBid(
        bytes32 indexed auctionId,
        address indexed bidder,
        uint256 indexed amount
    );
}
