import { ethers, upgrades } from 'hardhat';

async function main() {
  if (process.env.W3BSTREAM_PROJECT) {
    const W3bstreamProject = await ethers.getContractFactory('W3bstreamProject');
    await upgrades.forceImport(process.env.W3BSTREAM_PROJECT, W3bstreamProject);
    await upgrades.upgradeProxy(process.env.W3BSTREAM_PROJECT, W3bstreamProject, {});
    console.log(`Upgrade W3bstreamProject ${process.env.W3BSTREAM_PROJECT} successfull!`);
  }

  if (process.env.W3BSTREAM_ROUTER) {
    const W3bstreamRouter = await ethers.getContractFactory('W3bstreamRouter');
    await upgrades.forceImport(process.env.W3BSTREAM_ROUTER, W3bstreamRouter);
    await upgrades.upgradeProxy(process.env.W3BSTREAM_ROUTER, W3bstreamRouter, {});
    console.log(`Upgrade W3bstreamRouter ${process.env.W3BSTREAM_ROUTER} successfull!`);
  }

  if (process.env.W3BSTREAM_PROVER) {
    const W3bstreamProver = await ethers.getContractFactory('W3bstreamProver');
    await upgrades.forceImport(process.env.W3BSTREAM_PROVER, W3bstreamProver);
    await upgrades.upgradeProxy(process.env.W3BSTREAM_PROVER, W3bstreamProver, {});
    console.log(`Upgrade W3bstreamRouter ${process.env.W3BSTREAM_PROVER} successfull!`);
  }

  if (process.env.W3BSTREAM_VMTYPE) {
    const W3bstreamVMType = await ethers.getContractFactory('W3bstreamVMType');
    await upgrades.forceImport(process.env.W3BSTREAM_VMTYPE, W3bstreamVMType);
    await upgrades.upgradeProxy(process.env.W3BSTREAM_VMTYPE, W3bstreamVMType, {});
    console.log(`Upgrade W3bstreamVMType ${process.env.W3BSTREAM_VMTYPE} successfull!`);
  }

  if (process.env.W3BSTREAM_DEBITS) {
    const W3bstreamDebits = await ethers.getContractFactory('W3bstreamDebits');
    await upgrades.forceImport(process.env.W3BSTREAM_DEBITS, W3bstreamDebits);
    await upgrades.upgradeProxy(process.env.W3BSTREAM_DEBITS, W3bstreamDebits, {});
    console.log(`Upgrade W3bstreamDebits ${process.env.W3BSTREAM_DEBITS} successfull!`);
  }

  if (process.env.W3BSTREAM_TASK_MANAGER) {
    const W3bstreamTaskManager = await ethers.getContractFactory('W3bstreamTaskManager');
    await upgrades.forceImport(process.env.W3BSTREAM_TASK_MANAGER, W3bstreamTaskManager);
    await upgrades.upgradeProxy(process.env.W3BSTREAM_TASK_MANAGER, W3bstreamTaskManager, {});
    console.log(`Upgrade W3bstreamTaskManager ${process.env.W3BSTREAM_TASK_MANAGER} successfull!`);
  }

  if (process.env.W3BSTREAM_MINTER) {
    const W3bstreamBlockMinter = await ethers.getContractFactory('W3bstreamBlockMinter');
    await upgrades.forceImport(process.env.W3BSTREAM_MINTER, W3bstreamBlockMinter);
    await upgrades.upgradeProxy(process.env.W3BSTREAM_MINTER, W3bstreamBlockMinter, {});
    console.log(`Upgrade W3bstreamBlockMinter ${process.env.W3BSTREAM_MINTER} successfull!`);
  }
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
