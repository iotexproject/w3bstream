import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { ethers } from 'hardhat';

import { deployFleetManager } from './deployers';
import { registerOperator, registerProject } from './helpers';

describe('FleetManager', function () {
  it('should be initialized with project and node registry', async function () {
    const fleet = await loadFixture(deployFleetManager);

    expect(await fleet.projectRegistry()).to.not.eq(ethers.ZeroAddress);
    expect(await fleet.nodeRegistry()).to.not.eq(ethers.ZeroAddress);
  });

  const PROJECT_1 = {
    uri: 'project1',
    hash: '0x91f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b381389',
  };
  const PROJECT_ID_1 = 1;
  const NODE_ID_1 = 1;

  describe('isAllowed', function () {
    it('should return true if sender is allowed', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, projectOwner, node, operator] = await ethers.getSigners();

      await registerOperator(await fleet.nodeRegistry(), node, operator);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await fleet.connect(projectOwner).allow(PROJECT_ID_1, NODE_ID_1);

      expect(await fleet.isAllowed(operator.address, PROJECT_ID_1)).to.be.true;
    });
    it('returns false if unexisting project', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, node, operator] = await ethers.getSigners();

      await registerOperator(await fleet.nodeRegistry(), node, operator);

      expect(await fleet.isAllowed(operator.address, PROJECT_ID_1)).to.be.eq(false);
    });
    it('reverts if operator address is zero', async function () {
      const fleet = await loadFixture(deployFleetManager);

      await expect(fleet.isAllowed(ethers.ZeroAddress, 1)).to.be.revertedWithCustomError(
        fleet,
        'InvalidOperatorAddress',
      );
    });
  });
  describe('allow', function () {
    it('works if called by project owner', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, projectOwner, node, operator] = await ethers.getSigners();

      await registerOperator(await fleet.nodeRegistry(), node, operator);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.connect(projectOwner).allow(PROJECT_ID_1, NODE_ID_1))
        .to.emit(fleet, 'NodeAllowed')
        .withArgs(PROJECT_ID_1, NODE_ID_1);

      expect(await fleet.isAllowed(operator.address, PROJECT_ID_1)).to.be.true;
    });
    it('reverts if called by non-owner', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, projectOwner, node, operator] = await ethers.getSigners();

      await registerOperator(await fleet.nodeRegistry(), node, operator);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.allow(PROJECT_ID_1, NODE_ID_1)).to.be.revertedWithCustomError(fleet, 'NotProjectOwner');
    });
    it('reverts if operator is not registered', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, projectOwner] = await ethers.getSigners();

      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.connect(projectOwner).allow(PROJECT_ID_1, NODE_ID_1)).to.be.reverted;
    });
    it('reverts if operator is already allowed', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, projectOwner, node, operator] = await ethers.getSigners();

      await registerOperator(await fleet.nodeRegistry(), node, operator);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await fleet.connect(projectOwner).allow(PROJECT_ID_1, NODE_ID_1);
      await expect(fleet.connect(projectOwner).allow(PROJECT_ID_1, NODE_ID_1)).to.be.revertedWithCustomError(
        fleet,
        'NodeAlreadyAllowed',
      );
    });
  });
  describe('disallow', function () {
    it('works if called by project owner', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, projectOwner, node, operator] = await ethers.getSigners();

      await registerOperator(await fleet.nodeRegistry(), node, operator);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await fleet.connect(projectOwner).allow(PROJECT_ID_1, NODE_ID_1);
      await expect(fleet.connect(projectOwner).disallow(PROJECT_ID_1, NODE_ID_1))
        .to.emit(fleet, 'NodeDisallowed')
        .withArgs(PROJECT_ID_1, NODE_ID_1);

      expect(await fleet.isAllowed(operator.address, PROJECT_ID_1)).to.be.false;
    });
    it('reverts if called by non-owner', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, projectOwner, node, operator] = await ethers.getSigners();

      await registerOperator(await fleet.nodeRegistry(), node, operator);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await fleet.connect(projectOwner).allow(PROJECT_ID_1, NODE_ID_1);
      await expect(fleet.disallow(PROJECT_ID_1, NODE_ID_1)).to.be.revertedWithCustomError(fleet, 'NotProjectOwner');

      expect(await fleet.isAllowed(operator.address, PROJECT_ID_1)).to.be.true;
    });
    it('reverts if operator is not allowed', async function () {
      const fleet = await loadFixture(deployFleetManager);
      const [, projectOwner, node, operator] = await ethers.getSigners();

      await registerOperator(await fleet.nodeRegistry(), node, operator);
      await registerProject(await fleet.projectRegistry(), projectOwner, PROJECT_1.uri, PROJECT_1.hash);

      await expect(fleet.connect(projectOwner).disallow(PROJECT_ID_1, NODE_ID_1)).to.be.revertedWithCustomError(
        fleet,
        'NodeNotAllow',
      );
    });
  });
});
