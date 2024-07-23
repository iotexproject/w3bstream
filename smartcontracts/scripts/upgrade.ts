import { ethers, upgrades } from 'hardhat';

async function main() {
  if (process.env.W3BSTREAM_PROJECT) {
    const W3bstreamProject = await ethers.getContractFactory('W3bstreamProject');
    await upgrades.upgradeProxy(process.env.W3BSTREAM_PROJECT, W3bstreamProject, {});
    console.log(`Upgrade W3bstreamProject ${process.env.W3BSTREAM_PROJECT} successfull!`);
  }

  if (process.env.W3BSTREAM_ROUTER) {
    const W3bstreamRouter = await ethers.getContractFactory('W3bstreamRouter');
    await upgrades.upgradeProxy(process.env.W3BSTREAM_ROUTER, W3bstreamRouter, {});
    console.log(`Upgrade W3bstreamRouter ${process.env.W3BSTREAM_ROUTER} successfull!`);
  }

  if (process.env.W3BSTREAM_PROVER) {
    const W3bstreamProver = await ethers.getContractFactory('W3bstreamProver');
    await upgrades.upgradeProxy(process.env.W3BSTREAM_PROVER, W3bstreamProver, {});
    console.log(`Upgrade W3bstreamRouter ${process.env.W3BSTREAM_PROVER} successfull!`);
  }
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
