const FeralfileExhibitionV3_3 = artifacts.require("FeralfileExhibitionV3_3");

const IPFS_GATEWAY_PREFIX = "https://ipfs.bitmark.com/ipfs/";

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";

contract("FeralfileExhibitionV3_3", async (accounts) => {
    before(async function () {
        this.exhibition = await FeralfileExhibitionV3_3.new(
            "Feral File V3_4 Test 001",
            "FFV3",
            1000,
            "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
            "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
            IPFS_GATEWAY_PREFIX,
            true,
            true,
            accounts[1] // Set minter to accounts[1]
        );
    });

    it("can register artwork", async function () {
        let r = await this.exhibition.createArtworks([
            ["TestArtwork", "TestUser", "TestArtwork", 10, 2, 1],
        ]);
        this.artworkID = r.logs[0].args.artworkID;

        let artwork = await this.exhibition.artworks(this.artworkID);
        assert.equal(artwork.title, "TestArtwork");

        let artworkEditions = await this.exhibition.totalEditionOfArtwork(
            this.artworkID
        );
        assert.equal(artworkEditions, 0);
    });

    it("test mint edition from burn successfully", async function () {
        let artwork = await this.exhibition.artworks(this.artworkID);
        await this.exhibition.updateArtworkCIDs(
            [this.artworkID],
            ["QmTestCID"]
        );
        assert.equal(artwork.title, "TestArtwork");
        const r = await this.exhibition.mintArtworkEdition(
            this.artworkID,
            accounts[0],
            { from: accounts[1] }
        );
        const editionID = r.logs[1].args.editionID;
        const owner = await this.exhibition.ownerOf(editionID);
        assert.equal(accounts[0], owner);
    });

    it("test counter with mint edition multiple editions", async function () {
        const exh = await FeralfileExhibitionV3_3.new(
            "Feral File V3_4 Test 001",
            "FFV3",
            1000,
            "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
            "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
            IPFS_GATEWAY_PREFIX,
            true,
            true,
            accounts[2]
        );
        const ar = await exh.createArtworks([
            ["TestMultiple", "TestU", "TestMultiple", 10, 2, 1],
        ]);

        const artworkID = ar.logs[0].args.artworkID;
        const artwork = await exh.artworks(artworkID);
        assert.equal(artwork.title, "TestMultiple");

        await exh.updateArtworkCIDs([artworkID], ["QmTestCID"]);
        const mr1 = await exh.mintArtworkEdition(artworkID, accounts[0], {
            from: accounts[2],
        });
        const editionID1 = mr1.logs[1].args.editionID;
        const owner1 = await exh.ownerOf(editionID1);
        assert.equal(accounts[0], owner1);

        const mr2 = await exh.mintArtworkEdition(artworkID, accounts[1], {
            from: accounts[2],
        });
        const editionID2 = mr2.logs[1].args.editionID;
        const owner2 = await exh.ownerOf(editionID2);
        assert.equal(accounts[1], owner2);

        const mr3 = await exh.mintArtworkEdition(artworkID, accounts[0], {
            from: accounts[2],
        });
        const editionID3 = mr3.logs[1].args.editionID;
        const owner3 = await exh.ownerOf(editionID3);
        assert.equal(accounts[0], owner3);

        assert.equal(BigInt(0), BigInt(editionID1) - BigInt(artworkID));
        assert.equal(BigInt(1), BigInt(editionID2) - BigInt(artworkID));
        assert.equal(BigInt(2), BigInt(editionID3) - BigInt(artworkID));
    });

    it("test mint fail because wrong minter", async function () {
        const exh = await FeralfileExhibitionV3_3.new(
            "Feral File V3_4 Test 001",
            "FFV3",
            1000,
            "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
            "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
            IPFS_GATEWAY_PREFIX,
            true,
            true,
            accounts[1] // Set minter to accounts[1]
        );
        const ar = await exh.createArtworks([
            ["TestMultiple", "TestU", "TestMultiple", 10, 2, 1],
        ]);

        const artworkID = ar.logs[0].args.artworkID;
        const artwork = await exh.artworks(artworkID);
        assert.equal(artwork.title, "TestMultiple");

        await exh.updateArtworkCIDs([artworkID], ["QmTestCID"]);
        try {
            await exh.mintArtworkEdition(artworkID, accounts[0], {
                from: accounts[0],
            });
        } catch (error) {
            assert.include(error.message, "Only minter can call this function");
        }
    });

    it("test mint edition from burn fail because no artwork", async function () {
        const exh = await FeralfileExhibitionV3_3.new(
            "Feral File V3_4 Test 001",
            "FFV3",
            1000,
            "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
            "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
            IPFS_GATEWAY_PREFIX,
            true,
            true,
            accounts[1]
        );
        try {
            await exh.mintArtworkEdition(0, accounts[0], { from: accounts[1] });
        } catch (error) {
            assert.include(error.message, "Artwork does not exist");
        }
    });

    it("test mint edition fail because no artwork ipfs cid", async function () {
        const exh = await FeralfileExhibitionV3_3.new(
            "Feral File V3_4 Test 001",
            "FFV3",
            1000,
            "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
            "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
            IPFS_GATEWAY_PREFIX,
            true,
            true,
            accounts[1]
        );
        const ar = await exh.createArtworks([
            ["TestWithoutCID", "TestU", "TestWithoutCID", 10, 2, 1],
        ]);
        const artworkID = ar.logs[0].args.artworkID;
        const artwork = await exh.artworks(artworkID);
        assert.equal(artwork.title, "TestWithoutCID");
        try {
            await exh.mintArtworkEdition(artworkID, accounts[0], {
                from: accounts[1],
            });
        } catch (error) {
            assert.include(error.message, "Artwork CID does not exist");
        }
    });

    it("test get artwork id by token id successfully", async function () {
        const r = await this.exhibition.mintArtworkEdition(
            this.artworkID,
            accounts[0],
            { from: accounts[1] }
        );
        const editionID = r.logs[1].args.editionID;
        const artworkID =
            await this.exhibition.getArtworkIDByTokenID(editionID);
        assert.equal(this.artworkID.toString(), artworkID.toString());
    });

    it("test get artwork id by token id failed", async function () {
        const artworkID = await this.exhibition.getArtworkIDByTokenID(0);
        assert.equal("0", artworkID.toString());
    });

    it("test update artwork CIDs successfully by owner", async function () {
        await this.exhibition.updateArtworkCIDs(
            [this.artworkID],
            ["QmTestCID"]
        );

        const r = await this.exhibition.mintArtworkEdition(
            this.artworkID,
            accounts[0],
            { from: accounts[1] }
        );
        const editionID = r.logs[1].args.editionID;

        const owner = await this.exhibition.ownerOf(editionID);
        assert.equal(accounts[0], owner);

        const tokenURI = await this.exhibition.tokenURI(editionID);

        assert.equal(true, tokenURI.indexOf("QmTestCID") >= 0);
    });

    it("test update artwork CIDs successfully by trustee", async function () {
        await this.exhibition.addTrustee(accounts[2]);
        await this.exhibition.updateArtworkCIDs(
            [this.artworkID],
            ["QmTestCID"],
            {
                from: accounts[2],
            }
        );

        const r = await this.exhibition.mintArtworkEdition(
            this.artworkID,
            accounts[0],
            { from: accounts[1] }
        );
        const editionID = r.logs[1].args.editionID;

        const owner = await this.exhibition.ownerOf(editionID);
        assert.equal(accounts[0], owner);

        const tokenURI = await this.exhibition.tokenURI(editionID);

        assert.equal(true, tokenURI.indexOf("QmTestCID") >= 0);
    });

    it("test update artwork CIDs fail because invalid trustee", async function () {
        try {
            await this.exhibition.updateArtworkCIDs(
                [this.artworkID],
                ["QmTestCID"],
                {
                    from: accounts[3],
                }
            );
            assert.fail("Should throw error");
        } catch (error) {}
    });

    it("test get token uri from minted token", async function () {
        await this.exhibition.updateArtworkCIDs(
            [this.artworkID],
            ["QmTestCIDMinting"]
        );
        const r = await this.exhibition.mintArtworkEdition(
            this.artworkID,
            accounts[0],
            { from: accounts[1] }
        );
        const editionID = r.logs[1].args.editionID;
        const tokenURI = await this.exhibition.tokenURI(editionID);
        const expected =
            IPFS_GATEWAY_PREFIX + "QmTestCIDMinting/" + editionID.toString();
        assert.equal(expected, tokenURI);
    });

    it("test get token uri with not existed token", async function () {
        try {
            await this.exhibition.tokenURI(BigInt("121121235353535353535"));
        } catch (err) {
            assert.include(err.message, "URI query for nonexistent token");
        }
    });
});
