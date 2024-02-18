import { ethers } from 'hardhat';

export async function deployFleetManager() {
  const projectRegistry = await deployProjectRegistry();
  const projectRegistryAddress = await projectRegistry.getAddress();

  const operatorRegistry = await deployOperatorRegistry();
  const operatorRegistryAddress = await operatorRegistry.getAddress();

  const factory = await ethers.getContractFactory('FleetManager');
  const fleet = await factory.deploy();
  await fleet.initialize(projectRegistryAddress, operatorRegistryAddress);
  return fleet;
}

export async function deployProjectRegistry() {
  const factory = await ethers.getContractFactory('ProjectRegistrar');
  return factory.deploy();
}

export async function deployOperatorRegistry() {
  const factory = await ethers.getContractFactory('OperatorRegistry');
  return factory.deploy();
}

export async function deployW3bstreamRouter() {
  const factory = await ethers.getContractFactory('W3bstreamRouter');
  return factory.deploy();
}
