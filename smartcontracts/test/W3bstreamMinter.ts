import { expect } from 'chai';
import { keccak256 } from 'ethers';
import { ethers } from 'hardhat';

describe('W3bstream Minter', function () {
  let minter;
  let dao;
  let tm;
  const genesis = "0x0000000011111111222222223333333344444444555555556666666677777777";
  beforeEach(async function () {
    minter = await ethers.deployContract('W3bstreamMinter');
    dao = await ethers.deployContract('W3bstreamDAO');
    tm = await ethers.deployContract('W3bstreamTaskManager');
    await dao.initialize(genesis);
    await tm.initialize();
    await minter.initialize(dao.getAddress(), tm.getAddress());
    await dao.transferOwnership(minter.getAddress());
    await tm.addOperator(minter.getAddress());
  });
  it('mint block', async function () {
    const tip = await ethers.provider.getBlock('latest');
    const timestamp = tip.timestamp;
    const [owner, sequencer, prover] = await ethers.getSigners();
    const coinbase = {
      addr: sequencer.address,
      operator: sequencer.address,
      beneficiary: sequencer.address,
    };
    const merkleRoot = ethers.solidityPackedKeccak256(["uint256", "address", "address", "address"], [timestamp, coinbase.addr, coinbase.operator, coinbase.beneficiary]);
    const blockinfo = {
      meta: "0x00000000",
      prevhash: genesis,
      merkleRoot: merkleRoot,
      difficulty: "0xffffffff",
      nonce: "0x0000000000000000",
    };
    let tipinfo = await dao.tip();
    expect(tipinfo[0]).to.equal(0);
    expect(tipinfo[1]).to.equal(genesis);
    await minter.connect(sequencer).mint(
      blockinfo,
      timestamp,
      coinbase,
      [],
    );
    tipinfo = await dao.tip();
    // TODO: adjust timestamp & merkle root
    expect(tipinfo[0]).to.equal(1);
    await expect(minter.connect(sequencer).mint(
      blockinfo,
      timestamp,
      coinbase,
      [],
    )).to.be.revertedWith("invalid prevhash");
    blockinfo.prevhash = tipinfo[1];
    blockinfo.difficulty = "0x00000001";
    await expect(minter.connect(sequencer).mint(
      blockinfo,
      timestamp,
      coinbase,
      [],
    )).to.be.revertedWith("invalid proof of work");
    blockinfo.difficulty = "0xffffffff";
    await minter.connect(sequencer).mint(
      blockinfo,
      timestamp,
      coinbase,
      [],
    );
    tipinfo = await dao.tip();
    expect(tipinfo[0]).to.equal(2);
  });
});
