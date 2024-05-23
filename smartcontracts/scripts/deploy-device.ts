import { ethers, upgrades } from 'hardhat';

async function main() {
  if (!process.env.IOID_REGISTRY) {
    console.log(`Please provide ioID Registry address`);
    return;
  }
  if (!process.env.W3BSTREAM_PROJECT) {
    console.log(`Please provide w3bstream project address`);
    return;
  }

  const ProjectDevice = await ethers.getContractFactory('ProjectDevice');
  const pd = await upgrades.deployProxy(ProjectDevice, [process.env.IOID_REGISTRY, process.env.W3BSTREAM_PROJECT], {
    initializer: 'initialize',
  });
  await pd.waitForDeployment();
  console.log(`ProjectDevice deployed to ${pd.target}`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
