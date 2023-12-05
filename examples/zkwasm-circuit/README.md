ZKwasm locally verify proof
==================

## Build Verifier

``` shell
cargo build --release
```

## Verify
1. Get zkwasm proof
If you can send messages to znode successfully, then you can execute `ioctl ws message send --project-id 10002 --project-version "0.1" --data "{\"private_input\": [1, 1] , \"public_input\": [2] }"` to obtain a zkwasm proof, then put it in a file, like `zkwasm-proof.json`.

2. You can execute the binary file in the `target/release` directory. It's simple, just input the proof file and image-id. You can also use the help command to check how to use it.

``` shell
target/release/zkwasm-circuit verify -p zkwasm-proof.json
```

> **_NOTE:_**
> zkwasm just support single prove
