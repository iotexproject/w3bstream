import { ethers, upgrades } from 'hardhat';

async function main() {
  let projectAddr: string = ''
  let projectRegistrationFee: string = '0.0'

  if (process.env.PROJECT_ADDRESS) {
    projectAddr = process.env.PROJECT_ADDRESS
  } else {
    const MockProject = await ethers.deployContract('MockProject', []);
    await MockProject.waitForDeployment();
    console.log(`MockProject deployed to ${MockProject.target}`);

    projectAddr = MockProject.target.toString()
  }
  if (process.env.PROJECT_REGISTRATION_FEE) {
    projectRegistrationFee = process.env.PROJECT_REGISTRATION_FEE
  }

  const [deployer] = await ethers.getSigners();

  const W3bstreamProject = await ethers.getContractFactory('W3bstreamProject');
  const project = await upgrades.deployProxy(W3bstreamProject, [projectAddr], {
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
  tx = await projectRegistrar.setRegistrationFee(ethers.parseEther(projectRegistrationFee));
  await tx.wait();
  console.log(`ProjectRegistrar registration fee set to ${projectRegistrationFee}`);

  const W3bstreamProver = await ethers.getContractFactory('W3bstreamProver');
  const prover = await upgrades.deployProxy(W3bstreamProver, [], {
    initializer: 'initialize',
  });
  await prover.waitForDeployment();
  console.log(`W3bstreamProver deployed to ${prover.target}`);

  const W3bstreamVMType = await ethers.getContractFactory('W3bstreamVMType');
  const vmtype = await upgrades.deployProxy(W3bstreamVMType, ['W3bstream VmType', 'WVTN'], {
    initializer: 'initialize',
  });
  await vmtype.waitForDeployment();
  console.log(`W3bstreamVMType deployed to ${vmtype.target}`);

  const W3bstreamDebits = await ethers.getContractFactory('W3bstreamDebits');
  const debits = await upgrades.deployProxy(W3bstreamDebits, [], {
    initializer: 'initialize',
  });
  await debits.waitForDeployment();
  console.log(`W3bstreamDebits deployed to ${debits.target}`);

  const W3bstreamProjectReward = await ethers.getContractFactory('W3bstreamProjectReward');
  const projectReward = await upgrades.deployProxy(W3bstreamProjectReward, [project.target], {
    initializer: 'initialize',
  });
  await projectReward.waitForDeployment();
  console.log(`W3bstreamProjectReward deployed to ${projectReward.target}`);

  const W3bstreamTaskManager = await ethers.getContractFactory('W3bstreamTaskManager');
  const taskManager = await upgrades.deployProxy(W3bstreamTaskManager, [debits.target, projectReward.target, prover.target], {
    initializer: 'initialize',
  });
  await taskManager.waitForDeployment();
  console.log(`W3bstreamTaskManager deployed to ${taskManager.target}`);

  const W3bstreamRouter = await ethers.getContractFactory('W3bstreamRouter');
  const router = await upgrades.deployProxy(W3bstreamRouter, [taskManager.target, prover.target, project.target], {
    initializer: 'initialize',
  });
  await router.waitForDeployment();
  console.log(`W3bstreamRouter deployed to ${router.target}`);

  const W3bstreamRewardDistributor = await ethers.getContractFactory('W3bstreamRewardDistributor');
  const distributor = await upgrades.deployProxy(W3bstreamRewardDistributor, [], {
    initializer: 'initialize',
  });
  await distributor.waitForDeployment();
  console.log(`W3bstreamRewardDistributor deployed to ${distributor.target}`);

  const W3bstreamMinter = await ethers.getContractFactory('W3bstreamMinter');
  const minter = await upgrades.deployProxy(W3bstreamMinter, [taskManager.target, distributor.target], {
    initializer: 'initialize',
  });
  await minter.waitForDeployment();
  console.log(`W3bstreamMinter deployed to ${minter.target}`);

  tx = await distributor.setOperator(minter.target);
  await tx.wait();
  console.log(`W3bstreamRewardDistributor add operator to ${minter.target}`);

  tx = await taskManager.addOperator(minter.target);
  await tx.wait();
  console.log(`W3bstreamTaskManager add operator to ${minter.target}`);

  tx = await taskManager.addOperator(router.target);
  await tx.wait();
  console.log(`W3bstreamTaskManager add operator to ${router.target}`);

  tx = await distributor.setOperator(minter.target);
  await tx.wait();
  console.log(`W3bstreamRewardDistributor set operator to ${minter.target}`);

  tx = await debits.setOperator(minter.target);
  await tx.wait();
  console.log(`W3bstreamDebits set operator to ${minter.target}`);

  tx = await minter.setReward(0);
  await tx.wait();
  console.log(`W3bstreamMinter set reward to 0`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
