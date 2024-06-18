# Build circuit

## Get Repository
```bash
git clone https://github.com/machinefi/sprout.git
cd sprout
```

## Compile customized circuits

### Compile the customized Halo2 circuit

1. Install `wasm-pack`
```bash
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

2. Build wasm

```bash
cd examples/halo2-circuit/
wasm-pack build --target nodejs --out-dir pkg
```

you will find `halo2_wasm_bg.wasm` under the `pkg` folder.

3. (Optional) You can also write your circuit according to the [halo2 development documentation](https://zcash.github.io/halo2/user/simple-example.html), and put the circuit file in `src/circuits`; replace the `TODO` in `src/lib.rs` and build wasm with `wasm-pack build --target nodejs --out-dir pkg`.

More details and options for `Halo2 circuit` are given in [its README](./examples/halo2-circuit/README.md).

### Compile the customized Risc0 circuits

1. Build

```bash
cd examples/risc0-circuit/
cargo build --release
```

The path of `methods.rs` will be printed to the console, like this  

```bash
warning: methods_path is: "sprout/examples/risc0-circuits/target/release/build/risc0-circuits-5efc4ff59af940ab/out/methods.rs"
```

More details and options for `Risc0 circuit` are given in [its README](./examples/risc0-circuit/README.md).

### Compile the customized zkWasm circuits

1. Build

```bash
cd examples/zkwasm-circuit/
asc src/add.ts -O --noAssert -o zkwasm_demo.wasm
```

More details and options for `zkWasm circuit` are given in [its README](./examples/zkwasm-circuit/README.md).


## Convert compiled circuit to w3bstream project config

### Convert halo2 circuit to w3bstream project config

```bash
ioctl ws project config -s "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable" -t "halo2" -i "halo2_wasm_bg.wasm"
```

This command will generate a file named `halo2-config.json` in the current folder. 
Or you can run `ioctl ws project config -s "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable" -t "halo2" -i "halo2_wasm_bg.wasm" -o "path/filename.json"`

### Convert risc0 circuit to w3bstream project config

```bash
ioctl ws project config -s "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable" -t "risc0" -i "methods.rs" -e "{\"image_id\":\"RANGE_ID\", \"elf\":\"RANGE_ELF\"}
```
The values of `image_id` and `elf` are variable names, and will be found in the `methods.rs`.

This command will generate a file named `risc0-config.json` in the current folder.
Or you can run `ioctl ws project config -s "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable" -t "risc0" -i "methods.rs" -o "path/filename.json" -e "{\"image_id\":\"RANGE_ID\", \"elf\":\"RANGE_ELF\"}`

### Convert zkwasm circuit to w3bstream project config

```bash
ioctl ws project config -s "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable" -t "zkwasm" -i "zkwasm_demo.wasm"
```

This command will generate a file named `zkwasm-config.json` in the current folder.
Or you can run `ioctl ws project config -s "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable" -t "zkwasm" -i "zkwasm_demo.wasm" -o "path/filename.json"`
