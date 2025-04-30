import { ethers, upgrades } from 'hardhat';

async function main() {
  const GeodnetDapp = await ethers.deployContract('Verifier', []);
  await GeodnetDapp.waitForDeployment();
  console.log(`Geodnet batch deployed to ${GeodnetDapp.target}`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
