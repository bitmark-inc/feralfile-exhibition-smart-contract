const FeralFileAirdropV1 = artifacts.require("FeralFileAirdropV1");
const TOKEN_URI =
    "https://ipfs.bitmark.com/ipfs/QmNVdQSp1AvZonLwHzTbbZDPLgbpty15RMQrbPEWd4ooTU/{id}";
const TOKEN_ID = 999;
const TOKEN_TYPE_FUNGIBLE = 0;
const TOKEN_TYPE_NON_FUNGIBLE = 1;
const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";

contract("FeralFileAirdropV1", async (accounts) => {
    before(async () => {
        this.owner = accounts[0];
        this.trustee = accounts[1];
        this.contracts = [];

        // To isolate the test cases, we create multiple pairs of contracts
        // for each token type. The index of the contract pair will be used
        // as index of test cases.
        for (let i = 0; i < 5; i++) {
            let fungibleContract = await FeralFileAirdropV1.new(
                TOKEN_TYPE_FUNGIBLE,
                TOKEN_URI
            );
            await fungibleContract.addTrustee(this.trustee, {
                from: this.owner,
            });
            let nonFungibleContract = await FeralFileAirdropV1.new(
                TOKEN_TYPE_NON_FUNGIBLE,
                TOKEN_URI
            );
            await nonFungibleContract.addTrustee(this.trustee, {
                from: this.owner,
            });

            // Should follow the index of Type.Fungible and Type.NonFungible
            // so that we could leverage the index for token type check
            this.contracts[i] = [fungibleContract, nonFungibleContract];
        }
    });

    it("test constructor", async () => {
        let contracts = this.contracts[0];
        for (let i = 0; i < contracts.length; i++) {
            let contract = contracts[i];
            assert.equal(await contract.version(), "v1");
            assert.equal(await contract.uri(TOKEN_ID), TOKEN_URI);
            assert.equal(await contract.isEnded(), false);
            assert.equal(await contract.tokenType(), i);
        }
    });

    it("test mint", async () => {
        let contracts = this.contracts[1];
        for (let i = 0; i < contracts.length; i++) {
            let contract = contracts[i];

            // test unauthorized call
            try {
                await contract.mint(TOKEN_ID, 1, {
                    from: accounts[2],
                });
            } catch (e) {
                assert.ok(
                    e.hijackedStack.includes(
                        "VM Exception while processing transaction: revert"
                    )
                );
            }

            // test mint token with amount mismatch
            let wrongAmount = i == 0 ? 0 : 2;
            try {
                await contract.mint(TOKEN_ID, wrongAmount, {
                    from: this.owner,
                });
            } catch (e) {
                assert.equal(
                    e.reason,
                    "FeralFileTokenAirdrop: amount mismatch"
                );
            }

            // mint successfully
            let correctAmount = i == 0 ? 10 : 1;
            await contract.mint(TOKEN_ID, correctAmount, { from: this.owner });
            assert.equal(
                await contract.balanceOf(contract.address, TOKEN_ID),
                correctAmount
            );
        }
    });

    it("test airdrop", async () => {
        let contracts = this.contracts[2];
        for (let i = 0; i < contracts.length; i++) {
            let contract = contracts[i];

            // test unauthorized call
            try {
                await contract.airdrop(TOKEN_ID, [accounts[0]], {
                    from: accounts[2],
                });
            } catch (e) {
                assert.ok(
                    e.hijackedStack.includes(
                        "VM Exception while processing transaction: revert"
                    )
                );
            }

            // test airdrop to zero address
            try {
                await contract.airdrop(TOKEN_ID, [ZERO_ADDRESS], {
                    from: this.owner,
                });
            } catch (e) {
                assert.equal(e.reason, "ERC1155: transfer to the zero address");
            }

            // test insufficient balance
            let recipients =
                i == 0
                    ? [accounts[0], accounts[1], accounts[2]]
                    : [accounts[3]];
            try {
                await contract.airdrop(TOKEN_ID, recipients, {
                    from: this.owner,
                });
            } catch (e) {
                assert.equal(
                    e.reason,
                    "ERC1155: insufficient balance for transfer"
                );
            }

            // mint token to prepare for below tests
            let correctAmount = i == 0 ? 10 : 1;
            await contract.mint(TOKEN_ID, correctAmount, { from: this.owner });
            assert.equal(
                await contract.balanceOf(contract.address, TOKEN_ID),
                correctAmount
            );

            // test airdrop successfully
            let beforeBalance = await contract.balanceOf(
                contract.address,
                TOKEN_ID
            );
            await contract.airdrop(TOKEN_ID, recipients, {
                from: this.owner,
            });
            let afterBalance = await contract.balanceOf(
                contract.address,
                TOKEN_ID
            );

            // check contract token balance
            assert.equal(
                afterBalance.toNumber(),
                beforeBalance.toNumber() - recipients.length * 1 // always airdrop 1 token per address
            );

            // check recipients token balance
            for (let j = 0; j < recipients.length; j++) {
                assert.equal(
                    await contract.balanceOf(recipients[j], TOKEN_ID),
                    1
                );
            }

            // test already airdropped
            try {
                await contract.airdrop(TOKEN_ID, recipients, {
                    from: this.owner,
                });
            } catch (e) {
                assert.equal(
                    e.reason,
                    "FeralFileTokenAirdrop: already airdropped"
                );
            }
        }
    });

    it("test end airdrop", async () => {
        let contracts = this.contracts[3];
        for (let i = 0; i < contracts.length; i++) {
            let contract = contracts[i];

            // test unauthorized call
            try {
                await contract.burnRemaining(TOKEN_ID, {
                    from: accounts[2],
                });
            } catch (e) {
                assert.equal(e.reason, "Ownable: caller is not the owner");
            }

            // test end airdrop successfully
            await contract.end({ from: this.owner });
            assert.equal(await contract.isEnded(), true);

            // test already ended
            try {
                await contract.end({ from: this.owner });
            } catch (e) {
                assert.equal(e.reason, "FeralFileTokenAirdrop: already ended");
            }

            // test mint after ended
            try {
                await contract.mint(TOKEN_ID, 1, { from: this.trustee });
            } catch (e) {
                assert.equal(e.reason, "FeralFileTokenAirdrop: already ended");
            }

            // test airdrop after ended
            try {
                await contract.airdrop(TOKEN_ID, [accounts[0]], {
                    from: this.trustee,
                });
            } catch (e) {
                assert.equal(e.reason, "FeralFileTokenAirdrop: already ended");
            }
        }
    });

    it("test burn remaining", async () => {
        let contracts = this.contracts[4];
        for (let i = 0; i < contracts.length; i++) {
            let contract = contracts[i];

            // test unauthorized call
            try {
                await contract.burnRemaining(TOKEN_ID, {
                    from: accounts[2],
                });
            } catch (e) {
                assert.equal(e.reason, "Ownable: caller is not the owner");
            }

            // mint token to prepare for below tests
            let correctAmount = i == 0 ? 10 : 1;
            await contract.mint(TOKEN_ID, correctAmount, {
                from: this.trustee,
            });

            // test burn remaining successfully
            await contract.burnRemaining(TOKEN_ID, { from: this.owner });
            let afterBalance = await contract.balanceOf(
                contract.address,
                TOKEN_ID
            );
            assert.equal(afterBalance.toNumber(), 0);
        }
    });
});
