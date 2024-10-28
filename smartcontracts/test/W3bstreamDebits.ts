import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { Contract } from "ethers";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";

describe("W3bstreamDebits", function () {
    let w3bstreamDebits: Contract;
    let owner: SignerWithAddress;
    let operator: SignerWithAddress;
    let user1: SignerWithAddress;
    let user2: SignerWithAddress;
    let recipients: SignerWithAddress[];
    let token: Contract;
    let tokenAddress: string;
    let w3bstreamDebitsAddress: string;

    beforeEach(async function () {
        [owner, operator, user1, user2, ...recipients] = await ethers.getSigners();

        const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
        token = await ERC20Mock.deploy(
            "Mock Token",
            "MTK",
            owner.address,
            ethers.parseEther("10000")
        );
        await token.waitForDeployment();
        tokenAddress = token.target;

        const W3bstreamDebits = await ethers.getContractFactory("W3bstreamDebits");
        w3bstreamDebits = await upgrades.deployProxy(W3bstreamDebits, [], {
            initializer: "initialize",
        });
        await w3bstreamDebits.waitForDeployment();
        w3bstreamDebitsAddress = w3bstreamDebits.target;
    });

    describe("Deposit", function () {
        it("should allow users to deposit tokens", async function () {
            const depositAmount = ethers.parseEther("1000");

            await token.transfer(user1.address, depositAmount);
            await token.connect(user1).approve(w3bstreamDebitsAddress, depositAmount);

            await expect(w3bstreamDebits.connect(user1).deposit(tokenAddress, depositAmount))
                .to.emit(w3bstreamDebits, "Deposited")
                .withArgs(tokenAddress, user1.address, depositAmount);

            expect(await w3bstreamDebits.balanceOf(tokenAddress, user1.address)).to.equal(depositAmount);
        });
    });

    describe("Withhold", function () {
        beforeEach(async function () {
            await w3bstreamDebits.connect(owner).setOperator(operator.address);

            const depositAmount = ethers.parseEther("1000");
            await token.transfer(user1.address, depositAmount);
            await token.connect(user1).approve(w3bstreamDebitsAddress, depositAmount);
            await w3bstreamDebits.connect(user1).deposit(tokenAddress, depositAmount);
        });

        it("should allow operator to withhold tokens", async function () {
            const withholdAmount = ethers.parseEther("500");
            await expect(
                w3bstreamDebits.connect(operator).withhold(tokenAddress, user1.address, withholdAmount)
            )
                .to.emit(w3bstreamDebits, "Withheld")
                .withArgs(tokenAddress, user1.address, withholdAmount);

            expect(await w3bstreamDebits.balanceOf(tokenAddress, user1.address)).to.equal(
                ethers.parseEther("500")
            );
        });
    });

    describe("Redeem", function () {
        beforeEach(async function () {
            await w3bstreamDebits.connect(owner).setOperator(operator.address);

            const depositAmount = ethers.parseEther("1000");
            await token.transfer(user1.address, depositAmount);
            await token.connect(user1).approve(w3bstreamDebitsAddress, depositAmount);
            await w3bstreamDebits.connect(user1).deposit(tokenAddress, depositAmount);

            const withholdAmount = ethers.parseEther("1000");
            await w3bstreamDebits.connect(operator).withhold(tokenAddress, user1.address, withholdAmount);
        });

        it("should allow operator to redeem tokens", async function () {
            const redeemAmount = ethers.parseEther("500");
            await expect(
                w3bstreamDebits.connect(operator).redeem(tokenAddress, user1.address, redeemAmount)
            )
                .to.emit(w3bstreamDebits, "Redeemed")
                .withArgs(tokenAddress, user1.address, redeemAmount);

            expect(await w3bstreamDebits.balanceOf(tokenAddress, user1.address)).to.equal(
                ethers.parseEther("500")
            );
        });
    });

    describe("Withdraw", function () {
        beforeEach(async function () {
            const depositAmount = ethers.parseEther("1000");
            await token.transfer(user1.address, depositAmount);
            await token.connect(user1).approve(w3bstreamDebitsAddress, depositAmount);
            await w3bstreamDebits.connect(user1).deposit(tokenAddress, depositAmount);
        });

        it("should allow user to withdraw", async function () {
            const withdrawAmount = ethers.parseEther("500");

            await expect(w3bstreamDebits.connect(user1).withdraw(tokenAddress, withdrawAmount))
                .to.emit(w3bstreamDebits, "Withdrawn")
                .withArgs(tokenAddress, user1.address, withdrawAmount);

            expect(await token.balanceOf(user1.address)).to.equal(withdrawAmount);
        });
    });
});
