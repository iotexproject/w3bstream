# IoTeX Zero-Node (Sprout Release 🍀)

## Welcome to the Zero-Node Protocol Repository

- [Introduction](#-about-zero-node-protocol)
- [Getting Started](#quickstart)
- [Sending data](#send-testing-data-to-znode)
- [Contributing](#contributing)
- [Community & Support](#community-and-support)


#### 🌍 About Zero-Node Protocol

Zero-Node Protocol is an integral part of the [IoTeX network](https://iotex.io). It's a new protocol, dedicated to generating custom Zero-Knowledge (ZK) Proofs on top of machine data, forming a robust backbone for Decentralized Physical Infrastructures (**DePIN**) applications. These proofs are crucial in scaling DePIN data computation and storage, and are key in triggering token economies **based on verifiable proofs of real-world work**.

#### 🔗 Integrating with Blockchains

The Zero-Node Protocol sends these compact, verifiable proofs to various blockchains, activating DePIN token economies upon their verification. [Supported Blockchains →](#supported_blockchains)

#### 🛠 Custom Provers and VM Support

DePIN project owners can utilize native Halo2 circuits to create custom provers. The protocol currently supports RISC0 and zkWASM Virtual Machines for proof generation.

#### Supported Blockchains

Currently, all EVM blockchains are supported as the target for ZNode Proofs.

## Quickstart

### Installation

Install the node command line client `ioctl`:

```bash
brew tap iotexproject/ioctl-unstable
brew install iotexproject/ioctl-unstable/ioctl-unstable
alias ioctl=`which ioctl-unstable`
```

> **_NOTE_**
> `ioctl-unstable` is the latest command-line interface for interacting with IoTeX blockchain
> 
> We will soon release a stable version with full support for interacting with w3bstream node. (WIP)
> 
> For more, see [ioctl command line →](https://docs.iotex.io/the-iotex-stack/wallets/command-line-client)

### Send testing data to znode

The following example sends a message to an example project deployed on the node that makes use of a RISC0 prover, which has project ID 10000:
Project `10000` is an example of a Risc0 circuit, which can prove that an integer `private_input` is within an interval `public_input`.

```bash
ioctl ws message send --project-id 10000 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```

The following example sends a message to an example project deployed on the node that makes use of a Halo2 prover, which has project ID 10001:
Project `10001` is an example of Halo2 circuit, which can prove the product of two integers, `private_a` and `private_b`, multiplied by a constant of 4.

```bash
ioctl ws message send --project-id 10001 --project-version "0.1" --data "{\"private_a\": 3, \"private_b\": 4}"
```

The following example sends a message to an example project deployed on the node that makes use of a Zkwasm prover, which has project ID 10002, this may be slow and may take some time:
Project `10003` is an example of zkWasm circuit, which can prove that two integers 'private_input' are equal.

```bash
ioctl ws message send --project-id 10002 --project-version "0.1" --data "{\"private_input\": [1, 1] , \"public_input\": [] }"
```

### Retrieve ZKP

After znode received the message, a message id will return, like below:

```json
{
  "messageID": "4abbc43a-798f-49e8-bc05-b6baeafec630"
}
```

The following example queries the message status:

```shell
ioctl ws message query --message-id "4abbc43a-798f-49e8-bc05-b6baeafec630"
```

the query result like below:

```json
{
	"messageID": "4abbc43a-798f-49e8-bc05-b6baeafec630",
	"states": [{
		"state": "RECEIVED",
		"time": "2023-11-24T03:41:16.946333Z",
		"description": ""
	}, {
		"state": "FETCHED",
		"time": "2023-11-24T03:41:19.579558Z",
		"description": ""
	}, {
		"state": "PROVING",
		"time": "2023-11-24T03:41:19.59012Z",
		"description": ""
	}, {
		"state": "PROVED",
		"time": "2023-11-24T03:42:23.346377Z",
		"description": "your proof data"
	}, {
		"state": "OUTPUTTING",
		"time": "2023-11-24T03:42:23.357991Z",
		"description": "writing proof to chain"
	}, {
		"state": "OUTPUTTED",
		"time": "2023-11-24T03:42:26.013841Z",
		"description": "your transaction hash"
	}]
}
```

## Advancement
- [build circuit and run znode locally](RUN-LOCALLY.md)

## Contributing

We welcome contributions! Please read our [contributing guidelines](CONTRIBUTING.md) and submit pull requests to our GitHub repository.

## Community and support

We encourage you to seek support and ask questions in one of the following platforms:

#### Join Our Discord Community

For real-time discussions and community support, join our Discord server where we have a dedicated
Developers Lounge category. This is a great place to get quick help, discuss features, and connect with other community members:

[Join the IoTeX Discord →](https://iotex.io/devdiscord)

### Ask on Stack Overflow

For more structured and detailed questions, consider using **Stack Overflow**. Many of IoTeX's core and expert developers prefer this platform for its non-realtime format, which encourages well-structured and comprehensive questions. Ask your question here:

[Stack Overflow - IoTeX Tag →](https://stackoverflow.com/questions/tagged/iotex)

and make sure it's tagged [`IOTEX`].
