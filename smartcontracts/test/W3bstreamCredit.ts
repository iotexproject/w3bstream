import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('W3bstream Credit', function () {
  it('grant credit', async function () {
    const [owner, hacker] = await ethers.getSigners();

    const w3bstreamCredit = await ethers.deployContract('W3bstreamCredit');
    await w3bstreamCredit.initialize('W3bstream Credit', 'W3BC');
    await w3bstreamCredit.setMinter(owner.address);
    await expect(w3bstreamCredit.connect(hacker).grant(hacker.address, 1234567)).to.be.reverted;
    expect(await w3bstreamCredit.balanceOf(owner.address)).to.equal(0);
    await w3bstreamCredit.grant(owner.address, 1234567);

    const ownerBalance = await w3bstreamCredit.balanceOf(owner.address);
    expect(await w3bstreamCredit.totalSupply()).to.equal(ownerBalance);
  });
});
