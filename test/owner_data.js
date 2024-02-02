const OwnerData = artifacts.require("OwnerData");
const FeralfileExhibitionV4 = artifacts.require("FeralfileExhibitionV4");
const FeralfileVault = artifacts.require("FeralfileVault");

const CONTRACT_URI = "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

const bytesToString = (bytes) => {
    return web3.utils.toAscii(bytes).replace(/\u0000/g, "");
};

contract("OwnerData", async (accounts) => {
    before(async function () {
        this.signer = accounts[0];
        this.vault = await FeralfileVault.new(this.signer);
        this.ownerDataContract = await OwnerData.new();
        this.exhibitionContract = await FeralfileExhibitionV4.new(
            "Feral File V4 Test",
            "FFv4",
            true,
            true,
            this.signer,
            this.vault.address,
            this.signer,
            CONTRACT_URI,
            [1],
            [1000]
        );

        await this.exhibitionContract.mintArtworks([
            [1, 1, accounts[0]],
            [1, 2, accounts[1]],
            [1, 3, accounts[2]],
            [1, 4, accounts[0]],
            [1, 5, accounts[1]],
            [1, 6, accounts[2]],
        ]);
    });

    it("test adding data", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const tx = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            1,
            [accounts[0], cidBytes, "{duration: 1000}"]
        );
        const { logs } = tx;
        assert.equal(logs[0].event, "DataAdded");

        assert.equal(bytesToString(logs[0].args.data.dataHash), cid);
    });

    it("test getting data", async function () {
        const data = await this.ownerDataContract.get(
            this.exhibitionContract.address,
            1
        );
        assert.equal(data[0].owner, accounts[0]);
        assert.equal(
            bytesToString(data[0].dataHash),
            "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc"
        );
        assert.equal(data[0].metadata, "{duration: 1000}");
    });

    it("test remove data successfully", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const tx = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            1,
            [accounts[0], cidBytes, "{duration: 1000}"]
        );
        assert.equal(tx.logs[0].event, "DataAdded");

        const tx2 = await this.ownerDataContract.remove(
            this.exhibitionContract.address,
            1
        );
        assert.equal(tx2.logs[0].event, "DataRemoved");

        const data = await this.ownerDataContract.get(
            this.exhibitionContract.address,
            1
        );
        assert.equal(data.length, 0);
    });

    it("test remove data successfully in list data", async function () {
        const cid1 = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cid2 = "QmNVdQSp1AvZonLwHzTbbZDPLgbpty15RMQrbPEWd4ooTU";
        const cidBytes1 = web3.utils.fromAscii(cid1);
        const cidBytes2 = web3.utils.fromAscii(cid2);

        const tx1 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            4,
            [accounts[0], cidBytes1, "{duration: 1000}"],
            { from: accounts[0] }
        );
        assert.equal(bytesToString(tx1.logs[0].args.data.dataHash), cid1);

        // Transfer to account 2
        await this.exhibitionContract.transferFrom(
            accounts[0],
            accounts[2],
            4,
            { from: accounts[0] }
        );
        const owner = await this.exhibitionContract.ownerOf(4);
        assert.equal(owner, accounts[2]);

        const tx2 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            4,
            [accounts[2], cidBytes2, "{duration: 2000}"],
            { from: accounts[2] }
        );
        assert.equal(bytesToString(tx2.logs[0].args.data.dataHash), cid2);

        const data = await this.ownerDataContract.get(
            this.exhibitionContract.address,
            4
        );

        assert.equal(data.length, 2);
        assert.equal(data[0].owner, accounts[0]);
        assert.equal(bytesToString(data[0].dataHash), cid1);

        assert.equal(data[1].owner, accounts[2]);
        assert.equal(bytesToString(data[1].dataHash), cid2);

        const tx3 = await this.ownerDataContract.remove(
            this.exhibitionContract.address,
            4,
            { from: accounts[2] }
        );
        assert.equal(tx3.logs[0].event, "DataRemoved");

        const data2 = await this.ownerDataContract.get(
            this.exhibitionContract.address,
            4
        );
        assert.equal(data2.length, 1);
        assert.equal(data2[0].owner, accounts[0]);
    });

    it("test updating data", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const tx = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            3,
            [accounts[2], cidBytes, "{duration: 1000}"],
            { from: accounts[2] }
        );
        assert.equal(bytesToString(tx.logs[0].args.data.dataHash), cid);

        const updatedCid = "QmNVdQSp1AvZonLwHzTbbZDPLgbpty15RMQrbPEWd4ooTU";
        const updatedCidBytes = web3.utils.fromAscii(updatedCid);
        const tx2 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            3,
            [accounts[2], updatedCidBytes, "{duration: 2000}"],
            { from: accounts[2] }
        );
        assert.equal(bytesToString(tx2.logs[0].args.data.dataHash), updatedCid);

        const data = await this.ownerDataContract.get(
            this.exhibitionContract.address,
            3
        );
        assert.equal(data[0].owner, accounts[2]);
        assert.equal(bytesToString(data[0].dataHash), updatedCid);
    });

    it("test adding multiple data", async function () {
        const cid1 = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cid2 = "QmNVdQSp1AvZonLwHzTbbZDPLgbpty15RMQrbPEWd4ooTU";
        const cid3 = "QmR7jnM6jgoAE6VZf6uiRRF3e6JULgJ6GYMh1kRPRtr3CS";
        const cidBytes1 = web3.utils.fromAscii(cid1);
        const cidBytes2 = web3.utils.fromAscii(cid2);
        const cidBytes3 = web3.utils.fromAscii(cid3);

        const tx1 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            2,
            [accounts[1], cidBytes1, "{duration: 1000}"],
            { from: accounts[1] }
        );
        assert.equal(bytesToString(tx1.logs[0].args.data.dataHash), cid1);

        // Transfer to account 3
        await this.exhibitionContract.transferFrom(
            accounts[1],
            accounts[2],
            2,
            { from: accounts[1] }
        );
        const owner = await this.exhibitionContract.ownerOf(2);
        assert.equal(owner, accounts[2]);

        const tx2 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            2,
            [accounts[2], cidBytes2, "{duration: 2000}"],
            { from: accounts[2] }
        );
        assert.equal(bytesToString(tx2.logs[0].args.data.dataHash), cid2);

        // Transfer to account 4
        await this.exhibitionContract.transferFrom(
            accounts[2],
            accounts[4],
            2,
            { from: accounts[2] }
        );
        const owner2 = await this.exhibitionContract.ownerOf(2);
        assert.equal(owner2, accounts[4]);

        const tx3 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            2,
            [accounts[4], cidBytes3, "{duration: 3000}"],
            { from: accounts[4] }
        );
        assert.equal(bytesToString(tx3.logs[0].args.data.dataHash), cid3);

        const data = await this.ownerDataContract.get(
            this.exhibitionContract.address,
            2
        );

        assert.equal(data.length, 3);
        assert.equal(data[0].owner, accounts[1]);
        assert.equal(bytesToString(data[0].dataHash), cid1);

        assert.equal(data[1].owner, accounts[2]);
        assert.equal(bytesToString(data[1].dataHash), cid2);

        assert.equal(data[2].owner, accounts[4]);
        assert.equal(bytesToString(data[2].dataHash), cid3);
    });

    it("test adding fail with wrong owner address", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        try {
            await this.ownerDataContract.add(
                this.exhibitionContract.address,
                1,
                [accounts[1], cidBytes, "{duration: 1000}"]
            );
        } catch (error) {
            assert.equal(error.reason, "OwnerData: owner mismatch");
        }
    });

    it("test adding fail with invalid contract address", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        try {
            await this.ownerDataContract.add(accounts[1], 1, [
                accounts[0],
                cidBytes,
                "{duration: 1000}",
            ]);
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "VM Exception while processing transaction: revert"
                )
            );
        }
    });
});
