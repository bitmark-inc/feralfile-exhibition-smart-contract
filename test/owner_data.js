const OwnerData = artifacts.require("OwnerData");
const FeralfileExhibitionV4 = artifacts.require("FeralfileExhibitionV4");
const FeralfileVault = artifacts.require("FeralfileVault");

const CONTRACT_URI = "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
const COST_RECEIVER = "0x46f2B641d8702f29c45f6D06292dC34Eb9dB1801";

const { bufferToHex } = require("ethereumjs-util");

const bytesToString = (bytes) => {
    return web3.utils.toAscii(bytes).replace(/\u0000/g, "");
};

contract("OwnerData", async (accounts) => {
    before(async function () {
        this.signer = accounts[0];
        this.trustee = accounts[1];
        this.signTrustee = accounts[5];
        this.vault = await FeralfileVault.new(this.signer);
        this.ownerDataContract = await OwnerData.new(this.signTrustee);
        this.seriesIds = [1, 2];
        this.seriesMaxSupply = [100, 1];
        this.seriesArtworkMaxSupply = [1, 100];
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
            [1, 1, accounts[0]],
            [1, 2, accounts[1]],
            [1, 3, accounts[2]],
            [1, 4, accounts[0]],
            [1, 5, accounts[1]],
            [1, 6, accounts[2]],
            [1, 7, "0x23221e5403511CeC833294D2B1B006e9D639A61b"],
        ]);
        await this.exhibitionContract.addTrustee(this.trustee);
    });

    it("test adding data successfully", async function () {
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

    it("test adding data failed because already added", async function () {
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
        try {
            await this.ownerDataContract.add(
                this.exhibitionContract.address,
                3,
                [accounts[2], updatedCidBytes, "{duration: 2000}"],
                { from: accounts[2] }
            );
        } catch (error) {
            assert.equal(error.reason, "OwnerData: data already added");
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
        const acc2Owner = await this.exhibitionContract.ownerOf(2);
        assert.equal(acc2Owner, accounts[2]);

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
        const acc4Owner = await this.exhibitionContract.ownerOf(2);
        assert.equal(acc4Owner, accounts[4]);

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
            assert.equal(error.reason, "OwnerData: data owner mismatch");
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
        const trusteeSignature = await web3.eth.sign(hash, accounts[5]);
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
            7,
            data,
            signs
        );
        assert.equal(tx.logs[0].event, "DataAdded");
    });
});
