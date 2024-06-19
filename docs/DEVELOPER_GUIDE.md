# IoTeX W3bstream Project Developer Guide

W3bstream significantly enhances scalability and trust for dApps where the token economy depends on verifiable processing of real-world data. By providing a decentralized infrastructure to process raw data and generate custom Zero-Knowledge (ZK) Proofs, W3bstream ensures data authenticity and reliability in dApps' token economies.

## Integrate W3bstream in your dApp

Dapps looking to utilize W3bstream capabilities should:

1. [Create a W3bstream project](#create-a-w3bstream-project)
2. [Test the project](#test-your-w3bstream-project)
3. [Register it on the IoTeX blockchain](#registering-your-project)

### Prerequisites

- ioctl: The command-line interface for interacting with the IoTeX blockchain.

```bash
git clone https://github.com/iotexproject/iotex-core.git
cd iotex-core
make ioctl && mv bin/ioctl __YOUR_SYSTEM_PATH__
```

[More on the IoTeX ioctl client →](https://docs.iotex.io/the-iotex-stack/wallets/command-line-client)

### Create a W3bstream Project

A W3bstream project primarily includes the binary code of the a ZK Prover and the destination contract for dispatching proofs. The steps involve first compiling a zk circuit into a prover using one of the supported ZK frameworks, and then generating a W3bstream project file using **ioctl** command line client.

Start by cloning the W3bstream repository:

```bash
git clone https://github.com/machinefi/sprout.git
cd sprout
```

#### Create a W3bstream Project Using Halo2

>For more details on creating Halo2 circuits see the [Halo2 README](./examples/halo2-circuit/README.md).

Install `wasm-pack`

```bash
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

Write your own circuit following the [halo2 development documentation](https://zcash.github.io/halo2/user/simple-example.html), place it in `src/circuits`, adapt the code corresponding to `TODO` lines in `src/lib.rs` and build the wasm prover with:

```bash
cd examples/halo2-circuit/
wasm-pack build --target nodejs --out-dir pkg
```

The `halo2_simple_bg.wasm` will be located under the `pkg` folder.

Generate the W3bstream project file:

```bash
# Customize the output project file name "$ID" with a unique number
ioctl ws project config -t "halo2" -i "halo2_wasm_bg.wasm" -c "path/$ID"
```

Create the blockchain verifier (Solidity)

``` shell
target/release/halo2-simple-circuit solidity -f path/filename.sol
```

#### Create a W3bstream Project Using zkWASM

For more details on zkWASM circuits see the [zkWASM README](./examples/zkwasm-circuit/README.md).

Ensure you have AssemblyScript installed:

```bash
npm install -g assemblyscript
```

Build the circuit:

```bash
cd examples/zkwasm-circuit/
asc src/add.ts -O --noAssert -o zkwasm_demo.wasm
```

Create the verifier

```bash
# Work in progress
``````

Generate the W3bstream project:

```bash
# Customize the output project file name "$ID" with a unique number
ioctl ws project config -t "zkwasm" -i "zkwasm_demo.wasm" -o "path/ID""
```

#### Create a W3bstream Project Using RISC0

More details and options for `Risc0 circuit` are given in [its README](./examples/risc0-circuit/README.md).

Make sure you have cargo 1.72.0 or higher

   ```bash
   cargo version
   # Update with: rustup update
   ```

Install the rustzero toolchain

   ```bash
   cargo install cargo-risczero
   cargo risczero install
   ```

Build the circuit

```bash
cd examples/risc0-circuit/
cargo build --release
```

The path of `methods.rs` will be printed to the console, like in the output example below:  

```bash
warning: methods_path is: "sprout/examples/risc0-circuits/target/release/build/risc0-circuits-5efc4ff59af940ab/out/methods.rs"
```

Generate the W3bstream Project

```bash
ioctl ws project config -t "risc0" -i "methods.rs" -o "path/filename.json" -e "{\"image_id\":\"RANGE_ID\", \"elf\":\"RANGE_ELF\"}
```

The values of `image_id` and `elf` are variable names, and will be found in the `methods.rs` file.

### Test Your W3bstream Project

Once you have generated a W3bstream project file that includes a custom prover for your dApp, you might want to test it.

Please refer to the [OPERATOR GUIDE](./OPERATOR_GUIDE.md) for instructions on how to:

1. Run a W3bstream node locally.
2. Copy the W3bstream project file into the node's project directory (default location is ./test/project).
3. Run the node and send your test messages.

### Registering Your Project

To allow W3bstream node operators to download your project and compute ZK proofs for your dApp, you must register your W3bstream project on the IoTeX blockchain:

#### Acquire a Project ID

```bash
ioctl ioid register "your project name"
```

#### Register Project

```bash
ioctl ws project register --id "your project id"
```

#### Use the Project File Generated above and Update Project Config

```bash
ioctl ws project update --id "your project id" --path "path/to/project_file"
```

#### Start the Project

```bash
ioctl ws project resume --id "your project id"
```

#### Retrieve Project Info

```bash
ioctl ws project query --id "your project id"
```

#### Set Required Prover Amount of the Project

The default prover amount will process the project's tasks is one. And we can customize it by

```bash
ioctl ws project attributes set --id "your project id" --key "RequiredProverAmount" --val "your expected amount"
```

#### Stop the Project

If you want to stop the project's task process, can use this cmd

```bash
ioctl ws project pause --id "your project id"
```
