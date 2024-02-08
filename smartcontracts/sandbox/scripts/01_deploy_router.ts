import { ethers, upgrades } from 'hardhat';

async function main() {
  if (!process.env.PROJECT_REGISTRY) {
    console.log(`Please provide project registry address`);
    return;
  }

  const FleetManager = await ethers.getContractFactory('FleetManager');
  const fleetManager = await upgrades.deployProxy(FleetManager, [], {
    initializer: 'initialize',
  });
  console.log(`FleetManager deployed to ${fleetManager.target}`);

  const W3bstreamRouter = await ethers.getContractFactory('W3bstreamRouter');
  const router = await upgrades.deployProxy(W3bstreamRouter, [process.env.PROJECT_REGISTRY, fleetManager.target], {
    initializer: 'initialize',
  });
  console.log(`W3bstreamRouter deployed to ${router.target}`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
