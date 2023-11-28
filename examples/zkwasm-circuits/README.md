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
git clone git@github.com:machinefi/sprout.git && cd examples/zkwasm-circuits
```

2. build

``` shell
asc src/add.ts -O --noAssert -o demo.wasm
```
