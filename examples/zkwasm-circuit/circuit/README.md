zkwasm template
==================

## Preparation
Install AssemblyScript compiler

``` shell
npm install -g assemblyscript
```

## Build zkwasm program
1. get template

``` shell
git clone git@github.com:iotexproject/w3bstream.git && cd examples/zkwasm-circuit/circuit
```

2. build

``` shell
asc src/add.ts -O --noAssert -o add.wasm
```
