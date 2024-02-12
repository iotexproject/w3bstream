import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { ethers, network } from 'hardhat';

interface Operator {
  name: string;
  profile: string;
  node: string;
  rewards: string;
}

describe('OperatorRegistry', function () {
  const OPERATOR_1: Operator = {
    name: 'operator',
    profile: '0x67C3968A4F517D94EA319251c52CbAE0485bC6Bb',
    node: '0xFAa61239C2C95CF4Ae669BF1BF9Fd559207078f4',
    rewards: '0xFdeeBb88985ae60b5DaA95626953b538e552bD52',
  };

  const OPERATOR_2: Operator = {
    name: 'operator2',
    profile: '0x193e75b60A4Ca8BC842Dc28604Afc6c41aFE972A',
    node: '0x17e49637113da9A63004C545894b23A3434976b0',
    rewards: '0x0236cc9daBcD2c3CB2dAcE7a183EDfA553Ef4405',
  };
  const ID_0 = 0;
  const ID_1 = 1;

  describe('Operator registration', function () {
    it('should add an operator', async function () {
      const registry = await loadFixture(deployContractRegistry);

      let firstOperator = await registry.operators(ID_0);
      expect(firstOperator.name).to.eq('');

      await registry.registerOperator(OPERATOR_1);
      firstOperator = await registry.operators(ID_0);

      expect(firstOperator.name).to.not.eq('');
      expect(firstOperator.name).to.eq(OPERATOR_1.name);
      expect(firstOperator.profile).to.eq(OPERATOR_1.profile);
      expect(firstOperator.node).to.eq(OPERATOR_1.node);
      expect(firstOperator.rewards).to.eq(OPERATOR_1.rewards);
    });
    it('should register multiple operators', async function () {
      const registry = await loadFixture(deployContractRegistry);

      await registry.registerOperator(OPERATOR_1);

      let secondOperator = await registry.operators(ID_1);
      expect(secondOperator.name).to.be.eq('');

      await registry.registerOperator(OPERATOR_2);
      secondOperator = await registry.operators(ID_1);
      expect(secondOperator.name).to.be.eq(OPERATOR_2.name);
    });
    it('should emit an event', async function () {
      const registry = await loadFixture(deployContractRegistry);

      await expect(registry.registerOperator(OPERATOR_1))
        .to.emit(registry, 'OperatorRegistered')
        .withArgs(ID_0, OPERATOR_1.profile, OPERATOR_1.node, OPERATOR_1.rewards, OPERATOR_1.name);
    });
    it('should fail if name missing', async function () {
      const registry = await loadFixture(deployContractRegistry);
      const operatorWithoutProfileAddr: Operator = { ...OPERATOR_1, name: '' };

      await expect(registry.registerOperator(operatorWithoutProfileAddr)).to.be.revertedWith(
        "OperatorRegistry: name can't be empty",
      );
    });
    it('should fail if profileAddress missing', async function () {
      const registry = await loadFixture(deployContractRegistry);
      const operatorWithoutProfileAddr: Operator = { ...OPERATOR_1, profile: '' };

      await expect(registry.registerOperator(operatorWithoutProfileAddr)).to.be.rejected;
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
  });

  describe('Operator updates', function () {
    describe('Node address', function () {
      it('owner should update node address', async function () {
        const [profile] = await ethers.getSigners();
        const operator: Operator = { ...OPERATOR_1, profile: profile.address };

        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(operator);

        const oldNode = operator.node;
        expect((await registry.operators(ID_0)).node).to.be.eq(oldNode);

        await registry.updateNode(ID_0, OPERATOR_2.node);
        expect((await registry.operators(ID_0)).node).to.be.eq(OPERATOR_2.node);
      });
      it('reverts if not owner tries update operators node', async function () {
        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(OPERATOR_1);

        await expect(registry.updateNode(ID_0, OPERATOR_2.node)).to.be.revertedWith(
          'OperatorRegistry: Not operator owner',
        );
      });
      it('reverts if unexisting operator', async function () {
        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(OPERATOR_1);

        await expect(registry.updateNode(ID_1, OPERATOR_2.node)).to.be.revertedWith(
          'OperatorRegistry: unexistent operator',
        );
      });
      it('rejects if invalid node address', async function () {
        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(OPERATOR_1);

        await expect(registry.updateNode(ID_0, '')).to.be.rejected;
      });
    });
    describe('Rewards address', function () {
      it('owner should update rewards address', async function () {
        const [profile] = await ethers.getSigners();
        const operator: Operator = { ...OPERATOR_1, profile: profile.address };

        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(operator);

        const oldRewards = operator.rewards;
        expect((await registry.operators(ID_0)).rewards).to.be.eq(oldRewards);

        await registry.updateRewards(ID_0, OPERATOR_2.rewards);
        expect((await registry.operators(ID_0)).rewards).to.be.eq(OPERATOR_2.rewards);
      });
      it('reverts if not owner tries update operators rewards', async function () {
        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(OPERATOR_1);

        await expect(registry.updateRewards(ID_0, OPERATOR_2.rewards)).to.be.revertedWith(
          'OperatorRegistry: Not operator owner',
        );
      });
      it('reverts if unexisting operator', async function () {
        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(OPERATOR_1);

        await expect(registry.updateRewards(ID_1, OPERATOR_2.rewards)).to.be.revertedWith(
          'OperatorRegistry: unexistent operator',
        );
      });
      it('rejects if invalid rewards address', async function () {
        const registry = await loadFixture(deployContractRegistry);
        await registry.registerOperator(OPERATOR_1);

        await expect(registry.updateRewards(ID_0, '')).to.be.rejected;
      });
    });
  });
});

async function deployContractRegistry() {
  const OperatorRegistry = await ethers.getContractFactory('OperatorRegistry');
  return OperatorRegistry.deploy();
}
