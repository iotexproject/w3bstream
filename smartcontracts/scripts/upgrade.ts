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
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
