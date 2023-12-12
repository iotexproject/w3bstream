ZKwasm locally verify proof
==================

## Verify with docker image
1. Get zkwasm proof
   If you can send messages to znode successfully, then you can execute `ioctl ws message send --project-id 10002 --project-version "0.1" --data "{\"private_input\": [1, 1] , \"public_input\": [2] }"` to obtain a zkwasm proof, then put it in a file, like `zkwasm-proof.json`.

2. Execute the `docker run` command for local verification. Note that the directory where the proof is located needs to be mounted into the image.
   It's simple, just input the proof file. You can also use the help command to check how to use it.

```shell
docker run -v /host/zkwasm-proof.json:/zkwasm-proof.json iotexdev/zkverifier /verifier/zkwasm-circuit verify -p /zkwasm-proof.json
```

## Verify with binary file
> **_NOTE:_**
> Since a crate that `zkwasm-circuit` depends on is currently under development and is not suitable to be made public, 
> it is recommended to use a Docker image for local verification.

### Build Verifier

``` shell
cargo build --release
```

After this command is successful, a `zkwasm-circuit` executable file will be generated in the `target/release` directory.

### Verify
You can execute the binary file in the `target/release` directory. It's simple, just input the proof file. You can also use the help command to check how to use it.

``` shell
target/release/zkwasm-circuit verify -p zkwasm-proof.json
```

> **_NOTE:_**
> zkwasm just support single prove
