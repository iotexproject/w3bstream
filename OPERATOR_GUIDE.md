# IoTeX W3bstream (Node Operator Guide)

W3bstream is a permissionless, decentralized protocol within the IoTeX Network, where node operators contribute computing power to support verifiable computations for blockchain applications. These applications rely on insights from real-world data to trigger their token economies. Anyone can become a W3bstream Node Operator in the IoTeX Network, choosing which dApps to support in processing data and generating ZK (Zero Knowledge) Proofs. This guide covers how to operate a W3bstream node, register it in the IoTeX Network, join specific projects, and claim rewards.

## Run a W3bstream Node

The recommended method to run a W3bstream node is using official Docker images from IoTeX.

### Prerequisites

- Docker Engine (version 18.02 or higher):

    Check your Docker version:

  ```bash
    docker version
    ```

    Installation instructions → <https://docs.docker.com/engine/install/>

- Docker Compose Plugin
  
  Verify Docker Compose installation:

  ```bash
    docker compose version
    # Install with → sudo apt install docker-compose-plugin
  ```

- **Blockchain Wallet**: A funded wallet on the target blockchain is required for your W3bstream node to dispatch proofs to blockchain contracts. For IoTeX Testnet, see [create a wallet](https://docs.iotex.io/the-iotex-stack/wallets/metamask), and [claim test IOTX](https://docs.iotex.io/the-iotex-stack/iotx-faucets/testnet-tokens#the-iotex-developer-portal)

- **Bonsai API Key**: If you are joining a project requiring RISC0 snark proofs, as the W3bstream protocol currently utilizes the [Bonsai API](https://dev.risczero.com/api/bonsai/), obtain an [API key here](https://docs.google.com/forms/d/e/1FAIpQLSf9mu18V65862GS4PLYd7tFTEKrl90J5GTyzw_d14ASxrruFQ/viewform).

### Download Docker Images

Fetch the latest stable docker-compose.yaml:

```bash
curl https://raw.githubusercontent.com/machinefi/sprout/release/docker-compose.yaml > docker-compose.yaml
```

Pull the required images:

```bash
docker compose pull
```

### Configure your blockchain account

To enable your node to send proofs to the destination blockchain, set up a funded account on the target chain:

```bash
export PRIVATE_KEY=${your private key}
```

### Optional: Provide your BONSAI API Key

For projects using RISC0 Provers, supply your Bonsai API Key:

```bash
export BONSAI_KEY=${your bonsai key}
```

Refer to the W3bstream project documentation for the dApp you are joining to determine if Risc Zero proofs are required.

### Manage the node

To start W3bstream, run the following command in the directory containing `docker-compose.yaml`:

```bash
docker compose up -d
```

Monitor the W3bstream instance status:

```bash
docker-compose logs -f enode znode
```

To shut down the W3bstream instance:

```bash
docker-compose down
```

### Interacting with the node

Install **ioctl**: The command-line interface for interacting with the IoTeX blockchain.

```bash
brew tap iotexproject/ioctl-unstable
brew install iotexproject/ioctl-unstable/ioctl-unstable
alias ioctl=`which ioctl-unstable`
```

set the ioctl's wsEndpoint configuration option to your node endpoint:

```bash
ioctl config set wsEndpoint localhost:9000
```

[More on the IoTeX ioctl client →](https://docs.iotex.io/the-iotex-stack/wallets/command-line-client)

Test W3bstream projects are already registered into project contract.

#### Sending messages to the node

Send a message to a RISC0-based test project (ID 1):

```bash
ioctl ws message send --project-id 1 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```

Send a message to the Halo2-based test project (ID 2):

```bash
ioctl ws message send --project-id 2 --project-version "0.1" --data "{\"private_a\": 3, \"private_b\": 4}"
```

Send a message to a zkWasm-based test project (ID 3):

```bash
ioctl ws message send --project-id 3 --project-version "0.1" --data "{\"private_input\": [1, 1] , \"public_input\": [] }"
```

#### Query the status of a proof request

After sending a message, you'll receive a message ID as a response from the node, e.g.,

```json
{
 "messageID": "4abbc43a-798f-49e8-bc05-b6baeafec630"
}
```

you can quesry the history of the proof request with:

```bash
ioctl ws message query --message-id "4abbc43a-798f-49e8-bc05-b6baeafec630"
```

example result:

```json
{
 "messageID": "4abbc43a-798f-49e8-bc05-b6baeafec630",
 "states": [{
   "state": "received",
   "time": "2023-12-06T16:11:03.498785+08:00",
   "comment": ""
  },
  {
   "state": "fetched",
   "time": "2023-12-06T16:11:04.663608+08:00",
   "comment": ""
  },
  {
   "state": "proving",
   "time": "2023-12-06T16:11:04.664008+08:00",
   "comment": ""
  }
 ]
}
```

When the request is in "proved" state, you can check out the node logs to find out the hash of the blockchain transaction that wrote the proof to the destination chain.
