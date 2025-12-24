import { ethers, upgrades } from 'hardhat';

async function main() {
  const SumVerifier = await ethers.deployContract('SumVerifier', []);
  await SumVerifier.waitForDeployment();

  const DelenDapp = await ethers.deployContract('DelenDapp', [SumVerifier.target]);
  await DelenDapp.waitForDeployment();
  console.log(`DelenDapp deployed to ${DelenDapp.target}`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
