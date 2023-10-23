const FeralfileEnglishAuction = artifacts.require("FeralfileEnglishAuction");

contract("FeralfileEnglishAuction", async (accounts) => {
    before(async function () {
        this.signer = accounts[1];
        this.engAuction = await FeralfileEnglishAuction.new(this.signer);
    });

    // it("test get auction id", async function () {
    //     const expectedAuctionID =
    //         "0x42c6b6698b1a201b0e18f5d7c8a26d0469e15274b0623e2c06b9a3ea2825ac5f";
    //     const auc1 = await this.engAuction.getAuctionID(
    //         "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
    //         [3000000, 3000001, 4000000, 4000001, 4000002]
    //     );
    //     const auc2 = await this.engAuction.getAuctionID(
    //         "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
    //         [3000001, 3000000, 4000000, 4000001, 4000002]
    //     );
    //     const auc3 = await this.engAuction.getAuctionID(
    //         "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
    //         [4000000, 4000001, 4000002, 3000000, 3000001]
    //     );
    //     const auc4 = await this.engAuction.getAuctionID(
    //         "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
    //         [4000001, 3000001, 3000000, 4000000, 4000002]
    //     );
    //     assert.equal(expectedAuctionID, auc1);
    //     assert.equal(expectedAuctionID, auc2);
    //     assert.equal(expectedAuctionID, auc3);
    //     assert.equal(expectedAuctionID, auc4);
    // });

    // it("set auctions successfully", async function () {
    //     const aucID = await this.engAuction.getAuctionID(
    //         "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
    //         [3000000, 3000001, 4000000, 4000001, 4000002]
    //     );
    //     const startTime = parseInt(new Date().getTime() / 1000) + 5 * 60;
    //     const endTime = parseInt(new Date().getTime() / 1000) + 10 * 60;
    //     await this.engAuction.setEnglishAuctions([
    //         [
    //             aucID, // id
    //             startTime, // startAt
    //             endTime, // endAt
    //             300, //extendDuration
    //             10, // extendThreshold
    //             20, // minIncreaseFactor
    //             30, // minIncreaseAmount
    //             1e15, // minPrice
    //         ],
    //     ]);

    //     const auction = await this.engAuction.auctions(aucID);
    //     assert.equal(aucID, auction.id);
    //     assert.equal(startTime, auction.startAt);
    //     assert.equal(endTime, auction.endAt);
    // });

    // it("set auctions failed", async function () {
    //     const aucID = await this.engAuction.getAuctionID(
    //         "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
    //         [3000000, 3000001, 4000000, 4000001, 4000002]
    //     );
    //     const startTime = parseInt(new Date().getTime() / 1000) + 5 * 60;
    //     const endTime = parseInt(new Date().getTime() / 1000) + 10 * 60;
    //     try {
    //         await this.engAuction.setEnglishAuctions([
    //             [
    //                 aucID, // id
    //                 startTime, // startAt
    //                 endTime, // endAt
    //                 300, //extendDuration
    //                 10, // extendThreshold
    //                 20, // minIncreaseFactor
    //                 30, // minIncreaseAmount
    //                 1e15, // minPrice
    //             ],
    //         ]);
    //     } catch (error) {
    //         assert.ok(
    //             error.message.includes(
    //                 "FeralfileEnglishAuction: auction already exist"
    //             )
    //         );
    //     }
    // });

    it("check is auction on going", async function () {
        const auctionContract = await FeralfileEnglishAuction.new(this.signer);
        const aucID = await auctionContract.getAuctionID(
            "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
            [3000000, 3000001, 4000000, 4000001, 4000002]
        );
        const startTime = parseInt(new Date().getTime() / 1000) + 2;
        const endTime = parseInt(new Date().getTime() / 1000) + 10 * 60;
        await auctionContract.setEnglishAuctions([
            [
                aucID, // id
                startTime, // startAt
                endTime, // endAt // add 10 minutes
                300, //extendDuration
                10, // extendThreshold
                20, // minIncreaseFactor
                30, // minIncreaseAmount
                1e15, // minPrice
            ],
        ]);

        setTimeout(async () => {
            const isActionOngoing = await auctionContract.auctionOngoing(aucID);
            console.log(BigInt(isActionOngoing).toString());
            console.log(startTime, endTime);
            assert.equal(false, isActionOngoing);
        }, 3000);

        // const isActionOngoing = await auctionContract.auctionOngoing(aucID);
        // console.log(BigInt(isActionOngoing).toString());
        // console.log(startTime, endTime);
        // assert.equal(true, isActionOngoing);
    });

    // it("bid successfully", async function () {
    //     const aucID = await this.engAuction.getAuctionID(
    //         "0x6dEc00C910F0C666D0b9f4c531a7Ff3B07388C24",
    //         [3000000, 3000001, 4000000, 4000001, 4000002]
    //     );
    //     await this.engAuction.bid(aucID, { from: accounts[2], value: 1e15 });
    //     const auction = await this.engAuction.auctions(aucID);
    //     assert.equal(1e15, auction.highestBid);
    //     assert.equal(accounts[2], auction.highestBidder);
    // });
});
