import { ethers, upgrades } from 'hardhat';

async function main() {
  const pd = await ethers.deployContract('ProjectDevice', [
    '0x06b3Fcda51e01EE96e8E8873F0302381c955Fddd',
    '0x6AfCB0EB71B7246A68Bb9c0bFbe5cD7c11c4839f',
  ]);
  await pd.waitForDeployment();
  console.log(`ProjectDevice deployed to ${pd.target}`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
