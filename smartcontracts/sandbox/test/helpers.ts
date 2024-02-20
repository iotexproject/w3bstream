import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { SignerWithAddress } from '@nomicfoundation/hardhat-ethers/signers';

export async function registerOperator(operatorRegistry: string, node: Signer, operator: SignerWithAddress) {
  const registry = await ethers.getContractAt('NodeRegistry', operatorRegistry);
  await registry.connect(node).register(operator.address);
}

export async function registerProject(projectRegistry: string, projectOwner: Signer, uri: string, hash: string) {
  const registry = await ethers.getContractAt('ProjectRegistrar', projectRegistry);
  await registry.connect(projectOwner).createProject(uri, hash);
}
