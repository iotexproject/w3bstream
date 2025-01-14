import { ethers, upgrades } from 'hardhat';

async function main() {
  const GeodnetDapp = await ethers.deployContract('GeodnetDapp', [process.env.MOVEMENT_VERIFIER_ADDR, process.env.MARSHALDAO_TICKER_ADDR]);
  await GeodnetDapp.waitForDeployment();
  console.log(`GeodnetDapp deployed to ${GeodnetDapp.target}`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
