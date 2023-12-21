# ZKP circuit

Sprout already supports various zero-knowledge technologies, including Halo2, RISC0, and zkWASM, you can build and deploy them separately.

- [Halo2 ZKP]()
  - [Build Halo2 Circuit](#Build-Halo2-Circuit)
  - [Deploy Halo2 Circuit](#Deploy-Halo2-Circuit)
  - [Verify Halo2 proof](#Verify-Halo2-Proof)
- [RISC0 ZKP]()
  - [Build RISC0 Circuit](#Build-RISC0-Circuit)
  - [Deploy RISC0 circuit](#Deploy-RISC0-Circuit)
  - [Verify RISC0 Proof](#Verify-RISC0-Proof)
- [ZkWASM ZKP]()
  - [Build ZkWASM Circuit](#Build-ZkWASM-Circuit)
  - [Deploy ZkWASM circuit](#Deploy-ZkWASM-Circuit)
  - [Verify ZkWASM Proof](#Verify-ZkWASM-Proof)

# Halo2 ZKP
## Build Halo2 Circuit

```shell
git clone git@github.com:machinefi/sprout.git && cd examples/halo2-circuit
wasm-pack build --target nodejs --out-dir pkg
```
you will find `halo2_simple_bg.wasm` in the `pkg` folder.

If you need to develop and compile your own circuits, please refer to [Halo2 README](./halo2-circuit/README.md) 

## Deploy Halo2 Circuit
The deployment of circuits requires the circuit files to be first compressed using zlib and encoded into hex strings.
Here, we use `ioctl` command, please [install](https://github.com/iotexproject/iotex-core) it.

1. convert circuit to project config file
```shell
ioctl ws code convert -t "halo2" -i "halo2_simple_bg.wasm"
```
This command will generate a file named `halo2-config.json` in the current folder.

2. mv `halo2-config.json` to `test/project` and rename it to `10001`

3. restart docker images `docker-compose up`

## Verify Halo2 Proof
After sending data to w3bstream, you can get a proof. Then you can verify the proof.
The proof can be validated through smart contracts, or can be verified locally.

### Verified by smart contract
1. build
```shell
git clone git@github.com:machinefi/sprout.git && cd examples/halo2-circuit
cargo build --release
```
After this command is successful, a `halo2-simple-circuit` executable file(executable file corresponding to the [simple circuit](./src/circuits/simple.rs)) will be generated in the `target/release` directory.

2. generate verify smart contract
``` shell
target/release/halo2-simple-circuit solidity
```
You will find `Verifier.sol` under the current folder. Or you can run `target/release/halo2-simple-circuit solidity -f path/filename.sol`.
Then you can deploy the smart contract to IoTeX chain or other ETH-compatible chains.

3. get Halo2 proof and invoke smart contract.

### Verified by local
1. build
```shell
git clone git@github.com:machinefi/sprout.git && cd examples/halo2-circuit
cargo build --release
```
After this command is successful, a `halo2-simple-circuit` executable file(executable file corresponding to the [simple circuit](./src/circuits/simple.rs)) will be generated in the `target/release` directory.

2. get Halo2 proof
   if you can send messages to znode successfully, then you can execute `ioctl ws message send --project-id 10001 --project-version "0.1" --data "{\"private_a\": 3, \"private_b\": 4}"` to obtain a halo2 proof, then put it in a file, like `halo2-simple-proof.json`.

3. verify
   `--proof` is proof file, and `--public` is the public input

``` shell
target/release/halo2-simple-circuit verify --proof halo2-simple-proof.json --public 900
```

More details and options for `Halo2 circuit` are given in [its README](./halo2-circuit/README.md).


# RISC0 ZKP
## Build RISC0 Circuit

```shell
git clone git@github.com:machinefi/sprout.git && cd examples/risc0-circuit/method
cargo build --release
```
you will find `methods.rs` in a folder, the folder path will print on the console, like this

If you need to develop and compile your own circuits, please refer to [RISC0 README](./risc0-circuit/README.md).

```shell
warning: methods_path is: "sprout/examples/risc0-circuit/method/target/release/build/risc0-circuits-5efc4ff59af940ab/out/methods.rs"
```
## Deploy RISC0 circuit
The deployment of circuits requires the circuit files to be first compressed using zlib and encoded into hex strings.
Here, we use `ioctl` command, please [install](https://github.com/iotexproject/iotex-core) it.

1. convert circuit to project config file
```shell
ioctl ws code convert -t "risc0" -i "methods.rs"  -e "{\"image_id\":\"RANGE_ID\", \"elf\":\"RANGE_ELF\"}"
```
The values of `image_id` and `elf` are variable names, and will be found in the `methods.rs`.

This command will generate a file named `risc0-config.json` in the current folder.

2. mv `risc0-config.json` to `test/project` and rename it to `10000`

3. restart docker images `docker-compose up`

## Verify RISC0 Proof
After sending data to w3bstream, you can get a proof. Then you can verify the proof.
The proof can be validated through smart contracts, or can be verified locally.

### Verified by smart contract
The RISC0 verify proof has been deployed on the IoTeX mainnet, the address is [io108ecwt37dxmfd4ltltcjdygnpmapy2g7630lcz](https://iotexscan.io/address/io108ecwt37dxmfd4ltltcjdygnpmapy2g7630lcz?format=io#code)

The contract source code is [here](./risc0-circuit/contract).

### Verified by local
1. Get stark proof and image id
   If you have successfully built `methods.rs` in the `method` path, you can find `xx_ID`(this is `RANGE_ID` in example `methods.rs`), the corresponding [u32; 8] array for `xx_ID` without `[]` is `image-id` string.
   If you have successfully deployed `methods.rs` to w3bstream, and you can send messages to znode successfully, then you can execute `ioctl ws message send --project-id 10000 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Stark\"}"` to obtain a stark proof, then put it in a file, like `start-proof.json`.

2. Execute the `docker run` command for local verification. Note that the directory where the proof is located needs to be mounted into the image.
   It's simple, just input the proof file and image-id. You can also use the help command to check how to use it.

```shell
docker run -v /host/stark-proof.json:/stark-proof.json iotexdev/zkverifier:v0.0.1 /verifier/risc0-circuit verify -p /stark-proof.json -i "520991199, 1491489009, 3725421922, 2701107036, 261900524, 710029518, 655219346, 3077599842"
```

More details and options for `RISC0 circuit` are given in [its README](./risc0-circuit/README.md).


# ZkWASM ZKP
## Build ZkWASM Circuit

```shell
git clone git@github.com:machinefi/sprout.git && cd examples/zkwasm-circuit/circuit
asc src/add.ts -O --noAssert -o add.wasm
```

If you need to develop and compile your own circuits, please refer to [zkWASM README](./zkwasm-circuit/README.md)

## Deploy ZkWASM Circuit
The deployment of circuits requires the circuit files to be first compressed using zlib and encoded into hex strings.
Here, we use `ioctl` command, please [install](https://github.com/iotexproject/iotex-core) it.

1. convert circuit to project config file
```shell
ioctl ws code convert -t "zkwasm" -i "add.wasm"
```
This command will generate a file named `zkwasm-config.json` in the current folder.

2. mv `zkwasm-config.json` to `test/project` and rename it to `10002`

3. restart docker images `docker-compose up`

## Verify ZkWASM Proof
After sending data to w3bstream, you can get a proof. Then you can verify the proof.
The proof can be validated through smart contracts, or can be verified locally.

### Verified by smart contract
Developing

### Verified by local
1. Get zkwasm proof
   If you can send messages to znode successfully, then you can execute `ioctl ws message send --project-id 10002 --project-version "0.1" --data "{\"private_input\": [1, 1] , \"public_input\": [2] }"` to obtain a zkwasm proof, then put it in a file, like `zkwasm-proof.json`.

2. Execute the `docker run` command for local verification. Note that the directory where the proof is located needs to be mounted into the image.
   It's simple, just input the proof file. You can also use the help command to check how to use it.

```shell
docker run -v /host/zkwasm-proof.json:/zkwasm-proof.json iotexdev/zkverifier /verifier/zkwasm-circuit verify -p /zkwasm-proof.json
```

More details and options for `zkWasm circuit` are given in [its README](./zkwasm-circuit/README.md).
