import { ethers, upgrades } from 'hardhat';

async function main() {
    if (!process.env.PROJECT_REGISTRATION_FEE) {
        console.log(`Please provide project registration fee`);
        return;
    }
    if (!process.env.PROVER_REGISTRATION_FEE) {
        console.log(`Please provide prover registration fee`);
        return;
    }
    if (!process.env.MIN_STAKE) {
        console.log(`Please provide prover min stake`);
        return;
    }

    const [deployer] = await ethers.getSigners();

    const W3bstreamProject = await ethers.getContractFactory('W3bstreamProject');
    const project = await upgrades.deployProxy(
        W3bstreamProject,
        ['W3bstream Project', 'WPN'],
        {
            initializer: 'initialize',
        },
    );
    await project.waitForDeployment();
    console.log(`W3bstreamProject deployed to ${project.target}`);

    const ProjectRegistrar = await ethers.getContractFactory('ProjectRegistrar');
    const projectRegistrar = await upgrades.deployProxy(
        ProjectRegistrar,
        [project.target],
        {
            initializer: 'initialize',
        },
    );
    await projectRegistrar.waitForDeployment();
    console.log(`ProjectRegistrar deployed to ${projectRegistrar.target}`);
    let tx = await project.setMinter(projectRegistrar.target);
    await tx.wait();
    console.log(`W3bstreamProject minter set to ProjectRegistrar ${projectRegistrar.target}`);
    tx = await projectRegistrar.setRegistrationFee(ethers.parseEther(process.env.PROJECT_REGISTRATION_FEE));
    await tx.wait();
    console.log(`ProjectRegistrar registration fee set to ${process.env.PROJECT_REGISTRATION_FEE}`);
    
    const W3bstreamProver = await ethers.getContractFactory('W3bstreamProver');
    const prover = await upgrades.deployProxy(
        W3bstreamProver,
        ['W3bstream Prover', 'WPRN'],
        {
            initializer: 'initialize',
        },
    );
    await prover.waitForDeployment();
    console.log(`W3bstreamProver deployed to ${prover.target}`);

    const W3bstreamCredit = await ethers.getContractFactory('W3bstreamCredit');
    const credit = await upgrades.deployProxy(
        W3bstreamCredit,
        ['W3bstream Credit', 'WCT'],
        {
            initializer: 'initialize',
        },
    );
    await credit.waitForDeployment();
    console.log(`W3bstreamCredit deployed to ${credit.target}`);
    
    const FleetManagement = await ethers.getContractFactory('FleetManagement');
    const fleetManagement = await upgrades.deployProxy(
        FleetManagement,
        [ethers.parseEther(process.env.MIN_STAKE)],
        {
            initializer: 'initialize',
        },
    );
    await fleetManagement.waitForDeployment();
    console.log(`FleetManagement deployed to ${fleetManagement.target}`);
    
    tx = await credit.setMinter(fleetManagement.target);
    await tx.wait();
    console.log(`W3bstreamCredit minter set to ${fleetManagement.target}`);
    
    tx = await prover.setMinter(fleetManagement.target);
    await tx.wait();
    console.log(`W3bstreamProver minter set to ${fleetManagement.target}`);
    
    tx = await fleetManagement.setCreditCenter(credit.target);
    await tx.wait();
    console.log(`FleetManagement set CreditCenter to ${credit.target}`);
    
    // TODO setCoordinator,setStakingHub
    
    tx = await fleetManagement.setProverStore(prover.target);
    await tx.wait();
    console.log(`FleetManagement set ProverStore to ${prover.target}`);
    tx = await fleetManagement.setRegistrationFee(ethers.parseEther(process.env.PROVER_REGISTRATION_FEE));
    await tx.wait();
    console.log(`FleetManagement set prover registration fee to ${process.env.PROVER_REGISTRATION_FEE}`);

    const W3bstreamRouter = await ethers.getContractFactory('W3bstreamRouter');
    const router = await upgrades.deployProxy(
        W3bstreamRouter,
        [fleetManagement.target, project.target],
        {
            initializer: 'initialize',
        },
    );
    await router.waitForDeployment();
    console.log(`W3bstreamRouter deployed to ${router.target}`);
}

main().catch(err => {
    console.error(err);
    process.exitCode = 1;
});
