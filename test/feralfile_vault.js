const FeralfileVault = artifacts.require('FeralfileVault');

const IPFS_GATEWAY_PREFIX = 'https://cloudflare-ipfs.com/ipfs/';
const EXHIBITION_CONTRACT_ADDRESS = '0x46f2B641d8702f29c45f6D06292dC34Eb9dB1801';
const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

contract('FeralfileVault', async (accounts) => {
    before(async function () {
        this.vault = await FeralfileVault.new(accounts[1]);
    });

    it("check adding funds", async function () {
        const vaultBalanceBefore = await web3.eth.getBalance(this.vault.address);
        assert.equal(vaultBalanceBefore, 0);

        await web3.eth.sendTransaction({ from: accounts[0], to: this.vault.address, value: 0.5 * 1e18 });

        const vaultBalanceAfter = await web3.eth.getBalance(this.vault.address);
        assert.equal(vaultBalanceAfter, 0.5 * 1e18);
    });

    it("check payForSale", async function () {
        try {
            // Generate signature
            const expiryTime = (new Date().getTime() / 1000 + 300).toFixed(0);
            const saleData = [
                BigInt(0.1 * 1e18).toString(),
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
                true,
            ];
            const signParams = web3.eth.abi.encodeParameters(
                [
                    "uint",
                    "address",
                    "tuple(uint256,uint256,uint256,address,uint256[],tuple(address,uint256)[][],bool)",
                ],
                [
                    BigInt(await web3.eth.getChainId()).toString(),
                    EXHIBITION_CONTRACT_ADDRESS,
                    saleData,
                ]
            );
            const hash = web3.utils.keccak256(signParams);
            var sig = await web3.eth.sign(hash, accounts[1]);
            sig = sig.substr(2);
            const r = "0x" + sig.slice(0, 64);
            const s = "0x" + sig.slice(64, 128);
            const v = web3.utils.toDecimal("0x" + sig.slice(128, 130)) + 27;
            await this.vault.payForSale(EXHIBITION_CONTRACT_ADDRESS, r, s, v, saleData);
            const vaultBalanceAfter = await web3.eth.getBalance(this.vault.address);
            assert.equal(vaultBalanceAfter, 0.4 * 1e18);
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