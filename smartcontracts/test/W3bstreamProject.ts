import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('W3bstream Project', function () {
  let w3bstreamProject;
  beforeEach(async function() {
    w3bstreamProject = await ethers.deployContract('W3bstreamProject');
  })
  it('register project', async function () {
    const [owner, minter, projectOwner] = await ethers.getSigners();
    await w3bstreamProject.initialize('W3bstream Project', 'W3BProject');
    await w3bstreamProject.setMinter(minter.address);
    await w3bstreamProject.connect(minter).mint(projectOwner.address);
    const projectId = 1;
    // TODO: read project id from event
    expect(await w3bstreamProject.ownerOf(projectId)).to.equal(projectOwner.address);
    it('update config successfully', async function() {
      const uri = "http://test";
      const hash = "0x0000000011111111222222223333333344444444555555556666666677777777";
      await w3bstreamProject.connect(projectOwner).updateConfig(projectId, uri, hash);
      const cfg = await w3bstreamProject.config(projectId);
      expect(cfg[0]).to.equal(uri);
      expect(cfg[1]).to.equal(hash);
    });
  });
});
