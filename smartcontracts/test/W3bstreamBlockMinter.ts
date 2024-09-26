import { expect } from 'chai';
import { keccak256 } from 'ethers';
import { ethers } from 'hardhat';

describe('W3bstream Minter', function () {
  let minter;
  let dao;
  let tm;
  let brd;
  let scrypt;
  const genesis = "0x0000000000000000000000000000000000000000000000000000000000000000";
  beforeEach(async function () {
    minter = await ethers.deployContract('W3bstreamBlockMinter');
    dao = await ethers.deployContract('W3bstreamDAO');
    tm = await ethers.deployContract('W3bstreamTaskManager');
    brd = await ethers.deployContract('W3bstreamBlockRewardDistributor');
    scrypt = await ethers.deployContract('Scrypt');
    await dao.initialize(genesis);
    await tm.initialize();
    await brd.initialize();
    await minter.initialize(dao.getAddress(), tm.getAddress(), brd.getAddress(), scrypt.getAddress());
    // let tx = await minter.scrypt("0x000000020000000000000000000000000000000000000000000000000000000000000000ab04ea90eb7d931cbbaa94a11cb3907809c13262dd37acc526e4b4a628e43b111c7fffff00000089929df805");
    // console.log(tx);
    // exit(0);
    await dao.transferOwnership(minter.getAddress());
    await tm.addOperator(minter.getAddress());
  });
  it('mint block', async function () {
    const tip = await ethers.provider.getBlock('latest');
    const [owner, sequencer, prover] = await ethers.getSigners();
    const coinbase = {
      addr: sequencer.address,
      operator: sequencer.address,
      beneficiary: sequencer.address,
    };
    await minter.connect(owner).setAdhocNBits("0x1cffff00");
    let currentNBits = await minter.currentNBits();
    const merkleRoot = ethers.solidityPackedKeccak256(["address", "address", "address"], [coinbase.addr, coinbase.operator, coinbase.beneficiary]);
    let tipinfo = await dao.tip();
    expect(tipinfo[0]).to.equal(0);
    expect(tipinfo[1]).to.equal(genesis);
    let blockinfo = {
      meta: "0x00000000",
      prevhash: genesis,
      merkleRoot: merkleRoot,
      nbits: currentNBits,
      nonce: "0x0000000000000000",
    };
    const nbits = await minter.currentNBits();
    const currentTarget = await minter.nbitsToTarget(nbits);
    await minter.connect(sequencer).mint(
      blockinfo,
      coinbase,
      [],
    );
    tipinfo = await dao.tip();
    expect(tipinfo[0]).to.equal(1);
    await expect(minter.connect(sequencer).mint(
      blockinfo,
      coinbase,
      [],
    )).to.be.revertedWith("invalid prevhash");
    blockinfo.prevhash = tipinfo[1];
    blockinfo.nbits = "0x00008000";
    await expect(minter.connect(sequencer).mint(
      blockinfo,
      coinbase,
      [],
    )).to.be.revertedWith("invalid nbits");
    currentNBits = await minter.currentNBits();
    blockinfo.nbits = currentNBits;
    await minter.connect(sequencer).mint(
      blockinfo,
      coinbase,
      [],
    );
    tipinfo = await dao.tip();
    expect(tipinfo[0]).to.equal(2);
  });
});
