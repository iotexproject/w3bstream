import { expect } from 'chai';
import { ethers } from 'hardhat';

describe('Fleet Management', function () {
    let fleetManagement;
    let w3bstreamProver;
    beforeEach(async function() {
        fleetManagement = await ethers.deployContract('FleetManagement');
        await fleetManagement.initialize(100);
        w3bstreamProver = await ethers.deployContract('W3bstreamProver');
        await w3bstreamProver.initialize('W3bstream Prover', "W3BProver");
    });
    it('register', async function() {
        const [owner, prover] = await ethers.getSigners();
        await w3bstreamProver.setMinter(fleetManagement.getAddress());
        await fleetManagement.setProverStore(w3bstreamProver.getAddress());
        await fleetManagement.setRegistrationFee(12345);
        expect(await w3bstreamProver.count()).to.equal(0);
        await fleetManagement.connect(prover).register({value: 12345});
        expect(await w3bstreamProver.count()).to.equal(1);
    });

});
