# Run Locally

## Prerequisites

Ensure you have the following installed:

- **Docker Engine**: Version 18.02 or higher. Installation instructions can be found at https://docs.docker.com/engine/install/

- **Docker Compose Plugin**: Ensure you have the Compose plugin installed:

  ```bash
  docker compose version
  # Install with: sudo apt install docker-compose-plugin
  ```

- **Blockchain Wallet**: The ZNode will dispatch proofs to a blockchain contract, which requires a funded wallet account on the target blockchain (for IoTeX Testnet, see how to [create a wallet](https://docs.iotex.io/the-iotex-stack/wallets/metamask), and [claim test IOTX](https://docs.iotex.io/the-iotex-stack/iotx-faucets/testnet-tokens#the-iotex-developer-portal))

- **Bonsai API Key**: If you plan to generate RISC0 snark proofs, as the ZNode protocol currently relies on the [Bonsai API](https://dev.risczero.com/api/bonsai/) you'll need to get [their API key](https://docs.google.com/forms/d/e/1FAIpQLSf9mu18V65862GS4PLYd7tFTEKrl90J5GTyzw_d14ASxrruFQ/viewform).

## Get Repository
```bash
git clone https://github.com/machinefi/sprout.git
cd sprout
```

## Generate ZKP

### Compile the customized Halo2 circuit

1. Install `wasm-pack`
```bash
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

2. Build wasm

```bash
cd examples/halo2-circuits/
wasm-pack build --target nodejs --out-dir pkg
```

you will find `halo2_wasm_bg.wasm` under the `pkg` folder.

3. (Optional) You can also write your circuit according to the [halo2 development documentation](https://zcash.github.io/halo2/user/simple-example.html), and put the circuit file in `src/circuits`; replace the `TODO` in `src/lib.rs` and build wasm with `wasm-pack build --target nodejs --out-dir pkg`.

More details and options for `Halo2 circuit` are given in [its README](./examples/halo2-circuits/README.md).

### Compile the customized Risc0 circuits

1. Build

```bash
cd examples/risc0-circuits/
cargo build --release
```

The path of `methods.rs` will be printed to the console, like this  

```bash
warning: methods_path is: "sprout/examples/risc0-circuits/target/release/build/risc0-circuits-5efc4ff59af940ab/out/methods.rs"
```

More details and options for `Risc0 circuit` are given in [its README](./examples/risc0-circuits/README.md).

### Compile the customized zkWasm circuits

1. Build

```bash
cd examples/zkwasm-circuits/
asc src/add.ts -O --noAssert -o zkwasm_demo.wasm
```

More details and options for `zkWasm circuit` are given in [its README](./examples/zkwasm-circuits/README.md).


### Deploy Compiled circuit to W3bstream

#### Deploy halo2 circuit to W3bstream

```bash
ioctl ws code convert -t "halo2" -i "halo2_wasm_bg.wasm"
```

This command will generate a file named `halo2-config.json` in the current folder.
Or you can run `ioctl ws code convert -t "halo2" -i "halo2_wasm_bg.wasm" -o "path/filename.json"`

#### Deploy risc0 circuit to W3bstream

```bash
ioctl ws code convert -t "risc0" -i "methods.rs"  -e "{\"image_id\":\"RANGE_ID\", \"elf\":\"RANGE_ELF\"}"
```
The values of `image_id` and `elf` are variable names, and will be found in the `methods.rs`.

This command will generate a file named `risc0-config.json` in the current folder.
Or you can run `ioctl ws code convert -t "risc0" -i "methods.rs" -o "path/filename.json" -e "{\"image_id\":\"RANGE_ID\", \"elf\":\"RANGE_ELF\"}`

#### Deploy zkwasm circuit to W3bstream

```bash
ioctl ws code convert -t "zkwasm" -i "zkwasm_demo.wasm"
```

This command will generate a file named `zkwasm-config.json` in the current folder.
Or you can run `ioctl ws code convert -t "zkwasm" -i "zkwasm_demo.wasm" -o "path/filename.json"`


> **_NOTE:_**
> move `risc0-config.json` to `test/project, and then rename `risc0-config.json` to `20000`(`20000` is project id).  

> **_NOTE:_**
> move `halo2-config.json` to `test/project, and then rename `halo2-config.json` to `20001`(`20001` is project id).

> **_NOTE:_**
> move `zkwasm-config.json` to `test/project, and then rename `zkwasm-config.json` to `20002`(`20002` is project id).

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

## Running the node

Start the ZNode with the following command:

```bash
cd sprout
docker compose up -d
```
### Configure ioctl

Set up the `ioctl` w3bstream endpoint to your running node (`ioctl` settings are located in `$HOME/.config/ioctl/default/config.default`)

```bash
ioctl config set wsEndpoint 'localhost:9000'
```

After that, you can use ```ioctl config get wsEndpoint``` to make sure the config is effective.

### Monitoring and management

Monitor the node status with:

```bash
docker-compose logs -f w3bnode
```

Shut down the node with:

```bash
docker-compose down
```

## Send testing data to the server

znode projects are currently placed inside the folder `test/project`. Each project file name is a unique ID for the project. And each project file is composed of a JSON object definition that includes the binary code of the proover, vm type, and other parameters.

The following example sends a message to an example project deployed on the node that makes use of a RISC0 prover, which has project ID 20000, please change the project ID to yours if necessary:

```bash
ioctl ws message send --project-id 20000 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```

## w3bstream project management

You need deploy project register contract to **IOTX** before operating w3bstream projects.

### Create project

```sh
ioctl ws project --contract-address $PROJECT_REGISTER_CONTRACT_ADDR create --project-uri $PROJECT_URI --project-hash $PROJECT_HASH ## the project id will be retrieved. 
```

### Update project

```sh
ioctl ws project --contract-address $PROJECT_REGISTER_CONTRACT_ADDR update --project-id $PROJECT_ID --project-uri $PROJECT_URI --project-hash $PROJECT_HASH
```

### Query project

```sh
ioctl ws project --contract-address $PROJECT_REGISTER_CONTRACT_ADDR query --project-id $PROJECT_ID
```
