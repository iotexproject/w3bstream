import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('W3bstream Prover', function () {
  let w3bstreamProver;
  beforeEach(async function() {
    w3bstreamProver = await ethers.deployContract('W3bstreamProver');
  });
  it('register prover and update operator', async function () {
    const [owner, minter, prover, operator] = await ethers.getSigners();

    await w3bstreamProver.initialize('W3bstream Prover', 'W3BProver');
    await w3bstreamProver.setMinter(minter.address);
    await w3bstreamProver.connect(minter).mint(prover.address);
    const proverId = 1;
    expect(await w3bstreamProver.prover(proverId)).to.equal(prover.address);
    expect(await w3bstreamProver.operator(proverId)).to.equal(prover.address);
    await w3bstreamProver.connect(prover).changeOperator(proverId, operator.address);
    expect(await w3bstreamProver.operator(proverId)).to.equal(operator.address);
  });
});
