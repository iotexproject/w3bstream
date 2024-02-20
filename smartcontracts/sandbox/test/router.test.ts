import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { ethers } from 'hardhat';

import { deployFleetManager, deployW3bstreamRouter } from './deployers';
import { FleetManager, ProjectRegistrar, W3bstreamRouter, WSReceiver } from '../typechain-types';
import { registerOperator, registerProject } from './helpers';

describe('W3bstreamRouter', function () {
  const PROJECT_1_ID = 1;

  describe('data submission', function () {
    let router: W3bstreamRouter;
    let projectRegistry: ProjectRegistrar;
    let fleet: FleetManager;
    let receiver: WSReceiver;

    beforeEach(async function () {
      const [projectOwner, profile, node, rewards] = await ethers.getSigners();

      // 1. deploy operator registry and fleet manager (once)
      fleet = await loadFixture(deployFleetManager);
      // 2. register operator (each new operator) from Profile account
      await registerOperator(await fleet.operatorRegistry(), profile, node, rewards);
      // 3. deploy device registry and WSreceiver (each depin project)
      receiver = await loadFixture(deployWSReceiver);
      const deviceRegistry = await receiver.deviceNFTRegistry();
      // 4. SPROUT upload project config to ipfs (once from sprout node)
      // const { IpfsHash } = await uploadProjectConfig(deviceRegistry, await receiver.getAddress());
      // console.log(IpfsHash);
      const IpfsHash = 'QmcRTH9pmHYihK3ooZRXDSYXMhAX7GSGTF92Ab5Dc4fHPN';
      // 5. register new project, from ProjectOwner account
      const projectRegistryAddr = await fleet.projectRegistry();
      await registerProject(
        projectRegistryAddr,
        projectOwner,
        'ipfs://' + IpfsHash,
        '0x91f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b381389',
      );
      // 6. allow operator in fleet manager, from ProjectOwner account
      await fleet.connect(projectOwner).allow(PROJECT_1_ID, profile.address);
      // 7. deploy and initialize WSRouter
      router = await loadFixture(deployW3bstreamRouter);
      projectRegistry = await ethers.getContractAt('ProjectRegistrar', projectRegistryAddr);
      await router.initialize(await projectRegistry.getAddress(), await fleet.getAddress());
    });

    it('works', async function () {
      const [, , node] = await ethers.getSigners();

      const project = await projectRegistry.projects(PROJECT_1_ID);

      // const config = await getConfigByCID(project.uri.replace('ipfs://', ''));
      // const { id: tunnelId, targetContract: receiverAddr } = config.data.tunnels[0];

      await router.register(PROJECT_1_ID, await receiver.getAddress());

      await expect(
        router
          .connect(node)
          .submit(
            PROJECT_1_ID,
            await receiver.getAddress(),
            '0x91f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b38138992f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b38138993f11349770aadcc135213916bf429e39f7419b25d5fe6a2623115b35b381389',
          ),
      )
        .to.emit(router, 'DataReceived')
        .withArgs(node.address, true, '');

      //   expect(await receiver.getBatchHeight(123)).to.eq(1)
    });
  });
});

const deployDeviceRegistry = async () => {
  const DeviceRegistry = await ethers.getContractFactory('ProjectRegistrar');
  return DeviceRegistry.deploy();
};

const deployWSReceiver = async () => {
  const deviceRegistry = await deployDeviceRegistry();
  const WSReceiver = await ethers.getContractFactory('WSReceiver');
  return WSReceiver.deploy(deviceRegistry);
};
