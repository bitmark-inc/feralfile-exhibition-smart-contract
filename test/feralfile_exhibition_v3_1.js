const FeralfileExhibitionV3_1 = artifacts.require("FeralfileExhibitionV3_1");
const MockDecentraland = artifacts.require("MockDecentraland");

const axios = require("axios");

const IPFS_GATEWAY_PREFIX = "https://ipfs.bitmark.com/ipfs/";

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";

const originArtworkCID = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

const Base64 = {
    encode: (str) => Buffer.from(str).toString("base64"),
};

contract("FeralfileExhibitionV3_1", async (accounts) => {
    before(async function () {
        this.decentraland = await MockDecentraland.new();
        this.testThumbnailCid =
            "QmV68mphFwMraCE9J6KpQc89Sz8ppvJx5CP6XFruhGQrX8"; // AKG test IPFS thumbnail CID
        this.testArtworkCid = "QmTn3PfHHvoDHKawTPXutqxAk2k8ynFK9cZfsSwggryjkX"; // AKG test IPFS artwork CID

        this.exhibition = await FeralfileExhibitionV3_1.new(
            "Feral File V3 Test 002",
            "FFV3",
            1000,
            "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
            "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
            IPFS_GATEWAY_PREFIX,
            true,
            true
        );

        this.exhibition.updateDecentralandInfo(this.decentraland.address, 0);
    });

    it("create an artwork with an external IPFS link", async function () {
        let r = await this.exhibition.createArtworkWithIPFSCID(
            ["TestArtwork", "TestUser", "TestArtwork-fingerprint", 10, 2, 1],
            this.testThumbnailCid,
            this.testArtworkCid
        );
        this.artworkID = r.logs[0].args.artworkID;

        let artwork = await this.exhibition.artworks(this.artworkID);
        assert.equal(artwork.title, "TestArtwork");

        let artworkEditions = await this.exhibition.totalEditionOfArtwork(
            this.artworkID
        );
        assert.equal(artworkEditions, 0);

        let artworkCIDInContract = await this.exhibition.externalArtworkIPFSCID(
            this.artworkID
        );
        assert.equal(this.testThumbnailCid, artworkCIDInContract.thumbnailCID);
        assert.equal(this.testArtworkCid, artworkCIDInContract.artworkCID);
    });

    it("test batch mint 4 artworks", async function () {
        let editionIndex0 = 0;
        let editionIndex1 = 1;
        let editionIndex2 = 2;
        let editionIndex3 = 3;
        let editionIndex4 = 4;

        await this.exhibition.batchMint([
            [this.artworkID, editionIndex0, accounts[0], accounts[0], "test0"],
            [this.artworkID, editionIndex1, accounts[0], accounts[1], "test1"],
            [this.artworkID, editionIndex2, accounts[0], accounts[2], "test2"],
            [this.artworkID, editionIndex3, accounts[0], accounts[3], "test3"],
            [this.artworkID, editionIndex4, accounts[0], accounts[4], "test4"],
        ]);

        let editionID0 = BigInt(this.artworkID) + BigInt(editionIndex0);
        let ownerOfEdition0 = await this.exhibition.ownerOf(
            editionID0.toString()
        );
        assert.equal(ownerOfEdition0.toLowerCase(), accounts[0].toLowerCase());
        let editionNumber0 = await this.exhibition.tokenEditionNumber(
            editionID0.toString()
        );
        assert.equal(editionNumber0, "AE");

        let editionID1 = BigInt(this.artworkID) + BigInt(editionIndex1);
        let ownerOfEdition1 = await this.exhibition.ownerOf(
            editionID1.toString()
        );
        assert.equal(ownerOfEdition1.toLowerCase(), accounts[1].toLowerCase());
        let editionNumber1 = await this.exhibition.tokenEditionNumber(
            editionID1.toString()
        );
        assert.equal(editionNumber1, "AE");

        let editionID2 = BigInt(this.artworkID) + BigInt(editionIndex2);
        let ownerOfEdition2 = await this.exhibition.ownerOf(
            editionID2.toString()
        );
        assert.equal(ownerOfEdition2.toLowerCase(), accounts[2].toLowerCase());
        let editionNumber2 = await this.exhibition.tokenEditionNumber(
            editionID2.toString()
        );
        assert.equal(editionNumber2, "PP");

        let editionID3 = BigInt(this.artworkID) + BigInt(editionIndex3);
        let ownerOfEdition3 = await this.exhibition.ownerOf(
            editionID3.toString()
        );
        assert.equal(ownerOfEdition3.toLowerCase(), accounts[3].toLowerCase());
        let editionNumber3 = await this.exhibition.tokenEditionNumber(
            editionID3.toString()
        );
        assert.equal(editionNumber3, "#1");

        let editionID4 = BigInt(this.artworkID) + BigInt(editionIndex4);
        let ownerOfEdition4 = await this.exhibition.ownerOf(
            editionID4.toString()
        );
        assert.equal(ownerOfEdition4.toLowerCase(), accounts[4].toLowerCase());
        let editionNumber4 = await this.exhibition.tokenEditionNumber(
            editionID4.toString()
        );
        assert.equal(editionNumber4, "#2");
    });

    it("test build metadata for artwork with IPFS", async function () {
        let editionNumber1 = 1;
        let expectedAnimationUrl = `<!DOCTYPE html><html lang="en"><head><script> var defaultArtworkData= {landOwner:"0x0177ba2ad05cb12d6dd637963e040092676ee8de", ownerArray:["${accounts[0].toLowerCase()}","${accounts[1].toLowerCase()}","${accounts[2].toLowerCase()}","${accounts[3].toLowerCase()}","${accounts[4].toLowerCase()}",]}</script><script>let allowOrigins={"https://feralfile.com":!0};function resizeIframe(t){let e=document.getElementById("mainframe");t&&(e.style.minHeight=t+"px")}function initData(){pushArtworkDataToIframe(defaultArtworkData)}function pushArtworkDataToIframe(t){t&&document.getElementById("mainframe").contentWindow.postMessage(t,"*")}function updateArtowrkData(t){document.getElementById("mainframe").contentWindow.postMessage(t,"*")}window.addEventListener("message",function(t){allowOrigins[t.origin]?"update-artwork-data"===t.data.type&&updateArtowrkData(t.data.artworkData):"object"==typeof t.data&&"resize-iframe"===t.data.type&&resizeIframe(t.data.newHeight)});</script></head><body style="overflow-x:hidden;padding:0;margin:0;width: 100%;" onload="initData()"><iframe id="mainframe" style="display:block;padding:0;margin:0;border:none;width:100%;height:100vh;" src="https://ipfs.bitmark.com/ipfs/QmTn3PfHHvoDHKawTPXutqxAk2k8ynFK9cZfsSwggryjkX"></iframe> </body></html>`;
        let expectedAnimationUrlBase64 = Base64.encode(expectedAnimationUrl);

        let tokenURI = await this.exhibition.tokenURI(
            BigInt(this.artworkID) + BigInt(editionNumber1)
        );
        assert.equal(
            tokenURI.startsWith("data:application/json;base64,"),
            true
        );

        const metadataJSON = JSON.parse(
            Buffer.from(
                tokenURI.replace("data:application/json;base64,", ""),
                "base64"
            ).toString()
        );
        assert.equal(
            metadataJSON.image,
            "https://ipfs.bitmark.com/ipfs/" + this.testThumbnailCid
        );
        // TODO: use regular expression to validate the html data
        assert.equal(
            metadataJSON.animation_url,
            `data:text/html;base64,${expectedAnimationUrlBase64}`
        );
    });

    it("test update artwork external file CIDs", async function () {
        let newCid = "QmV68mphFwMraCE9J6KpQc89Sz8ppvJx5CP6XFruhGQrX8";

        await this.exhibition.updateExternalArtworkIPFSCID(
            this.artworkID,
            newCid,
            newCid
        );
        let artworkCIDInContract = await this.exhibition.externalArtworkIPFSCID(
            this.artworkID
        );
        assert.equal(newCid, artworkCIDInContract.thumbnailCID);
        assert.equal(newCid, artworkCIDInContract.artworkCID);
    });

    it("test remove artwork external file CIDs by using with empty CID", async function () {
        let newCid = "";
        let editionNumber1 = 1;

        await this.exhibition.updateExternalArtworkIPFSCID(
            this.artworkID,
            newCid,
            newCid
        );
        let artworkCIDInContract = await this.exhibition.externalArtworkIPFSCID(
            this.artworkID
        );
        assert.equal(newCid, artworkCIDInContract.thumbnailCID);
        assert.equal(newCid, artworkCIDInContract.artworkCID);

        let tokenURI = await this.exhibition.tokenURI(
            BigInt(this.artworkID) + BigInt(editionNumber1)
        );
        assert.equal(
            tokenURI.startsWith("https://ipfs.bitmark.com/ipfs/"),
            true
        );
    });

    it("test update artwork external file CIDs with non exist artwork ID", async function () {
        let newCid = "";
        let error;

        try {
            await this.exhibition.updateExternalArtworkIPFSCID(
                "0x0123",
                newCid,
                newCid
            );
        } catch (e) {
            error = e;
        }

        assert.include(error.message, "the target artwork is not existent");
    });
});
