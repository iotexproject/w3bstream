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
git clone git@github.com:machinefi/sprout.git && cd examples/zkwasm-circuit
```

2. build

``` shell
asc circuit/src/add.ts -O --noAssert -o demo.wasm
```
