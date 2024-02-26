import { ethers, upgrades } from 'hardhat';

async function main() {
  if (!process.env.RISC0_VERIFY) {
    console.log(`Please provide risc0 verify address`);
    return;
  }

  const MockRisc0SnarkReceiver = await ethers.getContractFactory('MockRisc0SnarkReceiver');
  const receiver = await upgrades.deployProxy(MockRisc0SnarkReceiver, [process.env.RISC0_VERIFY], {
    initializer: 'initialize',
  });
  console.log(`MockRisc0SnarkReceiver deployed to ${receiver.target}`);
}

main().catch(err => {
  console.error(err);
  process.exitCode = 1;
});
