import { ethers, upgrades } from 'hardhat';

async function main() {
  if (process.env.W3BSTREAM_PROJECT) {
    const W3bstreamProject = await ethers.getContractFactory('W3bstreamProject');
    await upgrades.upgradeProxy(process.env.W3BSTREAM_PROJECT, W3bstreamProject, {});
    console.log(`Upgrade W3bstreamProject ${process.env.W3BSTREAM_PROJECT} successfull!`);
  }
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
