# IoTeX W3bstream Project Developer Guide

W3bstream significantly enhances scalability and trust for dApps where the token economy depends on verifiable processing of real-world data. By providing a decentralized infrastructure to process raw data and generate custom Zero-Knowledge (ZK) Proofs, W3bstream ensures data authenticity and reliability in dApps' token economies.

## Integrate W3bstream in your dApp

Dapps looking to utilize W3bstream capabilities should:

1. [Create a W3bstream project](#create-a-w3bstream-project)
2. [Test the project](#test-your-w3bstream-project)
3. [Register it on the IoTeX blockchain](#registering-your-project)
4. [Verify proof on the IoTeX blockchain](#verify-proof-on-chain)

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
git clone https://github.com/iotexproject/w3bstream.git
cd sprout
```

#### Create a W3bstream Project Using Halo2

>For more details on creating Halo2 circuits see the [Halo2 README](./examples/halo2-circuit/README.md).

NOTE: If you want to develop your circuit, please refer [halo2 development documentation](https://zcash.github.io/halo2/user/simple-example.html)

Install `wasm-pack`

```bash
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

build the wasm prover with:

```bash
cd examples/halo2-circuit/
wasm-pack build --target nodejs --out-dir pkg
```

The `halo2_simple_bg.wasm` will be located under the `pkg` folder.

Generate the W3bstream project file:

```bash
# Customize the output project file name "$ID" with a unique number
ioctl ws project config -t 2 -i "halo2_wasm_bg.wasm" -c "path/$ID"
```

Create the blockchain verifier (Solidity)

``` shell
target/release/halo2-simple-circuit solidity -f path/filename.sol
```

#### Create a W3bstream Project Using zkWASM

For more details on zkWASM circuits see the [zkWASM README](./examples/zkwasm-circuit/README.md).

NOTE: If you want to develop your circuit, please refer [zkwasm project bootstrap](https://github.com/DelphinusLab/zkWasm?tab=readme-ov-file#project-bootstrap)

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
ioctl ws project config -t 3 -i "zkwasm_demo.wasm" -o "path/ID"
```

#### Create a W3bstream Project Using RISC0

More details and options for `Risc0 circuit` are given in [its README](./examples/risc0-circuit/README.md).

NOTE: If you want to develop your circuit, please refer [more risc0 guest examples](https://github.com/risc0/risc0/tree/main/examples)

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
ioctl ws project config -t 1 -i "methods.rs" -o "path/filename.json" -e "{\"image_id\":\"RANGE_ID\", \"elf\":\"RANGE_ELF\"}"
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

### Bind New ZK VmType

The zk vm types registered in the contract are `risc0`, `halo2`, `zkwasm` and `wasm`, and the `id` of risc0 is `1`, halo2 is `2`, zkwasm is `3`, and wasm is `4`.

If you want to use a new zk vm type in the prover, you need to register vm type to the vm type contract.

``` bash
ws vmtype register --vm-type "new vm type name"
```
the command will return a `id`, and the `id` is the "new vm type" id. And you can use the `id` to query the name of the `id`.

``` bash
ioctl ws vmtype query --id "vm type id"
```

If you don't want to use a vm type, you need to `pause` it.

``` bash
ioctl ws vmtype pause --id "vm type id"
```

You can also use `resume` to `resume` it.

``` bash
ioctl ws vmtype resume --id "vm type id"
```

### Verify Proof On Chain

If you want to verify a proof on the chain, 
first, you should set the `output` from `stdout` to `ethereumContract` in the **project config file**, like this

### Set Output to EthereumContract
``` bash
{
   "output": {
     "type": "ethereumContract",
     "ethereum": {
       "chainEndpoint": "https://babel-api.testnet.iotex.io",
       "contractAddress": "0x3841A746F811c244292194825C5e528e61F890F8",
       "contractMethod": "route",
       "contractAbiJSON": "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\", \"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"}],\"name\":\"DappBound\",\"type\":\"event\"},{\"anonymous\":false,  \"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}], \"name\":\"DappUnbound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,  \"internalType\":\"uint256\",\"name\":\"router\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false, \"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"}],\"name\":\"DataProcessed\",  \"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":  [{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dapp\",\"type\":\"address\"}],\"name\":\"bindDapp\",\"outputs\":[],  \"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dapp\",\"outputs\":[{\"internalType\":\"address\", \"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fleetManagement\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",   \"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fleetManagement\",\"type\":\"address\"},{\"internalType\":\"address\", \"name\":\"_projectStore\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectStore\",\"outputs\":  [{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\", \"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proverId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\", \"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"route\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",   \"type\":\"uint256\"}],\"name\":\"unbindDapp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
     }
   }
}
```
`chainEndpoint` is the target chain endpoint, 
`contractAddress` is the router contract address, 
`contractMethod` is the function that you want to call, 
`contractAbiJSON` is the ABI of router contract.

second, you should bind the Dapp of your project with the `Router`.

### Bind Your Dapp Contract

If you want to verify zk proof in your Dapp contract, you need to bind the project with the Dapp.

``` bash
ioctl ws router bind --project-id "your project id" --dapp "your dapp contract address"
```

If you want to unbind the project with the Dapp.

``` bash
ioctl ws router unbind --project-id "your project id"
```

### Dapp Contract
The Dapp Contract that bound with the project will implement the interface `process`, then invoke the verifcation contract, and verify proof. 

The example of Halo2Dapp 

``` bash
function process(uint256 _projectId, uint256 _proverId, string memory _clientId, bytes calldata _data) public {
    require(halo2Verifier != address(0), "verifier address not set");
    
    (uint256 publicInput, uint256 taskID, bytes memory _proof) = abi.decode(_data, (uint256, uint256, bytes));
    bytes32 _publicInput = uint256ToFr(publicInput);
    bytes32 _taskID = uint256ToFr(taskID);
    bytes32 _projectID = uint256ToFr(projectId);
    bytes memory callData = abi.encodePacked(_publicInput, _projectID, _taskID, _proof);
    
    (bool success,) = halo2Verifier.staticcall(callData);
    require(success, "Failed to verify proof");
    // TODO
}
```