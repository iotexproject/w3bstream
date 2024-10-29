// // test/W3bstreamTaskManager.test.ts

// import { expect } from "chai";
// import { ethers } from "hardhat";
// import { Contract, BigNumber, HDNodeWallet } from "ethers";
// import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";

// describe("W3bstreamTaskManager", function () {
//     let w3bstreamDebits: Contract;
//     let w3bstreamReward: Contract;
//     let w3bstreamProver: Contract;
//     let w3bstreamTaskManager: Contract;
//     let owner: SignerWithAddress;
//     let operator: SignerWithAddress;
//     let user1: SignerWithAddress;
//     let user2: SignerWithAddress;
//     let token: Contract;
//     let project: Contract;
//     let w3bstreamProject: Contract;
//     let binder: SignerWithAddress;
//     let projectOwner: SignerWithAddress;
//     let projectId: number;
//     let taskOwner: SignerWithAddress;
//     let taskOwnerWallet: HDNodeWallet
//     let sequencer: SignerWithAddress;
//     let prover: SignerWithAddress;

//     beforeEach(async function () {
//         [owner, taskOwner, operator, user1, user2, binder, projectOwner, sequencer, prover] = await ethers.getSigners();
//         const accounts = config.networks.hardhat.accounts;
//         taskOwnerWallet = ethers.HDNodeWallet.fromMnemonic(
//             ethers.Mnemonic.fromPhrase(accounts.mnemonic),
//             "m/44'/60'/0'/0/1",
//         );

//         // Deploy debits contract
//         const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
//         token = await ERC20Mock.deploy(
//             "Mock Token",
//             "MTK",
//             owner.address,
//             ethers.parseEther("10000")
//         );
//         await token.waitForDeployment();
//         const W3bstreamDebits = await ethers.getContractFactory("W3bstreamDebits");
//         w3bstreamDebits = await upgrades.deployProxy(W3bstreamDebits, [], {
//             initializer: "initialize",
//         });
//         await w3bstreamDebits.waitForDeployment();

//         // Deploy project reward contract
//         project = await ethers.deployContract('MockProject');
//         w3bstreamProject = await ethers.deployContract('W3bstreamProject');
//         await project.connect(projectOwner).register();
//         await w3bstreamProject.initialize(project.target);
//         await w3bstreamProject.setBinder(binder.address);
//         projectId = 1;
//         await w3bstreamProject.connect(binder).bind(projectId);
//         await w3bstreamProject.connect(projectOwner).resume(projectId);
//         w3bstreamReward = await ethers.deployContract('W3bstreamProjectReward');
//         await w3bstreamReward.initialize(w3bstreamProject.target);

//         // Deploy prover contract
//         const W3bstreamProver = await ethers.getContractFactory("W3bstreamProver");
//         w3bstreamProver = await W3bstreamProver.deploy();
//         await w3bstreamProver.initialize();

//         // Deploy task manager contract
//         const W3bstreamTaskManager = await ethers.getContractFactory("W3bstreamTaskManager");
//         w3bstreamTaskManager = await W3bstreamTaskManager.deploy();
//         await w3bstreamTaskManager.initialize(
//             w3bstreamDebits.target,
//             w3bstreamReward.target,
//             w3bstreamProver.target
//         );
//         await w3bstreamDebits.connect(owner).setOperator(w3bstreamTaskManager.target);
//     });

//     describe("Initialization", function () {
//         it("should initialize with correct addresses", async function () {
//             expect(await w3bstreamTaskManager.debits()).to.equal(w3bstreamDebits.target);
//             expect(await w3bstreamTaskManager.projectReward()).to.equal(w3bstreamReward.target);
//             expect(await w3bstreamTaskManager.proverStore()).to.equal(w3bstreamProver.target);
//         });
//     });

//     describe("Operator Management", function () {
//         it("should allow owner to add operator", async function () {
//             await expect(w3bstreamTaskManager.connect(owner).addOperator(user1.address))
//                 .to.emit(w3bstreamTaskManager, "OperatorAdded")
//                 .withArgs(user1.address);
//             expect(await w3bstreamTaskManager.operators(user1.address)).to.be.true;
//         });

//         it("should not allow non-owner to add operator", async function () {
//             await expect(w3bstreamTaskManager.connect(user1).addOperator(user2.address))
//                 .to.be.revertedWith("Ownable: caller is not the owner");
//         });

//         it("should allow owner to remove operator", async function () {
//             await w3bstreamTaskManager.connect(owner).addOperator(user1.address);
//             await expect(w3bstreamTaskManager.connect(owner).removeOperator(user1.address))
//                 .to.emit(w3bstreamTaskManager, "OperatorRemoved")
//                 .withArgs(user1.address);
//             expect(await w3bstreamTaskManager.operators(user1.address)).to.be.false;
//         });

//         it("should not allow non-owner to remove operator", async function () {
//             await w3bstreamTaskManager.connect(owner).addOperator(user1.address);
//             await expect(w3bstreamTaskManager.connect(user1).removeOperator(user1.address))
//                 .to.be.revertedWith("Ownable: caller is not the owner");
//         });
//     });

//     describe("Task Assignment", function () {
//         let taskId: string;
//         let taskHash: string;
//         let taskSignature: string;
//         let rewardAmount: BigNumber;

//         beforeEach(async function () {
//             // Set up task assignment
//             taskId = ethers.encodeBytes32String("task1");
//             rewardAmount = ethers.parseEther("1000");

//             // Transfer tokens to taskOwner and deposit into debits contract
//             await token.connect(owner).transfer(taskOwner.address, rewardAmount);
//             await token.connect(taskOwner).approve(w3bstreamDebits.target, rewardAmount);
//             await w3bstreamDebits.connect(taskOwner).deposit(token.target, rewardAmount);

//             // Set reward token and amount
//             await w3bstreamReward.connect(projectOwner).setRewardToken(projectId, token.target);
//             await w3bstreamReward.connect(taskOwner).setReward(projectId, rewardAmount);

//             // Prover registers and sets rebate ratio
//             await w3bstreamProver.connect(prover).register();
//             await w3bstreamProver.connect(prover).setRebateRatio(1000); // 10%

//             // Generate task hash and signature
//             taskHash = ethers.solidityPackedKeccak256(
//                 ["uint256", "bytes32", "address"],
//                 [projectId, taskId, taskOwner.address]
//             );
//             taskSignature = taskOwnerWallet.signingKey.sign(ethers.getBytes(taskHash)).serialized;

//             // Owner adds operator
//             await w3bstreamTaskManager.connect(owner).addOperator(operator.address);
//         });

//         it("should allow operator to assign a task", async function () {
//             const deadline = (await ethers.provider.getBlockNumber()) + 100;

//             const assignment = {
//                 projectId: projectId,
//                 taskId: taskId,
//                 hash: taskHash,
//                 signature: taskSignature,
//                 prover: prover.address,
//             };


//             await expect(
//                 w3bstreamTaskManager.connect(operator)["assign((uint256,bytes32,bytes32,bytes,address),address,uint256)"](assignment, sequencer.address, deadline)
//             )
//                 .to.emit(w3bstreamTaskManager, "TaskAssigned")
//                 .withArgs(projectId, taskId, prover.address, deadline);

//             // Verify record
//             const record = await w3bstreamTaskManager.records(projectId, taskId);
//             expect(record.prover).to.equal(prover.address);
//             expect(record.sequencer).to.equal(sequencer.address);
//             expect(record.deadline).to.equal(deadline);
//             expect(record.owner).to.equal(taskOwner.address);
//             expect(record.settled).to.be.false;
//         });

//         it("should not allow non-operator to assign a task", async function () {
//             const deadline = (await ethers.provider.getBlockNumber()) + 100;

//             const assignment = {
//                 projectId: projectId,
//                 taskId: taskId,
//                 hash: taskHash,
//                 signature: taskSignature,
//                 prover: prover.address,
//             };

//             await expect(
//                 w3bstreamTaskManager.connect(user1)["assign((uint256,bytes32,bytes32,bytes,address),address,uint256)"](assignment, sequencer.address, deadline)
//             )
//                 .to.be.revertedWith("not task manager operator");
//         });
//     });

//     describe("Task Settlement", function () {
//         let taskId: string;
//         let taskHash: string;
//         let taskSignature: string;
//         let rewardAmount: BigNumber;

//         beforeEach(async function () {
//             // Set up task assignment
//             taskId = ethers.encodeBytes32String("task1");
//             rewardAmount = ethers.parseEther("1000");

//             // Transfer tokens to taskOwner and deposit into debits contract
//             await token.connect(owner).transfer(taskOwner.address, rewardAmount);
//             await token.connect(taskOwner).approve(w3bstreamDebits.target, rewardAmount);
//             await w3bstreamDebits.connect(taskOwner).deposit(token.target, rewardAmount);

//             // Set reward token and amount
//             await w3bstreamReward.connect(projectOwner).setRewardToken(projectId, token.target);
//             await w3bstreamReward.connect(taskOwner).setReward(projectId, rewardAmount);

//             // Prover registers and sets rebate ratio
//             await w3bstreamProver.connect(prover).register();
//             await w3bstreamProver.connect(prover).setRebateRatio(1000); // 10%

//             // Generate task hash and signature
//             taskHash = ethers.solidityPackedKeccak256(
//                 ["uint256", "bytes32", "address"],
//                 [projectId, taskId, taskOwner.address]
//             );
//             taskSignature = taskOwnerWallet.signingKey.sign(ethers.getBytes(taskHash)).serialized;

//             // Owner adds operator
//             await w3bstreamTaskManager.connect(owner).addOperator(operator.address);

//             // Assign the task
//             const deadline = (await ethers.provider.getBlockNumber()) + 100;
//             const assignment = {
//                 projectId: projectId,
//                 taskId: taskId,
//                 hash: taskHash,
//                 signature: taskSignature,
//                 prover: prover.address,
//             };

//             await w3bstreamTaskManager.connect(operator)["assign((uint256,bytes32,bytes32,bytes,address),address,uint256)"](assignment, sequencer.address, deadline);
//         });

//         it("should allow operator to settle a task", async function () {
//             await expect(
//                 w3bstreamTaskManager.connect(operator).settle(projectId, taskId, prover.address)
//             )
//                 .to.emit(w3bstreamTaskManager, "TaskSettled")
//                 .withArgs(projectId, taskId, prover.address);

//             // Verify record is marked as settled
//             const record = await w3bstreamTaskManager.records(projectId, taskId);
//             expect(record.settled).to.be.true;
//         });

//         it("should not allow non-operator to settle a task", async function () {
//             await expect(
//                 w3bstreamTaskManager.connect(user1).settle(projectId, taskId, prover.address)
//             )
//                 .to.be.revertedWith("not task manager operator");
//         });

//         it("should not allow settling with invalid prover", async function () {
//             await expect(
//                 w3bstreamTaskManager.connect(operator).settle(projectId, taskId, user1.address)
//             )
//                 .to.be.revertedWith("invalid prover");
//         });
//     });

//     describe("Task Recall", function () {
//         let taskId: string;
//         let taskHash: string;
//         let taskSignature: string;
//         let rewardAmount: BigNumber;

//         beforeEach(async function () {
//             // Set up task assignment
//             taskId = ethers.encodeBytes32String("task1");
//             rewardAmount = ethers.parseEther("1000");

//             // Transfer tokens to taskOwner and deposit into debits contract
//             await token.connect(owner).transfer(taskOwner.address, rewardAmount);
//             await token.connect(taskOwner).approve(w3bstreamDebits.target, rewardAmount);
//             await w3bstreamDebits.connect(taskOwner).deposit(token.target, rewardAmount);

//             // Set reward token and amount
//             await w3bstreamReward.connect(projectOwner).setRewardToken(projectId, token.target);
//             await w3bstreamReward.connect(taskOwner).setReward(projectId, rewardAmount);

//             // Prover registers and sets rebate ratio
//             await w3bstreamProver.connect(prover).register();
//             await w3bstreamProver.connect(prover).setRebateRatio(1000); // 10%

//             // Generate task hash and signature
//             taskHash = ethers.solidityPackedKeccak256(
//                 ["uint256", "bytes32", "address"],
//                 [projectId, taskId, taskOwner.address]
//             );
//             taskSignature = taskOwnerWallet.signingKey.sign(ethers.getBytes(taskHash)).serialized;

//             // Owner adds operator
//             await w3bstreamTaskManager.connect(owner).addOperator(operator.address);

//             // Assign the task
//             const deadline = (await ethers.provider.getBlockNumber()) + 100;
//             const assignment = {
//                 projectId: projectId,
//                 taskId: taskId,
//                 hash: taskHash,
//                 signature: taskSignature,
//                 prover: prover.address,
//             };

//             await w3bstreamTaskManager.connect(operator)["assign((uint256,bytes32,bytes32,bytes,address),address,uint256)"](assignment, sequencer.address, deadline);
//         });

//         it("should allow task owner to recall task after deadline", async function () {
//             // Advance block number beyond deadline
//             await ethers.provider.send("hardhat_mine", ["0x65"]); // Mine 101 blocks

//             await expect(
//                 w3bstreamTaskManager.connect(taskOwner).recall(projectId, taskId)
//             )
//                 .to.emit(w3bstreamDebits, "Redeemed")
//                 .withArgs(
//                     token.target,
//                     taskOwner.address,
//                     rewardAmount
//                 );

//             // Verify record is updated
//             const record = await w3bstreamTaskManager.records(projectId, taskId);
//             expect(record.prover).to.equal(ethers.ZeroAddress);
//             expect(record.deadline).to.equal(0);
//         });

//         it("should not allow task owner to recall task before deadline", async function () {
//             await expect(
//                 w3bstreamTaskManager.connect(taskOwner).recall(projectId, taskId)
//             )
//                 .to.be.revertedWith("task assignement not expired");
//         });

//         it("should not allow non-owner to recall task", async function () {
//             // Advance block number beyond deadline
//             await ethers.provider.send("hardhat_mine", ["0x65"]); // Mine 101 blocks

//             await expect(
//                 w3bstreamTaskManager.connect(user1).recall(projectId, taskId)
//             )
//                 .to.be.revertedWith("not owner");
//         });
//     });
// });
