const FeralFileFundDispatcher = artifacts.require('FeralFileFundDispatcher');

const MockERC20 = artifacts.require('MockERC20');

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

contract('FeralFileFundDispatcher', async (accounts) => {
  before(async function () {
    this.ffFundDispatcher = await FeralFileFundDispatcher.new();

    this.mockERC20 = await MockERC20.new('Test', 'TT');
  });

  it('test update operator successfully', async function () {
    const result = await this.ffFundDispatcher.updateOperator(accounts[1]);
    const operator = result.logs[0].args.operator;

    assert.equal(accounts[1], operator);
  });

  it('test update operator failed', async function () {
    try {
      await this.ffFundDispatcher.updateOperator(ZERO_ADDRESS);
    } catch (e) {
      assert.equal(
        e.message,
        'Returned error: VM Exception while processing transaction: revert Invalid operator address'
      );
    }
  });

  it('test batch transfer token from contract successfully', async function () {
    const mockERC20Success = await MockERC20.new('Test', 'TT');
    await mockERC20Success.transfer(
      this.ffFundDispatcher.address,
      BigInt(10 * 1e18).toString()
    );

    const balanceBefore = await mockERC20Success.balanceOf(
      this.ffFundDispatcher.address
    );

    assert.equal(balanceBefore.toString(), BigInt(10 * 1e18).toString());

    await this.ffFundDispatcher.batchTransferToken(
      mockERC20Success.address,
      [accounts[2], accounts[3]],
      [BigInt(1 * 1e18).toString(), BigInt(2 * 1e18).toString()],
      { from: accounts[1] }
    );

    const balanceAfter = await mockERC20Success.balanceOf(
      this.ffFundDispatcher.address
    );
    assert.equal(balanceAfter.toString(), BigInt(7 * 1e18).toString());

    const balance2 = await mockERC20Success.balanceOf(accounts[2]);
    assert.equal(balance2.toString(), BigInt(1 * 1e18).toString());

    const balance3 = await mockERC20Success.balanceOf(accounts[3]);
    assert.equal(balance3.toString(), BigInt(2 * 1e18).toString());
  });

  it('test batch transfer token from contract failed because no balance', async function () {
    const balanceBefore = await this.mockERC20.balanceOf(
      this.ffFundDispatcher.address
    );

    assert.equal(balanceBefore.toString(), BigInt(0).toString());
    try {
      await this.ffFundDispatcher.batchTransferToken(
        this.mockERC20.address,
        [accounts[2], accounts[3]],
        [BigInt(1 * 1e18).toString(), BigInt(2 * 1e18).toString()],
        { from: accounts[1] }
      );
    } catch (e) {
      assert.equal(
        e.message,
        'Returned error: VM Exception while processing transaction: revert insufficient balance'
      );
    }
  });

  it('test batch transfer token from contract failed because balance does not enough', async function () {
    await this.mockERC20.transfer(
      this.ffFundDispatcher.address,
      BigInt(3 * 1e18).toString()
    );

    const balanceBefore = await this.mockERC20.balanceOf(
      this.ffFundDispatcher.address
    );

    assert.equal(balanceBefore.toString(), BigInt(3 * 1e18).toString());

    try {
      await this.ffFundDispatcher.batchTransferToken(
        this.mockERC20.address,
        [accounts[2], accounts[3]],
        [BigInt(3 * 1e18).toString(), BigInt(2 * 1e18).toString()],
        { from: accounts[1] }
      );
    } catch (e) {
      assert.equal(
        e.message,
        'Returned error: VM Exception while processing transaction: revert insufficient balance'
      );
    }
  });

  it('test batch transfer token from address successfully', async function () {
    await this.mockERC20.transfer(accounts[4], BigInt(20 * 1e18).toString());

    const balanceBefore = await this.mockERC20.balanceOf(accounts[4]);

    assert.equal(balanceBefore.toString(), BigInt(20 * 1e18).toString());

    await this.mockERC20.approve(
      this.ffFundDispatcher.address,
      BigInt(10 * 1e18).toString(),
      {
        from: accounts[4],
      }
    );

    const allowanceBefore = await this.mockERC20.allowance(
      accounts[4],
      this.ffFundDispatcher.address
    );

    assert.equal(allowanceBefore.toString(), BigInt(10 * 1e18).toString());

    await this.ffFundDispatcher.batchTransferTokenFrom(
      this.mockERC20.address,
      accounts[4],
      [accounts[5], accounts[6]],
      [BigInt(1.8 * 1e18).toString(), BigInt(2.6 * 1e18).toString()],
      { from: accounts[1] }
    );

    const balanceAfter = await this.mockERC20.balanceOf(accounts[4]);
    assert.equal(balanceAfter.toString(), BigInt(15.6 * 1e18).toString());

    const allowanceAfter = await this.mockERC20.allowance(
      accounts[4],
      this.ffFundDispatcher.address
    );

    assert.equal(allowanceAfter.toString(), BigInt(5.6 * 1e18).toString());

    const balance2 = await this.mockERC20.balanceOf(accounts[5]);
    assert.equal(balance2.toString(), BigInt(1.8 * 1e18).toString());

    const balance3 = await this.mockERC20.balanceOf(accounts[6]);
    assert.equal(balance3.toString(), BigInt(2.6 * 1e18).toString());
  });

  it('test batch transfer token from address failed because no balance', async function () {
    const mockERC20ForFail = await MockERC20.new('Test', 'TT');

    const balanceBefore = await mockERC20ForFail.balanceOf(accounts[4]);

    assert.equal(balanceBefore.toString(), BigInt(0).toString());

    try {
      await this.ffFundDispatcher.batchTransferTokenFrom(
        mockERC20ForFail.address,
        accounts[4],
        [accounts[5], accounts[6]],
        [BigInt(1.8 * 1e18).toString(), BigInt(2.6 * 1e18).toString()],
        { from: accounts[1] }
      );
    } catch (e) {
      assert.equal(
        e.message,
        'Returned error: VM Exception while processing transaction: revert insufficient balance'
      );
    }
  });

  it('test batch transfer token from address failed because no approval', async function () {
    const mockERC20ForFail = await MockERC20.new('Test', 'TT');

    await mockERC20ForFail.transfer(accounts[4], BigInt(20 * 1e18).toString());

    const balanceBefore = await mockERC20ForFail.balanceOf(accounts[4]);

    assert.equal(balanceBefore.toString(), BigInt(20 * 1e18).toString());

    try {
      await this.ffFundDispatcher.batchTransferTokenFrom(
        mockERC20ForFail.address,
        accounts[4],
        [accounts[5], accounts[6]],
        [BigInt(1.8 * 1e18).toString(), BigInt(2.6 * 1e18).toString()],
        { from: accounts[1] }
      );
    } catch (e) {
      assert.equal(
        e.message,
        'Returned error: VM Exception while processing transaction: revert ERC20: insufficient allowance'
      );
    }
  });

  it('test batch transfer token from address failed because balance not enough', async function () {
    const mockERC20ForFail = await MockERC20.new('Test', 'TT');

    await mockERC20ForFail.transfer(accounts[4], BigInt(20 * 1e18).toString());

    const balanceBefore = await mockERC20ForFail.balanceOf(accounts[4]);

    assert.equal(balanceBefore.toString(), BigInt(20 * 1e18).toString());

    await mockERC20ForFail.approve(
      this.ffFundDispatcher.address,
      BigInt(20 * 1e18).toString(),
      {
        from: accounts[4],
      }
    );

    const allowanceBefore = await mockERC20ForFail.allowance(
      accounts[4],
      this.ffFundDispatcher.address
    );

    assert.equal(allowanceBefore.toString(), BigInt(20 * 1e18).toString());

    try {
      await this.ffFundDispatcher.batchTransferTokenFrom(
        mockERC20ForFail.address,
        accounts[4],
        [accounts[5], accounts[6]],
        [BigInt(15.2 * 1e18).toString(), BigInt(11.4788 * 1e18).toString()],
        { from: accounts[1] }
      );
    } catch (e) {
      assert.equal(
        e.message,
        'Returned error: VM Exception while processing transaction: revert insufficient balance'
      );
    }
  });

  it('test batch transfer token from address failed because allowance not enough', async function () {
    const mockERC20ForFail = await MockERC20.new('Test', 'TT');

    await mockERC20ForFail.transfer(accounts[4], BigInt(20 * 1e18).toString());

    const balanceBefore = await mockERC20ForFail.balanceOf(accounts[4]);

    assert.equal(balanceBefore.toString(), BigInt(20 * 1e18).toString());

    await mockERC20ForFail.approve(
      this.ffFundDispatcher.address,
      BigInt(10 * 1e18).toString(),
      {
        from: accounts[4],
      }
    );

    const allowanceBefore = await mockERC20ForFail.allowance(
      accounts[4],
      this.ffFundDispatcher.address
    );

    assert.equal(allowanceBefore.toString(), BigInt(10 * 1e18).toString());

    try {
      await this.ffFundDispatcher.batchTransferTokenFrom(
        mockERC20ForFail.address,
        accounts[4],
        [accounts[5], accounts[6]],
        [BigInt(10 * 1e18).toString(), BigInt(10 * 1e18).toString()],
        { from: accounts[1] }
      );
    } catch (e) {
      assert.equal(
        e.message,
        'Returned error: VM Exception while processing transaction: revert ERC20: insufficient allowance'
      );
    }
  });

  it('test batch transfer ethereum from contract successfully', async function () {
    const balanceBefore2 = await web3.eth.getBalance(accounts[2]);
    const balanceBefore3 = await web3.eth.getBalance(accounts[3]);

    await this.ffFundDispatcher.batchTransferETH(
      [accounts[2], accounts[3]],
      [BigInt(0.15 * 1e18).toString(), BigInt(0.1 * 1e18).toString()],
      { from: accounts[1], value: 0.3 * 1e18 }
    );

    const balanceAfter2 = await web3.eth.getBalance(accounts[2]);
    const balanceAfter3 = await web3.eth.getBalance(accounts[3]);

    assert.equal(
      (balanceAfter2 - balanceBefore2).toString(),
      BigInt(0.15 * 1e18).toString()
    );
    assert.equal(
      (balanceAfter3 - balanceBefore3).toString(),
      BigInt(0.1 * 1e18).toString()
    );

    const balanceBefore = await web3.eth.getBalance(
      this.ffFundDispatcher.address
    );

    assert.equal(balanceBefore.toString(), BigInt(0.05 * 1e18).toString());
  });

  it('test batch transfer ethereum from contract failed because over balance', async function () {
    const balanceBefore = await web3.eth.getBalance(accounts[1]);
    try {
      await this.ffFundDispatcher.batchTransferETH(
        [accounts[2], accounts[3]],
        [BigInt(0.15 * 1e18).toString(), BigInt(0.1 * 1e18).toString()],
        { from: accounts[1], value: (balanceBefore * 2).toString() }
      );
    } catch (e) {
      assert.equal(
        true,
        e.message.indexOf(
          `Returned error: sender doesn't have enough funds to send tx. The upfront cost is:`
        ) >= 0
      );
      assert.equal(
        true,
        e.message.indexOf(
          `and the sender's account only has: ${balanceBefore.toString()}`
        ) >= 0
      );
    }
  });

  it('test batch transfer ethereum from contract failed because insufficient balance input', async function () {
    try {
      await this.ffFundDispatcher.batchTransferETH(
        [accounts[2], accounts[3]],
        [BigInt(0.15 * 1e18).toString(), BigInt(0.1 * 1e18).toString()],
        { from: accounts[1], value: BigInt(0.2 * 1e18).toString() }
      );
    } catch (e) {
      assert.equal(
        e.message,
        'Returned error: VM Exception while processing transaction: revert insufficient balance'
      );
    }
  });
});
