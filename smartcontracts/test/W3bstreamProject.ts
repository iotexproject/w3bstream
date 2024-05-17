import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('W3bstream Project', function () {
  let project;
  let w3bstreamProject;
  beforeEach(async function () {
    project = await ethers.deployContract('MockProject');
    w3bstreamProject = await ethers.deployContract('W3bstreamProject');
  });
  it('bind project', async function () {
    const [owner, binder, projectOwner] = await ethers.getSigners();

    await project.connect(projectOwner).register();

    await w3bstreamProject.initialize(project.target);
    await w3bstreamProject.setBinder(binder.address);
    const projectId = 1;

    await w3bstreamProject.connect(binder).bind(projectId);
    // TODO: read project id from event
    expect(await w3bstreamProject.ownerOf(projectId)).to.equal(projectOwner.address);
    it('update config successfully', async function () {
      const uri = 'http://test';
      const hash = '0x0000000011111111222222223333333344444444555555556666666677777777';
      await w3bstreamProject.connect(projectOwner).updateConfig(projectId, uri, hash);
      const cfg = await w3bstreamProject.config(projectId);
      expect(cfg[0]).to.equal(uri);
      expect(cfg[1]).to.equal(hash);
    });
  });
});
