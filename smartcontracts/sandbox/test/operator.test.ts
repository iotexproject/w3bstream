import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { ethers } from 'hardhat';

interface Operator {
  node: string;
  rewards: string;
}

describe('OperatorRegistry', function () {
  const OPERATOR_1: Operator = {
    node: '0xFAa61239C2C95CF4Ae669BF1BF9Fd559207078f4',
    rewards: '0xFdeeBb88985ae60b5DaA95626953b538e552bD52',
  };

  const OPERATOR_2: Operator = {
    node: '0x17e49637113da9A63004C545894b23A3434976b0',
    rewards: '0x0236cc9daBcD2c3CB2dAcE7a183EDfA553Ef4405',
  };

  describe('registering new operator', function () {
    it('works with a single operator', async function () {
      const registry = await loadFixture(deployContractRegistry);
      const [profile] = await ethers.getSigners();

      let firstOperator = await registry.getOperator(profile.address);
      expect(firstOperator.node).to.eq(ethers.ZeroAddress);

      await expect(registry.registerOperator(OPERATOR_1))
        .to.emit(registry, 'OperatorRegistered')
        .withArgs(profile.address, OPERATOR_1.node, OPERATOR_1.rewards);

      firstOperator = await registry.getOperator(profile.address);
      expect(firstOperator.node).to.eq(OPERATOR_1.node);
      expect(firstOperator.rewards).to.eq(OPERATOR_1.rewards);
    });
    it('works with multiple operators', async function () {
      const registry = await loadFixture(deployContractRegistry);
      const [firstProfile, secondProfile] = await ethers.getSigners();

      await registry.connect(firstProfile).registerOperator(OPERATOR_1);

      let secondOperator = await registry.getOperator(secondProfile.address);
      expect(secondOperator.node).to.be.eq(ethers.ZeroAddress);

      await registry.connect(secondProfile).registerOperator(OPERATOR_2);
      secondOperator = await registry.getOperator(secondProfile.address);
      expect(secondOperator.node).to.be.eq(OPERATOR_2.node);
    });
    it('should fail if nodeAddress missing', async function () {
      const registry = await loadFixture(deployContractRegistry);
      const operatorWithoutProfileAddr: Operator = { ...OPERATOR_1, node: '' };

      await expect(registry.registerOperator(operatorWithoutProfileAddr)).to.be.rejected;
    });
    it('should fail if rewardAddress missing', async function () {
      const registry = await loadFixture(deployContractRegistry);
      const operatorWithoutProfileAddr: Operator = { ...OPERATOR_1, rewards: '' };

      await expect(registry.registerOperator(operatorWithoutProfileAddr)).to.be.rejected;
    });
    it('reverts if sender alerady has a registered operator', async function () {
      const registry = await loadFixture(deployContractRegistry);
      await registry.registerOperator(OPERATOR_1);

      await expect(registry.registerOperator(OPERATOR_2)).to.be.revertedWithCustomError(
        registry,
        'OperatorAlreadyRegistered',
      );
    });
  });

  describe('operator updates', function () {
    describe('updating node address', function () {
      it('works when called by owner', async function () {
        const registry = await loadFixture(deployContractRegistry);

        const [profile] = await ethers.getSigners();
        await registry.registerOperator(OPERATOR_1);

        const oldNode = OPERATOR_1.node;
        expect((await registry.getOperator(profile.address)).node).to.be.eq(oldNode);

        await expect(registry.updateNode(OPERATOR_2.node))
          .to.emit(registry, 'OperatorNodeUpdated')
          .withArgs(profile.address, OPERATOR_2.node);

        expect((await registry.getOperator(profile.address)).node).to.be.eq(OPERATOR_2.node);
      });
      it('reverts if unexisting operator', async function () {
        const registry = await loadFixture(deployContractRegistry);

        await expect(registry.updateNode(OPERATOR_2.node)).to.be.revertedWithCustomError(
          registry,
          'UnexistentOperator',
        );
      });
      it('rejects if invalid node address', async function () {
        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(OPERATOR_1);

        await expect(registry.updateNode('')).to.be.rejected;
      });
    });
    describe('updating rewards address', function () {
      it('owner should update rewards address', async function () {
        const registry = await loadFixture(deployContractRegistry);
        const [profile] = await ethers.getSigners();

        await registry.registerOperator(OPERATOR_1);

        expect((await registry.getOperator(profile.address)).rewards).to.be.eq(OPERATOR_1.rewards);

        await expect(registry.updateRewards(OPERATOR_2.rewards))
          .to.emit(registry, 'OperatorRewardsUpdated')
          .withArgs(profile.address, OPERATOR_2.rewards);

        expect((await registry.getOperator(profile.address)).rewards).to.be.eq(OPERATOR_2.rewards);
      });
      it('reverts if unexisting operator', async function () {
        const registry = await loadFixture(deployContractRegistry);

        await expect(registry.updateRewards(OPERATOR_2.rewards)).to.be.revertedWithCustomError(
          registry,
          'UnexistentOperator',
        );
      });
      it('rejects if invalid rewards address', async function () {
        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(OPERATOR_1);

        await expect(registry.updateRewards('')).to.be.rejected;
      });
    });
  });

  describe.skip('Staking', function () {
    it('should stake', async function () {
      const registry = loadFixture(deployContractRegistry);
      // await registry.stake()
    });
  });
  describe.skip('Unstaking', function () {
    it('should unstake', async function () {
      const registry = loadFixture(deployContractRegistry);
      // await registry.unstake()
    });
  });
});

async function deployContractRegistry() {
  const OperatorRegistry = await ethers.getContractFactory('OperatorRegistry');
  return OperatorRegistry.deploy();
}
