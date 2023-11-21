Halo2 wasm template
==================

## Preparation
install `wasm-pack`
``` shell
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

install `solc`
Reference [install guid](https://docs.soliditylang.org/en/v0.8.9/installing-solidity.html)

## Build halo2 wasm program
1. get template 

``` shell
git clone git@github.com:machinefi/vm-template.git && cd halo2-wasm
```

2. build wasm

``` shell
wasm-pack build --target nodejs --out-dir pkg
```

you will find `xx_bg.wasm` in the `pkg` 

## Advanced
You can also develop your own halo2 circuit program.

1. Write a circuit according to the [halo2 development documentation](https://zcash.github.io/halo2/user/simple-example.html), and put the circuit file in `src/circuits`.
2. Replace the `TODO` in `src/lib.rs`.
3. Build wasm with `wasm-pack build --target nodejs --out-dir pkg`.

## Generate verify smart contract

### build 

``` shell
cargo build --release
```

### run 

``` shell
target/release/halo2-wasm
```
you will find `Verifier.sol` under the current folder. Or you can run `target/release/halo2-wasm path/filename.sol`.
