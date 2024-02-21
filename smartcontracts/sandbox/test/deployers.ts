import { ethers } from 'hardhat';

export async function deployFleetManager() {
  const projectRegistrar = await deployProjectRegistrar();
  const projectRegistrarAddress = await projectRegistrar.getAddress();

  const nodeRegistry = await deployNodeRegistry();
  const nodeRegistryAddress = await nodeRegistry.getAddress();

  const factory = await ethers.getContractFactory('FleetManager');
  const fleet = await factory.deploy();
  await fleet.initialize(projectRegistrarAddress, nodeRegistryAddress);
  return fleet;
}

export async function deployProjectRegistrar() {
  const factory = await ethers.getContractFactory('ProjectRegistrar');
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
