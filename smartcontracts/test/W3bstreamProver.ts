import { expect } from "chai";
import { ethers } from "hardhat";
import { Contract } from "ethers";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";

describe("W3bstreamProver", function () {
    let w3bstreamProver: Contract;
    let owner: SignerWithAddress;
    let prover1: SignerWithAddress;
    let prover2: SignerWithAddress;
    let beneficiary: SignerWithAddress;

    beforeEach(async function () {
        // Get signers
        [owner, prover1, prover2, beneficiary] = await ethers.getSigners();

        // Deploy contract
        const W3bstreamProver = await ethers.getContractFactory("W3bstreamProver");
        w3bstreamProver = await W3bstreamProver.deploy();
        await w3bstreamProver.initialize();
    });

    describe("Initialization", function () {
        it("should set the correct owner", async function () {
            expect(await w3bstreamProver.owner()).to.equal(owner.address);
        });
    });

    describe("Registration", function () {
        it("should allow a prover to register", async function () {
            await expect(w3bstreamProver.connect(prover1).register())
                .to.emit(w3bstreamProver, "BeneficiarySet")
                .withArgs(prover1.address, prover1.address);

            expect(await w3bstreamProver.beneficiary(prover1.address)).to.equal(prover1.address);
        });

        it("should not allow double registration", async function () {
            await w3bstreamProver.connect(prover1).register();
            await expect(w3bstreamProver.connect(prover1).register())
                .to.be.revertedWith("already registered");
        });
    });

    describe("VM Types", function () {
        beforeEach(async function () {
            await w3bstreamProver.connect(prover1).register();
        });

        it("should allow adding VM types", async function () {
            const vmType = 1;
            await expect(w3bstreamProver.connect(prover1).addVMType(vmType))
                .to.emit(w3bstreamProver, "VMTypeAdded")
                .withArgs(prover1.address, vmType);

            expect(await w3bstreamProver.isVMTypeSupported(prover1.address, vmType)).to.be.true;
        });

        it("should allow deleting VM types", async function () {
            const vmType = 1;
            await w3bstreamProver.connect(prover1).addVMType(vmType);
            await expect(w3bstreamProver.connect(prover1).delVMType(vmType))
                .to.emit(w3bstreamProver, "VMTypeDeleted")
                .withArgs(prover1.address, vmType);

            expect(await w3bstreamProver.isVMTypeSupported(prover1.address, vmType)).to.be.false;
        });
    });

    describe("Rebate Ratio", function () {
        it("should allow setting rebate ratio", async function () {
            const ratio = 500; // 5%
            await expect(w3bstreamProver.connect(prover1).setRebateRatio(ratio))
                .to.emit(w3bstreamProver, "RebateRatioSet")
                .withArgs(prover1.address, ratio);

            expect(await w3bstreamProver.rebateRatio(prover1.address)).to.equal(ratio);
        });
    });

    describe("Beneficiary Management", function () {
        beforeEach(async function () {
            await w3bstreamProver.connect(prover1).register();
        });

        it("should allow changing beneficiary", async function () {
            await expect(w3bstreamProver.connect(prover1).changeBeneficiary(beneficiary.address))
                .to.emit(w3bstreamProver, "BeneficiarySet")
                .withArgs(prover1.address, beneficiary.address);

            expect(await w3bstreamProver.beneficiary(prover1.address)).to.equal(beneficiary.address);
        });

        it("should not allow setting zero address as beneficiary", async function () {
            await expect(w3bstreamProver.connect(prover1).changeBeneficiary(ethers.ZeroAddress))
                .to.be.revertedWith("zero address");
        });
    });

    describe("Pause/Resume", function () {
        beforeEach(async function () {
            await w3bstreamProver.connect(prover1).register();
        });

        it("should allow pausing", async function () {
            await expect(w3bstreamProver.connect(prover1).pause())
                .to.emit(w3bstreamProver, "ProverPaused")
                .withArgs(prover1.address);

            expect(await w3bstreamProver.isPaused(prover1.address)).to.be.true;
        });

        it("should not allow double pausing", async function () {
            await w3bstreamProver.connect(prover1).pause();
            await expect(w3bstreamProver.connect(prover1).pause())
                .to.be.revertedWith("already paused");
        });

        it("should allow resuming", async function () {
            await w3bstreamProver.connect(prover1).pause();
            await expect(w3bstreamProver.connect(prover1).resume())
                .to.emit(w3bstreamProver, "ProverResumed")
                .withArgs(prover1.address);

            expect(await w3bstreamProver.isPaused(prover1.address)).to.be.false;
        });

        it("should not allow resuming when not paused", async function () {
            await expect(w3bstreamProver.connect(prover1).resume())
                .to.be.revertedWith("already actived");
        });
    });

    describe("Multiple Provers", function () {
        it("should handle multiple provers independently", async function () {
            // Register both provers
            await w3bstreamProver.connect(prover1).register();
            await w3bstreamProver.connect(prover2).register();

            // Set different VM types
            await w3bstreamProver.connect(prover1).addVMType(1);
            await w3bstreamProver.connect(prover2).addVMType(2);

            // Set different rebate ratios
            await w3bstreamProver.connect(prover1).setRebateRatio(500);
            await w3bstreamProver.connect(prover2).setRebateRatio(1000);

            // Verify independent state
            expect(await w3bstreamProver.isVMTypeSupported(prover1.address, 1)).to.be.true;
            expect(await w3bstreamProver.isVMTypeSupported(prover1.address, 2)).to.be.false;
            expect(await w3bstreamProver.isVMTypeSupported(prover2.address, 1)).to.be.false;
            expect(await w3bstreamProver.isVMTypeSupported(prover2.address, 2)).to.be.true;
            expect(await w3bstreamProver.rebateRatio(prover1.address)).to.equal(500);
            expect(await w3bstreamProver.rebateRatio(prover2.address)).to.equal(1000);
        });
    });
});