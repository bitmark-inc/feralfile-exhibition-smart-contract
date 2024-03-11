const OwnerData = artifacts.require("OwnerData");
const FeralfileExhibitionV4 = artifacts.require("FeralfileExhibitionV4");
const FeralFileAirdropV1 = artifacts.require("FeralFileAirdropV1");
const FeralfileVault = artifacts.require("FeralfileVault");

const CONTRACT_URI = "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
const COST_RECEIVER = "0x46f2B641d8702f29c45f6D06292dC34Eb9dB1801";
const TOKEN_URI =
    "https://ipfs.bitmark.com/ipfs/QmNVdQSp1AvZonLwHzTbbZDPLgbpty15RMQrbPEWd4ooTU/{id}";
const TOKEN_TYPE_FUNGIBLE = 0;

const { bufferToHex } = require("ethereumjs-util");

const bytesToString = (bytes) => {
    return web3.utils.toAscii(bytes).replace(/\u0000/g, "");
};

contract("OwnerData", async (accounts) => {
    before(async function () {
        this.signer = accounts[0];
        this.trustee = accounts[1];
        this.ownerDataSigner = accounts[2];
        this.vault = await FeralfileVault.new(this.signer);
        this.ownerDataContract = await OwnerData.new(
            this.ownerDataSigner,
            COST_RECEIVER,
            web3.utils.toWei("0.015", "ether"),
            200001
        );
        this.seriesIds = [1, 2, 3, 4, 5];
        this.seriesMaxSupply = [100, 1, 1, 1, 1];
        this.exhibitionContract = await FeralfileExhibitionV4.new(
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

        await this.exhibitionContract.mintArtworks([
            [this.seriesIds[0], 100001, accounts[0]],
            [this.seriesIds[0], 100002, accounts[1]],
            [this.seriesIds[0], 100003, accounts[2]],
            [this.seriesIds[0], 100004, accounts[0]],
            [this.seriesIds[0], 100005, accounts[1]],
            [this.seriesIds[0], 100006, accounts[2]],
            [
                this.seriesIds[0],
                100007,
                "0x23221e5403511CeC833294D2B1B006e9D639A61b",
            ],
            [this.seriesIds[1], 200001, accounts[0]],
            [this.seriesIds[2], 300001, accounts[0]],
            [this.seriesIds[3], 400001, accounts[0]],
            [this.seriesIds[4], 500001, accounts[0]],
        ]);

        await this.exhibitionContract.addTrustee(this.trustee);

        this.fungibleContract = await FeralFileAirdropV1.new(
            TOKEN_TYPE_FUNGIBLE,
            TOKEN_URI,
            CONTRACT_URI,
            true,
            true
        );
        await this.fungibleContract.mint(999, 10);
        await this.fungibleContract.airdrop(999, [accounts[0], accounts[1]]);
    });

    it("test adding data successfully", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const tx = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            100001,
            [accounts[0], cidBytes, "{duration: 1000}"],
            { from: accounts[0], value: 0 }
        );
        const { logs } = tx;
        assert.equal(logs[0].event, "DataAdded");

        assert.equal(bytesToString(logs[0].args.data.dataHash), cid);
    });

    it("test getting data", async function () {
        const data = await this.ownerDataContract.get(
            this.exhibitionContract.address,
            100001,
            0,
            100
        );
        assert.equal(data[0].owner, accounts[0]);
        assert.equal(
            bytesToString(data[0].dataHash),
            "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc"
        );
        assert.equal(data[0].metadata, "{duration: 1000}");
    });

    it("test adding data failed because already added", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const tx = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            100003,
            [accounts[2], cidBytes, "{duration: 1000}"],
            { from: accounts[2], value: 0 }
        );
        assert.equal(bytesToString(tx.logs[0].args.data.dataHash), cid);

        const updatedCid = "QmNVdQSp1AvZonLwHzTbbZDPLgbpty15RMQrbPEWd4ooTU";
        const updatedCidBytes = web3.utils.fromAscii(updatedCid);
        try {
            await this.ownerDataContract.add(
                this.exhibitionContract.address,
                100003,
                [accounts[2], updatedCidBytes, "{duration: 2000}"],
                { from: accounts[2], value: 0 }
            );
        } catch (error) {
            assert.equal(error.reason, "Custom error (could not decode)");
        }
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
            100002,
            [accounts[1], cidBytes1, "{duration: 1000}"],
            { from: accounts[1], value: 0 }
        );
        assert.equal(bytesToString(tx1.logs[0].args.data.dataHash), cid1);

        // Transfer to account 3
        await this.exhibitionContract.transferFrom(
            accounts[1],
            accounts[2],
            100002,
            { from: accounts[1] }
        );
        const acc2Owner = await this.exhibitionContract.ownerOf(100002);
        assert.equal(acc2Owner, accounts[2]);

        const tx2 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            100002,
            [accounts[2], cidBytes2, "{duration: 2000}"],
            { from: accounts[2], value: 0 }
        );
        assert.equal(bytesToString(tx2.logs[0].args.data.dataHash), cid2);

        // Transfer to account 4
        await this.exhibitionContract.transferFrom(
            accounts[2],
            accounts[4],
            100002,
            { from: accounts[2] }
        );
        const acc4Owner = await this.exhibitionContract.ownerOf(100002);
        assert.equal(acc4Owner, accounts[4]);

        const tx3 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            100002,
            [accounts[4], cidBytes3, "{duration: 3000}"],
            { from: accounts[4], value: 0 }
        );
        assert.equal(bytesToString(tx3.logs[0].args.data.dataHash), cid3);

        const data = await this.ownerDataContract.get(
            this.exhibitionContract.address,
            100002,
            0,
            100
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
                100001,
                [accounts[1], cidBytes, "{duration: 1000}"],
                { from: accounts[0], value: 0 }
            );
        } catch (error) {
            assert.equal(error.reason, "Custom error (could not decode)");
        }
    });

    it("test adding fail with invalid contract address", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        try {
            await this.ownerDataContract.add(
                accounts[1],
                100001,
                [accounts[0], cidBytes, "{duration: 1000}"],
                { from: accounts[0], value: 0 }
            );
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "VM Exception while processing transaction: revert"
                )
            );
        }
    });

    it("test removing data successfully", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const ownerDataContract = await OwnerData.new(
            this.ownerDataSigner,
            COST_RECEIVER,
            web3.utils.toWei("0.015", "ether"),
            300001
        );
        await ownerDataContract.add(
            this.exhibitionContract.address,
            300001,
            [accounts[0], cidBytes, "{duration: 1000}"],
            { from: accounts[0], value: web3.utils.toWei("0.015", "ether") }
        );
        await ownerDataContract.add(
            this.exhibitionContract.address,
            300001,
            [accounts[1], cidBytes, "{duration: 1000}"],
            { from: accounts[1], value: web3.utils.toWei("0.015", "ether") }
        );
        await ownerDataContract.add(
            this.exhibitionContract.address,
            300001,
            [accounts[2], cidBytes, "{duration: 1000}"],
            { from: accounts[2], value: web3.utils.toWei("0.015", "ether") }
        );
        await ownerDataContract.add(
            this.exhibitionContract.address,
            300001,
            [accounts[1], cidBytes, "{duration: 1000}"],
            { from: accounts[1], value: web3.utils.toWei("0.015", "ether") }
        );
        await ownerDataContract.add(
            this.exhibitionContract.address,
            300001,
            [accounts[0], cidBytes, "{duration: 1000}"],
            { from: accounts[0], value: web3.utils.toWei("0.015", "ether") }
        );
        const res = await ownerDataContract.get(
            this.exhibitionContract.address,
            300001,
            0,
            100
        );
        assert.equal(res.length, 5);
        const tx = await ownerDataContract.remove(
            this.exhibitionContract.address,
            300001,
            [4, 2],
            { from: accounts[0] }
        );

        const res2 = await ownerDataContract.get(
            this.exhibitionContract.address,
            300001,
            0,
            100
        );
        assert.equal(res2.length, 5);
        assert.equal(bytesToString(res2[0].dataHash), cid);
        assert.equal(bytesToString(res2[1].dataHash), cid);
        assert.equal(bytesToString(res2[2].dataHash), "");
        assert.equal(bytesToString(res2[3].dataHash), cid);
        assert.equal(bytesToString(res2[4].dataHash), "");
    });

    it("test adding with signed add function", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const data = [
            "0x23221e5403511CeC833294D2B1B006e9D639A61b",
            cidBytes,
            "{duration: 1000}",
        ];

        const msg = `Authorize to write your data to the contract ${this.ownerDataContract.address.toLowerCase()}.`;
        const msgHash = bufferToHex(Buffer.from(msg, "utf-8"));
        const privateKey =
            "0x5cd8bcda59dd3a9988bd20bdbdea7225a4a57949d12b9a527caf3ff819941d7f";
        const { signature } = await web3.eth.accounts.sign(msgHash, privateKey);

        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const chainId = await web3.eth.getChainId();
        const signedParams = web3.eth.abi.encodeParameters(
            ["uint", "address", "bytes", "uint256"],
            [
                BigInt(chainId).toString(),
                this.ownerDataContract.address,
                signature,
                expiryTime,
            ]
        );

        const hash = web3.utils.keccak256(signedParams);
        const trusteeSignature = await web3.eth.sign(
            hash,
            this.ownerDataSigner
        );
        const sig = trusteeSignature.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);

        // sign params
        const signs = [
            signature,
            expiryTime,
            r,
            s,
            web3.utils.toDecimal(v) + 27,
        ];

        const tx = await this.ownerDataContract.signedAdd(
            this.exhibitionContract.address,
            100007,
            data,
            signs
        );
        assert.equal(tx.logs[0].event, "DataAdded");
    });

    it("test adding data for public token", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);

        const tx1 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            200001,
            [accounts[0], cidBytes, "{duration: 1000}"],
            { from: accounts[0], value: web3.utils.toWei("0.015", "ether") }
        );
        assert.equal(tx1.logs[0].event, "DataAdded");

        const tx2 = await this.ownerDataContract.add(
            this.exhibitionContract.address,
            200001,
            [accounts[2], cidBytes, "{duration: 1000}"],
            { from: accounts[2], value: web3.utils.toWei("0.015", "ether") }
        );
        assert.equal(tx2.logs[0].event, "DataAdded");
    });

    it("test adding data for public token using signed add function", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const data = [
            "0x23221e5403511CeC833294D2B1B006e9D639A61b",
            cidBytes,
            "{duration: 1000}",
        ];

        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const chainId = await web3.eth.getChainId();
        const signedParams = web3.eth.abi.encodeParameters(
            ["uint", "address", "bytes", "uint256"],
            [
                BigInt(chainId).toString(),
                this.ownerDataContract.address,
                web3.utils.fromAscii(""),
                expiryTime,
            ]
        );

        const hash = web3.utils.keccak256(signedParams);
        const trusteeSignature = await web3.eth.sign(
            hash,
            this.ownerDataSigner
        );
        const sig = trusteeSignature.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);

        // sign params
        const signs = [
            web3.utils.fromAscii(""),
            expiryTime,
            r,
            s,
            web3.utils.toDecimal(v) + 27,
        ];

        const tx = await this.ownerDataContract.signedAdd(
            this.exhibitionContract.address,
            200001,
            data,
            signs
        );
        assert.equal(tx.logs[0].event, "DataAdded");
    });

    it("test add sound for ERC1155 token successfully", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        const tx = await this.ownerDataContract.add(
            this.fungibleContract.address,
            999,
            [accounts[0], cidBytes, "{duration: 1000}"],
            { from: accounts[0], value: 0 }
        );
        assert.equal(tx.logs[0].event, "DataAdded");
    });

    it("test add sound for ERC1155 token failed because wrong token owner", async function () {
        const cid = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
        const cidBytes = web3.utils.fromAscii(cid);
        try {
            const tx = await this.ownerDataContract.add(
                this.fungibleContract.address,
                999,
                [accounts[0], cidBytes, "{duration: 1000}"],
                { from: accounts[2], value: 0 }
            );
        } catch (error) {
            assert.equal(error.reason, "Custom error (could not decode)");
        }
    });

    it("test sort data by timestamp", async function () {
        const cid1 = "123";
        const cid2 = "456";
        const cid3 = "789";
        const cid4 = "555";

        const cidBytes1 = web3.utils.fromAscii(cid1);
        const cidBytes2 = web3.utils.fromAscii(cid2);
        const cidBytes3 = web3.utils.fromAscii(cid3);
        const cidBytes4 = web3.utils.fromAscii(cid4);

        const ownerDataContract = await OwnerData.new(
            this.ownerDataSigner,
            COST_RECEIVER,
            web3.utils.toWei("0.015", "ether"),
            400001
        );

        await ownerDataContract.add(
            this.exhibitionContract.address,
            400001,
            [accounts[1], cidBytes1, "{duration: 1000}"],
            { from: accounts[1], value: web3.utils.toWei("0.015", "ether") }
        );

        await ownerDataContract.add(
            this.exhibitionContract.address,
            400001,
            [accounts[1], cidBytes2, "{duration: 1000}"],
            { from: accounts[1], value: web3.utils.toWei("0.015", "ether") }
        );

        await ownerDataContract.add(
            this.exhibitionContract.address,
            400001,
            [accounts[1], cidBytes3, "{duration: 1000}"],
            { from: accounts[1], value: web3.utils.toWei("0.015", "ether") }
        );

        await ownerDataContract.add(
            this.exhibitionContract.address,
            400001,
            [accounts[1], cidBytes4, "{duration: 1000}"],
            { from: accounts[1], value: web3.utils.toWei("0.015", "ether") }
        );

        await ownerDataContract.remove(
            this.exhibitionContract.address,
            400001,
            [1],
            { from: accounts[0] }
        );

        const data = await ownerDataContract.get(
            this.exhibitionContract.address,
            400001,
            0,
            100
        );

        assert.equal(data.length, 4);
        assert.equal(bytesToString(data[0].dataHash), cid1);
        assert.equal(bytesToString(data[1].dataHash), "");
        assert.equal(bytesToString(data[2].dataHash), cid3);
        assert.equal(bytesToString(data[3].dataHash), cid4);
    });

    it("test get data by owner", async function () {
        const cid1 = "123";
        const cid2 = "456";
        const cid3 = "789";
        const cid4 = "555";

        const cidBytes1 = web3.utils.fromAscii(cid1);
        const cidBytes2 = web3.utils.fromAscii(cid2);
        const cidBytes3 = web3.utils.fromAscii(cid3);
        const cidBytes4 = web3.utils.fromAscii(cid4);

        const ownerDataContract = await OwnerData.new(
            this.ownerDataSigner,
            COST_RECEIVER,
            web3.utils.toWei("0.015", "ether"),
            500001
        );

        await ownerDataContract.add(
            this.exhibitionContract.address,
            500001,
            [accounts[1], cidBytes1, "{duration: 1000}"],
            { from: accounts[1], value: web3.utils.toWei("0.015", "ether") }
        );

        await ownerDataContract.add(
            this.exhibitionContract.address,
            500001,
            [accounts[0], cidBytes2, "{duration: 1000}"],
            { from: accounts[0], value: web3.utils.toWei("0.015", "ether") }
        );

        await ownerDataContract.add(
            this.exhibitionContract.address,
            500001,
            [accounts[0], cidBytes3, "{duration: 1000}"],
            { from: accounts[0], value: web3.utils.toWei("0.015", "ether") }
        );

        await ownerDataContract.add(
            this.exhibitionContract.address,
            500001,
            [accounts[1], cidBytes4, "{duration: 1000}"],
            { from: accounts[1], value: web3.utils.toWei("0.015", "ether") }
        );

        const data = await ownerDataContract.getByOwner(
            this.exhibitionContract.address,
            500001,
            accounts[0]
        );

        assert.equal(data.length, 2);
    });

    it("test set cost negative value", async function () {
        try {
            await this.ownerDataContract.setServiceFee(
                web3.utils.toWei("-0.015", "ether")
            );
        } catch (error) {
            assert.equal(error.reason, "value out-of-bounds");
        }
    });
});
