# IoTeX W3bstream (Sprout Release 🍀)

W3bstream is an integral part of the IoTeX network. It is a Layer-2 protocol orchestrated by the IoTeX Blockchain, dedicated to facilitating decentralized data processing in blockchain ecosystems. W3bstream nodes fetch raw data messages from supported data infrastructures and process them through project-specific ZK circuits to generate Zero-Knowledge (ZK) Proofs. These proofs are then dispatched to dApps on different blockchains and verified on-chain, enabling dApps to securely act on validated data insights. W3bstream enhances scalability and trust in blockchain applications, particularly where actions depend on the integrity of real-world data, such as in supply chain management, IoT, or any DePIN network where data authenticity triggers significant incentives on the blockchain.
## Architecture

<p align="center">
  <img src="./docs/architecture.drawio.png"/>
</p>

The diagram represents the main components of the software and how they interact with each other. Note that this reflects a single entity running an enode and a znode. However, there are many entities running nodes in the network. More on this later.

- Enode: short for Edge node. An enode contains a sequencer and task dispatcher. It receives messages from users, persists them in DA, and packs them into tasks. The enode is defined by the project, and how messages are packed into tasks is also defined by the project.
- Znode: short for ZK node. A znode receives the task, constructs a ZK runtime instance, and generates a ZK proof. It contains a task processor, ZK runtime manager, project manager, and output module. Anyone can stake IOTX and obtain permission to run a znode.
- Data availability: W3bstream uses data availability to ensure that messages and task lifecycles persist.
- P2P network: In W3bstream, all enodes and znodes interact with each other over the P2P network, including dispatching, receiving, and reporting task status. Every node needs to join the project topic and then process the information related to the project.
- IPFS: Project config data is stored on IPFS. Users who want to publish a new project can use ioctl to push the project config file to IPFS.
- Chain contract: Project metadata and znode information are stored on the chain contract. If a project needs to be loaded by a znode, the znode needs to first fetch the project metadata from the chain contract, and then fetch the project configuration file from IPFS.
- ZK runtime: Currently, W3bstream supports three ZK runtimes: Halo2, ZkWasm, and Risc0. The project configuration defines which runtime will be used by the project.

## Running

Just want to give it a try, see Quick Start →

For the initial setup and operation of a W3bstream node, please refer to the OPERATOR_GUIDE →

Developers looking to build circuits and deploy W3bstream projects should consult the DEVELOPER_GUIDE →

## Contributing

We welcome contributions! Please read our contributing guidelines and submit pull requests to our GitHub repository.

## Community and Support

We encourage you to seek support and ask questions on one of the following platforms:

### Join Our Discord Community

For real-time discussions and community support, join our Discord server where we have a dedicated Developers Lounge category. This is a great place to get quick help, discuss features, and connect with other community members:

[Join the IoTeX Discord →](https://iotex.io/devdiscord)

### Ask on Stack Overflow

For more structured and detailed questions, consider using **Stack Overflow**. Many of IoTeX's core and expert developers prefer this platform for its non-real-time format, which encourages well-structured and comprehensive questions. Ask your question here:

[Stack Overflow - IoTeX Tag →](https://stackoverflow.com/questions/tagged/iotex) and make sure it's tagged [`IOTEX`].
