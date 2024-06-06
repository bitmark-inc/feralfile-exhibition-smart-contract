const FeralfileExhibitionV4_2 = artifacts.require("FeralfileExhibitionV4_2");
const FeralfileVault = artifacts.require("FeralfileVault");

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";
const COST_RECEIVER = "0x46f2B641d8702f29c45f6D06292dC34Eb9dB1801";
const VAULT_ADDRESS = "0x7a15b36cb834aea88553de69077d3777460d73ac";
const CONTRACT_URI = "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";

contract("FeralfileExhibitionV4_2", async (accounts) => {
    before(async function () {
        this.signer = accounts[0];
        this.vault = await FeralfileVault.new(this.signer);
        this.seriesIds = [0, 1, 2, 3, 4];
        this.seriesMaxSupply = [1, 1, 100, 1000, 10000];
        this.seriesNextPurchasableTokenIds = [1, 1000001, 2000001, 3000001, 4000001]
        // Deploy multiple contracts
        this.contracts = [];
        for (let i = 0; i < 12; i++) {
            let contract = await FeralfileExhibitionV4_2.new(
                "Feral File V4 Test",
                "FFv4",
                true,
                true,
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                CONTRACT_URI,
                this.seriesIds,
                this.seriesMaxSupply,
                this.seriesNextPurchasableTokenIds,
            );
            this.contracts.push(contract);
        }
    });

    it("test contract constructor", async function () {
        const contract = this.contracts[0];

        const burnable = await contract.burnable();
        assert.ok(burnable);

        const bridgeable = await contract.bridgeable();
        assert.ok(bridgeable);

        const mintable = await contract.mintable();
        assert.ok(mintable);

        const vaultAddress = await contract.vault();
        assert.equal(vaultAddress, this.vault.address);

        const signer = await contract.signer();
        assert.equal(signer, signer);

        const seriesMaxSupply1 = await contract.seriesMaxSupply(
            this.seriesIds[0]
        );
        assert.equal(seriesMaxSupply1, this.seriesMaxSupply[0]);

        const seriesMaxSupply2 = await contract.seriesMaxSupply(
            this.seriesIds[1]
        );
        assert.equal(seriesMaxSupply2, this.seriesMaxSupply[1]);

        const seriesTotalSupply1 = await contract.seriesTotalSupply(
            this.seriesIds[0]
        );
        assert.equal(seriesTotalSupply1, 0);
    });

    it("test duplicate series in constructor", async function () {
        // Deploy contract with duplicate series defined
        let seriesIds = [0, 1, 2, 3, 1];
        let seriesMaxSupply = [1, 1, 100, 1000, 10000];
        let seriesNextPurchasableTokenIds = [1, 1000001, 2000001, 3000001, 1000001]
        try {
            await FeralfileExhibitionV4_2.new(
                "Feral File V4 Test",
                "FFv4",
                true,
                true,
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                CONTRACT_URI,
                seriesIds,
                seriesMaxSupply,
                seriesNextPurchasableTokenIds
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
        const contract = this.contracts[0];

        // 1. mint successful
        const owner = accounts[1];
        const tokenId1 = this.seriesIds[3] * 1000000 + 0;
        const tokenId2 = this.seriesIds[4] * 1000000 + 0;
        let data = [
            [this.seriesIds[2], tokenId1, owner],
            [this.seriesIds[3], tokenId2, owner],
        ];
        const tx = await contract.mintArtworks(data);

        // total supply
        const totalSupply = await contract.totalSupply();
        assert.equal(totalSupply, 2);

        // series total supply
        const totalSupply1 = await contract.seriesTotalSupply(
            this.seriesIds[2]
        );
        const totalSupply2 = await contract.seriesTotalSupply(
            this.seriesIds[3]
        );
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
        assert.equal(artwork.seriesId, this.seriesIds[2]);

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
        data = [[this.seriesIds[3], 999999, ZERO_ADDRESS]];
        try {
            await contract.mintArtworks(data);
        } catch (error) {
            assert.ok(
                error.message.includes("ERC721: mint to the zero address")
            );
        }

        // token already minted
        data = [[this.seriesIds[3], tokenId1, accounts[2]]];
        try {
            await contract.mintArtworks(data);
        } catch (error) {
            assert.ok(error.message.includes("ERC721: token already minted"));
        }

        // no more slots
        data = [[this.seriesIds[2], 999999, accounts[2]]];
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
        const contract = this.contracts[1];

        // 1. Prepare data
        const owner = accounts[1];
        const tokenId1 = this.seriesIds[0] * 1000000 + 0;
        const tokenId2 = this.seriesIds[1] * 1000000 + 0;
        const tokenId3 = this.seriesIds[2] * 1000000 + 0;
        let mintData = [
            [this.seriesIds[0], tokenId1, owner],
            [this.seriesIds[1], tokenId2, owner],
            [this.seriesIds[2], tokenId3, owner],
        ];
        let tx = await contract.mintArtworks(mintData);

        // 2. Burn successfully
        let burnData = [tokenId1, tokenId2];
        tx = await contract.burnArtworks(burnData, { from: owner });

        // total supply
        const totalSupply = await contract.totalSupply();
        assert.equal(totalSupply, 1);

        // series total supply
        const totalSupply1 = await contract.seriesTotalSupply(
            this.seriesIds[0]
        );
        const totalSupply2 = await contract.seriesTotalSupply(
            this.seriesIds[1]
        );
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

    it("test start/stop and burn, pause/resume sale", async function () {
        const contract = this.contracts[4];

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
            [this.seriesIds[4], tokenIDs[0], owner],
            [this.seriesIds[4], tokenIDs[1], owner],
            [this.seriesIds[4], tokenIDs[2], owner],
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
            console.log(error);
            assert.fail();
        }

        // Stop sale and burn remaining tokens
        try {
            await contract.stopSaleAndBurn();
        } catch (error) {
            console.log(error);
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
        const contract = this.contracts[5];

        // Mint new tokens
        const tokenIDs = [3000997, 3000998, 4000998, 4000999];
        await contract.mintArtworks([
            [this.seriesIds[3], tokenIDs[0], contract.address],
            [this.seriesIds[3], tokenIDs[1], contract.address],
            [this.seriesIds[4], tokenIDs[2], contract.address],
            [this.seriesIds[4], tokenIDs[3], contract.address],
        ]);

        // Start the sale
        await contract.startSale();

        const owner1 = accounts[1];
        const owner2 = accounts[2];

        // 1. Stop the sale and transfer remaining tokens failed
        try {
            await contract.stopSaleAndTransfer(
                [this.seriesIds[3], this.seriesIds[4]],
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
                [this.seriesIds[1], this.seriesIds[2]],
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
                [this.seriesIds[1], this.seriesIds[2]],
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
                [this.seriesIds[3], this.seriesIds[4]],
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
        const contract = this.contracts[6];

        // mint tokens
        const owner1 = accounts[1];
        const owner2 = accounts[2];
        const tokenId1 = this.seriesIds[3] * 1000000 + 555;
        const tokenId2 = this.seriesIds[3] * 1000000 + 556;
        const tokenId3 = this.seriesIds[4] * 1000000 + 555;
        const tokenId4 = this.seriesIds[4] * 1000000 + 556;
        let data = [
            [this.seriesIds[3], tokenId1, owner1],
            [this.seriesIds[3], tokenId2, owner1],
            [this.seriesIds[4], tokenId3, owner2],
            [this.seriesIds[4], tokenId4, owner2],
        ];

        try {
            await contract.mintArtworks(data);
        } catch (error) {
            console.log(error);
            assert.fail();
        }

        // transfer tokens
        // 1. Transfer to other address rather than contract
        try {
            await contract.transferFrom(owner1, owner2, tokenId1, {
                from: owner1,
            });
        } catch (error) {
            console.log(error);
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
        const contract = this.contracts[7];
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
            assert.equal(advanceAmount0, advanceAmounts[1]);
        } catch (error) {
            console.log(error);
            assert.fail();
        }
    });

    it("test set advance setting failed because address in use", async function () {
        const contract = this.contracts[8];
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
            console.log(error);
            assert.equal(error.reason, "Custom error (could not decode)");
        }
    });

    it("test replace advance addresses successfully", async function () {
        const contract = this.contracts[9];
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

    it("should buy bulk artworks successfully", async function () {
        const contract = this.contract[3];

        // Mint artworks to the contract
        const owner = contract.address;
        await contract.mintArtworks([
            [this.seriesIds[3], 3000000, owner],
            [this.seriesIds[3], 3000001, owner],
            [this.seriesIds[4], 4000000, owner],
            [this.seriesIds[4], 4000001, owner],
            [this.seriesIds[4], 4000002, owner],
        ]);

        // Generate signature
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const saleData = {
            price: BigInt(0.25 * 1e18).toString(),
            cost: BigInt(0.02 * 1e18).toString(),
            expiryTime,
            destination: accounts[2],
            nonce: 0,
            seriesID: this.seriesIds[3],
            quantity: 2,
            revenueShares: [
                { recipient: accounts[3], bps: 8000 },
                { recipient: accounts[4], bps: 2000 }
            ],
            payByVaultContract: false
        };

        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256,uint256,tuple(address,uint256)[],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                contract.address,
                saleData
            ]
        );

        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);

        try {
            const acc3BalanceBefore = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceBefore = await web3.eth.getBalance(accounts[4]);
            const accCostReceiverBalanceBefore = await web3.eth.getBalance(COST_RECEIVER);

            await contract.startSale();
            await contract.buyBulkArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27, // magic 27
                saleData,
                { from: accounts[5], value: 0.25 * 1e18 }
            );

            // Validate new owner of tokens
            const ownerOfToken1 = await contract.ownerOf(3000000);
            const ownerOfToken2 = await contract.ownerOf(3000001);
            assert.equal(ownerOfToken1, accounts[2]);
            assert.equal(ownerOfToken2, accounts[2]);

            // Validate balances
            const acc3BalanceAfter = await web3.eth.getBalance(accounts[3]);
            const acc4BalanceAfter = await web3.eth.getBalance(accounts[4]);
            const accCostReceiverBalanceAfter = await web3.eth.getBalance(COST_RECEIVER);

            assert.equal(
                (BigInt(acc3BalanceAfter) - BigInt(acc3BalanceBefore)).toString(),
                BigInt((0.23 * 1e18 * 80) / 100).toString()
            );
            assert.equal(
                (BigInt(acc4BalanceAfter) - BigInt(acc4BalanceBefore)).toString(),
                BigInt((0.23 * 1e18 * 20) / 100).toString()
            );
            assert.equal(
                (BigInt(accCostReceiverBalanceAfter) - BigInt(accCostReceiverBalanceBefore)).toString(),
                BigInt(0.02 * 1e18).toString()
            );
        } catch (err) {
            console.log(err);
            assert.fail();
        }
    });

    it("should fail to buy bulk artworks with invalid signature", async function () {
        const contract = this.contract[10];

        // Mint artworks to the contract
        const owner = contract.address;
        await contract.mintArtworks([
            [this.seriesIds[3], 3000000, owner],
            [this.seriesIds[3], 3000001, owner],
            [this.seriesIds[4], 4000000, owner],
            [this.seriesIds[4], 4000001, owner],
            [this.seriesIds[4], 4000002, owner],
        ]);

        // Generate signature
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const saleData = {
            price: BigInt(0.25 * 1e18).toString(),
            cost: BigInt(0.02 * 1e18).toString(),
            expiryTime,
            destination: accounts[2],
            nonce: 0,
            seriesID: this.seriesIds[3],
            quantity: 2,
            revenueShares: [
                { recipient: accounts[3], bps: 8000 },
                { recipient: accounts[4], bps: 2000 }
            ],
            payByVaultContract: false
        };

        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256,uint256,tuple(address,uint256)[],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                contract.address,
                saleData
            ]
        );

        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, accounts[1]); // Invalid signer
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);

        try {
            await contract.startSale();
            await contract.buyBulkArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27, // magic 27
                saleData,
                { from: accounts[5], value: 0.25 * 1e18 }
            );
            assert.fail("Expected revert not received");
        } catch (error) {
            assert.ok(error.message.includes("FeralfileExhibitionV4: invalid signature"));
        }
    });

    it("should fail to buy bulk artworks with incorrect payment", async function () {
        const contract = this.contract[11];

        // Mint artworks to the contract
        const owner = contract.address;
        await contract.mintArtworks([
            [this.seriesIds[3], 3000000, owner],
            [this.seriesIds[3], 3000001, owner],
            [this.seriesIds[4], 4000000, owner],
            [this.seriesIds[4], 4000001, owner],
            [this.seriesIds[4], 4000002, owner],
        ]);

        // Generate signature
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const saleData = {
            price: BigInt(0.25 * 1e18).toString(),
            cost: BigInt(0.02 * 1e18).toString(),
            expiryTime,
            destination: accounts[2],
            nonce: 0,
            seriesID: this.seriesIds[3],
            quantity: 2,
            revenueShares: [
                { recipient: accounts[3], bps: 8000 },
                { recipient: accounts[4], bps: 2000 }
            ],
            payByVaultContract: false
        };

        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256,uint256,tuple(address,uint256)[],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                contract.address,
                saleData
            ]
        );

        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);

        try {
            await contract.startSale();
            await contract.buyBulkArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27, // magic 27
                saleData,
                { from: accounts[5], value: 0.2 * 1e18 } // Incorrect payment amount
            );
            assert.fail("Expected revert not received");
        } catch (error) {
            assert.ok(error.message.includes("FeralfileExhibitionV4: invalid payment amount"));
        }
    });

    it("should fail to buy bulk artworks with invalid series ID", async function () {
        const contract = this.contract;

        // Mint artworks to the contract
        const owner = contract.address;
        await contract.mintArtworks([
            [this.seriesIds[3], 3000000, owner],
            [this.seriesIds[3], 3000001, owner],
            [this.seriesIds[4], 4000000, owner],
            [this.seriesIds[4], 4000001, owner],
            [this.seriesIds[4], 4000002, owner],
        ]);

        // Generate signature
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const saleData = {
            price: BigInt(0.25 * 1e18).toString(),
            cost: BigInt(0.02 * 1e18).toString(),
            expiryTime,
            destination: accounts[2],
            nonce: 0,
            seriesID: 999999, // Invalid series ID
            quantity: 2,
            revenueShares: [
                { recipient: accounts[3], bps: 8000 },
                { recipient: accounts[4], bps: 2000 }
            ],
            payByVaultContract: false
        };

        const signParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256,uint256,tuple(address,uint256)[],bool)",
            ],
            [
                BigInt(await web3.eth.getChainId()).toString(),
                contract.address,
                saleData
            ]
        );

        const hash = web3.utils.keccak256(signParams);
        var sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        const r = "0x" + sig.slice(0, 64);
        const s = "0x" + sig.slice(64, 128);
        const v = "0x" + sig.slice(128, 130);

        try {
            await contract.startSale();
            await contract.buyBulkArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27, // magic 27
                saleData,
                { from: accounts[5], value: 0.25 * 1e18 }
            );
            assert.fail("Expected revert not received");
        } catch (error) {
            assert.ok(error.message.includes("TokenIDNotFound"));
        }
    });
});
