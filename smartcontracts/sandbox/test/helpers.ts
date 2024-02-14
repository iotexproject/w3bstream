import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { SignerWithAddress } from '@nomicfoundation/hardhat-ethers/signers';

export async function registerOperator(
  operatorRegistry: string,
  profile: Signer,
  node: SignerWithAddress,
  rewards: SignerWithAddress,
) {
  const registry = await ethers.getContractAt('OperatorRegistry', operatorRegistry);
  await registry.connect(profile).registerOperator({
    node: node.address,
    rewards: rewards.address,
  });
}

export async function registerProject(projectRegistry: string, projectOwner: Signer, uri: string, hash: string) {
  const registry = await ethers.getContractAt('ProjectRegistry', projectRegistry);
  await registry.connect(projectOwner).createProject(uri, hash);
}
