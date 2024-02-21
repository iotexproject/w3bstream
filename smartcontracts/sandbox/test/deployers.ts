import { ethers } from 'hardhat';

export async function deployFleetManager() {
  const projectRegistry = await deployProjectRegistry();
  const projectRegistryAddress = await projectRegistry.getAddress();

  const nodeRegistry = await deployNodeRegistry();
  const nodeRegistryAddress = await nodeRegistry.getAddress();

  const factory = await ethers.getContractFactory('FleetManager');
  const fleet = await factory.deploy();
  await fleet.initialize(projectRegistryAddress, nodeRegistryAddress);
  return fleet;
}

export async function deployProjectRegistry() {
  const factory = await ethers.getContractFactory('ProjectRegistry');
  return factory.deploy();
}

export async function deployNodeRegistry() {
  const factory = await ethers.getContractFactory('NodeRegistry');
  return factory.deploy();
}

export async function deployW3bstreamRouter() {
  const factory = await ethers.getContractFactory('W3bstreamRouter');
  return factory.deploy();
}
