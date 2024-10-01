import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('W3bstream Minter', function () {
  let minter;
  let dao;
  let tm;
  let brd;
  let bhv;
  let scrypt;
  const genesis = "0x0000000000000000000000000000000000000000000000000000000000000000";
  beforeEach(async function () {
    minter = await ethers.deployContract('W3bstreamBlockMinter');
    dao = await ethers.deployContract('W3bstreamDAO');
    tm = await ethers.deployContract('W3bstreamTaskManager');
    brd = await ethers.deployContract('W3bstreamBlockRewardDistributor');
    scrypt = await ethers.deployContract('MockScrypt');
    bhv = await ethers.deployContract('W3bstreamBlockHeaderValidator', [scrypt.getAddress()]);
    await dao.initialize(genesis);
    await tm.initialize();
    await brd.initialize();
    await minter.initialize(dao.getAddress(), tm.getAddress(), brd.getAddress(), bhv.getAddress());
    await dao.transferOwnership(minter.getAddress());
    await tm.addOperator(minter.getAddress());
    await brd.setOperator(minter.getAddress());
    await bhv.setOperator(minter.getAddress());
    await minter.setBlockReward(0);
  });
  it('mint block', async function () {
    const tip = await ethers.provider.getBlock('latest');
    const [owner, sequencer, prover] = await ethers.getSigners();
    const coinbase = {
      addr: sequencer.address,
      operator: sequencer.address,
      beneficiary: sequencer.address,
    };
    await bhv.connect(owner).setAdhocNBits("0x1cffff00");
    let currentNBits = await bhv.currentNBits();
    const merkleRoot = ethers.keccak256(ethers.AbiCoder.defaultAbiCoder().encode(
      ["address", "address", "address"],
      [coinbase.addr, coinbase.operator, coinbase.beneficiary]
    ));
    let tipinfo = await dao.tip();
    expect(tipinfo[0]).to.equal(0);
    expect(tipinfo[1]).to.equal(genesis);
    await scrypt.setHash("0x00000000ffff0000000000000000000000000000000000000000000000000000");
    let blockinfo = {
      meta: "0x00000000",
      prevhash: genesis,
      merkleRoot: merkleRoot,
      nbits: currentNBits,
      nonce: "0x0000000000000000",
    };
    const nbits = await bhv.currentNBits();
    const currentTarget = await bhv.nbitsToTarget(nbits);
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
    currentNBits = await bhv.currentNBits();
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
