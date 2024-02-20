import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { ethers } from 'hardhat';

import { deployOperatorRegistry } from './deployers';

describe('NodeRegistry', function () {
  const TOKEN_ID_1 = 1;

  describe('registering new node', function () {
    it('works with a single operator', async function () {
      const registry = await loadFixture(deployOperatorRegistry);
      const [, node, operator] = await ethers.getSigners();

      let firstNode = await registry.getNode(0);
      expect(firstNode.active).to.eq(false);

      await expect(registry.connect(node).register(operator.address))
        .to.emit(registry, 'NodeRegistered')
        .withArgs(node.address, TOKEN_ID_1, operator.address);

      firstNode = await registry.getNode(TOKEN_ID_1);
      expect(firstNode.active).to.eq(true);
      expect(firstNode.operator).to.eq(operator.address);
    });
    it('works with multiple operators', async function () {
      const registry = await loadFixture(deployOperatorRegistry);
      const [, node1, operator1, node2, operator2] = await ethers.getSigners();

      await registry.connect(node1).register(operator1.address);

      await expect(registry.getNodeByOperator(operator2.address)).to.be.revertedWithCustomError(
        registry,
        'OperatorUnregister',
      );

      await registry.connect(node2).register(operator2.address);
      const res = await registry.getNodeByOperator(operator2);
      expect(res[1].active).to.be.eq(true);
    });
    it('reverts if sender alerady has a registered operator', async function () {
      const [, node, operator] = await ethers.getSigners();
      const registry = await loadFixture(deployOperatorRegistry);
      await registry.connect(node).register(operator.address);

      await expect(registry.register(operator.address)).to.be.revertedWithCustomError(
        registry,
        'OperatorAlreadyRegistered',
      );
    });
  });

  describe('operator updates', function () {
    describe('updating node address', function () {
      it('works when called by owner', async function () {
        const registry = await loadFixture(deployOperatorRegistry);

        const [, node, operator, operator2] = await ethers.getSigners();
        await registry.connect(node).register(operator.address);

        const oldOperator = operator.address;
        expect((await registry.getNode(TOKEN_ID_1)).operator).to.be.eq(oldOperator);

        await expect(registry.connect(node).updateOperator(TOKEN_ID_1, operator2.address))
          .to.emit(registry, 'NodeUpdated')
          .withArgs(TOKEN_ID_1, operator2.address);

        expect((await registry.getNode(TOKEN_ID_1)).operator).to.be.eq(operator2.address);
      });
      it('reverts if unexisting operator', async function () {
        const registry = await loadFixture(deployOperatorRegistry);
        const [, operator, operator2] = await ethers.getSigners();

        await expect(registry.updateOperator(TOKEN_ID_1, operator2.address)).to.be.reverted;
      });
      it('reverts if invalid operator address', async function () {
        const registry = await loadFixture(deployOperatorRegistry);
        const [, node, operator] = await ethers.getSigners();
        await registry.connect(node).register(operator.address);

        await expect(
          registry.connect(node).updateOperator(TOKEN_ID_1, ethers.ZeroAddress),
        ).to.be.revertedWithCustomError(registry, 'InvalidAddress');
      });
      it('reverts if called not by owner', async function () {
        const registry = await loadFixture(deployOperatorRegistry);
        const [, node, operator] = await ethers.getSigners();
        await registry.connect(node).register(operator.address);

        await expect(registry.updateOperator(TOKEN_ID_1, ethers.ZeroAddress)).to.be.revertedWithCustomError(
          registry,
          'NotNodeOwner',
        );
      });
    });
  });

  describe.skip('Staking', function () {
    it('should stake', async function () {
      const registry = loadFixture(deployOperatorRegistry);
      // await registry.stake()
    });
  });
  describe.skip('Unstaking', function () {
    it('should unstake', async function () {
      const registry = loadFixture(deployOperatorRegistry);
      // await registry.unstake()
    });
  });
});
