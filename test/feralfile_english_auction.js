const { time, BN } = require("@openzeppelin/test-helpers");

const FeralfileEnglishAuction = artifacts.require("FeralfileEnglishAuction");
const FeralfileExhibitionV4 = artifacts.require("FeralfileExhibitionV4");
const FeralfileVault = artifacts.require("FeralfileVault");

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";
const COST_RECEIVER = "0x46f2B641d8702f29c45f6D06292dC34Eb9dB1801";
const CONTRACT_URI = "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

contract("FeralfileEnglishAuction", async (accounts) => {
    before(async function () {
        this.signer = accounts[0];
        this.vault = await FeralfileVault.new(this.signer);
        this.seriesIds = [0, 1, 2, 3];
        this.seriesMaxSupply = [1, 1, 2, 2];
        this.engAuction = await FeralfileEnglishAuction.new(this.signer);
        this.ff4Contract = await FeralfileExhibitionV4.new(
            "Feral File V4 Test",
            "FFv4",
            true,
            true,
            this.signer,
            this.vault.address,
            COST_RECEIVER,
            CONTRACT_URI,
            this.seriesIds,
            this.seriesMaxSupply
        );

        // mint tokens
        this.tokenId1 = this.seriesIds[0] * 1000000 + 0;
        this.tokenId2 = this.seriesIds[1] * 1000000 + 0;
        this.tokenId3 = this.seriesIds[3] * 1000000 + 0;
        await this.ff4Contract.mintArtworks([
            [this.seriesIds[0], this.tokenId1, this.ff4Contract.address],
            [this.seriesIds[1], this.tokenId2, this.ff4Contract.address],
            [this.seriesIds[2], this.tokenId3, this.ff4Contract.address],
        ]);
        await this.ff4Contract.startSale();
        const tokenOwner1 = await this.ff4Contract.ownerOf(this.tokenId1);
        assert.equal(tokenOwner1, this.ff4Contract.address);
        const tokenOwner2 = await this.ff4Contract.ownerOf(this.tokenId2);
        assert.equal(tokenOwner2, this.ff4Contract.address);
        const tokenOwner3 = await this.ff4Contract.ownerOf(this.tokenId3);
        assert.equal(tokenOwner3, this.ff4Contract.address);
        const artwork1 = await this.ff4Contract.getArtwork(this.tokenId1);
        assert.equal(artwork1.tokenId, this.tokenId1);
        assert.equal(artwork1.seriesId, this.seriesIds[0]);
        const artwork2 = await this.ff4Contract.getArtwork(this.tokenId2);
        assert.equal(artwork2.tokenId, this.tokenId2);
        assert.equal(artwork2.seriesId, this.seriesIds[1]);
        const artwork3 = await this.ff4Contract.getArtwork(this.tokenId3);
        assert.equal(artwork3.tokenId, this.tokenId3);
        assert.equal(artwork3.seriesId, this.seriesIds[2]);

        // auctionId for token 1
        this.aucID1 = await this.engAuction.getAuctionID(
            this.ff4Contract.address,
            [this.tokenId1]
        );

        // auctionId for token 2
        this.aucID2 = await this.engAuction.getAuctionID(
            this.ff4Contract.address,
            [this.tokenId2]
        );

        // auctionId for token 3
        this.aucID3 = await this.engAuction.getAuctionID(
            this.ff4Contract.address,
            [this.tokenId3]
        );
    });

    it("test get auction id", async function () {
        const expectedAuctionID =
            "0x42c6b6698b1a201b0e18f5d7c8a26d0469e15274b0623e2c06b9a3ea2825ac5f";
        const auc1 = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000, 3000001, 4000000, 4000001, 4000002]
        );
        const auc2 = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000001, 3000000, 4000000, 4000001, 4000002]
        );
        const auc3 = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [4000000, 4000001, 4000002, 3000000, 3000001]
        );
        const auc4 = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [4000001, 3000001, 3000000, 4000000, 4000002]
        );
        assert.equal(expectedAuctionID, auc1);
        assert.equal(expectedAuctionID, auc2);
        assert.equal(expectedAuctionID, auc3);
        assert.equal(expectedAuctionID, auc4);
    });

    it("test auction is on going", async function () {
        const auctionContract = await FeralfileEnglishAuction.new(this.signer);
        const aucID = await auctionContract.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000, 3000001, 4000000, 4000001, 4000002]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 seconds
        const endTime = latestTime.add(new BN(10 * 60)).toString(); // end in 10 minutes
        await auctionContract.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                300, //extendDuration
                10, // extendThreshold
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        const isActionOngoing = await auctionContract.auctionOngoing(aucID);
        assert.equal(true, isActionOngoing);
    });

    it("test auction is not started yet", async function () {
        const auctionContract = await FeralfileEnglishAuction.new(this.signer);
        const aucID = await auctionContract.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000, 3000001, 4000000, 4000001, 4000002]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(10 * 60)).toString(); // start in 10 minutes
        const endTime = latestTime.add(new BN(20 * 60)).toString(); // end in 20 minutes
        await auctionContract.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                300, //extendDuration
                10, // extendThreshold
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        const isActionOngoing = await auctionContract.auctionOngoing(aucID);
        assert.equal(false, isActionOngoing);
    });

    it("test auction ended", async function () {
        const auctionContract = await FeralfileEnglishAuction.new(this.signer);
        const aucID = await auctionContract.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000, 3000001, 4000000, 4000001, 4000002]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 second
        const endTime = latestTime.add(new BN(10)).toString(); // end in 10 seconds
        await auctionContract.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                300, //extendDuration
                10, // extendThreshold
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 15 seconds to end auction
        await time.increase(time.duration.seconds(15));

        const isActionOngoing = await auctionContract.auctionOngoing(aucID);
        assert.equal(false, isActionOngoing);
    });

    it("test set auctions successfully", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000, 3000001, 4000000, 4000001, 4000002]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(5 * 60)).toString(); // start in 5 minutes
        const endTime = latestTime.add(new BN(10 * 60)).toString(); // end in 10 minutes
        await this.engAuction.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt
                300, //extendDuration
                10, // extendThreshold
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        const auction = await this.engAuction.auctions(aucID);
        assert.equal(aucID, auction.id);
        assert.equal(startTime, auction.startAt);
        assert.equal(endTime, auction.endAt);
    });

    it("test set auctions failed because existed", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000, 3000001, 4000000, 4000001, 4000002]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(5 * 60)).toString(); // start in 5 minutes
        const endTime = latestTime.add(new BN(10 * 60)).toString(); // end in 10 minutes
        try {
            await this.engAuction.setEnglishAuctions([
                [
                    aucID, // id
                    startTime, // startAt
                    endTime, // endAt
                    300, //extendDuration
                    10, // extendThreshold
                    20, // minIncreaseFactor
                    30, // minIncreaseAmount
                    web3.utils.toWei("0.1", "ether"), // minPrice
                    false,
                ],
            ]);
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileEnglishAuction: auction already exist"
                )
            );
        }
    });

    it("test bid on auction by crypto successfully", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 seconds
        const endTime = latestTime.add(new BN(60 * 60)).toString(); // end in 1 hour

        await this.engAuction.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                600, //extendDuration // 10 minutes
                300, // extendThreshold // 5 minutes
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        const isActionOngoing = await this.engAuction.auctionOngoing(aucID);
        assert.equal(true, isActionOngoing);

        await this.engAuction.bidOnAuction(aucID, {
            from: accounts[2],
            value: web3.utils.toWei("0.2", "ether"),
        });
        const latestBid = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.2", "ether"), latestBid.amount);
        assert.equal(accounts[2], latestBid.bidder);
    });

    it("test bid on auction by Feralfile successfully", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000002]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 second
        const endTime = latestTime.add(new BN(10)).toString(); // end in 10 seconds

        await this.engAuction.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                600, //extendDuration // 10 minutes
                300, // extendThreshold // 5 minutes
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        const biddingTime = latestTime.add(new BN(10)).toString();

        // Generate signature
        const signParams = web3.eth.abi.encodeParameters(
            ["uint", "bytes32", "address", "uint256", "uint256"],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                aucID,
                accounts[2],
                web3.utils.toWei("0.25", "ether"),
                biddingTime,
            ]
        );
        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);
        // Generate signature

        await this.engAuction.bidOnAuctionByFeralFile(
            aucID,
            accounts[2],
            web3.utils.toWei("0.25", "ether"),
            biddingTime,
            r,
            s,
            web3.utils.toDecimal(v) + 27, // magic 27
            {
                from: accounts[5], // any address from biconomy
            }
        );
        const latestBid = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.25", "ether"), latestBid.amount);
        assert.equal(accounts[2], latestBid.bidder);
        assert.equal(true, latestBid.fromFeralFile);
    });

    it("test bid on auction failed because price under last bid", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000]
        );

        try {
            await this.engAuction.bidOnAuction(aucID, {
                from: accounts[3],
                value: (web3.utils.toWei("0.1", "ether") * 1.1).toFixed(0), // 10% increase, expect 20% increase
            });
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileEnglishAuction: bid amount should be greater than highest bid amount"
                )
            );
        }
    });

    it("test bid on auction failed because under minimum increment", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000]
        );

        try {
            await this.engAuction.bidOnAuction(aucID, {
                from: accounts[3],
                value: (web3.utils.toWei("0.2", "ether") * 1.1).toFixed(0), // 10% increase, expect 20% increase
            });
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileEnglishAuction: bid amount should follow the minimum increment"
                )
            );
        }
    });

    it("test extends bidding time", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000001]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 second
        const endTime = latestTime.add(new BN(10)).toString(); // end in 10 seconds

        await this.engAuction.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                600, //extendDuration // 10 minutes
                300, // extendThreshold // 5 minutes
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        const timestamp = await time.latest();
        await this.engAuction.bidOnAuction(aucID, {
            from: accounts[2],
            value: web3.utils.toWei("0.2", "ether"),
        });
        const latestBid = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.2", "ether"), latestBid.amount);
        assert.equal(accounts[2], latestBid.bidder);

        const auction = await this.engAuction.auctions(aucID);
        assert.equal(
            timestamp.add(new BN(600)).toString(),
            auction.endAt.toString()
        ); // extend 10 minutes
    });

    it("test refund to previous crypto bidder by new bidder by crypto", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000006]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 seconds
        const endTime = latestTime.add(new BN(60 * 60)).toString(); // end in 1 hour

        await this.engAuction.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                600, //extendDuration // 10 minutes
                300, // extendThreshold // 5 minutes
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        const isActionOngoing = await this.engAuction.auctionOngoing(aucID);
        assert.equal(true, isActionOngoing);

        // Bid by account[2]
        await this.engAuction.bidOnAuction(aucID, {
            from: accounts[2],
            value: web3.utils.toWei("0.2", "ether"),
        });

        const latestBidAcc2 = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.2", "ether"), latestBidAcc2.amount);
        assert.equal(accounts[2], latestBidAcc2.bidder);

        const acc2BalanceBeforeRefund = await web3.eth.getBalance(accounts[2]);

        // Bid by accounts[3]
        await this.engAuction.bidOnAuction(aucID, {
            from: accounts[3],
            value: web3.utils.toWei("0.3", "ether"),
        });
        const latestBidAcc3 = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.3", "ether"), latestBidAcc3.amount);
        assert.equal(accounts[3], latestBidAcc3.bidder);

        const acc2BalanceAfterRefund = await web3.eth.getBalance(accounts[2]);

        assert.equal(
            latestBidAcc2.amount,
            acc2BalanceAfterRefund - acc2BalanceBeforeRefund
        );
    });

    it("test refund to previous crypto bidder by new bidder from Feralfile", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000007]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 seconds
        const endTime = latestTime.add(new BN(60 * 60)).toString(); // end in 1 hour

        await this.engAuction.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                600, //extendDuration // 10 minutes
                300, // extendThreshold // 5 minutes
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        const isActionOngoing = await this.engAuction.auctionOngoing(aucID);
        assert.equal(true, isActionOngoing);

        // Bid by account[2]
        await this.engAuction.bidOnAuction(aucID, {
            from: accounts[2],
            value: web3.utils.toWei("0.2", "ether"),
        });

        const latestBidAcc2 = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.2", "ether"), latestBidAcc2.amount);
        assert.equal(accounts[2], latestBidAcc2.bidder);
        assert.equal(false, latestBidAcc2.fromFeralFile);

        const acc2BalanceBeforeRefund = await web3.eth.getBalance(accounts[2]);

        // Bid by accounts[3]
        const biddingTime = latestTime.add(new BN(10)).toString();
        // Generate signature
        const biddingParams = web3.eth.abi.encodeParameters(
            ["uint", "bytes32", "address", "uint256", "uint256"],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                aucID,
                accounts[3],
                web3.utils.toWei("0.3", "ether"),
                biddingTime,
            ]
        );
        const biddingHash = web3.utils.keccak256(biddingParams);
        var biddingSig = await web3.eth.sign(biddingHash, this.signer);
        biddingSig = biddingSig.substr(2);
        const bidR = "0x" + biddingSig.slice(0, 64);
        const bidS = "0x" + biddingSig.slice(64, 128);
        const bidV = "0x" + biddingSig.slice(128, 130);
        // Generate signature

        // Bid by account[2]
        await this.engAuction.bidOnAuctionByFeralFile(
            aucID,
            accounts[3],
            web3.utils.toWei("0.3", "ether"),
            biddingTime,
            bidR,
            bidS,
            web3.utils.toDecimal(bidV) + 27, // magic 27
            {
                from: accounts[7], // any address from biconomy
            }
        );

        const latestBidAcc3 = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.3", "ether"), latestBidAcc3.amount);
        assert.equal(accounts[3], latestBidAcc3.bidder);
        assert.equal(true, latestBidAcc3.fromFeralFile);

        const acc2BalanceAfterRefund = await web3.eth.getBalance(accounts[2]);

        assert.equal(
            latestBidAcc2.amount,
            acc2BalanceAfterRefund - acc2BalanceBeforeRefund
        );
    });

    it("test refund to previous FF bidder by new bidder by crypto", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000008]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 seconds
        const endTime = latestTime.add(new BN(60 * 60)).toString(); // end in 1 hour

        await this.engAuction.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                600, //extendDuration // 10 minutes
                300, // extendThreshold // 5 minutes
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        const isActionOngoing = await this.engAuction.auctionOngoing(aucID);
        assert.equal(true, isActionOngoing);

        // Bid by account[2]
        const biddingTime = latestTime.add(new BN(10)).toString();
        // Generate signature
        const biddingParams = web3.eth.abi.encodeParameters(
            ["uint", "bytes32", "address", "uint256", "uint256"],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                aucID,
                accounts[2],
                web3.utils.toWei("0.2", "ether"),
                biddingTime,
            ]
        );
        const biddingHash = web3.utils.keccak256(biddingParams);
        var biddingSig = await web3.eth.sign(biddingHash, this.signer);
        biddingSig = biddingSig.substr(2);
        const bidR = "0x" + biddingSig.slice(0, 64);
        const bidS = "0x" + biddingSig.slice(64, 128);
        const bidV = "0x" + biddingSig.slice(128, 130);
        // Generate signature
        await this.engAuction.bidOnAuctionByFeralFile(
            aucID,
            accounts[2],
            web3.utils.toWei("0.2", "ether"),
            biddingTime,
            bidR,
            bidS,
            web3.utils.toDecimal(bidV) + 27, // magic 27
            {
                from: accounts[7], // any address from biconomy
            }
        );

        const latestBidAcc2 = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.2", "ether"), latestBidAcc2.amount);
        assert.equal(accounts[2], latestBidAcc2.bidder);
        assert.equal(true, latestBidAcc2.fromFeralFile);

        const acc2BalanceBeforeRefund = await web3.eth.getBalance(accounts[2]);

        // Bid by accounts[3]
        await this.engAuction.bidOnAuction(aucID, {
            from: accounts[3],
            value: web3.utils.toWei("0.3", "ether"),
        });
        const latestBidAcc3 = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.3", "ether"), latestBidAcc3.amount);
        assert.equal(accounts[3], latestBidAcc3.bidder);

        const acc2BalanceAfterRefund = await web3.eth.getBalance(accounts[2]);

        assert.equal(acc2BalanceBeforeRefund, acc2BalanceAfterRefund); // balance no change
    });

    it("test refund to previous FF bidder by new FF bidder", async function () {
        const aucID = await this.engAuction.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000009]
        );
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 seconds
        const endTime = latestTime.add(new BN(60 * 60)).toString(); // end in 1 hour

        await this.engAuction.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                600, //extendDuration // 10 minutes
                300, // extendThreshold // 5 minutes
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        const isActionOngoing = await this.engAuction.auctionOngoing(aucID);
        assert.equal(true, isActionOngoing);

        // Bid by account[2]
        const biddingTime2 = latestTime.add(new BN(10)).toString();
        // Generate signature
        const biddingParams2 = web3.eth.abi.encodeParameters(
            ["uint", "bytes32", "address", "uint256", "uint256"],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                aucID,
                accounts[2],
                web3.utils.toWei("0.2", "ether"),
                biddingTime2,
            ]
        );
        const biddingHash2 = web3.utils.keccak256(biddingParams2);
        var biddingSig2 = await web3.eth.sign(biddingHash2, this.signer);
        biddingSig2 = biddingSig2.substr(2);
        const bidR2 = "0x" + biddingSig2.slice(0, 64);
        const bidS2 = "0x" + biddingSig2.slice(64, 128);
        const bidV2 = "0x" + biddingSig2.slice(128, 130);
        // Generate signature

        await this.engAuction.bidOnAuctionByFeralFile(
            aucID,
            accounts[2],
            web3.utils.toWei("0.2", "ether"),
            biddingTime2,
            bidR2,
            bidS2,
            web3.utils.toDecimal(bidV2) + 27, // magic 27
            {
                from: accounts[7], // any address from biconomy
            }
        );

        const latestBidAcc2 = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.2", "ether"), latestBidAcc2.amount);
        assert.equal(accounts[2], latestBidAcc2.bidder);
        assert.equal(true, latestBidAcc2.fromFeralFile);

        const acc2BalanceBeforeRefund = await web3.eth.getBalance(accounts[2]);

        // Bid by accounts[3]
        const biddingTime3 = (await time.latest()).add(new BN(2)).toString();
        // Generate signature
        const biddingParams = web3.eth.abi.encodeParameters(
            ["uint", "bytes32", "address", "uint256", "uint256"],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                aucID,
                accounts[3],
                web3.utils.toWei("0.3", "ether"),
                biddingTime3,
            ]
        );
        const biddingHash3 = web3.utils.keccak256(biddingParams);
        var biddingSig3 = await web3.eth.sign(biddingHash3, this.signer);
        biddingSig3 = biddingSig3.substr(2);
        const bidR3 = "0x" + biddingSig3.slice(0, 64);
        const bidS3 = "0x" + biddingSig3.slice(64, 128);
        const bidV3 = "0x" + biddingSig3.slice(128, 130);
        // Generate signature

        // Bid by account[2]
        await this.engAuction.bidOnAuctionByFeralFile(
            aucID,
            accounts[3],
            web3.utils.toWei("0.3", "ether"),
            biddingTime3,
            bidR3,
            bidS3,
            web3.utils.toDecimal(bidV3) + 27, // magic 27
            {
                from: accounts[4], // any address from biconomy
            }
        );

        const latestBidAcc3 = await this.engAuction.auctionLatestBid(aucID);
        assert.equal(web3.utils.toWei("0.3", "ether"), latestBidAcc3.amount);
        assert.equal(accounts[3], latestBidAcc3.bidder);
        assert.equal(true, latestBidAcc3.fromFeralFile);

        const acc2BalanceAfterRefund = await web3.eth.getBalance(accounts[2]);

        assert.equal(acc2BalanceBeforeRefund, acc2BalanceAfterRefund); // balance no change
    });

    it("test settle auction purchase by crypto", async function () {
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 second
        const endTime = latestTime.add(new BN(10)).toString(); // end in 10 seconds
        await this.engAuction.setEnglishAuctions([
            [
                this.aucID1, // id
                startTime, // startAt
                endTime, // endAt
                0, //extendDuration
                0, // extendThreshold
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        const contractBalanceBeforeBid = await web3.eth.getBalance(
            this.engAuction.address
        );

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        await this.engAuction.bidOnAuction(this.aucID1, {
            from: accounts[2],
            value: web3.utils.toWei("0.2", "ether"),
        });

        const latestBid = await this.engAuction.auctionLatestBid(this.aucID1);
        assert.equal(web3.utils.toWei("0.2", "ether"), latestBid.amount);
        assert.equal(accounts[2], latestBid.bidder);

        // Waiting for auction to end...
        await time.increase(time.duration.seconds(10));

        const contractBalanceBeforeSettle = await web3.eth.getBalance(
            this.engAuction.address
        );
        assert.equal(
            latestBid.amount,
            contractBalanceBeforeSettle - contractBalanceBeforeBid
        );

        // Generate signature
        const latestTimeSettle = await time.latest();
        const expiryTime = latestTimeSettle.add(new BN(300)).toString();
        const saleData = [
            latestBid.amount,
            "0",
            expiryTime,
            accounts[2],
            [this.tokenId1],
            [
                [
                    [accounts[3], 8000],
                    [accounts[4], 2000],
                ],
            ],
            true,
        ];
        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                this.ff4Contract.address,
                saleData,
            ]
        );
        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);
        // Generate signature

        await this.engAuction.settleAuction(
            this.ff4Contract.address,
            this.vault.address,
            saleData,
            r,
            s,
            web3.utils.toDecimal(v) + 27 // magic 27
        );

        const owner = await this.ff4Contract.ownerOf(this.tokenId1);
        assert.equal(owner, accounts[2]);

        const contractBalanceAfter = await web3.eth.getBalance(
            this.engAuction.address
        );
        assert.equal(contractBalanceAfter, contractBalanceBeforeBid);

        const auction = await this.engAuction.auctions(this.aucID1);
        assert.equal(true, auction.isSettled);
    });

    it("test settle auction purchase by credit card", async function () {
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 second
        const endTime = latestTime.add(new BN(10)).toString(); // end in 10 seconds
        await this.engAuction.setEnglishAuctions([
            [
                this.aucID2, // id
                startTime, // startAt
                endTime, // endAt
                0, //extendDuration
                0, // extendThreshold
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        const biddingTime = latestTime.add(new BN(10)).toString();
        // Generate signature
        const biddingParams = web3.eth.abi.encodeParameters(
            ["uint", "bytes32", "address", "uint256", "uint256"],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                this.aucID2,
                accounts[3],
                web3.utils.toWei("0.25", "ether"),
                biddingTime,
            ]
        );
        const biddingHash = web3.utils.keccak256(biddingParams);
        var biddingSig = await web3.eth.sign(biddingHash, this.signer);
        biddingSig = biddingSig.substr(2);
        const bidR = "0x" + biddingSig.slice(0, 64);
        const bidS = "0x" + biddingSig.slice(64, 128);
        const bidV = "0x" + biddingSig.slice(128, 130);
        // Generate signature

        await this.engAuction.bidOnAuctionByFeralFile(
            this.aucID2,
            accounts[3],
            web3.utils.toWei("0.25", "ether"),
            biddingTime,
            bidR,
            bidS,
            web3.utils.toDecimal(bidV) + 27, // magic 27
            {
                from: accounts[6], // any address from biconomy
            }
        );

        const latestBid = await this.engAuction.auctionLatestBid(this.aucID2);
        assert.equal(web3.utils.toWei("0.25", "ether"), latestBid.amount);
        assert.equal(accounts[3], latestBid.bidder);
        assert.equal(true, latestBid.fromFeralFile);

        // Waiting for auction to end...
        await time.increase(time.duration.seconds(10));

        // Generate signature
        const latestTimeSettle = await time.latest();
        const expiryTime = latestTimeSettle.add(new BN(300)).toString();
        const saleData = [
            latestBid.amount,
            "0",
            expiryTime,
            accounts[3],
            [this.tokenId2],
            [
                [
                    [accounts[3], 8000],
                    [accounts[4], 2000],
                ],
            ],
            true,
        ];
        const settleSignParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                this.ff4Contract.address,
                saleData,
            ]
        );
        const settleHash = web3.utils.keccak256(settleSignParams);
        var settleSig = await web3.eth.sign(settleHash, this.signer);
        settleSig = settleSig.substr(2);
        const r = "0x" + settleSig.slice(0, 64);
        const s = "0x" + settleSig.slice(64, 128);
        const v = "0x" + settleSig.slice(128, 130);
        // Generate signature

        await web3.eth.sendTransaction({
            to: this.vault.address,
            from: accounts[8],
            value: latestBid.amount,
        });

        await this.engAuction.settleAuction(
            this.ff4Contract.address,
            this.vault.address,
            saleData,
            r,
            s,
            web3.utils.toDecimal(v) + 27 // magic 27
        );

        const owner = await this.ff4Contract.ownerOf(this.tokenId2);
        assert.equal(owner, accounts[3]);
    });

    it("test settle auction with FF cost", async function () {
        const latestTime = await time.latest();
        const startTime = latestTime.add(new BN(1)).toString(); // start in 1 second
        const endTime = latestTime.add(new BN(10)).toString(); // end in 10 seconds
        await this.engAuction.setEnglishAuctions([
            [
                this.aucID3, // id
                startTime, // startAt
                endTime, // endAt
                0, //extendDuration
                0, // extendThreshold
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                web3.utils.toWei("0.1", "ether"), // minPrice
                false,
            ],
        ]);

        const contractBalanceBeforeBid = await web3.eth.getBalance(
            this.engAuction.address
        );

        // wait 2 seconds to start auction
        await time.increase(time.duration.seconds(2));

        await this.engAuction.bidOnAuction(this.aucID3, {
            from: accounts[2],
            value: web3.utils.toWei("0.2", "ether"),
        });

        const latestBid = await this.engAuction.auctionLatestBid(this.aucID3);
        assert.equal(web3.utils.toWei("0.2", "ether"), latestBid.amount);
        assert.equal(accounts[2], latestBid.bidder);

        // Waiting for auction to end...
        await time.increase(time.duration.seconds(10));

        const contractBalanceBeforeSettle = await web3.eth.getBalance(
            this.engAuction.address
        );
        assert.equal(
            latestBid.amount,
            contractBalanceBeforeSettle - contractBalanceBeforeBid
        );

        // Generate signature
        const latestTimeSettle = await time.latest();
        const expiryTime = latestTimeSettle.add(new BN(300)).toString();
        const saleData = [
            latestBid.amount,
            web3.utils.toWei("0.1", "ether").toString(),
            expiryTime,
            accounts[2],
            [this.tokenId3],
            [
                [
                    [accounts[7], 8000],
                    [accounts[8], 2000],
                ],
            ],
            true,
        ];
        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                this.ff4Contract.address,
                saleData,
            ]
        );
        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);
        // Generate signature

        // Get balance before settle
        const costReceiverBalanceBeforeSettle = await web3.eth.getBalance(
            COST_RECEIVER
        );
        const acc7BalanceBeforeSettle = await web3.eth.getBalance(accounts[7]);
        const acc8BalanceBeforeSettle = await web3.eth.getBalance(accounts[8]);

        await this.engAuction.settleAuction(
            this.ff4Contract.address,
            this.vault.address,
            saleData,
            r,
            s,
            web3.utils.toDecimal(v) + 27 // magic 27
        );

        const owner = await this.ff4Contract.ownerOf(this.tokenId1);
        assert.equal(owner, accounts[2]);

        const contractBalanceAfter = await web3.eth.getBalance(
            this.engAuction.address
        );
        assert.equal(contractBalanceAfter, contractBalanceBeforeBid);

        const auction = await this.engAuction.auctions(this.aucID3);
        assert.equal(true, auction.isSettled);

        // Assert cost receiver balance
        const costReceiverBalanceAfterSettle = await web3.eth.getBalance(
            COST_RECEIVER
        );
        assert.equal(
            web3.utils.toWei("0.1", "ether"),
            costReceiverBalanceAfterSettle - costReceiverBalanceBeforeSettle
        );

        // Assert revenues balances
        const acc7BalanceAfterSettle = await web3.eth.getBalance(accounts[7]);
        const acc8BalanceAfterSettle = await web3.eth.getBalance(accounts[8]);
        assert.equal(
            web3.utils.toWei("0.08", "ether"),
            acc7BalanceAfterSettle - acc7BalanceBeforeSettle
        );
        assert.equal(
            web3.utils.toWei("0.02", "ether"),
            acc8BalanceAfterSettle - acc8BalanceBeforeSettle
        );
    });
});
