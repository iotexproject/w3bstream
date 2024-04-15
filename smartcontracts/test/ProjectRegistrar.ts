import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('Project Registrar', function () {
    let projectRegistrar;
    let w3bstreamProject;
    beforeEach(async function() {
        w3bstreamProject = await ethers.deployContract('W3bstreamProject');
        await w3bstreamProject.initialize('W3bstream Project', 'W3BProject');
        projectRegistrar = await ethers.deployContract('ProjectRegistrar');
        await projectRegistrar.initialize(w3bstreamProject.getAddress());
        await projectRegistrar.setRegistrationFee(12345);
    });
    it('register', async function() {
        const [owner, projectOwner] = await ethers.getSigners();
        await w3bstreamProject.setMinter(projectRegistrar.getAddress());
        expect(await w3bstreamProject.count()).to.equal(0);
        await projectRegistrar.connect(projectOwner).register({value: 12345});
        expect(await w3bstreamProject.count()).to.equal(1);
        expect(await w3bstreamProject.ownerOf(1)).to.equal(projectOwner.address);
    });
});
