import { ethers, upgrades } from 'hardhat';

async function main() {
  if (!process.env.PROJECT_ADDRESS) {
    console.log(`Please provide project address`);
    return;
  }
  if (!process.env.PROJECT_REGISTRATION_FEE) {
    console.log(`Please provide project registration fee`);
    return;
  }
  if (!process.env.PROVER_REGISTRATION_FEE) {
    console.log(`Please provide prover registration fee`);
    return;
  }
  if (!process.env.MIN_STAKE) {
    console.log(`Please provide prover min stake`);
    return;
  }

  const [deployer] = await ethers.getSigners();

  const W3bstreamProject = await ethers.getContractFactory('W3bstreamProject');
  const project = await upgrades.deployProxy(W3bstreamProject, [process.env.PROJECT_ADDRESS], {
    initializer: 'initialize',
  });
  await project.waitForDeployment();
  console.log(`W3bstreamProject deployed to ${project.target}`);

  const ProjectRegistrar = await ethers.getContractFactory('ProjectRegistrar');
  const projectRegistrar = await upgrades.deployProxy(ProjectRegistrar, [project.target], {
    initializer: 'initialize',
  });
  await projectRegistrar.waitForDeployment();
  console.log(`ProjectRegistrar deployed to ${projectRegistrar.target}`);
  let tx = await project.setBinder(projectRegistrar.target);
  await tx.wait();
  console.log(`W3bstreamProject binder set to ProjectRegistrar ${projectRegistrar.target}`);
  tx = await projectRegistrar.setRegistrationFee(ethers.parseEther(process.env.PROJECT_REGISTRATION_FEE));
  await tx.wait();
  console.log(`ProjectRegistrar registration fee set to ${process.env.PROJECT_REGISTRATION_FEE}`);

  const W3bstreamProver = await ethers.getContractFactory('W3bstreamProver');
  const prover = await upgrades.deployProxy(W3bstreamProver, ['W3bstream Prover', 'WPRN'], {
    initializer: 'initialize',
  });
  await prover.waitForDeployment();
  console.log(`W3bstreamProver deployed to ${prover.target}`);

  const W3bstreamCredit = await ethers.getContractFactory('W3bstreamCredit');
  const credit = await upgrades.deployProxy(W3bstreamCredit, ['W3bstream Credit', 'WCT'], {
    initializer: 'initialize',
  });
  await credit.waitForDeployment();
  console.log(`W3bstreamCredit deployed to ${credit.target}`);

  const FleetManagement = await ethers.getContractFactory('FleetManagement');
  const fleetManagement = await upgrades.deployProxy(FleetManagement, [ethers.parseEther(process.env.MIN_STAKE)], {
    initializer: 'initialize',
  });
  await fleetManagement.waitForDeployment();
  console.log(`FleetManagement deployed to ${fleetManagement.target}`);

  tx = await credit.setMinter(fleetManagement.target);
  await tx.wait();
  console.log(`W3bstreamCredit minter set to ${fleetManagement.target}`);

  tx = await prover.setMinter(fleetManagement.target);
  await tx.wait();
  console.log(`W3bstreamProver minter set to ${fleetManagement.target}`);

  tx = await fleetManagement.setCreditCenter(credit.target);
  await tx.wait();
  console.log(`FleetManagement set CreditCenter to ${credit.target}`);

  let coordinator = process.env.COORDINATOR;
  if (!coordinator) {
    coordinator = deployer.address;
  }
  tx = await fleetManagement.setCoordinator(coordinator);
  await tx.wait();
  console.log(`FleetManagement set coordinator to ${coordinator}`);

  // TODO deploy mock StakingHub
  const stakingHub = await ethers.deployContract('MockStakingHub', [ethers.parseEther(process.env.MIN_STAKE)]);
  await stakingHub.waitForDeployment();
  console.log(`MockStakingHub deployed to ${stakingHub.target}`);
  tx = await fleetManagement.setStakingHub(stakingHub.target);
  await tx.wait();
  console.log(`FleetManagement set stakingHub to ${stakingHub.target}`);

  tx = await fleetManagement.setProverStore(prover.target);
  await tx.wait();
  console.log(`FleetManagement set ProverStore to ${prover.target}`);
  tx = await fleetManagement.setRegistrationFee(ethers.parseEther(process.env.PROVER_REGISTRATION_FEE));
  await tx.wait();
  console.log(`FleetManagement set prover registration fee to ${process.env.PROVER_REGISTRATION_FEE}`);

  const W3bstreamRouter = await ethers.getContractFactory('W3bstreamRouter');
  const router = await upgrades.deployProxy(W3bstreamRouter, [fleetManagement.target, project.target], {
    initializer: 'initialize',
  });
  await router.waitForDeployment();
  console.log(`W3bstreamRouter deployed to ${router.target}`);

  const W3bstreamVMType = await ethers.getContractFactory('W3bstreamVMType');
  const vmtype = await upgrades.deployProxy(W3bstreamVMType, ['W3bstream VmType', 'WVTN'], {
    initializer: 'initialize',
  });
  await vmtype.waitForDeployment();
  console.log(`W3bstreamVMType deployed to ${vmtype.target}`);

  const W3bstreamDAO = await ethers.getContractFactory('W3bstreamDAO');
  const dao = await upgrades.deployProxy(W3bstreamDAO, ['0x0000000000000000000000000000000000000000000000000000000000000000'], {
    initializer: 'initialize',
  });
  await dao.waitForDeployment();
  console.log(`W3bstreamDAO deployed to ${dao.target}`);

  const W3bstreamTaskManager = await ethers.getContractFactory('W3bstreamTaskManager');
  const taskManager = await upgrades.deployProxy(W3bstreamTaskManager, [], {
    initializer: 'initialize',
  });
  await taskManager.waitForDeployment();
  console.log(`W3bstreamTaskManager deployed to ${taskManager.target}`);

  const W3bstreamBlockRewardDistributor = await ethers.getContractFactory('W3bstreamBlockRewardDistributor');
  const distributor = await upgrades.deployProxy(W3bstreamBlockRewardDistributor, [], {
    initializer: 'initialize',
  });
  await distributor.waitForDeployment();
  console.log(`W3bstreamBlockRewardDistributor deployed to ${distributor.target}`);

  const scrypt = await ethers.deployContract('Scrypt');
  await scrypt.waitForDeployment();
  console.log(`Scrypt deployed to ${scrypt.target}`);

  const headerValidator = await ethers.deployContract('W3bstreamBlockHeaderValidator', [scrypt.target]);
  await headerValidator.waitForDeployment();
  console.log(`W3bstreamBlockHeaderValidator deployed to ${headerValidator.target}`);

  const W3bstreamBlockMinter = await ethers.getContractFactory('W3bstreamBlockMinter');
  const minter = await upgrades.deployProxy(W3bstreamBlockMinter, [dao.target, taskManager.target, distributor.target, headerValidator.target], {
    initializer: 'initialize',
  });
  await minter.waitForDeployment();
  console.log(`W3bstreamBlockMinter deployed to ${minter.target}`);
  tx = await dao.transferOwnership(minter.target);
  await tx.wait();
  console.log(`W3bstreamDAO ownership transferred to ${minter.target}`);

  tx = await headerValidator.setOperator(minter.target);
  await tx.wait();
  console.log(`W3bstreamBlockHeaderValidator add operator to ${minter.target}`);

  tx = await distributor.setOperator(minter.target);
  await tx.wait();
  console.log(`W3bstreamBlockRewardDistributor add operator to ${minter.target}`);

  tx = await taskManager.addOperator(minter.target);
  await tx.wait();
  console.log(`W3bstreamTaskManager add operator to ${minter.target}`);

  tx = await distributor.setOperator(minter.target);
  await tx.wait();
  console.log(`W3bstreamBlockRewardDistributor set operator to ${minter.target}`);

  tx = await minter.setBlockReward(0);
  await tx.wait();
  console.log(`W3bstreamBlockMinter set block reward to 0`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
