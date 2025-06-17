const FeralfileExhibitionV2 = artifacts.require("FeralfileExhibitionV2");

const IPFS_GATEWAY_PREFIX = "https://cloudflare-ipfs.com/ipfs/";

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";

const originArtworkCID = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

contract("FeralfileExhibitionV2", async (accounts) => {
    before(async function () {
        this.exhibition = await FeralfileExhibitionV2.new(
            "Feral File V2 Test 001",
            "FFV2",
            50,
            1000,
            "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
            "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
            IPFS_GATEWAY_PREFIX
        );
    });

    it("create an artwork correctly with no editions on it", async function () {
        let r = await this.exhibition.createArtwork(
            "TestArtwork",
            "TestArtwork",
            "TestUser",
            5
        );
        this.artworkID = r.logs[0].args.artworkID;

        let artwork = await this.exhibition.artworks(this.artworkID);
        assert.equal(artwork.title, "TestArtwork");

        let artworkEditions = await this.exhibition.totalEditionOfArtwork(
            this.artworkID
        );
        assert.equal(artworkEditions, 0);
    });

    it("fail to create an artwork if the edition size is greater than exhibition limits", async function () {
        try {
            let r = await this.exhibition.createArtwork(
                "TestArtworkFail",
                "TestArtworkFail",
                "TestUser",
                100
            );
        } catch (error) {
            assert.include(
                error.message,
                "artwork edition size exceeds the maximum edition size of the exhibition"
            );
        }
    });

    it("fail to create an artwork with duplicated fingerprint", async function () {
        try {
            let r = await this.exhibition.createArtwork(
                "TestArtwork",
                "TestArtwork",
                "TestUser",
                5
            );
        } catch (error) {
            assert.include(
                error.message,
                "an artwork with the same fingerprint has already registered"
            );
        }
    });

    it("fail to create an artwork with empty fingerprint", async function () {
        try {
            let r = await this.exhibition.createArtwork(
                "",
                "TestArtwork",
                "TestUser",
                5
            );
        } catch (error) {
            assert.include(error.message, "fingerprint can not be empty");
        }
    });

    it("fail to create an artwork with empty title", async function () {
        try {
            let r = await this.exhibition.createArtwork(
                "TestArtwork",
                "",
                "TestUser",
                5
            );
        } catch (error) {
            assert.include(error.message, "title can not be empty");
        }
    });

    it("fail to create an artwork with empty artist name", async function () {
        try {
            let r = await this.exhibition.createArtwork(
                "TestArtwork",
                "TestArtwork",
                "",
                5
            );
        } catch (error) {
            assert.include(error.message, "artist can not be empty");
        }
    });

    it("can swap tokens from bitmark blockchain correctly", async function () {
        let tokenOwner = "0xbbed1c61b6e68c397b021c1274080a2005042c08";

        let bitmarkID = Buffer.from(
            "793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e382",
            "hex"
        );

        let testEditionNumber = 3;

        await this.exhibition.swapArtworkFromBitmark(
            this.artworkID,
            bitmarkID,
            testEditionNumber,
            tokenOwner,
            originArtworkCID
        );

        let editionID = BigInt(this.artworkID) + BigInt(testEditionNumber);
        let artworkEdition = await this.exhibition.artworkEditions(editionID);

        assert.equal(artworkEdition.editionID, editionID);
        assert.equal(artworkEdition.ipfsCID, originArtworkCID);
    });

    it("can update the IPFS to a new one", async function () {
        let newCID = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXcNew";

        let testEditionNumber = 3;
        let editionID = BigInt(this.artworkID) + BigInt(testEditionNumber);
        await this.exhibition.updateArtworkEditionIPFSCid(editionID, newCID);

        let artworkEdition = await this.exhibition.artworkEditions(editionID);

        assert.equal(artworkEdition.ipfsCID, newCID);
    });

    it("cannot update the IPFS cid using an existent IPFS cid", async function () {
        let testEditionNumber = 3;
        let editionID = BigInt(this.artworkID) + BigInt(testEditionNumber);
        try {
            await this.exhibition.updateArtworkEditionIPFSCid(
                editionID,
                originArtworkCID
            );
        } catch (error) {
            assert.equal(error.message, "error herer");
        }
    });

    it("cannot update the IPFS cid for a non-existent token", async function () {
        let newCID = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXcNew";

        let editionID = 1;
        try {
            await this.exhibition.updateArtworkEditionIPFSCid(
                editionID,
                newCID
            );
        } catch (error) {
            assert.include(error.message, "artwork edition is not found");
        }
    });

    it("can not set royalty payout address to zero address", async function () {
        try {
            await this.exhibition.setRoyaltyPayoutAddress(ZERO_ADDRESS);
        } catch (error) {
            assert.include(error.message, "invalid royalty payout address");
        }
    });

    it("return the correct amount of royalty corresponded to the percentage we set", async function () {
        let testEditionNumber = 3;
        let editionID = BigInt(this.artworkID) + BigInt(testEditionNumber);
        let info = await this.exhibition.royaltyInfo(editionID, 1000000000);

        assert.equal(Number(info.royaltyAmount), 100000000);
    });
});
