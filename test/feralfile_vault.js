const FeralfileVault = artifacts.require('FeralfileVault');

const IPFS_GATEWAY_PREFIX = 'https://cloudflare-ipfs.com/ipfs/';

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

contract('FeralfileVault', async (accounts) => {
    before(async function () {
        this.vault = await FeralfileVault.new();
    });

    it("check adding funds", async function () {
        const vaultBalanceBefore = await web3.eth.getBalance(this.vault.address);
        assert.equal(vaultBalanceBefore, 0);

        await web3.eth.sendTransaction({ from: accounts[0], to: this.vault.address, value: 0.5 * 1e18 });

        const vaultBalanceAfter = await web3.eth.getBalance(this.vault.address);
        assert.equal(vaultBalanceAfter, 0.5 * 1e18);
    });

    it("check call pay by not exhibition contract", async function () {
        try {
            await this.vault.pay(BigInt(0.5 * 1e18).toString());
        } catch (error) {
            assert.equal(error.reason, "Only exhibition contract can call this function.")
        }
    });

    it("check call pay by exhibition contract", async function () {
        try {
            await this.vault.setExhibitionContract(accounts[0])
            await this.vault.pay(BigInt(0.01 * 1e18).toString());
            const vaultBalanceAfter = await web3.eth.getBalance(this.vault.address);
            assert.equal(vaultBalanceAfter, 0.49 * 1e18);
        } catch (error) {
            console.log(error)
            assert.fail();
        }
    });

    it("check withdraw all funds by owner", async function () {
        try {
            await this.vault.withdrawFunds();
            const vaultBalanceAfter = await web3.eth.getBalance(this.vault.address);
            assert.equal(vaultBalanceAfter, 0);
        } catch (error) {
            console.log(error)
            assert.fail();
        }
    });
});