Halo2 wasm template
==================

## Preparation
install `wasm-pack`
``` shell
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

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
