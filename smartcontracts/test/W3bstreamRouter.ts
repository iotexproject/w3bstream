import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('W3bstream Router', function () {
  let w3bstreamRouter;
  let fleetManagement;
  let w3bstreamCredit;
  let w3bstreamProject;
  let w3bstreamProver;
  let mockProcessor;
  let mockStakingHub;
  beforeEach(async function() {
    w3bstreamRouter = await ethers.deployContract('W3bstreamRouter');
    fleetManagement = await ethers.deployContract('FleetManagement');
    w3bstreamCredit = await ethers.deployContract('W3bstreamCredit');
    w3bstreamProject = await ethers.deployContract('W3bstreamProject');
    w3bstreamProver = await ethers.deployContract('W3bstreamProver');
    mockProcessor = await ethers.deployContract('MockProcessor');
    mockStakingHub = await ethers.deployContract('MockStakingHub');
    await w3bstreamRouter.initialize(fleetManagement.getAddress(), w3bstreamProject.getAddress());
    await fleetManagement.initialize(100);
    await w3bstreamProject.initialize('W3bstream Project', 'W3BProject');
    await w3bstreamProver.initialize('W3bstream Prover', "W3BProver");
    await w3bstreamCredit.initialize('W3bstreamCredit', 'W3BC');
  });
  it('route', async function() {
    const [owner, minter, coordinator, projectOwner, prover] = await ethers.getSigners();
    await w3bstreamCredit.setMinter(fleetManagement.getAddress());
    await w3bstreamProver.setMinter(fleetManagement.getAddress());
    await fleetManagement.setCoordinator(coordinator.address);
    await fleetManagement.setStakingHub(mockStakingHub.getAddress());
    await fleetManagement.setProverStore(w3bstreamProver.getAddress());
    await fleetManagement.setRegistrationFee(12345);
    await fleetManagement.connect(prover).register({value: 12345});
    await mockStakingHub.setAmount(100);
    await w3bstreamProject.setMinter(minter.address);
    await w3bstreamProject.connect(minter).mint(projectOwner.address);
    const projectId = 1;
    const proverId = 1;
    await w3bstreamProver.connect(prover).resume(proverId);
    await w3bstreamProject.connect(projectOwner).resume(projectId);
    await w3bstreamRouter.connect(projectOwner).bindDapp(projectId, mockProcessor.getAddress());
    await w3bstreamRouter.connect(coordinator).route(projectId, proverId, "0x0000");
  });
});
