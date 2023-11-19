# IoTeX Zero-Node (Sprout Release 🍀)

## Welcome to the Zero-Node Protocol Repository

- [Introduction](#-about-zero-node-protocol)
- [Getting Started](#getting-started)
- [Running the Node](#running-the-node)
- [Sending data](#sending-data)
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

## Getting Started

### Prerequisites

Ensure you have the following installed:

- **Golang**: Version 1.21 or higher. Download and Install instructions can be found at https://go.dev/doc/install.

- **Docker Engine**: Version 18.02 or higher. Installation instructions can be found at https://docs.docker.com/engine/install/

- **Docker Compose Plugin**: Ensure you have the Compose plugin installed:

  ```bash
  docker compose version
  # Install with: sudo apt install docker-compose-plugin
  ```

- **Blockchain Wallet**: The ZNode will dispatch proofs to a blockchain contract, which requires a funded wallet account on the target blockchain (for IoTeX Testnet, see how to [create a wallet](https://docs.iotex.io/the-iotex-stack/wallets/metamask), and [claim test IOTX](https://docs.iotex.io/the-iotex-stack/iotx-faucets/testnet-tokens#the-iotex-developer-portal))

- **Bonsai API Key**: If you plan to generate RISC0 snark proofs, as the ZNode protocol currently relies on the [Bonsai API](https://dev.risczero.com/api/bonsai/) you'll need to get [their API key](https://docs.google.com/forms/d/e/1FAIpQLSf9mu18V65862GS4PLYd7tFTEKrl90J5GTyzw_d14ASxrruFQ/viewform).

### Installation

1. Download the latest release from [releases page](https://github.com/machinefi/sprout/releases) and unpack the release code (replace with your specific file name):

    ```bash
    tar xzf sprout-x.y.z.tar.gz
    cd sprout-x.y.z.tar.gz
    ```

2. Alternatively, ```git clone https://github.com/machinefi/sprout.git```
    
3. Install the node command line client `wsctl`:
    ```bash
    curl https://raw.githubusercontent.com/machinefi/sprout/main/scripts/install_wsctl.sh | bash
    ```


### Configure the node

> **_NOTE:_**
>
> - RISC Zero is currently the only supported zkVM.
> - EVM chains are currently the only supported target for proofs.

#### Set your blockchain account

1. To enable the node to send proofs to the destination blockchain, configure a funded account on the target chain:

    ```bash
    export PRIVATE_KEY=${your private key}
    ```

2. To use RISC0 Provers for proof generation, provide your Bonsai API Key (see prerequisites above):

    ```bash
    export BONSAI_KEY=${your bonsai key}
    ```

3. Docker Compose will mount the current work directory under the `/data` volume. You can edit the file `docker-compose.yaml` to set `PROJECT_FILE_DIRECTORY` tp the appropriate path where the project configuration file (which includes the prover code) is stored.

## Running the node

Start the ZNode with the following command:

```bash
cd sprout
docker compose up -d
```
### Configure wsctl

Set up the `wsctl` endpoint to your running node (`wsctl`settings are located in `~/.w3bstream/config.yaml``)

```bash
wsctl config set endpoint localhost:9000
```

After that, you can use ```wsctl config get endpoint``` to make sure the config is effective.

### Monitoring and management

Monitor the node status with:

```bash
docker-compose logs -f w3bnode
```

Shut down the node with:

```bash
docker-compose down
```

## Generate ZKP

### Compile the customized Halo2 circuit

1. Install `wasm-pack`
``` shell
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

2. Build wasm

``` shell
cd examples/halo2-circuits/
wasm-pack build --target nodejs --out-dir pkg
```

you will find `halo2_wasm_bg.wasm` under the `pkg` folder. 

3. (Optional) You can also write your circuit according to the [halo2 development documentation](https://zcash.github.io/halo2/user/simple-example.html), and put the circuit file in `src/circuits`; replace the `TODO` in `src/lib.rs` and build wasm with `wasm-pack build --target nodejs --out-dir pkg`.

### Deploy Compiled circuit to W3bstream

TBD

### Send testing data to the server

znode projects are currently placed inside the folder `test/data`. Each project file is composed of a JSON object definition that includes a unique ID for the project, the binary code of the proover, and other parameters.

The following example sends a message to an example project deployed on the node that makes use of a RISC0 prover, which has project ID 10000:

```bash
wsctl message send --project-id 10000 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```

The following example sends a message to an example project deployed on the node that makes use of a Halo2 prover, which has project ID 10001:

```bash
wsctl message send --project-id 10001 --project-version "0.1" --data "{\"private_a\": 3, \"private_b\": 4}"
```

### Retrieve ZKP

TBD

## Contributing

We welcome contributions! Please read our [contributing guidelines](CONTRIBUTING.md) and submit pull requests to our GitHub repository.

After making changes to the code, you can rebuild the Docker image with:

```bash
make docker
```

Shut down the node and ensure you replace the `w3bnode` image name inside `docker-compose.yaml`` with the name:tag of your local image before running the node again.


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
