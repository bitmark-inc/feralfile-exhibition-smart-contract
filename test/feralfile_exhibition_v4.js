const FeralfileExhibitionV4 = artifacts.require("FeralfileExhibitionV4");
const FeralfileVault = artifacts.require("FeralfileVault");

const CONTRACT_URI =
    "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ";

const IPFS_GATEWAY_PREFIX =
    "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";
const VAULT_CONTRACT_ADDRESS = "0x0000000000000000000000000000000000000000";

const originArtworkCID = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

contract("FeralfileExhibitionV4", async (accounts) => {
    before(async function () {
        this.vault = await FeralfileVault.new();

        this.exhibition = await FeralfileExhibitionV4.new(
            "Feral File V4 Test 001",
            "FFV4",
            CONTRACT_URI,
            IPFS_GATEWAY_PREFIX,
            accounts[1],
            this.vault.address,
            true,
            true
        );

        await this.vault.setExhibitionContract(this.exhibition.address);
    });

    it("check contract is burnable", async function () {
        const isBurnable = await this.exhibition.isBurnable();
        assert.equal(isBurnable, true);
    });

    it("check contract is bridgeable", async function () {
        const isBridgeable = await this.exhibition.isBridgeable();
        assert.equal(isBridgeable, true);
    });

    it("check contract is NOT burnable", async function () {
        const exh = await FeralfileExhibitionV4.new(
            "Feral File V4 Test 002",
            "FFV4",
            CONTRACT_URI,
            IPFS_GATEWAY_PREFIX,
            accounts[1],
            VAULT_CONTRACT_ADDRESS,
            false,
            true
        );
        const isBurnable = await exh.isBurnable();
        assert.equal(isBurnable, false);
    });

    it("check contract is NOT bridgeable", async function () {
        const exh = await FeralfileExhibitionV4.new(
            "Feral File V4 Test 002",
            "FFV4",
            CONTRACT_URI,
            IPFS_GATEWAY_PREFIX,
            accounts[1],
            VAULT_CONTRACT_ADDRESS,
            true,
            false
        );
        const isBridgeable = await exh.isBridgeable();
        assert.equal(isBridgeable, false);
    });

    it("mint artworks successfully", async function () {
        try {
            // Mint for buy by crypto
            await this.exhibition.mintArtworks([
                [1, 0],
                [1, 1],
                [2, 0],
                [2, 1],
                [2, 2],
            ]);
            // Mint for credit card
            await this.exhibition.mintArtworks([
                [1, 2],
                [1, 3],
                [2, 3],
                [2, 4],
            ]);
            const totalSupply = await this.exhibition.totalSupply();
            assert.equal(totalSupply, 9);

            const ownerOfToken0 = await this.exhibition.ownerOf(1000000);
            const ownerOfToken1 = await this.exhibition.ownerOf(1000001);
            const ownerOfToken2 = await this.exhibition.ownerOf(2000000);
            const ownerOfToken3 = await this.exhibition.ownerOf(2000001);
            const ownerOfToken4 = await this.exhibition.ownerOf(2000002);

            const ownerOfToken5 = await this.exhibition.ownerOf(1000002);
            const ownerOfToken6 = await this.exhibition.ownerOf(1000003);
            const ownerOfToken7 = await this.exhibition.ownerOf(2000003);
            const ownerOfToken8 = await this.exhibition.ownerOf(2000004);

            assert.equal(ownerOfToken0, this.exhibition.address);
            assert.equal(ownerOfToken1, this.exhibition.address);
            assert.equal(ownerOfToken2, this.exhibition.address);
            assert.equal(ownerOfToken3, this.exhibition.address);
            assert.equal(ownerOfToken4, this.exhibition.address);

            assert.equal(ownerOfToken5, this.exhibition.address);
            assert.equal(ownerOfToken6, this.exhibition.address);
            assert.equal(ownerOfToken7, this.exhibition.address);
            assert.equal(ownerOfToken8, this.exhibition.address);

            const tokenURIOfToken0 = await this.exhibition.tokenURI(1000000);
            const tokenURIOfToken1 = await this.exhibition.tokenURI(1000001);
            const tokenURIOfToken2 = await this.exhibition.tokenURI(2000000);
            const tokenURIOfToken3 = await this.exhibition.tokenURI(2000001);
            const tokenURIOfToken4 = await this.exhibition.tokenURI(2000002);

            const tokenURIOfToken5 = await this.exhibition.tokenURI(1000002);
            const tokenURIOfToken6 = await this.exhibition.tokenURI(1000003);
            const tokenURIOfToken7 = await this.exhibition.tokenURI(2000003);
            const tokenURIOfToken8 = await this.exhibition.tokenURI(2000004);

            assert.equal(tokenURIOfToken0, IPFS_GATEWAY_PREFIX + "/1000000");
            assert.equal(tokenURIOfToken1, IPFS_GATEWAY_PREFIX + "/1000001");
            assert.equal(tokenURIOfToken2, IPFS_GATEWAY_PREFIX + "/2000000");
            assert.equal(tokenURIOfToken3, IPFS_GATEWAY_PREFIX + "/2000001");
            assert.equal(tokenURIOfToken4, IPFS_GATEWAY_PREFIX + "/2000002");

            assert.equal(tokenURIOfToken5, IPFS_GATEWAY_PREFIX + "/1000002");
            assert.equal(tokenURIOfToken6, IPFS_GATEWAY_PREFIX + "/1000003");
            assert.equal(tokenURIOfToken7, IPFS_GATEWAY_PREFIX + "/2000003");
            assert.equal(tokenURIOfToken8, IPFS_GATEWAY_PREFIX + "/2000004");
        } catch (err) {
            console.log(err);
            assert.fail();
        }
    });

    it("test buy artworks successfully", async function () {
        // Generate signature
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool)",
            ],
            [
                "1",
                this.exhibition.address,
                [
                    BigInt(0.25 * 1e18).toString(),
                    BigInt(0.02 * 1e18).toString(),
                    expiryTime,
                    accounts[2],
                    [1000000, 1000001, 2000000, 2000001, 2000002],
                    [
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                    ],
                    false,
                ],
            ]
        );
        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, accounts[1]);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);
        // Generate signature

        try {
            const acc3BalanceBefore = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceBefore = await web3.eth.getBalance(accounts[4]);

            await this.exhibition.startSale();
            await this.exhibition.buyArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27, // magic 27
                [
                    BigInt(0.25 * 1e18).toString(),
                    BigInt(0.02 * 1e18).toString(),
                    expiryTime,
                    accounts[2],
                    [1000000, 1000001, 2000000, 2000001, 2000002],
                    [
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                    ],
                    false,
                ],
                { from: accounts[5], value: 0.25 * 1e18 }
            );
            const ownerOfToken0 = await this.exhibition.ownerOf(1000000);
            const ownerOfToken1 = await this.exhibition.ownerOf(1000001);
            const ownerOfToken2 = await this.exhibition.ownerOf(2000000);
            const ownerOfToken3 = await this.exhibition.ownerOf(2000001);
            const ownerOfToken4 = await this.exhibition.ownerOf(2000002);
            assert.equal(ownerOfToken0, accounts[2]);
            assert.equal(ownerOfToken1, accounts[2]);
            assert.equal(ownerOfToken2, accounts[2]);
            assert.equal(ownerOfToken3, accounts[2]);
            assert.equal(ownerOfToken4, accounts[2]);

            const acc3BalanceAfter = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceAfter = await web3.eth.getBalance(accounts[4]);

            assert.equal(
                (
                    BigInt(acc3BalanceAfter) - BigInt(acc3BalanceBefore)
                ).toString(),
                BigInt((0.23 * 1e18 * 80) / 100).toString()
            );
            assert.equal(
                (
                    BigInt(acc4BalanceAfter) - BigInt(acc4BalanceBefore)
                ).toString(),
                BigInt((0.23 * 1e18 * 20) / 100).toString()
            );
        } catch (err) {
            console.log(err);
            assert.fail();
        }
    });

    it("test buy artworks successfully with vault contract", async function () {
        await web3.eth.sendTransaction({
            to: this.vault.address,
            from: accounts[8],
            value: BigInt(0.5 * 1e18).toString(),
        });
        // Generate signature
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool)",
            ],
            [
                "1",
                this.exhibition.address,
                [
                    BigInt(0.22 * 1e18).toString(),
                    BigInt(0.02 * 1e18).toString(),
                    expiryTime,
                    accounts[2],
                    [1000002, 1000003, 2000003, 2000004],
                    [
                        [
                            [accounts[3], 5000],
                            [accounts[4], 5000],
                        ],
                        [
                            [accounts[3], 5000],
                            [accounts[4], 5000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                    ],
                    true,
                ],
            ]
        );
        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, accounts[1]);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);
        // Generate signature
        try {
            const acc3BalanceBefore = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceBefore = await web3.eth.getBalance(accounts[4]);

            await this.exhibition.startSale();
            await this.exhibition.buyArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27, // magic 27
                [
                    BigInt(0.22 * 1e18).toString(),
                    BigInt(0.02 * 1e18).toString(),
                    expiryTime,
                    accounts[2],
                    [1000002, 1000003, 2000003, 2000004],
                    [
                        [
                            [accounts[3], 5000],
                            [accounts[4], 5000],
                        ],
                        [
                            [accounts[3], 5000],
                            [accounts[4], 5000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                        [
                            [accounts[3], 8000],
                            [accounts[4], 2000],
                        ],
                    ],
                    true,
                ],
                { from: accounts[5], value: 0 } // itx relay
            );
            const ownerOfToken0 = await this.exhibition.ownerOf(1000002);
            const ownerOfToken1 = await this.exhibition.ownerOf(1000003);
            const ownerOfToken2 = await this.exhibition.ownerOf(2000003);
            const ownerOfToken3 = await this.exhibition.ownerOf(2000004);
            assert.equal(ownerOfToken0, accounts[2]);
            assert.equal(ownerOfToken1, accounts[2]);
            assert.equal(ownerOfToken2, accounts[2]);
            assert.equal(ownerOfToken3, accounts[2]);

            const acc3BalanceAfter = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceAfter = await web3.eth.getBalance(accounts[4]);
            const vaultBalanceAfter = await web3.eth.getBalance(
                this.vault.address
            );

            assert.equal(
                (
                    BigInt(acc3BalanceAfter) - BigInt(acc3BalanceBefore)
                ).toString(),
                BigInt(((0.2 / 4) * 2 * 1e18 * 130) / 100).toString()
            );
            assert.equal(
                (
                    BigInt(acc4BalanceAfter) - BigInt(acc4BalanceBefore)
                ).toString(),
                BigInt(((0.2 / 4) * 2 * 1e18 * 70) / 100).toString()
            );

            // Check vault contract balance
            assert.equal(
                BigInt(0.5 * 1e18 - 0.22 * 1e18).toString(),
                BigInt(vaultBalanceAfter).toString()
            );
        } catch (err) {
            console.log(err);
            assert.fail();
        }
    });
});
