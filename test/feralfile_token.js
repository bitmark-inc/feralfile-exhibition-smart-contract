const FeralfileToken = artifacts.require('FeralfileToken');

contract('FeralfileToken', async (accounts) => {
    before(async function () {
        this.contract = await FeralfileToken.new("Feral File Token Test", "FFToken",);
    });

    it("test contract constructor", async function () {
        const name = await this.contract.name();
        assert.equal(name, "Feral File Token Test");

        const symbol = await this.contract.symbol();
        assert.equal(symbol, "FFToken");

        const totalSupply = await this.contract.totalSupply();
        assert.equal(totalSupply, 0);

    });

    it("test mint token", async function () {
        const owner = accounts[1];
        const amount = 24;

        await this.contract.mint(owner, amount, { from: accounts[0] });

        // total supply
        const totalSupply = await this.contract.totalSupply();
        assert.equal(totalSupply, amount);

        // balance
        const balance = await this.contract.balanceOf(owner);
        assert.equal(balance, amount);
    });

    it("test mint tokens for multiple owners", async function () {
        const owners = [accounts[2], accounts[3], accounts[4]];
        const amounts = [50, 100, 150];

        await this.contract.batchMint(owners, amounts, { from: accounts[0] });

        const totalSupply = await this.contract.totalSupply();
        assert.equal(totalSupply.toNumber(), 300, "Total supply does not match minted amount for multiple users");

        for (let i = 0; i < owners.length; i++) {
            const balance = await this.contract.balanceOf(owners[i]);
            assert.equal(balance.toNumber(), amounts[i], `Balance of owner ${i + 1} does not match minted amount`);
        }
    });

    it("test mint tokens for multiple owners mismatched array lengths", async function () {
        const owners = [accounts[2], accounts[3]];
        const amounts = [50, 100, 150];

        try {
            await this.contract.batchMint(owners, amounts, { from: accounts[0] });
            assert.fail("Minting with mismatched array lengths should have thrown an error");
        } catch (error) {
            assert(error.message.includes("owners and amounts length mismatch"), "Expected Owners and amounts length mismatch error");
        }
    });
});