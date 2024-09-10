import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('W3bstream Router', function () {
  let w3bstreamRouter;
  let fleetManagement;
  let w3bstreamCredit;
  let project;
  let w3bstreamProject;
  let w3bstreamProver;
  let mockProcessor;
  let mockStakingHub;
  beforeEach(async function () {
    w3bstreamRouter = await ethers.deployContract('W3bstreamRouter');
    fleetManagement = await ethers.deployContract('FleetManagement');
    w3bstreamCredit = await ethers.deployContract('W3bstreamCredit');
    project = await ethers.deployContract('MockProject');
    w3bstreamProject = await ethers.deployContract('W3bstreamProject');
    w3bstreamProver = await ethers.deployContract('W3bstreamProver');
    mockProcessor = await ethers.deployContract('MockProcessor');
    mockStakingHub = await ethers.deployContract('MockStakingHub');
    await w3bstreamRouter.initialize(fleetManagement.getAddress(), w3bstreamProject.getAddress());
    await fleetManagement.initialize(100);
    await w3bstreamProject.initialize(project.target);
    await w3bstreamProver.initialize('W3bstream Prover', 'W3BProver');
    await w3bstreamCredit.initialize('W3bstreamCredit', 'W3BC');
  });
  it('route', async function () {
    const [owner, binder, coordinator, projectOwner, prover] = await ethers.getSigners();
    await w3bstreamCredit.setMinter(fleetManagement.getAddress());
    await w3bstreamProver.setMinter(fleetManagement.getAddress());
    await fleetManagement.setCoordinator(coordinator.address);
    await fleetManagement.setStakingHub(mockStakingHub.getAddress());
    await fleetManagement.setProverStore(w3bstreamProver.getAddress());
    await fleetManagement.setCreditCenter(w3bstreamCredit.getAddress());
    await fleetManagement.setRegistrationFee(12345);
    await fleetManagement.connect(prover).register({ value: 12345 });
    await mockStakingHub.setAmount(100);
    await w3bstreamProject.setBinder(binder.address);

    await project.connect(projectOwner).register();

    const projectId = 1;
    await w3bstreamProject.connect(binder).bind(projectId);

    const proverId = 1;
    const clientId = '_testClientId';
    await w3bstreamProver.connect(prover).resume(proverId);
    await w3bstreamProject.connect(projectOwner).resume(projectId);
    await w3bstreamRouter.connect(projectOwner).bindDapp(projectId, mockProcessor.getAddress());
    await w3bstreamRouter.connect(coordinator).route(projectId, proverId, clientId, '0x0000');
  });
});
