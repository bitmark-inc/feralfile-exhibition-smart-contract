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

        const tx = await this.contract.mintToken(amount, owner);

        // total supply
        const totalSupply = await this.contract.totalSupply();
        assert.equal(totalSupply, amount);

        // balance
        const balance = await this.contract.balanceOf(owner);
        assert.equal(balance, amount);
    });
});