import { ethers } from 'hardhat';

export async function deployFleetManager() {
  const projectRegistry = await deployProjectRegistry();
  const projectRegistryAddress = await projectRegistry.getAddress();

  const operatorRegistry = await deployOperatorRegistry();
  const operatorRegistryAddress = await operatorRegistry.getAddress();

  const factory = await ethers.getContractFactory('FleetManager');
  return factory.deploy(projectRegistryAddress, operatorRegistryAddress);
}

export async function deployProjectRegistry() {
  const factory = await ethers.getContractFactory('ProjectRegistry');
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