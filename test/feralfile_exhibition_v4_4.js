const FeralfileExhibitionV4_4 = artifacts.require("FeralfileExhibitionV4_4");
const FeralfileVault = artifacts.require("FeralfileVault");
const { BN, expectEvent } = require("@openzeppelin/test-helpers");
const abi = require("../build/contracts/FeralfileExhibitionV4_4.json").abi;
const { expectCustomError } = require("./helper/expectCustomError.js");
const { randomString } = require("./helper/array.js");
const { expect } = require("chai");
const Base64 = {
    encode: (str) => Buffer.from(str).toString("base64"),
};

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";
const COST_RECEIVER = "0x46f2B641d8702f29c45f6D06292dC34Eb9dB1801";
const CONTRACT_URI = "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

contract("FeralfileExhibitionV4_4", async (accounts) => {
    let contract;
    let vault;
    const signer = accounts[0];
    const seriesIds = [0, 1, 2, 3, 4];
    const seriesMaxSupply = [1, 1, 100, 1000, 10000];

    beforeEach(async function () {
        vault = await FeralfileVault.new(signer);
        contract = await FeralfileExhibitionV4_4.new(
            "Feral File V4 Test",
            "FFv4",
            true,
            true,
            signer,
            vault.address,
            COST_RECEIVER,
            CONTRACT_URI,
            seriesIds,
            seriesMaxSupply
        );
    });

    it("test contract constructor", async function () {
        const burnable = await contract.burnable();
        assert.ok(burnable);

        const bridgeable = await contract.bridgeable();
        assert.ok(bridgeable);

        const mintable = await contract.mintable();
        assert.ok(mintable);

        const vaultAddress = await contract.vault();
        assert.equal(vaultAddress, vault.address);

        const signer = await contract.signer();
        assert.equal(signer, signer);

        const seriesMaxSupply1 = await contract.seriesMaxSupply(seriesIds[0]);
        assert.equal(seriesMaxSupply1, seriesMaxSupply[0]);

        const seriesMaxSupply2 = await contract.seriesMaxSupply(seriesIds[1]);
        assert.equal(seriesMaxSupply2, seriesMaxSupply[1]);

        const seriesTotalSupply1 = await contract.seriesTotalSupply(
            seriesIds[0]
        );
        assert.equal(seriesTotalSupply1, 0);
    });

    it("test duplicate series in constructor", async function () {
        // Deploy contract with duplicate series defined
        let seriesIds = [0, 1, 2, 3, 1];
        let seriesMaxSupply = [1, 1, 100, 1000, 10000];
        try {
            await FeralfileExhibitionV4_4.new(
                "Feral File V4 Test",
                "FFv4",
                true,
                true,
                signer,
                vault.address,
                COST_RECEIVER,
                CONTRACT_URI,
                seriesIds,
                seriesMaxSupply
            );
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: duplicate seriesId"
                )
            );
        }
    });

    it("test mint artwork", async function () {
        // 1. mint successful
        const owner = accounts[1];
        const tokenId1 = seriesIds[3] * 1000000 + 0;
        const tokenId2 = seriesIds[4] * 1000000 + 0;
        let data = [
            [seriesIds[2], tokenId1, owner],
            [seriesIds[3], tokenId2, owner],
        ];
        const tx = await contract.mintArtworks(data);

        // total supply
        const totalSupply = await contract.totalSupply();
        assert.equal(totalSupply, 2);

        // series total supply
        const totalSupply1 = await contract.seriesTotalSupply(seriesIds[2]);
        const totalSupply2 = await contract.seriesTotalSupply(seriesIds[3]);
        assert.equal(totalSupply1, 1);
        assert.equal(totalSupply2, 1);

        // owner
        const tokenOwner = await contract.ownerOf(tokenId1);
        assert.equal(tokenOwner, owner);

        // balance
        const balance = await contract.balanceOf(owner);
        assert.equal(balance, 2);

        // get artwork
        const artwork = await contract.getArtwork(tokenId1);
        assert.equal(artwork.tokenId, tokenId1);
        assert.equal(artwork.seriesId, seriesIds[2]);

        // get token ID from owner
        const tokenIds = await contract.tokensOfOwner(owner);
        assert.equal(tokenIds.length, 2);
        assert.equal(tokenIds[0], tokenId1);
        assert.equal(tokenIds[1], tokenId2);

        // get token ID from owner & index
        const tokenId = await contract.tokenOfOwnerByIndex(owner, 0);
        assert.equal(tokenId, tokenId1);

        // event emitted
        const { logs } = tx;
        assert.ok(Array.isArray(logs));
        assert.equal(logs.length, 4);

        const log0 = logs[0];
        const log1 = logs[1];
        assert.equal(log0.event, "Transfer");
        assert.equal(log1.event, "NewArtwork");

        // 2. mint failed
        // series doesn't exist
        data = [[999999, 0, owner]];
        try {
            await contract.mintArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: seriesId doesn't exist: 999999"
                )
            );
        }

        // mint to zero address
        data = [[seriesIds[3], 999999, ZERO_ADDRESS]];
        try {
            await contract.mintArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes("ERC721: mint to the zero address")
            );
        }

        // token already minted
        data = [[seriesIds[3], tokenId1, accounts[2]]];
        try {
            await contract.mintArtworks(data);
        } catch (error) {
            assert.ok(error.message.includes("ERC721: token already minted"));
        }

        // no more slots
        data = [[seriesIds[2], 999999, accounts[2]]];
        try {
            await contract.mintArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: no slots available"
                )
            );
        }
    });

    it("test burn artwork", async function () {
        // 1. Prepare data
        const owner = accounts[1];
        const tokenId1 = seriesIds[0] * 1000000 + 0;
        const tokenId2 = seriesIds[1] * 1000000 + 0;
        const tokenId3 = seriesIds[2] * 1000000 + 0;
        let mintData = [
            [seriesIds[0], tokenId1, owner],
            [seriesIds[1], tokenId2, owner],
            [seriesIds[2], tokenId3, owner],
        ];
        let tx = await contract.mintArtworks(mintData);

        // 2. Burn successfully
        let burnData = [tokenId1, tokenId2];
        tx = await contract.burnArtworks(burnData, { from: owner });

        // total supply
        const totalSupply = await contract.totalSupply();
        assert.equal(totalSupply, 1);

        // series total supply
        const totalSupply1 = await contract.seriesTotalSupply(seriesIds[0]);
        const totalSupply2 = await contract.seriesTotalSupply(seriesIds[1]);
        assert.equal(totalSupply1, 0);
        assert.equal(totalSupply2, 0);

        // balance
        const balance = await contract.balanceOf(owner);
        assert.equal(balance, 1);

        // token of owner
        const tokenIds = await contract.tokensOfOwner(owner);
        assert.equal(tokenIds.length, 1);

        // get artwork
        try {
            await contract.getArtwork(tokenId1);
        } catch (error) {
            assert.ok(error.message.includes("ERC721: invalid token ID"));
        }

        // event emitted
        const { logs } = tx;
        assert.ok(Array.isArray(logs));
        assert.equal(logs.length, 4);

        const log1 = logs[0];
        const log2 = logs[1];
        assert.equal(log1.event, "Transfer");
        assert.equal(log2.event, "BurnArtwork");

        // 3. Burn failed
        // token doesn't exist
        try {
            await contract.burnArtworks([111111, 222222]);
        } catch (error) {
            assert.ok(error.message.includes("ERC721: invalid token ID"));
        }

        // Wrong owner or approval
        const wrongOwner = accounts[5];
        try {
            await contract.burnArtworks([tokenId3], {
                from: wrongOwner,
            });
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "ERC721: caller is not token owner or approved"
                )
            );
        }
    });

    it("test buy artworks successfully", async function () {
        // Mint for buy by crypto
        let owner = contract.address;
        await contract.mintArtworks([
            [seriesIds[3], 3000000, owner],
            [seriesIds[3], 3000001, owner],
            [seriesIds[4], 4000000, owner],
            [seriesIds[4], 4000001, owner],
            [seriesIds[4], 4000002, owner],
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
                contract.address,
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
        var sig = await web3.eth.sign(hash, signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = web3.utils.toDecimal("0x" + sig.slice(128, 130));
        // Generate signature

        try {
            const acc3BalanceBefore = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceBefore = await web3.eth.getBalance(accounts[4]);
            const accCostReceiverBalanceBefore =
                await web3.eth.getBalance(COST_RECEIVER);

            await contract.startSale();
            await contract.buyArtworks(
                r,
                s,
                v < 2 ? v + 27 : v, // magic 27
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
            const ownerOfToken0 = await contract.ownerOf(3000000);
            const ownerOfToken1 = await contract.ownerOf(3000001);
            const ownerOfToken2 = await contract.ownerOf(4000000);
            const ownerOfToken3 = await contract.ownerOf(4000001);
            const ownerOfToken4 = await contract.ownerOf(4000002);
            assert.equal(ownerOfToken0, accounts[2]);
            assert.equal(ownerOfToken1, accounts[2]);
            assert.equal(ownerOfToken2, accounts[2]);
            assert.equal(ownerOfToken3, accounts[2]);
            assert.equal(ownerOfToken4, accounts[2]);

            const acc3BalanceAfter = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceAfter = await web3.eth.getBalance(accounts[4]);
            const accCostReceiverBalanceAfter =
                await web3.eth.getBalance(COST_RECEIVER);

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
            assert.fail();
        }
    });

    it("test buy artworks successfully with vault contract", async function () {
        // Mint for credit card
        let owner = contract.address;
        await contract.mintArtworks([
            [seriesIds[3], 3000002, owner],
            [seriesIds[3], 3000003, owner],
            [seriesIds[4], 4000003, owner],
            [seriesIds[4], 4000004, owner],
        ]);

        await web3.eth.sendTransaction({
            to: vault.address,
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
                contract.address,
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
        var sig = await web3.eth.sign(hash, signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = web3.utils.toDecimal("0x" + sig.slice(128, 130));
        // Generate signature
        try {
            const acc3BalanceBefore = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceBefore = await web3.eth.getBalance(accounts[4]);
            const vaultBalanceBefore = await web3.eth.getBalance(vault.address);
            const accCostReceiverBalanceBefore =
                await web3.eth.getBalance(COST_RECEIVER);

            await contract.startSale();
            await contract.buyArtworks(
                r,
                s,
                v < 2 ? v + 27 : v, // magic 27
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
            const ownerOfToken0 = await contract.ownerOf(3000002);
            const ownerOfToken1 = await contract.ownerOf(3000003);
            const ownerOfToken2 = await contract.ownerOf(4000003);
            const ownerOfToken3 = await contract.ownerOf(4000004);
            assert.equal(ownerOfToken0, accounts[2]);
            assert.equal(ownerOfToken1, accounts[2]);
            assert.equal(ownerOfToken2, accounts[2]);
            assert.equal(ownerOfToken3, accounts[2]);

            const acc3BalanceAfter = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceAfter = await web3.eth.getBalance(accounts[4]);
            const vaultBalanceAfter = await web3.eth.getBalance(vault.address);
            const accCostReceiverBalanceAfter =
                await web3.eth.getBalance(COST_RECEIVER);

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
            assert.fail();
        }
    });

    it("test buy artworks failed with total bps over 10k", async function () {
        // Mint for buy by crypto
        let owner = contract.address;
        await contract.mintArtworks([
            [seriesIds[3], 3000000, owner],
            [seriesIds[3], 3000001, owner],
            [seriesIds[4], 4000000, owner],
            [seriesIds[4], 4000001, owner],
            [seriesIds[4], 4000002, owner],
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
                contract.address,
                [
                    BigInt(0.25 * 1e18).toString(),
                    BigInt(0.02 * 1e18).toString(),
                    expiryTime,
                    accounts[2],
                    [3000000, 3000001, 4000000, 4000001, 4000002],
                    [
                        [
                            [accounts[3], 8001],
                            [COST_RECEIVER, 2000],
                        ],
                        [
                            [accounts[3], 8001],
                            [COST_RECEIVER, 2000],
                        ],
                        [
                            [accounts[3], 9000],
                            [COST_RECEIVER, 2000],
                        ],
                        [
                            [accounts[3], 9000],
                            [COST_RECEIVER, 2000],
                        ],
                        [
                            [accounts[3], 8500],
                            [COST_RECEIVER, 2000],
                        ],
                    ],
                    false,
                ],
            ]
        );
        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = web3.utils.toDecimal("0x" + sig.slice(128, 130));
        // Generate signature

        try {
            await contract.startSale();
            await contract.buyArtworks(
                r,
                s,
                v < 2 ? v + 27 : v, // magic 27
                [
                    BigInt(0.25 * 1e18).toString(),
                    BigInt(0.02 * 1e18).toString(),
                    expiryTime,
                    accounts[2],
                    [3000000, 3000001, 4000000, 4000001, 4000002],
                    [
                        [
                            [accounts[3], 8001],
                            [COST_RECEIVER, 2000],
                        ],
                        [
                            [accounts[3], 8001],
                            [COST_RECEIVER, 2000],
                        ],
                        [
                            [accounts[3], 9000],
                            [COST_RECEIVER, 2000],
                        ],
                        [
                            [accounts[3], 9000],
                            [COST_RECEIVER, 2000],
                        ],
                        [
                            [accounts[3], 8500],
                            [COST_RECEIVER, 2000],
                        ],
                    ],
                    false,
                ],
                { from: accounts[5], value: 0.25 * 1e18 }
            );
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: total bps over 10,000"
                )
            );
        }
    });

    it("test start/stop and burn, pause/resume sale", async function () {
        // Resume the sale is not allowed before starting the sale
        try {
            await contract.resumeSale();
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: mintable required to be false"
                )
            );
        }

        // Pause the sale is not allowed before starting the sale
        try {
            await contract.pauseSale();
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: mintable required to be false"
                )
            );
        }

        // Mint new tokens
        const owner = contract.address;
        const tokenIDs = [4000997, 4000998, 4000999];
        await contract.mintArtworks([
            [seriesIds[4], tokenIDs[0], owner],
            [seriesIds[4], tokenIDs[1], owner],
            [seriesIds[4], tokenIDs[2], owner],
        ]);

        // Start the sale
        await contract.startSale();

        // Validate mintable flag to be false
        const mintable = await contract.mintable();
        assert.equal(mintable, false);

        // Validate selling flag to be true
        let selling = await contract.selling();
        assert.equal(selling, true);

        // Pause and resume the sale
        try {
            await contract.pauseSale();

            // Validate selling flag to be false
            selling = await contract.selling();
            assert.equal(selling, false);

            await contract.resumeSale();

            // Validate selling flag to be true
            selling = await contract.selling();
            assert.equal(selling, true);
        } catch (error) {
            assert.fail();
        }

        // Stop sale and burn remaining tokens
        try {
            await contract.stopSaleAndBurn();
        } catch (error) {
            assert.fail();
        }

        // Validate selling flag to be false
        selling = await contract.selling();
        assert.equal(selling, false);

        // Validate tokens were burned
        for (let i = 0; i < tokenIDs.length; i++) {
            try {
                await contract.getArtwork(tokenIDs[i]);
            } catch (error) {
                assert.ok(error.message.includes("ERC721: invalid token ID"));
            }
        }

        // Resume the sale is not allowed since there is no token left for sale
        try {
            await contract.resumeSale();
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: No token owned by the contract"
                )
            );
        }
    });

    it("test start/stop and transfer remaining token", async function () {
        // Mint new tokens
        const tokenIDs = [3000997, 3000998, 4000998, 4000999];
        await contract.mintArtworks([
            [seriesIds[3], tokenIDs[0], contract.address],
            [seriesIds[3], tokenIDs[1], contract.address],
            [seriesIds[4], tokenIDs[2], contract.address],
            [seriesIds[4], tokenIDs[3], contract.address],
        ]);

        // Start the sale
        await contract.startSale();

        const owner1 = accounts[1];
        const owner2 = accounts[2];

        // 1. Stop the sale and transfer remaining tokens failed
        try {
            await contract.stopSaleAndTransfer(
                [seriesIds[3], seriesIds[4]],
                [owner1]
            );
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: seriesIds length is different from recipientAddresses"
                )
            );
        }

        try {
            await contract.stopSaleAndTransfer(
                [seriesIds[1], seriesIds[2]],
                []
            );
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: seriesIds or recipientAddresses length is zero"
                )
            );
        }

        try {
            await contract.stopSaleAndTransfer(
                [seriesIds[1], seriesIds[2]],
                [owner1, owner2]
            );
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: Token for sale balance has to be zero"
                )
            );
        }

        // 2. Stop the sale and transfer remaining token successful
        try {
            await contract.stopSaleAndTransfer(
                [seriesIds[3], seriesIds[4]],
                [owner1, owner2]
            );
        } catch (error) {
            assert.fail();
        }

        // Validate balances
        const contractBal = await contract.balanceOf(contract.address);
        assert.equal(contractBal, 0);

        const owner1Bal = await contract.balanceOf(owner1);
        assert.equal(owner1Bal, 2);

        const owner2Bal = await contract.balanceOf(owner2);
        assert.equal(owner2Bal, 2);
    });

    it("test transfer token", async function () {
        // mint tokens
        const owner1 = accounts[1];
        const owner2 = accounts[2];
        const tokenId1 = seriesIds[3] * 1000000 + 555;
        const tokenId2 = seriesIds[3] * 1000000 + 556;
        const tokenId3 = seriesIds[4] * 1000000 + 555;
        const tokenId4 = seriesIds[4] * 1000000 + 556;
        let data = [
            [seriesIds[3], tokenId1, owner1],
            [seriesIds[3], tokenId2, owner1],
            [seriesIds[4], tokenId3, owner2],
            [seriesIds[4], tokenId4, owner2],
        ];

        try {
            await contract.mintArtworks(data);
        } catch (error) {
            assert.fail();
        }

        // transfer tokens
        // 1. Transfer to other address rather than contract
        try {
            await contract.transferFrom(owner1, owner2, tokenId1, {
                from: owner1,
            });
        } catch (error) {
            assert.fail();
        }

        // Validate new owner of token
        let owner = await contract.ownerOf(tokenId1);
        assert.equal(owner, owner2);

        // 2. Transfer to contract address
        try {
            await contract.transferFrom(owner2, owner1, tokenId1, {
                from: owner2,
            });
        } catch (error) {
            assert.ok(
                error.message.includes(
                    "FeralfileExhibitionV4: Contract isn't allowed to receive token"
                )
            );
        }

        // 3. Transfer as operator
        // TODO add later

        // 4. Transfer as approval
        // TODO add later
    });

    it("test set advance setting successfully", async function () {
        const advanceAddresses = [accounts[3], accounts[4]];
        const advanceAmounts = [
            web3.utils.toWei("0.3", "ether"),
            web3.utils.toWei("0.8", "ether"),
        ];

        // 1. Set advance setting
        try {
            await contract.setAdvanceSetting(advanceAddresses, advanceAmounts);

            const advanceAmount0 = await contract.advances(advanceAddresses[0]);
            assert.equal(advanceAmount0, advanceAmounts[0]);
            const advanceAmount1 = await contract.advances(advanceAddresses[1]);
            assert.equal(advanceAmount1, advanceAmounts[1]);
        } catch (error) {
            assert.fail();
        }
    });

    it("test set advance setting failed because address in use", async function () {
        const advanceAddresses = [accounts[3], accounts[4]];
        const advanceAmounts = [
            web3.utils.toWei("0.3", "ether"),
            web3.utils.toWei("0.8", "ether"),
        ];

        // 1. Set advance setting
        await contract.setAdvanceSetting(advanceAddresses, advanceAmounts);

        const advanceAmount0 = await contract.advances(advanceAddresses[0]);
        assert.equal(advanceAmount0, advanceAmounts[0]);
        const advanceAmount1 = await contract.advances(advanceAddresses[1]);
        assert.equal(advanceAmount1, advanceAmounts[1]);

        const updatedAddresses = [advanceAddresses[0], accounts[5]];
        const updatedAmounts = [
            web3.utils.toWei("0.5", "ether"),
            web3.utils.toWei("1", "ether"),
        ];

        // 2. Update advance setting
        try {
            await contract.setAdvanceSetting(updatedAddresses, updatedAmounts);
        } catch (error) {
            assert.equal(error.reason, "Custom error (could not decode)");
        }
    });

    it("test replace advance addresses successfully", async function () {
        const advanceAddresses = [accounts[3], accounts[4]];
        const advanceAmounts = [
            web3.utils.toWei("0.3", "ether"),
            web3.utils.toWei("0.8", "ether"),
        ];

        // 1. Set advance setting
        await contract.setAdvanceSetting(advanceAddresses, advanceAmounts);

        const advanceAmount0 = await contract.advances(advanceAddresses[0]);
        assert.equal(advanceAmount0, advanceAmounts[0]);
        const advanceAmount1 = await contract.advances(advanceAddresses[1]);
        assert.equal(advanceAmount1, advanceAmounts[1]);

        const oldAddresses = [advanceAddresses[0]];
        const newAddresses = [accounts[7]];

        // 2. Replace advance addresses
        await contract.replaceAdvanceAddresses(oldAddresses, newAddresses);

        const newAdvanceAmount0 = await contract.advances(newAddresses[0]);
        assert.equal(newAdvanceAmount0, advanceAmounts[0]);
    });

    it("test advance amount", async function () {
        const advanceAddresses = [accounts[3], accounts[4]];
        const advanceAmounts = [
            web3.utils.toWei("0.3", "ether"),
            web3.utils.toWei("0.8", "ether"),
        ];

        // 0. set advance setting
        await contract.setAdvanceSetting(advanceAddresses, advanceAmounts);

        // 1. mint artworks
        const tokenID1 = seriesIds[2] * 1000000 + 0;
        const tokenID2 = seriesIds[2] * 1000000 + 1;
        const tokenID3 = seriesIds[3] * 1000000 + 0;
        const tokenID4 = seriesIds[3] * 1000000 + 1;
        const data = [
            [seriesIds[2], tokenID1, contract.address],
            [seriesIds[2], tokenID2, contract.address],
            [seriesIds[3], tokenID3, contract.address],
            [seriesIds[3], tokenID4, contract.address],
        ];
        await contract.mintArtworks(data);

        // 2. buy artworks
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                contract.address,
                [
                    web3.utils.toWei("1", "ether"),
                    "0",
                    expiryTime,
                    accounts[2],
                    [tokenID1, tokenID2, tokenID3, tokenID4],
                    [
                        [
                            [accounts[3], 7000],
                            [COST_RECEIVER, 2000],
                            [accounts[5], 1000], // Curator
                        ],
                        [
                            [accounts[3], 7000],
                            [COST_RECEIVER, 2000],
                            [accounts[5], 1000],
                        ],
                        [
                            [accounts[4], 7000],
                            [COST_RECEIVER, 2000],
                            [accounts[5], 1000],
                        ],
                        [
                            [accounts[4], 7000],
                            [COST_RECEIVER, 2000],
                            [accounts[5], 1000],
                        ],
                    ],
                    false,
                ],
            ]
        );
        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = web3.utils.toDecimal("0x" + sig.slice(128, 130));

        try {
            const acc3BalanceBefore = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceBefore = await web3.eth.getBalance(accounts[4]);
            const acc5BalanceBefore = await web3.eth.getBalance(accounts[5]);
            const accCostReceiverBalanceBefore =
                await web3.eth.getBalance(COST_RECEIVER);

            await contract.startSale();
            await contract.buyArtworks(
                r,
                s,
                v < 2 ? v + 27 : v, // magic 27
                [
                    web3.utils.toWei("1", "ether"),
                    "0",
                    expiryTime,
                    accounts[2],
                    [tokenID1, tokenID2, tokenID3, tokenID4],
                    [
                        [
                            [accounts[3], 7000],
                            [COST_RECEIVER, 2000],
                            [accounts[5], 1000], // Curator
                        ],
                        [
                            [accounts[3], 7000],
                            [COST_RECEIVER, 2000],
                            [accounts[5], 1000],
                        ],
                        [
                            [accounts[4], 7000],
                            [COST_RECEIVER, 2000],
                            [accounts[5], 1000],
                        ],
                        [
                            [accounts[4], 7000],
                            [COST_RECEIVER, 2000],
                            [accounts[5], 1000],
                        ],
                    ],
                    false,
                ],
                { from: accounts[2], value: 1 * 1e18 }
            );
            const ownerOfToken1 = await contract.ownerOf(tokenID1);
            const ownerOfToken2 = await contract.ownerOf(tokenID2);
            const ownerOfToken3 = await contract.ownerOf(tokenID3);
            const ownerOfToken4 = await contract.ownerOf(tokenID4);
            assert.equal(ownerOfToken1, accounts[2]);
            assert.equal(ownerOfToken2, accounts[2]);
            assert.equal(ownerOfToken3, accounts[2]);
            assert.equal(ownerOfToken4, accounts[2]);

            const acc3BalanceAfter = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceAfter = await web3.eth.getBalance(accounts[4]);
            const acc5BalanceAfter = await web3.eth.getBalance(accounts[5]);
            const accCostReceiverBalanceAfter =
                await web3.eth.getBalance(COST_RECEIVER);

            // Revenue per item is 1 / 4 = 0.25
            // Revenues is 0.25 * 2 then deduct advance amount 0.3 => received 0.2 * 70% = 0.14
            assert.equal(
                (
                    BigInt(acc3BalanceAfter) - BigInt(acc3BalanceBefore)
                ).toString(),
                web3.utils.toWei("0.14", "ether")
            );
            // Revenue is 0.25 * 2 then deduct advance amount 0.8 => received 0
            assert.equal(
                (
                    BigInt(acc4BalanceAfter) - BigInt(acc4BalanceBefore)
                ).toString(),
                "0"
            );
            // revenue is 10% of price 0.2 * 10% = 0.02
            assert.equal(
                (
                    BigInt(acc5BalanceAfter) - BigInt(acc5BalanceBefore)
                ).toString(),
                web3.utils.toWei("0.02", "ether")
            );
            // revenue is 20% of price 0.2 = 0.04 and 0.8 advance amount
            assert.equal(
                (
                    BigInt(accCostReceiverBalanceAfter) -
                    BigInt(accCostReceiverBalanceBefore)
                ).toString(),
                web3.utils.toWei("0.84", "ether")
            );
        } catch (err) {
            assert.fail();
        }
    });

    it("test set series renderer", async function () {
        // 1. Test set series renderer
        const chunkSize = 24000;
        const rendererBlobs = [randomString(10), randomString(25000)];
        const seriesID = seriesIds[0];

        for (const rendererBlob of rendererBlobs) {
            // Set the renderer
            const rendererBlobBytes = web3.utils.utf8ToHex(rendererBlob);
            const tx = await contract.setSeriesRenderer(
                seriesID,
                rendererBlobBytes
            );

            // Verify the event
            const { logs } = tx;
            let expectedLog;
            for (const log of logs) {
                if (log.event === "NewSeriesRenderer") {
                    expectedLog = log;
                    break;
                }
            }
            expect(expectedLog.event).not.null;

            let ptrs = Math.floor(rendererBlob.length / chunkSize);
            if (rendererBlob.length % chunkSize !== 0) {
                ptrs++;
            }

            expect(expectedLog.args[0]).to.be.a.bignumber.that.equal(
                new BN(seriesID)
            );
            expect(expectedLog.args[1].length).to.equal(ptrs);

            // Verify the renderer content
            const renderer = await contract.getSeriesRenderer(seriesID);
            assert.equal(renderer, rendererBlob);
        }

        // 2. Test revert due to empty renderer
        await expectCustomError(
            contract.setSeriesRenderer(seriesID, "0x"),
            abi,
            "ErrEmptyBytes",
            []
        );

        // 3. Test unauthorized call
        try {
            await contract.setSeriesRenderer(seriesID, "0x", {
                from: accounts[1],
            });
        } catch (error) {
            expect(error.message).to.include(
                "Ownable: caller is not the owner"
            );
        }

        // 4. Test series does not exist
        await expectCustomError(
            contract.setSeriesRenderer(1000000, "0x"),
            abi,
            "ErrSeriesDoesNotExist",
            [1000000]
        );
    });

    it("test set series name", async function () {
        // 1. Test set series name
        // 1.1. Set series renderer
        const rendererSeriesIDs = [seriesIds[0], seriesIds[1]];
        const nonRendererSeriesIDs = [seriesIds[2], seriesIds[3]];
        const rendererBlobs = [randomString(250), randomString(250)];
        for (let i = 0; i < rendererBlobs.length; i++) {
            const rendererBlobBytes = web3.utils.utf8ToHex(rendererBlobs[i]);
            await contract.setSeriesRenderer(
                rendererSeriesIDs[i],
                rendererBlobBytes
            );
        }

        // 1.2. Set series name
        const names = [randomString(10), randomString(10)];
        const tx = await contract.setSeriesNames(rendererSeriesIDs, names);
        for (let i = 0; i < rendererSeriesIDs.length; i++) {
            expectEvent(tx, "NewSeriesName", {
                seriesId: new BN(rendererSeriesIDs[i]),
                name: names[i],
            });
        }

        // 2. Test get series name
        const name0 = await contract.seriesNames(rendererSeriesIDs[0]);
        const name1 = await contract.seriesNames(rendererSeriesIDs[1]);
        expect(name0).to.equal(names[0]);
        expect(name1).to.equal(names[1]);

        // 3. Test length mismatch
        await expectCustomError(
            contract.setSeriesNames(rendererSeriesIDs, [names[0]]),
            abi,
            "ErrLengthMismatch",
            []
        );

        // 4. Test invalid series ID
        await expectCustomError(
            contract.setSeriesNames([], []),
            abi,
            "ErrEmptyArray",
            []
        );

        // 5. Test series does not exist
        await expectCustomError(
            contract.setSeriesNames([rendererSeriesIDs[0], 1000000], names),
            abi,
            "ErrSeriesDoesNotExist",
            [1000000]
        );

        // 6. Test series has no renderer
        await expectCustomError(
            contract.setSeriesNames(nonRendererSeriesIDs, names),
            abi,
            "ErrSeriesHasNoRenderer",
            []
        );

        // 7. Test unauthorized call
        try {
            await contract.setSeriesNames(rendererSeriesIDs, names, {
                from: accounts[1],
            });
        } catch (error) {
            expect(error.message).to.include(
                "Ownable: caller is not the owner"
            );
        }
    });

    it("test set renderer token data", async function () {
        // 1. Test set renderer token data
        // 1.1. Mint tokens
        const rendererTokenIDs = [1, 2];
        const rendererSeriesID = seriesIds[2];
        const nonRendererTokenIDs = [3, 4];
        const nonRendererSeriesID = seriesIds[3];
        await contract.mintArtworks([
            [rendererSeriesID, rendererTokenIDs[0], contract.address],
            [rendererSeriesID, rendererTokenIDs[1], contract.address],
        ]);
        await contract.mintArtworks([
            [nonRendererSeriesID, nonRendererTokenIDs[0], contract.address],
            [nonRendererSeriesID, nonRendererTokenIDs[1], contract.address],
        ]);

        // 1.2. Set renderer blob
        const rendererBlob = randomString(250);
        const rendererBlobBytes = web3.utils.utf8ToHex(rendererBlob);
        await contract.setSeriesRenderer(rendererSeriesID, rendererBlobBytes);

        // 1.3. Set renderer token data
        const data = [
            [
                "https://example.com/image-0.png",
                "https://example.com/texture-0.json",
                new BN(0),
            ],
            [
                "https://example.com/image-1.png",
                "https://example.com/texture-1.json",
                new BN(1),
            ],
        ];
        const tx = await contract.setRendererTokenData(rendererTokenIDs, data);
        const { logs } = tx;
        let expectedLogs = [];
        for (let i = 0; i < rendererTokenIDs.length; i++) {
            if (logs[i].event === "NewRendererTokenData") {
                expectedLogs.push(logs[i]);
            }
        }
        expect(expectedLogs.length).to.equal(rendererTokenIDs.length);
        for (let i = 0; i < expectedLogs.length; i++) {
            expect(expectedLogs[i].args[0]).to.be.a.bignumber.that.equal(
                new BN(rendererTokenIDs[i])
            );
            expect(expectedLogs[i].args[1].imageURI).to.equal(data[i][0]);
            expect(expectedLogs[i].args[1].textureURI).to.equal(data[i][1]);
            expect(expectedLogs[i].args[1].index).to.be.a.bignumber.that.equal(
                new BN(data[i][2])
            );
        }

        // 2. Test get renderer token data
        for (let i = 0; i < rendererTokenIDs.length; i++) {
            const rendererTokenData = await contract.rendererTokenData(
                rendererTokenIDs[i]
            );
            expect(rendererTokenData.imageURI).to.equal(data[i][0]);
            expect(rendererTokenData.textureURI).to.equal(data[i][1]);
            expect(rendererTokenData.index).to.be.a.bignumber.that.equal(
                new BN(data[i][2])
            );
        }

        // 3. Test unauthorized call
        try {
            await contract.setRendererTokenData(rendererTokenIDs, data, {
                from: accounts[1],
            });
        } catch (error) {
            expect(error.message).to.include(
                "Ownable: caller is not the owner"
            );
        }

        // 4. Test length mismatch
        await expectCustomError(
            contract.setRendererTokenData(rendererTokenIDs, [data[0]]),
            abi,
            "ErrLengthMismatch",
            []
        );

        // 5. Test invalid token ID
        await expectCustomError(
            contract.setRendererTokenData([99999, 100000], data),
            abi,
            "ErrTokenDoesNotExist",
            [99999]
        );

        // 6. Test series has no renderer
        await expectCustomError(
            contract.setRendererTokenData(nonRendererTokenIDs, data),
            abi,
            "ErrSeriesHasNoRenderer",
            []
        );
    });

    it("test tokenURI", async function () {
        // 1. Test tokenURI for non-renderer token
        // 1.1 Mint token
        const nonRendererSeriesID = seriesIds[0];
        const nonRendererTokenID = 0;
        await contract.mintArtworks([
            [nonRendererSeriesID, nonRendererTokenID, contract.address],
        ]);

        // 1.2 Test tokenURI
        const baseTokenURI = "ipfs://baseTokenCID";
        await contract.setTokenBaseURI(baseTokenURI);
        let tokenURI = await contract.tokenURI(nonRendererTokenID);
        expect(tokenURI).to.equal(`${baseTokenURI}/${nonRendererTokenID}`);

        // 2. Test tokenURI for renderer token
        // 2.1 Mint token
        const rendererSeriesID = seriesIds[1];
        const rendererTokenID = 1;
        await contract.mintArtworks([
            [rendererSeriesID, rendererTokenID, contract.address],
        ]);

        // 2.2 Set renderer blob
        const placeholder = "{{TEXTURE_JSON}}";
        const rendererBlob =
            randomString(100) + placeholder + randomString(100);
        const rendererBlobBytes = web3.utils.utf8ToHex(rendererBlob);
        await contract.setSeriesRenderer(rendererSeriesID, rendererBlobBytes);

        // 2.3 Set renderer series name
        const name = "Test Series";
        await contract.setSeriesNames([rendererSeriesID], [name]);

        // 2.4 Set renderer token data
        const textureURI = "https://example.com/texture-0.json";
        const imageURI = "https://example.com/image-0.png";
        const index = new BN(0);
        const data = [[imageURI, textureURI, new BN(0)]];
        await contract.setRendererTokenData([rendererTokenID], data);

        // 2.5 Test tokenURI
        tokenURI = await contract.tokenURI(rendererTokenID);
        const expectedJson = `{"animation_url":"data:text/html;base64,${Base64.encode(rendererBlob.replace(placeholder, textureURI))}","image":"${imageURI}","name":"${name} #${index.toString()}"}`;
        expect(tokenURI).to.equal(
            `data:application/json;base64,${Base64.encode(expectedJson)}`
        );
    });
});
