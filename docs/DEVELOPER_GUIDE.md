# IoTeX W3bstream Project Developer Guide

## Prerequisites

Ensure you have the following tools installed:

- [ioctl](https://docs.iotex.io/builders/reference-docs/ioctl-client#install-latest-release-build): For interacting with the IoTeX blockchain.

## Create a W3bstream Prover Configuration

The W3bstream Prover Configuration consists primarily of the ZK Prover binary code and its revisions. To create a configuration file, first compile a ZK circuit into a prover using one of the supported ZK frameworks, then generate a W3bstream configuration file using the ioctl command-line client.

Begin by cloning the W3bstream repository:

```bash
git clone https://github.com/iotexproject/w3bstream.git
cd w3bstream
```

<details>
<summary>
Halo2 Circuits
</summary>
<blockquote>
💡 NOTE: For circuit development with Halo2, refer to the <a href="https://zcash.github.io/halo2/user/simple-example.html">halo2 development documentation</a>.
</blockquote>

#### Step 1: Install `wasm-pack`

```bash
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

#### Step 2: build the wasm prover

```bash
cd examples/halo2-circuit/
wasm-pack build --target nodejs --out-dir pkg
```

The `halo2_simple_bg.wasm` will be located under the `pkg` folder.

#### Step 3: Generate the W3bstream config file

Customize the output project file name:

```bash
ioctl ws project config -t 2 -i "halo2_wasm_bg.wasm" -c "path/config_file_name.json"
```

#### Step 4: Create the blockchain verifier (Solidity)

``` shell
target/release/halo2-simple-circuit solidity -f path/filename.sol
```

</details>

<details>
<summary>
zkWASM Circuits
</summary>
<blockquote>
💡 NOTE: For zkWASM circuit development, refer to the <a href="https://github.com/DelphinusLab/zkWasm?tab=readme-ov-file#project-bootstrap">zkWASM development documentation</a>.

</blockquote>

#### Step 1: Install AssemblyScript

Ensure you have AssemblyScript installed:

```bash
npm install -g assemblyscript
```

#### Step 2: Build the circuit

```bash
cd examples/zkwasm-circuit/
asc src/add.ts -O --noAssert -o zkwasm_demo.wasm
```

#### Step 3: Create the verifier

```bash
# Work in progress
```

#### Step 4: Generate the W3bstream config file

```bash
ioctl ws project config -t 3 -i "zkwasm_demo.wasm" -o "path/config_file_name.json"
```

</details>

<details>
<summary>
Risc0 Circuits
</summary>

<blockquote>
💡 NOTE: For Risc0 circuit development, refer to the <a href="https://github.com/risc0/risc0/tree/main/examples">Risc0 development documentation</a>.

</blockquote>

#### Step 1: Install the Toolchain

Make sure cargo version 1.72.0 or higher is installed:

```bash
# Check with 
cargo version

# Install with
curl https://sh.rustup.rs -sSf | sh

# Update with
rustup update
```

Install the rustzero toolchain

```bash
cargo install cargo-risczero
cargo risczero install
```

#### Step 2: Build the circuit

```bash
cd examples/risc0-circuit/
cargo build --release
```

The path of `methods.rs` will be printed in the output, as sown in the example below:  

```bash
warning: methods_path is: "w3bstream/examples/risc0-circuits/target/release/build/risc0-circuits-5efc4ff59af940ab/out/methods.rs"
```

#### Step 3: Generate the W3bstream Config File

```bash
ioctl ws project config -t 1 -i "methods.rs" -o "path/filename.json" -e "{\"image_id\":\"RANGE_ID\", \"elf\":\"RANGE_ELF\"}"
```

Customize the values for `image_id` and `elf` with the actual values found in your `methods.rs` file.

</details>

## Test Your W3bstream Project

Refer to the [Quick Start](QUICK_START.md) guide to learn how to **register** your project on-chain, **deploy** your W3bstream config file, and **send messages** to the demo API node.

## Advanced Concepts

### Bind a New ZK VmType

Registered zk VM types in W3bstream include `risc0`, `halo2`, `zkwasm` and `wasm`. Their corresponding IDs are `1` for risc0, `2`for halo2, `3` for zkWASM, and `4` for WASM.

To use a different zk VM type in the prover, you’ll need to register it in the VM type contract:

```bash
ws vmtype register --vm-type "new vm type name"
```

This command returns an id for the new VM type, which you can use to query the name associated with it:

```bash
ioctl ws vmtype query --id "vm type id"
```

To disable a VM type, use `pause`:

```bash
ioctl ws vmtype pause --id "vm type id"
```

To re-enable it, use `resume`:

``` bash
ioctl ws vmtype resume --id "vm type id"
```

### Verifying Proofs in your Dapp

To receive and verify a W3bstream proof in your DApp, set the `output` field in the configuration file to `ethereumContract`, as shown below:

<details>
<summary>Proof Output set to EthereumContract</summary>
<pre>
{
  "defaultVersion": "0.1",
  "versions": [
    {
      "version": "v1.0.0",
      "vmTypeID": 1,
      "output": {
         "type": "ethereumContract",
        "ethereum": {
          "chainEndpoint": "https://babel-api.testnet.iotex.io",
          "contractAddress": "0x28E0A99A76a467E7418019cBBbF79E4599C73B5B",
          "contractMethod": "route",
          "contractAbiJSON": "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"}],\"name\":\"DappBound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"DappUnbound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"router\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"}],\"name\":\"DataProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dapp\",\"type\":\"address\"}],\"name\":\"bindDapp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dapp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fleetManagement\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fleetManagement\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_projectStore\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectStore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proverId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"route\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"unbindDapp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
        }
      },
      "codeExpParam": " ..."
      ...
      </pre>
      </details>

### Dapp Contracts

The DApp contract linked to the W3bstream project must implement the [IDapp](../smartcontracts/contracts/interfaces/IDapp.sol) interface. This allows it to call the verification contract with the appropriate data to verify proofs before performing any additional computations.

**For Risc0** proofs: A verifier is already deployed on-chain. You can refer to a minimal[Dapp Example](../examples/risc0-circuit/contract/Risc0Dapp.sol).

**For Halo2** proofs: A corresponding Solidity verifier contract for the prover must be created during the prover compilation process and deployed on-chain. See the minimal [Dapp Example](../examples/halo2-circuit/contracts/Halo2Dapp.sol).
