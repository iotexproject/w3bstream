import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { SignerWithAddress } from '@nomicfoundation/hardhat-ethers/signers';

describe('FleetManager', function () {
  it('should be initialized with project and operator registry', async function () {
    const fleet = await loadFixture(deployFleetManager);

    expect(await fleet.projectRegistry()).to.not.eq(ethers.ZeroAddress);
    expect(await fleet.operatorRegistry()).to.not.eq(ethers.ZeroAddress);
  });

  const PROJECT_1 = {
    uri: 'project1',
    hash: '0x91f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b381389',
  };
  const PROJECT_ID_1 = 1;

  describe('isAllowed', function () {
    it('should return true if sender is allowed', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [projectOwner, profile, node, rewards] = await ethers.getSigners();

      await registerOperator(await fleet.operatorRegistry(), profile, node, rewards);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await fleet.connect(projectOwner).allow(PROJECT_ID_1, profile.address);

      expect(await fleet.isAllowed(node.address, PROJECT_ID_1)).to.be.true;
    });
    it('returns false if unexisting project', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, node] = await ethers.getSigners();

      expect(await fleet.isAllowed(node.address, 1)).to.be.eq(false);
    });
    it('reverts if node address is zero', async function () {
      const fleet = await loadFixture(deployFleetManager);

      await expect(fleet.isAllowed(ethers.ZeroAddress, 1)).to.be.revertedWithCustomError(fleet, 'InvalidNodeAddress');
    });
  });
  describe('allow', function () {
    it('works if called by project owner', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [projectOwner, profile, node, rewards] = await ethers.getSigners();

      await registerOperator(await fleet.operatorRegistry(), profile, node, rewards);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.connect(projectOwner).allow(PROJECT_ID_1, profile.address))
        .to.emit(fleet, 'OperatorAdded')
        .withArgs(PROJECT_ID_1, profile.address);

      expect(await fleet.isAllowed(node.address, PROJECT_ID_1)).to.be.true;
    });
    it('reverts if called by non-owner', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [projectOwner, profile] = await ethers.getSigners();

      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.connect(profile).allow(PROJECT_ID_1, profile.address)).to.be.revertedWithCustomError(
        fleet,
        'NotProjectOwner',
      );
    });
    it('reverts if operator is not registered', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [projectOwner, profile] = await ethers.getSigners();

      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.connect(projectOwner).allow(PROJECT_ID_1, profile.address)).to.be.revertedWithCustomError(
        fleet,
        'OperatorNotRegistered',
      );
    });
    it('reverts if operator is already allowed', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [projectOwner, profile, node, rewards] = await ethers.getSigners();

      await registerOperator(await fleet.operatorRegistry(), profile, node, rewards);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await fleet.connect(projectOwner).allow(PROJECT_ID_1, profile.address);
      await expect(fleet.connect(projectOwner).allow(PROJECT_ID_1, profile.address)).to.be.revertedWithCustomError(
        fleet,
        'OperatorAlreadyAllowed',
      );
    });
  });
  describe('disallow', function () {
    it('works if called by project owner', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [projectOwner, profile, node, rewards] = await ethers.getSigners();

      await registerOperator(await fleet.operatorRegistry(), profile, node, rewards);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await fleet.connect(projectOwner).allow(PROJECT_ID_1, profile.address);
      await expect(fleet.connect(projectOwner).disallow(PROJECT_ID_1, profile.address))
        .to.emit(fleet, 'OperatorRemoved')
        .withArgs(PROJECT_ID_1, profile.address);

      expect(await fleet.isAllowed(node.address, PROJECT_ID_1)).to.be.false;
    });
    it('reverts if called by non-owner', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [projectOwner, profile] = await ethers.getSigners();

      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.connect(profile).disallow(PROJECT_ID_1, profile.address)).to.be.revertedWithCustomError(
        fleet,
        'NotProjectOwner',
      );
    });
    it('reverts if operator is not allowed', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [projectOwner, profile] = await ethers.getSigners();

      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.connect(projectOwner).disallow(PROJECT_ID_1, profile.address)).to.be.revertedWithCustomError(
        fleet,
        'OperatorNotFound',
      );
    });
  });
});

async function deployFleetManager() {
  const projectRegistry = await deployProjectRegistry();
  const projectRegistryAddress = await projectRegistry.getAddress();

  const operatorRegistry = await deployOperatorRegistry();
  const operatorRegistryAddress = await operatorRegistry.getAddress();

  const factory = await ethers.getContractFactory('FleetManager');
  return factory.deploy(projectRegistryAddress, operatorRegistryAddress);
}

async function deployProjectRegistry() {
  const factory = await ethers.getContractFactory('ProjectRegistry');
  return factory.deploy();
}

async function deployOperatorRegistry() {
  const factory = await ethers.getContractFactory('OperatorRegistry');
  return factory.deploy();
}

async function registerOperator(
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

async function registerProject(projectRegistry: string, projectOwner: Signer, uri: string, hash: string) {
  const registry = await ethers.getContractAt('ProjectRegistry', projectRegistry);
  await registry.connect(projectOwner).createProject(uri, hash);
}
