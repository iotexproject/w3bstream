import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { ethers } from 'hardhat';

describe.only('ProjectRegistry', function () {
  it('should be initialized', async function () {
    const registry = await loadFixture(deployProjectRegistry);

    expect(await registry.getAddress()).to.not.eq(ethers.ZeroAddress);
  });

  const PROJECT_1 = {
    uri: 'project1',
    hash: '0x91f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b381389',
  };
  const PROJECT_2 = {
    uri: 'project2',
    hash: '0x91f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b381388',
  };

  const ID_1 = 1;

  const OPERATOR_1 = '0x0236cc9daBcD2c3CB2dAcE7a183EDfA553Ef4405';
  const OPERATOR_2 = '0x17e49637113da9A63004C545894b23A3434976b0';
  const OPERATOR_3 = '0x193e75b60A4Ca8BC842Dc28604Afc6c41aFE972A';

  describe('project creation', function () {
    it('works with a single project', async function () {
      const registry = await loadFixture(deployProjectRegistry);
      await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

      const firstProject = await registry.projects(ID_1);

      expect(firstProject.uri).to.eq(PROJECT_1.uri);
      expect(firstProject.hash).to.eq(PROJECT_1.hash);
      expect(firstProject.paused).to.eq(false);
    });
    it('works with multiple projects', async function () {
      const registry = await loadFixture(deployProjectRegistry);
      await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);
      await registry.createProject(PROJECT_2.uri, PROJECT_2.hash);

      const firstProject = await registry.projects(ID_1);
      const secondProject = await registry.projects(ID_1 + 1);

      expect(firstProject.uri).to.eq(PROJECT_1.uri);
      expect(secondProject.uri).to.eq(PROJECT_2.uri);
    });
  });
  describe('project update: ', function () {
    describe('adding operator', function () {
      it('works with a single operator', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await expect(registry.addOperator(ID_1, OPERATOR_1))
          .to.emit(registry, 'OperatorAdded')
          .withArgs(ID_1, OPERATOR_1);

        const operators = await registry.getOperators(ID_1);

        expect(operators[0]).to.eq(OPERATOR_1);
      });
      it('works with multiple operators', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await registry.addOperator(ID_1, OPERATOR_1);
        await registry.addOperator(ID_1, OPERATOR_2);

        const operators = await registry.getOperators(ID_1);

        expect(operators[0]).to.eq(OPERATOR_1);
        expect(operators[1]).to.eq(OPERATOR_2);
      });
      it('reverts if not project owner', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        const [, notOwner] = await ethers.getSigners();

        await expect(registry.connect(notOwner).addOperator(ID_1, OPERATOR_1)).to.be.revertedWith(
          'ProjectRegistry: Only the owner can perform this action',
        );
      });
      it('reverts if already was added', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await registry.addOperator(ID_1, OPERATOR_1);

        await expect(registry.addOperator(ID_1, OPERATOR_1)).to.be.revertedWith(
          'ProjectRegistry: Operator already added',
        );
      });
      it('reverts if project doesnt exist', async function () {
        const registry = await loadFixture(deployProjectRegistry);

        await expect(registry.addOperator(ID_1, OPERATOR_1)).to.be.revertedWith('ERC721: invalid token ID');
      });
    });
    describe('removing operator', function () {
      it('works with a single removal', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await registry.addOperator(ID_1, OPERATOR_1);
        await registry.addOperator(ID_1, OPERATOR_2);

        await expect(registry.removeOperator(ID_1, OPERATOR_1))
          .to.emit(registry, 'OperatorRemoved')
          .withArgs(ID_1, OPERATOR_1);

        const operators = await registry.getOperators(ID_1);

        expect(operators[0]).to.eq(OPERATOR_2);
      });
      it('works with multiple removals', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await registry.addOperator(ID_1, OPERATOR_1);
        await registry.addOperator(ID_1, OPERATOR_2);
        await registry.addOperator(ID_1, OPERATOR_3);

        await registry.removeOperator(ID_1, OPERATOR_1);
        await registry.removeOperator(ID_1, OPERATOR_2);

        const operators = await registry.getOperators(ID_1);

        expect(operators[0]).to.eq(OPERATOR_3);
      });
      it('reverts when operators array is empty', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await expect(registry.removeOperator(ID_1, OPERATOR_1)).to.be.revertedWith(
          'ProjectRegistry: Operator not found',
        );
      });
      it('reverts if not project owner', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        const [, notOwner] = await ethers.getSigners();

        await expect(registry.connect(notOwner).removeOperator(ID_1, OPERATOR_1)).to.be.revertedWith(
          'ProjectRegistry: Only the owner can perform this action',
        );
      });
      it('reverts if project doesnt exist', async function () {
        const registry = await loadFixture(deployProjectRegistry);

        await expect(registry.removeOperator(ID_1, OPERATOR_1)).to.be.revertedWith('ERC721: invalid token ID');
      });
      it('reverts if operator is not added', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await expect(registry.removeOperator(ID_1, OPERATOR_1)).to.be.revertedWith(
          'ProjectRegistry: Operator not found',
        );
      });
    });
    describe('pausing project', function () {
      it('works when called by owner', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        const project = await registry.projects(ID_1);
        expect(project.paused).to.eq(false);

        await expect(registry.pauseProject(ID_1)).to.emit(registry, 'ProjectPaused').withArgs(ID_1);

        const firstProject = await registry.projects(ID_1);
        expect(firstProject.paused).to.eq(true);
      });
      it('works when called by operator', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        const [, operator] = await ethers.getSigners();

        await registry.addOperator(ID_1, operator.address);

        const project = await registry.projects(ID_1);
        expect(project.paused).to.eq(false);

        await expect(registry.connect(operator).pauseProject(ID_1)).to.emit(registry, 'ProjectPaused').withArgs(ID_1);

        const firstProject = await registry.projects(ID_1);
        expect(firstProject.paused).to.eq(true);
      });
      it('reverts if not project operator', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        const [, notOperator] = await ethers.getSigners();

        await expect(registry.connect(notOperator).pauseProject(ID_1)).to.be.revertedWith(
          'ProjectRegistry: Not authorized to operate this project',
        );
      });
      it('reverts if project doesnt exist', async function () {
        const registry = await loadFixture(deployProjectRegistry);

        await expect(registry.pauseProject(ID_1)).to.be.revertedWith('ERC721: invalid token ID');
      });
      it('reverts if already paused', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await registry.pauseProject(ID_1);

        await expect(registry.pauseProject(ID_1)).to.be.revertedWith('ProjectRegistry: Project already paused');
      });
    });
    describe('unpausing project', function () {
      it('works when called by owner', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await registry.pauseProject(ID_1);

        const project = await registry.projects(ID_1);
        expect(project.paused).to.eq(true);

        await expect(registry.unpauseProject(ID_1)).to.emit(registry, 'ProjectUnpaused').withArgs(ID_1);

        const firstProject = await registry.projects(ID_1);
        expect(firstProject.paused).to.eq(false);
      });
      it('works when called by operator', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await registry.pauseProject(ID_1);

        const [, operator] = await ethers.getSigners();

        await registry.addOperator(ID_1, operator.address);

        const project = await registry.projects(ID_1);
        expect(project.paused).to.eq(true);

        await expect(registry.connect(operator).unpauseProject(ID_1))
          .to.emit(registry, 'ProjectUnpaused')
          .withArgs(ID_1);

        const firstProject = await registry.projects(ID_1);
        expect(firstProject.paused).to.eq(false);
      });
      it('reverts if not project operator', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await registry.pauseProject(ID_1);

        const [, notOperator] = await ethers.getSigners();

        await expect(registry.connect(notOperator).unpauseProject(ID_1)).to.be.revertedWith(
          'ProjectRegistry: Not authorized to operate this project',
        );
      });
      it('reverts if project doesnt exist', async function () {
        const registry = await loadFixture(deployProjectRegistry);

        await expect(registry.unpauseProject(ID_1)).to.be.revertedWith('ERC721: invalid token ID');
      });
      it('reverts if already unpaused', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await expect(registry.unpauseProject(ID_1)).to.be.revertedWith('ProjectRegistry: Project is not paused');
      });
    });
    describe('updating project uri', function () {
      const NEW_URI = 'newUri';
      const NEW_HASH = '0x91f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b381388';

      it('works when called by owner', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await expect(registry.updateProject(ID_1, NEW_URI, NEW_HASH))
          .to.emit(registry, 'ProjectUpserted')
          .withArgs(ID_1, NEW_URI, NEW_HASH);

        const firstProject = await registry.projects(ID_1);
        expect(firstProject.uri).to.eq(NEW_URI);
      });
      it('works when called by operator', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        const [, operator] = await ethers.getSigners();

        await registry.addOperator(ID_1, operator.address);

        await expect(registry.connect(operator).updateProject(ID_1, NEW_URI, NEW_HASH))
          .to.emit(registry, 'ProjectUpserted')
          .withArgs(ID_1, NEW_URI, NEW_HASH);

        const firstProject = await registry.projects(ID_1);
        expect(firstProject.uri).to.eq(NEW_URI);
      });
      it('reverts if not project operator', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        const [, notOperator] = await ethers.getSigners();

        await expect(registry.connect(notOperator).updateProject(ID_1, NEW_URI, NEW_HASH)).to.be.revertedWith(
          'ProjectRegistry: Not authorized to operate this project',
        );
      });
      it('reverts if project doesnt exist', async function () {
        const registry = await loadFixture(deployProjectRegistry);

        await expect(registry.updateProject(ID_1, NEW_URI, NEW_HASH)).to.be.revertedWith('ERC721: invalid token ID');
      });
      it('reverts if uri is empty', async function () {
        const registry = await loadFixture(deployProjectRegistry);
        await registry.createProject(PROJECT_1.uri, PROJECT_1.hash);

        await expect(registry.updateProject(ID_1, '', NEW_HASH)).to.be.revertedWith('ProjectRegistry: Invalid URI');
      });
    });
  });
});

async function deployProjectRegistry() {
  const factory = await ethers.getContractFactory('ProjectRegistry');
  return factory.deploy();
}
