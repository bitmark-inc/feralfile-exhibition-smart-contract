const FeralfileExhibitionV5 = artifacts.require("FeralfileExhibitionV5");
const FeralfileVault = artifacts.require("FeralfileVaultV2");

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";
const COST_RECEIVER = "0x46f2B641d8702f29c45f6D06292dC34Eb9dB1801";
const VAULT_ADDRESS = "0x9c90C920370D4Cd0097786d25812b2aEab5e9E3c";
const CONTRACT_URI = "ipfs://QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc";
const TOKEN_URI = "ipfs://QmZt5r6bU7r3BvCp6b6Z1yYK5Hd9M9WnQV2g8VfNnMmM3T/{id}";

contract("FeralfileExhibitionV5_0", async (accounts) => {
    before(async function () {
        this.owner = accounts[0];
        this.trustee = accounts[1];
        this.signer = accounts[2];
        this.vault = await FeralfileVault.new(this.signer);
        this.seriesIds = [1, 2, 3, 4, 5];
        this.seriesMaxSupply = [1, 5, 10, 1, 1];
        this.seriesArtworkMaxSupply = [1, 1, 1, 100, 100];
        this.contracts = [];

        // Deploy multiple contracts
        for (let i = 0; i < 7; i++) {
            let contract = await FeralfileExhibitionV5.new(
                TOKEN_URI,
                CONTRACT_URI,
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                true,
                this.seriesIds,
                this.seriesMaxSupply,
                this.seriesArtworkMaxSupply
            );
            await contract.addTrustee(this.trustee);
            this.contracts.push(contract);
        }
    });

    it("test contract constructor", async function () {
        for (let i = 0; i < this.contracts.length; i++) {
            let contract = this.contracts[i];

            // burnable
            const burnable = await contract.burnable();
            assert.ok(burnable);

            // mintable
            const mintable = await contract.mintable();
            assert.ok(mintable);

            // selling
            const selling = await contract.selling();
            assert.ok(!selling);

            // vault
            const vaultAddress = await contract.vault();
            assert.equal(this.vault.address, vaultAddress);

            // signer
            const signer = await contract.signer();
            assert.equal(this.signer, signer);

            // series max supply
            const seriesMaxSupply1 = await contract.seriesMaxSupply(
                this.seriesIds[0]
            );
            assert.equal(seriesMaxSupply1, this.seriesMaxSupply[0]);

            const seriesMaxSupply2 = await contract.seriesMaxSupply(
                this.seriesIds[1]
            );
            assert.equal(seriesMaxSupply2, this.seriesMaxSupply[1]);

            // series total supply
            const seriesTotalSupply = await contract.seriesTotalSupply(
                this.seriesIds[0]
            );
            assert.equal(seriesTotalSupply, 0);

            // artwork total supply
            const artworkTotalSupply = await contract.artworkTotalSupply(
                this.seriesIds[4]
            );
            assert.equal(artworkTotalSupply, 0);

            // series artwork max supply
            const seriesArtworkMaxSupply =
                await contract.seriesArtworkMaxSupply(this.seriesIds[4]);
            assert.equal(seriesArtworkMaxSupply, 100);
        }
    });

    it("test failed constructor", async function () {
        const data = [
            [
                "",
                CONTRACT_URI,
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                true,
                [1, 2, 3],
                [1, 1, 1],
                [1, 1, 1],
                "FeralfileExhibitionV5: baseTokenURI_ is empty",
            ], // empty base token URI
            [
                TOKEN_URI,
                "",
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                true,
                [1, 2, 3],
                [1, 1, 1],
                [1, 1, 1],
                "FeralfileExhibitionV5: contractURI_ is empty",
            ], // empty contract URI
            [
                TOKEN_URI,
                CONTRACT_URI,
                ZERO_ADDRESS,
                this.vault.address,
                COST_RECEIVER,
                true,
                [1, 2, 3],
                [1, 1, 1],
                [1, 1, 1],
                "ECDSASign: signer_ is zero address",
            ], // zero signer
            [
                TOKEN_URI,
                CONTRACT_URI,
                this.signer,
                ZERO_ADDRESS,
                COST_RECEIVER,
                true,
                [1, 2, 3],
                [1, 1, 1],
                [1, 1, 1],
                "FeralfileExhibitionV5: vault_ is zero address",
            ], // zero vault address
            [
                TOKEN_URI,
                CONTRACT_URI,
                this.signer,
                this.vault.address,
                ZERO_ADDRESS,
                true,
                [1, 2, 3],
                [1, 1, 1],
                [1, 1, 1],
                "FeralfileExhibitionV5: costReceiver_ is zero address",
            ], // zero cost receiver
            [
                TOKEN_URI,
                CONTRACT_URI,
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                true,
                [0, 1, 2],
                [1, 1, 1],
                [1, 1, 1],
                "FeralfileExhibitionV5: zero seriesId",
            ], // zero series ID
            [
                TOKEN_URI,
                CONTRACT_URI,
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                true,
                [1, 2, 3],
                [1, 0, 1],
                [1, 1, 1],
                "FeralfileExhibitionV5: zero series max supply",
            ], // zero series max supply
            [
                TOKEN_URI,
                CONTRACT_URI,
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                true,
                [1, 2, 3],
                [1, 1, 1],
                [0, 1, 1],
                "FeralfileExhibitionV5: zero artwork max supply",
            ], // zero artwork max supply
            [
                TOKEN_URI,
                CONTRACT_URI,
                this.signer,
                this.vault.address,
                COST_RECEIVER,
                true,
                [1, 2, 1],
                [1, 1, 1],
                [1, 1, 1],
                "FeralfileExhibitionV5: duplicate seriesId",
            ], // duplicate series ID
        ];

        for (let i = 0; i < data.length; i++) {
            try {
                await FeralfileExhibitionV5.new(
                    data[i][0],
                    data[i][1],
                    data[i][2],
                    data[i][3],
                    data[i][4],
                    data[i][5],
                    data[i][6],
                    data[i][7],
                    data[i][8]
                );
            } catch (error) {
                assert.equal(data[i][9], error.reason);
            }
        }
    });

    it("test mint artwork", async function () {
        const contract = this.contracts[0];
        const owner = accounts[1];
        const seriesId = this.seriesIds[0];

        // 1. unauthorized call
        try {
            await contract.mintArtworks([[seriesId, 1, owner, 1]], {
                from: accounts[2],
            });
        } catch (error) {
            assert.equal(
                "VM Exception while processing transaction: revert",
                error.message
            );
        }

        // 2. pre-condition check failed
        const data = [
            [
                [0, 1, owner, 1],
                "FeralfileExhibitionV5: seriesId doesn't exist: 0",
            ], // series ID doesn't exist
            [
                [seriesId, 0, owner, 1],
                "FeralfileExhibitionV5: token ID is zero",
            ], // token ID is zero
            [
                [seriesId, 1, ZERO_ADDRESS, 1],
                "ERC1155: mint to the zero address",
            ], // mint to zero address
            [[seriesId, 1, owner, 0], "FeralfileExhibitionV5: amount is zero"], // amount is zero
            [
                [seriesId, 1, owner, 2],
                "FeralfileExhibitionV5: no more artwork slots available",
            ], // exceed artwork max supply
        ];
        for (let i = 0; i < data.length; i++) {
            try {
                await contract.mintArtworks([data[i][0]]);
            } catch (error) {
                assert.equal(data[i][1], error.reason);
            }
        }

        const testData = [
            [this.seriesIds[0], 1, accounts[1], 1], // unique artwork
            [this.seriesIds[4], 2, accounts[2], 50], // edition
        ];

        for (let i = 0; i < testData.length; i++) {
            const seriesId = testData[i][0];
            const tokenId = testData[i][1];
            const owner = testData[i][2];
            const amount = testData[i][3];

            // 3. mint successfully
            const tx = await contract.mintArtworks([
                [seriesId, tokenId, owner, amount],
            ]);

            // check artworkOf()
            const artworks = await contract.artworkOf(owner);
            assert.equal(artworks.length, 1);
            assert.equal(artworks[0].seriesId, seriesId);
            assert.equal(artworks[0].tokenId, tokenId);
            assert.equal(artworks[0].amount, amount);

            // check seriesTotalSupply()
            const seriesTotalSupply = await contract.seriesTotalSupply(
                seriesId
            );
            assert.equal(seriesTotalSupply, 1);

            // check artworkTotalSupply()
            const artworkTotalSupply = await contract.artworkTotalSupply(
                tokenId
            );
            assert.equal(artworkTotalSupply, amount);

            // check balanceOf()
            const balance = await contract.balanceOf(owner, tokenId);
            assert.equal(balance, amount);

            // get artwork
            const artwork = await contract.getArtwork(tokenId);
            assert.equal(artwork.tokenId, tokenId);
            assert.equal(artwork.seriesId, seriesId);
            assert.equal(artwork.amount, amount);

            // event emitted
            const { logs } = tx;
            assert.ok(Array.isArray(logs));
            assert.equal(logs.length, 2);

            const log0 = logs[0];
            const log1 = logs[1];
            assert.equal(log0.event, "TransferSingle");
            assert.equal(log1.event, "NewArtwork");

            // 4. mint failed due to series max supply
            try {
                await contract.mintArtworks([
                    [seriesId, tokenId * 123, owner, 1],
                ]);
            } catch (error) {
                assert.equal(
                    "FeralfileExhibitionV5: no more series slots available",
                    error.reason
                );
            }

            // 5. mint failed due to artwork max supply
            try {
                await contract.mintArtworks([
                    [seriesId, tokenId, owner, amount * 123],
                ]);
            } catch (error) {
                assert.equal(
                    "FeralfileExhibitionV5: no more artwork slots available",
                    error.reason
                );
            }
        }
    });

    it("test burn artwork", async function () {
        const contract = this.contracts[1];
        const owner = accounts[1];

        // 1. pre-condition check failed
        const data = [
            [
                [ZERO_ADDRESS, 1, 1],
                "FeralfileExhibitionV5: from is zero address",
            ], // from is zero address
            [[owner, 0, 1], "FeralfileExhibitionV5: token ID is zero"], // token ID is zero
            [[owner, 1, 0], "FeralfileExhibitionV5: amount is zero"], // amount is zero
            [[owner, 1, 1], "FeralfileExhibitionV5: artwork doesn't exist"], // exceed artwork max supply
        ];
        for (let i = 0; i < data.length; i++) {
            try {
                await contract.burnArtworks([data[i][0]]);
            } catch (error) {
                assert.equal(data[i][1], error.reason);
            }
        }

        // 2. burn successfully
        const testData = [
            [this.seriesIds[0], 1, accounts[1], 1], // unique artwork
            [this.seriesIds[4], 2, accounts[2], 50], // edition
        ];
        for (let i = 0; i < testData.length; i++) {
            const seriesId = testData[i][0];
            const tokenId = testData[i][1];
            const owner = testData[i][2];
            const amount = testData[i][3];
            const unique = amount == 1;

            // mint artworks
            await contract.mintArtworks([[seriesId, tokenId, owner, amount]]);

            // check artworkOf()
            let artworks = await contract.artworkOf(owner);
            assert.equal(artworks.length, 1);
            assert.equal(artworks[0].seriesId, seriesId);
            assert.equal(artworks[0].tokenId, tokenId);
            assert.equal(artworks[0].amount, amount);

            // check seriesTotalSupply()
            let seriesTotalSupply = await contract.seriesTotalSupply(seriesId);
            assert.equal(seriesTotalSupply, 1);

            // check artworkTotalSupply()
            let artworkTotalSupply = await contract.artworkTotalSupply(tokenId);
            assert.equal(artworkTotalSupply, amount);

            // check balanceOf()
            let balance = await contract.balanceOf(owner, tokenId);
            assert.equal(balance, amount);

            // get artwork
            let artwork = await contract.getArtwork(tokenId);
            assert.equal(artwork.tokenId, tokenId);
            assert.equal(artwork.seriesId, seriesId);
            assert.equal(artwork.amount, amount);

            // burn artwork
            let burnAmount = unique ? 1 : amount / 2; // burn 1 for unique artwork, burn half for edition
            const tx = await contract.burnArtworks([
                [owner, tokenId, burnAmount],
            ]);

            // check artworkOf()
            artworks = await contract.artworkOf(owner);
            let artworkLen = unique ? 0 : 1;
            assert.equal(artworks.length, artworkLen);

            // check seriesTotalSupply()
            seriesTotalSupply = await contract.seriesTotalSupply(seriesId);
            let sts = unique ? 0 : 1;
            assert.equal(seriesTotalSupply, sts);

            // check artworkTotalSupply()
            artworkTotalSupply = await contract.artworkTotalSupply(tokenId);
            let ats = unique ? 0 : amount - burnAmount;
            assert.equal(artworkTotalSupply, ats);

            // check balanceOf()
            balance = await contract.balanceOf(owner, tokenId);
            let b = unique ? 0 : amount - burnAmount;
            assert.equal(balance, b);

            // get artwork
            artwork = await contract.getArtwork(tokenId);
            let ti = unique ? 0 : tokenId;
            assert.equal(artwork.tokenId, ti);

            // event emitted
            const { logs } = tx;
            assert.ok(Array.isArray(logs));
            assert.equal(logs.length, 2);

            const log0 = logs[0];
            const log1 = logs[1];
            assert.equal(log0.event, "TransferSingle");
            assert.equal(log1.event, "BurnArtwork");
        }
    });

    it("test buy artworks not ok", async function () {
        const contract = this.contracts[2];
        const price = 0.25 * 1e18;
        const cost = 0.02 * 1e18;
        const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
        const recipient = accounts[3];
        const buyer = accounts[2];
        const seriesIds = [this.seriesIds[0], this.seriesIds[3]];
        const tokenIds = [1, 2];
        const biddingUnixNano = BigInt(new Date().getTime() * 1e6).toString();
        const royaltyShares = [
            [
                [accounts[4], 8000],
                [accounts[5], 2000],
            ],
            [
                [accounts[4], 8000],
                [accounts[5], 2000],
            ],
        ];
        const chainId = await web3.eth.getChainId();
        const saleData = [
            BigInt(price).toString(),
            BigInt(cost).toString(),
            expiryTime,
            recipient,
            tokenIds,
            royaltyShares,
            false,
            biddingUnixNano,
        ];

        // 1. the sale is not started yet
        try {
            await contract.buyArtworks("0x0", "0x0", 0 + 27, saleData, {
                from: buyer,
                value: price,
            });
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: sale is not started",
                error.reason
            );
        }

        // mint artworks and start sale
        await contract.mintArtworks([
            [seriesIds[0], tokenIds[0], contract.address, 1],
            [seriesIds[1], tokenIds[1], contract.address, 100],
        ]);
        await contract.startSale();

        // 2. empty token IDs
        saleData[4] = [];
        try {
            await contract.buyArtworks("0x0", "0x0", 0 + 27, saleData, {
                from: buyer,
                value: price,
            });
        } catch (error) {
            assert.equal("FeralfileSaleData: tokenIds is empty", error.reason);
        }
        saleData[4] = tokenIds;

        // 3. tokenIds and royaltyShares length mismatch
        saleData[4] = [1];
        try {
            await contract.buyArtworks("0x0", "0x0", 0 + 27, saleData, {
                from: buyer,
                value: price,
            });
        } catch (error) {
            assert.equal(
                "FeralfileSaleData: tokenIds and revenueShares length mismatch",
                error.reason
            );
        }
        saleData[4] = tokenIds;

        // 4. expired sale
        saleData[2] = (new Date().getTime() / 1000 - 300).toFixed(0);
        try {
            await contract.buyArtworks("0x0", "0x0", 0 + 27, saleData, {
                from: buyer,
                value: price,
            });
        } catch (error) {
            assert.equal("FeralfileSaleData: sale is expired", error.reason);
        }
        saleData[2] = expiryTime;

        // 5. invalid signature
        try {
            await contract.buyArtworks("0x0", "0x0", 0 + 27, saleData, {
                from: buyer,
                value: price,
            });
        } catch (error) {
            assert.equal("ECDSA: invalid signature", error.reason);
        }

        // 6. wrong signature
        // encode params
        let signedParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool,uint256)",
            ],
            [BigInt(chainId).toString(), contract.address, saleData]
        );

        // sign params
        let hash = web3.utils.keccak256(signedParams);
        let sig = await web3.eth.sign(hash, accounts[0]);
        sig = sig.substr(2);
        let r = "0x" + sig.slice(0, 64);
        let s = "0x" + sig.slice(64, 128);
        let v = "0x" + sig.slice(128, 130);

        try {
            await contract.buyArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27,
                saleData,
                {
                    from: buyer,
                    value: price,
                }
            );
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: invalid signature",
                error.reason
            );
        }

        // 7. over max bps
        saleData[5] = [
            [
                [accounts[4], 8000],
                [accounts[5], 2001],
            ],
            [
                [accounts[4], 8000],
                [accounts[5], 2000],
            ],
        ];

        // encode params
        signedParams = web3.eth.abi.encodeParameters(
            [
                "uint",
                "address",
                "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool,uint256)",
            ],
            [BigInt(chainId).toString(), contract.address, saleData]
        );

        // sign params
        hash = web3.utils.keccak256(signedParams);
        sig = await web3.eth.sign(hash, this.signer);
        sig = sig.substr(2);
        r = "0x" + sig.slice(0, 64);
        s = "0x" + sig.slice(64, 128);
        v = "0x" + sig.slice(128, 130);
        try {
            await contract.buyArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27,
                saleData,
                {
                    from: buyer,
                    value: price,
                }
            );
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: total bps over 10,000",
                error.reason
            );
        }
        saleData[5] = royaltyShares;
    });

    it("test buy artworks ok", async function () {
        const testData = [
            [
                [1, 2], // token IDs
                [this.seriesIds[0], this.seriesIds[3]], // series ID, both unique and edition
                [1, 100], // amounts
                accounts[2], // buyer
                accounts[3], // recipient
                [0.8, 0.2], // royalties shares
                [accounts[4], accounts[5]], // royalty payees
                true, // with funds
                this.contracts[3],
            ],
            [
                [3, 4], // token IDs
                [this.seriesIds[1], this.seriesIds[4]], // series ID, both unique and edition
                [1, 100], // amounts
                accounts[6], // buyer
                accounts[7], // recipient
                [0.8, 0.2], // royalties shares
                [accounts[8], accounts[9]], // royalty payees
                false, // without funds
                this.contracts[4],
            ],
        ];

        for (let i = 0; i < testData.length; i++) {
            const tokenIDs = testData[i][0];
            const seriesIds = testData[i][1];
            const amounts = testData[i][2];
            const buyer = testData[i][3];
            const recipient = testData[i][4];
            const royalties = testData[i][5];
            const royaltyPayees = testData[i][6];
            const withFunds = testData[i][7];
            const contract = testData[i][8];
            const bps = 10000;
            const royaltyShares = [
                [
                    [royaltyPayees[0], royalties[0] * bps],
                    [royaltyPayees[1], royalties[1] * bps],
                ],
                [
                    [royaltyPayees[0], royalties[0] * bps],
                    [royaltyPayees[1], royalties[1] * bps],
                ],
            ];
            const price = 0.25 * 1e18;
            const cost = 0.02 * 1e18;
            const chainId = await web3.eth.getChainId();
            const payByVault = !withFunds;

            if (payByVault) {
                // deposit for vault contract
                await web3.eth.sendTransaction({
                    to: this.vault.address,
                    from: accounts[0],
                    value: BigInt(price).toString(),
                });
            }

            // mint artwork
            await contract.mintArtworks(
                [
                    [seriesIds[0], tokenIDs[0], contract.address, amounts[0]],
                    [seriesIds[1], tokenIDs[1], contract.address, amounts[1]],
                ],
                { from: this.trustee }
            );
            const balanceOfToken1 = await contract.balanceOf(
                contract.address,
                tokenIDs[0]
            );
            assert.equal(balanceOfToken1, amounts[0]);
            const balanceOfToken2 = await contract.balanceOf(
                contract.address,
                tokenIDs[1]
            );
            assert.equal(balanceOfToken2, amounts[1]);

            // start sale
            await contract.startSale();

            // encode params
            const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
            const biddingUnixNano = BigInt(
                new Date().getTime() * 1e6
            ).toString();
            const saleData = [
                BigInt(price).toString(),
                BigInt(cost).toString(),
                expiryTime,
                recipient,
                tokenIDs,
                royaltyShares,
                payByVault,
                biddingUnixNano,
            ];
            const signedParams = web3.eth.abi.encodeParameters(
                [
                    "uint",
                    "address",
                    "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool,uint256)",
                ],
                [BigInt(chainId).toString(), contract.address, saleData]
            );

            // sign params
            const hash = web3.utils.keccak256(signedParams);
            var sig = await web3.eth.sign(hash, this.signer);
            sig = sig.substr(2);
            const r = "0x" + sig.slice(0, 64);
            const s = "0x" + sig.slice(64, 128);
            const v = "0x" + sig.slice(128, 130);

            // get ETH balances for stakeholders
            const payee1BalanceBefore = await web3.eth.getBalance(
                royaltyPayees[0]
            );
            const payee2BalanceBefore = await web3.eth.getBalance(
                royaltyPayees[1]
            );
            const costReceiverBalanceBefore = await web3.eth.getBalance(
                COST_RECEIVER
            );

            // check token balance before sale
            for (let i = 0; i < tokenIDs.length; i++) {
                assert.equal(
                    await contract.balanceOf(recipient, tokenIDs[i]),
                    0
                );
                assert.equal(
                    await contract.balanceOf(contract.address, tokenIDs[i]),
                    amounts[i]
                );
            }

            // check seller artworks before sale
            const artworkOfSender = await contract.artworkOf(contract.address);
            assert.equal(artworkOfSender.length, 2);
            for (let i = 0; i < artworkOfSender.length; i++) {
                assert.equal(artworkOfSender[i].seriesId, seriesIds[i]);
                assert.equal(artworkOfSender[i].tokenId, tokenIDs[i]);
                assert.equal(artworkOfSender[i].amount, amounts[i]);
            }

            // check recipient artworks before sale
            assert.equal((await contract.artworkOf(recipient)).length, 0);

            // buy artworks
            const funds = withFunds ? price : 0;
            await contract.buyArtworks(
                r,
                s,
                web3.utils.toDecimal(v) + 27, // magic 27
                saleData,
                { from: buyer, value: funds }
            );

            // check token balance after sale
            for (let i = 0; i < tokenIDs.length; i++) {
                assert.equal(
                    await contract.balanceOf(contract.address, tokenIDs[i]),
                    amounts[i] - 1
                );
                assert.equal(
                    await contract.balanceOf(recipient, tokenIDs[i]),
                    1
                );
            }

            // check recipient artworks after sale
            const artworksOfRecipientAfter = await contract.artworkOf(
                recipient
            );
            assert.equal(artworksOfRecipientAfter.length, 2);
            for (let i = 0; i < artworksOfRecipientAfter.length; i++) {
                assert.equal(
                    artworksOfRecipientAfter[i].seriesId,
                    seriesIds[i]
                );
                assert.equal(artworksOfRecipientAfter[i].tokenId, tokenIDs[i]);
                assert.equal(artworksOfRecipientAfter[i].amount, amounts[i]);
            }

            // get ETH balances for stakeholders after sale
            const payee1BalanceAfter = await web3.eth.getBalance(
                royaltyPayees[0]
            );
            const payee2BalanceAfter = await web3.eth.getBalance(
                royaltyPayees[1]
            );
            const costReceiverBalanceAfter = await web3.eth.getBalance(
                COST_RECEIVER
            );

            // check the royalty distribution
            assert.equal(
                (
                    BigInt(payee1BalanceAfter) - BigInt(payee1BalanceBefore)
                ).toString(),
                BigInt((price - cost) * royalties[0]).toString()
            );
            assert.equal(
                (
                    BigInt(payee2BalanceAfter) - BigInt(payee2BalanceBefore)
                ).toString(),
                BigInt((price - cost) * royalties[1]).toString()
            );
            assert.equal(
                (
                    BigInt(costReceiverBalanceAfter) -
                    BigInt(costReceiverBalanceBefore)
                ).toString(),
                BigInt(cost).toString()
            );
        }
    });

    it("test burn unsold artworks", async function () {
        let contract = this.contracts[5];

        // 1. unauthorized call
        try {
            await contract.burnUnsoldArtworks(10, { from: this.trustee });
        } catch (error) {
            assert.equal("Ownable: caller is not the owner", error.reason);
        }

        // 2. limit is zero
        try {
            await contract.burnUnsoldArtworks(0);
        } catch (error) {
            assert.equal("FeralfileExhibitionV5: limit_ is zero", error.reason);
        }

        // 3. mintable is true
        try {
            await contract.burnUnsoldArtworks(10);
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: mintable required to be false",
                error.reason
            );
        }

        // mint artworks and start sale
        assert.equal(await contract.selling(), false);
        const tokenIDs = [1, 2, 3, 4, 5];
        const amounts = [1, 1, 1, 100, 100];
        await contract.mintArtworks([
            [this.seriesIds[0], tokenIDs[0], contract.address, amounts[0]],
            [this.seriesIds[1], tokenIDs[1], contract.address, amounts[1]],
            [this.seriesIds[2], tokenIDs[2], contract.address, amounts[2]],
            [this.seriesIds[3], tokenIDs[3], contract.address, amounts[3]],
            [this.seriesIds[4], tokenIDs[4], contract.address, amounts[4]],
        ]);
        await contract.startSale();
        assert.equal(await contract.selling(), true);

        // 4. selling is true
        try {
            await contract.burnUnsoldArtworks(10);
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: selling required to be false",
                error.reason
            );
        }

        // 4. burn unsold artworks successfully
        await contract.pauseSale();
        assert.equal(await contract.selling(), false);
        assert.equal(await contract.mintable(), false);
        const tokensBefore = await contract.tokenOf(contract.address);
        assert.equal(tokensBefore.length, 5);
        const artworkBefore = await contract.artworkOf(contract.address);
        assert.equal(artworkBefore.length, 5);

        // burn with limit = 1
        await contract.burnUnsoldArtworks(1);
        let tokensAfter = await contract.tokenOf(contract.address);
        assert.equal(tokensAfter.length, 4);
        let artworkAfter = await contract.artworkOf(contract.address);
        assert.equal(artworkAfter.length, 4);

        // burn with large limitation
        await contract.burnUnsoldArtworks(10);
        tokensAfter = await contract.tokenOf(contract.address);
        assert.equal(tokensAfter.length, 0);
        artworkAfter = await contract.artworkOf(contract.address);
        assert.equal(artworkAfter.length, 0);

        // series total supply
        assert.equal(await contract.seriesTotalSupply(this.seriesIds[0]), 0);
        assert.equal(await contract.seriesTotalSupply(this.seriesIds[1]), 0);
        assert.equal(await contract.seriesTotalSupply(this.seriesIds[2]), 0);
        assert.equal(await contract.seriesTotalSupply(this.seriesIds[3]), 0);
        assert.equal(await contract.seriesTotalSupply(this.seriesIds[4]), 0);

        // artwork total supply
        assert.equal(await contract.artworkTotalSupply(tokenIDs[0]), 0);
        assert.equal(await contract.artworkTotalSupply(tokenIDs[1]), 0);
        assert.equal(await contract.artworkTotalSupply(tokenIDs[2]), 0);
        assert.equal(await contract.artworkTotalSupply(tokenIDs[3]), 0);
        assert.equal(await contract.artworkTotalSupply(tokenIDs[4]), 0);

        // get artwork
        for (let i = 0; i < tokenIDs.length; i++) {
            const artwork = await contract.getArtwork(tokenIDs[i]);
            assert.equal(artwork.tokenId, 0);
            assert.equal(artwork.seriesId, 0);
            assert.equal(artwork.amount, 0);
        }
    });

    it("test transfer unsold artworks", async function () {
        let contract = this.contracts[6];

        // 1. unauthorized call
        try {
            await contract.transferUnsoldArtworks([1, 2, 3], accounts[0], {
                from: this.trustee,
            });
        } catch (error) {
            assert.equal("Ownable: caller is not the owner", error.reason);
        }

        // 2. address is zero
        try {
            await contract.transferUnsoldArtworks([1, 2, 3], ZERO_ADDRESS),
                { from: accounts[0] };
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: to_ is zero address",
                error.reason
            );
        }

        // 3. token IDs is empty
        try {
            await contract.transferUnsoldArtworks([], accounts[0], {
                from: accounts[0],
            });
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: tokenIds_ is empty",
                error.reason
            );
        }

        // 4. mintable is true
        try {
            await contract.transferUnsoldArtworks([1, 2, 3], accounts[0], {
                from: accounts[0],
            });
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: mintable required to be false",
                error.reason
            );
        }

        // mint artworks and start sale
        assert.equal(await contract.selling(), false);
        const tokenIDs = [1, 2, 3, 4, 5];
        const amounts = [1, 1, 1, 100, 100];
        await contract.mintArtworks([
            [this.seriesIds[0], tokenIDs[0], contract.address, amounts[0]],
            [this.seriesIds[1], tokenIDs[1], contract.address, amounts[1]],
            [this.seriesIds[2], tokenIDs[2], contract.address, amounts[2]],
            [this.seriesIds[3], tokenIDs[3], contract.address, amounts[3]],
            [this.seriesIds[4], tokenIDs[4], contract.address, amounts[4]],
        ]);
        await contract.startSale();
        assert.equal(await contract.selling(), true);

        // 5. selling is true
        try {
            await contract.transferUnsoldArtworks([1, 2, 3], accounts[0], {
                from: accounts[0],
            });
        } catch (error) {
            assert.equal(
                "FeralfileExhibitionV5: selling required to be false",
                error.reason
            );
        }

        // 6. transfer unsold artworks successfully
        await contract.pauseSale();
        assert.equal(await contract.selling(), false);
        assert.equal(await contract.mintable(), false);

        // verify token balance before transfer
        const tokensOfContractBefore = await contract.tokenOf(contract.address);
        assert.equal(tokensOfContractBefore.length, tokenIDs.length);
        const recipient = accounts[0];
        const tokensOfRecipientBefore = await contract.tokenOf(recipient);
        assert.equal(tokensOfRecipientBefore.length, 0);

        // transfer unsold artworks
        await contract.transferUnsoldArtworks(tokenIDs, recipient, {
            from: accounts[0],
        });

        // verify token balance after transfer
        const tokensOfContractAfter = await contract.tokenOf(contract.address);
        assert.equal(tokensOfContractAfter.length, 0);
        const tokensOfRecipientAfter = await contract.tokenOf(recipient);
        assert.equal(tokensOfRecipientAfter.length, tokenIDs.length);
    });
});
