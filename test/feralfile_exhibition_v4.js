const FeralfileExhibitionV4 = artifacts.require("FeralfileExhibitionV4");
const FeralfileVault = artifacts.require("FeralfileVault");

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";
const COST_RECEIVER = "0x46f2B641d8702f29c45f6D06292dC34Eb9dB1801";
const VAULT_ADDRESS = "0x7a15b36cb834aea88553de69077d3777460d73ac";
const TOKEN_BASE_URI = "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

contract("FeralfileExhibitionV4_0", async (accounts) => {
    before(async function () {
        this.signer = accounts[0];
        this.vault = await FeralfileVault.new();
        this.exhibition = await FeralfileExhibitionV4.new(
            "Feral File V4 Test",
            "FFv4",
            this.signer,
            true,
            true,
            this.vault.address,
            COST_RECEIVER,
            TOKEN_BASE_URI,
            [0, 1, 2, 3, 4],
            [1, 1, 100, 1000, 10000]
        );
        await this.vault.setExhibitionContract(this.exhibition.address);
    });

    it("test contract constructor", async function () {
        let burnable = await this.exhibition.burnable();
        assert.ok(burnable);

        let bridgeable = await this.exhibition.bridgeable();
        assert.ok(bridgeable);

        let signer = await this.exhibition.signer();
        assert.equal(signer, this.signer);

        let seriesMaxSupply1 = await this.exhibition.seriesMaxSupply(0);
        assert.equal(seriesMaxSupply1, 1);

        let seriesMaxSupply2 = await this.exhibition.seriesMaxSupply(1);
        assert.equal(seriesMaxSupply2, 1);

        let seriesTotalSupply1 = await this.exhibition.seriesTotalSupply(0);
        assert.equal(seriesTotalSupply1, 0);
    });

    it("test mint artwork", async function () {
        // 1. mint successful
        const owner = accounts[1];
        const seriesId1 = 0;
        const seriesId2 = 1;
        const tokenId1 = seriesId1 * 1000000 + 0;
        const tokenId2 = seriesId2 * 1000000 + 0;
        let data = [
            [seriesId1, tokenId1, owner],
            [seriesId2, tokenId2, owner],
        ];
        let tx = await this.exhibition.mintArtworks(data);

        // total supply
        let totalSupply = await this.exhibition.totalSupply();
        assert.equal(totalSupply, 2);

        // series total supply
        let totalSupply1 = await this.exhibition.seriesTotalSupply(seriesId1);
        let totalSupply2 = await this.exhibition.seriesTotalSupply(seriesId2);
        assert.equal(totalSupply1, 1);
        assert.equal(totalSupply2, 1);

        // owner
        let tokenOwner = await this.exhibition.ownerOf(tokenId1);
        assert.equal(tokenOwner, owner);

        // balance
        let balance = await this.exhibition.balanceOf(owner);
        assert.equal(balance, 2);

        // get artwork
        let artwork = await this.exhibition.getArtwork(tokenId1);
        assert.equal(artwork.tokenId, tokenId1);
        assert.equal(artwork.seriesId, seriesId1);

        // event emitted
        let { logs } = tx;
        assert.ok(Array.isArray(logs));
        assert.equal(logs.length, 4);

        let log0 = logs[0];
        let log1 = logs[1];
        assert.equal(log0.event, "Transfer");
        assert.equal(log1.event, "NewArtwork");

        // 2. mint failed
        // series doesn't exist
        data = [[999999, 0, owner]];
        try {
            await this.exhibition.mintArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: seriesId doesn't exist: 999999"
                )
            );
        }

        // mint to zero address
        data = [[seriesId1, 999999, ZERO_ADDRESS]];
        try {
            await this.exhibition.mintArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes("ERC721: mint to the zero address")
            );
        }

        // token already minted
        data = [[seriesId1, tokenId1, accounts[2]]];
        try {
            await this.exhibition.mintArtworks(data);
        } catch (error) {
            assert.ok(error.message.includes("ERC721: token already minted"));
        }

        // no more slots
        data = [[seriesId1, 999999, accounts[2]]];
        try {
            await this.exhibition.mintArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: no slots available"
                )
            );
        }

        // not mintable
        await this.exhibition.setMintable(false);
        try {
            await this.exhibition.mintArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: contract doesn't allow to mint"
                )
            );
        }
    });

    it("test burn artwork", async function () {
        // 1. Burn successfully
        const owner = accounts[0];
        const seriesId1 = 0;
        const seriesId2 = 1;
        const tokenId1 = seriesId1 * 1000000 + 0;
        const tokenId2 = seriesId2 * 1000000 + 0;
        let data = [tokenId1, tokenId2];
        let tx = await this.exhibition.burnArtworks(data);

        // total supply
        let totalSupply = await this.exhibition.totalSupply();
        assert.equal(totalSupply, 0);

        // series total supply
        let totalSupply1 = await this.exhibition.seriesTotalSupply(seriesId1);
        let totalSupply2 = await this.exhibition.seriesTotalSupply(seriesId2);
        assert.equal(totalSupply1, 0);
        assert.equal(totalSupply2, 0);

        // balance
        let balance = await this.exhibition.balanceOf(owner);
        assert.equal(balance, 0);

        // get artwork
        try {
            await this.exhibition.getArtwork(tokenId1);
        } catch (error) {
            assert.ok(error.message.includes("ERC721: token doesn't exist"));
        }

        // event emitted
        let { logs } = tx;
        assert.ok(Array.isArray(logs));
        assert.equal(logs.length, 4);

        let log0 = logs[0];
        let log1 = logs[1];
        assert.equal(log0.event, "Transfer");
        assert.equal(log1.event, "BurnArtwork");

        // 2. Burn failed

        // token doesn't exist
        data = [111111, 222222];
        try {
            await this.exhibition.burnArtworks(data);
        } catch (error) {
            assert.ok(error.message.includes("ERC721: token doesn't exist"));
        }

        // not burnable
        await this.exhibition.setBurnable(false);
        try {
            await this.exhibition.burnArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: token is not burnable"
                )
            );
        }
    });

    it("test buy artworks successfully", async function () {
        await this.exhibition.setMintable(true);
        // Mint for buy by crypto
        let owner = this.exhibition.address;
        await this.exhibition.mintArtworks([
            [3, 3000000, owner],
            [3, 3000001, owner],
            [4, 4000000, owner],
            [4, 4000001, owner],
            [4, 4000002, owner],
        ]);
        // Mint for credit card
        await this.exhibition.mintArtworks([
            [3, 3000002, owner],
            [3, 3000003, owner],
            [4, 4000003, owner],
            [4, 4000004, owner],
        ]);

        // Generate signature
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                this.exhibition.address,
                [
                    BigInt(0.25 * 1e18).toString(),
                    BigInt(0.02 * 1e18).toString(),
                    expiryTime,
                    accounts[2],
                    [3000000, 3000001, 4000000, 4000001, 4000002],
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
        var sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);
        // Generate signature

        try {
            const acc3BalanceBefore = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceBefore = await web3.eth.getBalance(accounts[4]);
            const accCostReceiverBalanceBefore = await web3.eth.getBalance(
                COST_RECEIVER
            );

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
                    [3000000, 3000001, 4000000, 4000001, 4000002],
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
            const ownerOfToken0 = await this.exhibition.ownerOf(3000000);
            const ownerOfToken1 = await this.exhibition.ownerOf(3000001);
            const ownerOfToken2 = await this.exhibition.ownerOf(4000000);
            const ownerOfToken3 = await this.exhibition.ownerOf(4000001);
            const ownerOfToken4 = await this.exhibition.ownerOf(4000002);
            assert.equal(ownerOfToken0, accounts[2]);
            assert.equal(ownerOfToken1, accounts[2]);
            assert.equal(ownerOfToken2, accounts[2]);
            assert.equal(ownerOfToken3, accounts[2]);
            assert.equal(ownerOfToken4, accounts[2]);

            const acc3BalanceAfter = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceAfter = await web3.eth.getBalance(accounts[4]);
            const accCostReceiverBalanceAfter = await web3.eth.getBalance(
                COST_RECEIVER
            );

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
            assert.equal(
                (
                    BigInt(accCostReceiverBalanceAfter) -
                    BigInt(accCostReceiverBalanceBefore)
                ).toString(),
                BigInt(0.02 * 1e18).toString()
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
                BigInt(await web3.eth.getChainId()).toString(),
                this.exhibition.address,
                [
                    BigInt(0.22 * 1e18).toString(),
                    BigInt(0.02 * 1e18).toString(),
                    expiryTime,
                    accounts[2],
                    [3000002, 3000003, 4000003, 4000004],
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
        var sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);
        // Generate signature
        try {
            const acc3BalanceBefore = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceBefore = await web3.eth.getBalance(accounts[4]);
            const vaultBalanceBefore = await web3.eth.getBalance(
                this.vault.address
            );
            const accCostReceiverBalanceBefore = await web3.eth.getBalance(
                COST_RECEIVER
            );

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
                    [3000002, 3000003, 4000003, 4000004],
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
            const ownerOfToken0 = await this.exhibition.ownerOf(3000002);
            const ownerOfToken1 = await this.exhibition.ownerOf(3000003);
            const ownerOfToken2 = await this.exhibition.ownerOf(4000003);
            const ownerOfToken3 = await this.exhibition.ownerOf(4000004);
            assert.equal(ownerOfToken0, accounts[2]);
            assert.equal(ownerOfToken1, accounts[2]);
            assert.equal(ownerOfToken2, accounts[2]);
            assert.equal(ownerOfToken3, accounts[2]);

            const acc3BalanceAfter = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceAfter = await web3.eth.getBalance(accounts[4]);
            const vaultBalanceAfter = await web3.eth.getBalance(
                this.vault.address
            );
            const accCostReceiverBalanceAfter = await web3.eth.getBalance(
                COST_RECEIVER
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
            assert.equal(
                (
                    BigInt(vaultBalanceBefore) - BigInt(vaultBalanceAfter)
                ).toString(),
                BigInt(0.22 * 1e18).toString()
            );
            assert.equal(
                (
                    BigInt(accCostReceiverBalanceAfter) -
                    BigInt(accCostReceiverBalanceBefore)
                ).toString(),
                BigInt(0.02 * 1e18).toString()
            );
        } catch (err) {
            console.log(err);
            assert.fail();
        }
    });
});
