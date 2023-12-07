# Run Locally

## Prerequisites

- **Build Circuit**: [build circuit](BUILD-CIRCUIT.md) first  

Ensure you have the following installed:

- **Docker Engine**: Version 18.02 or higher. Installation instructions can be found at https://docs.docker.com/engine/install/

- **Docker Compose Plugin**: Ensure you have the Compose plugin installed:

```bash
docker compose version
# Install with: sudo apt install docker-compose-plugin
```

- **Blockchain Wallet**: W3bstream will dispatch proofs to a blockchain contract, which requires a funded wallet account on the target blockchain (for IoTeX Testnet, see how to [create a wallet](https://docs.iotex.io/the-iotex-stack/wallets/metamask), and [claim test IOTX](https://docs.iotex.io/the-iotex-stack/iotx-faucets/testnet-tokens#the-iotex-developer-portal))

- **Bonsai API Key**: If you plan to generate RISC0 snark proofs, as the ZNode protocol currently relies on the [Bonsai API](https://dev.risczero.com/api/bonsai/) you'll need to get [their API key](https://docs.google.com/forms/d/e/1FAIpQLSf9mu18V65862GS4PLYd7tFTEKrl90J5GTyzw_d14ASxrruFQ/viewform).

## Create project config file
after [build circuit](BUILD-CIRCUIT.md)
- move `risc0-config.json` to `test/project, and then rename `risc0-config.json` to `20000`(`20000` is project id).  
- move `halo2-config.json` to `test/project, and then rename `halo2-config.json` to `20001`(`20001` is project id).
- move `zkwasm-config.json` to `test/project, and then rename `zkwasm-config.json` to `20002`(`20002` is project id).

## Configure the node

### Set your blockchain account

1. To enable the node to send proofs to the destination blockchain, configure a funded account on the target chain:

```bash
export PRIVATE_KEY=${your private key}
```

2. To use RISC0 Provers for proof generation, provide your Bonsai API Key (see prerequisites above):

```bash
export BONSAI_KEY=${your bonsai key}
```

3. Docker Compose will mount the current work directory under the `/data` volume. You can edit the file `docker-compose.yaml` to set `PROJECT_FILE_DIRECTORY` tp the appropriate path where the project configuration file (which includes the prover code) is stored.

## Running W3bstream

Start the W3bstream with the following command:

```bash
cd sprout
docker compose up -d
```
### Configure ioctl

Set up the `ioctl` w3bstream endpoint to your running W3bstream instance (`ioctl` settings are located in `$HOME/.config/ioctl/default/config.default`)

 ```bash
 ioctl config set wsEndpoint 'localhost:9000'
 ```

After that, you can use ```ioctl config get wsEndpoint``` to make sure the config is effective.

### Monitoring and management

Monitor the W3bstream instance status with:

```bash
docker-compose logs -f w3bznode
```

Shut down the W3bstream instance with:

```bash
docker-compose down
```

## Send testing data to the W3bstream instance

W3bstream projects are currently placed inside the folder `test/project`. Each project file name is a unique ID for the project. And each project file is composed of a JSON object definition that includes the binary code of the proover, vm type, and other parameters.

The following example sends a message to an example project deployed on the W3bstream instance that makes use of a RISC0 prover, which has project ID 20000, please change the project ID to yours if necessary:

```bash
ioctl ws message send --project-id 20000 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```
