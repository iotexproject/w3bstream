Risc0 local verify proof
==================

## Verify with docker image
1. Get stark proof and image id
   If you have successfully built `methods.rs` in the `method` path, you can find `xx_ID`(this is `RANGE_ID` in example `methods.rs`), the corresponding [u32; 8] array for `xx_ID` without `[]` is `image-id` string.
   If you have successfully deployed `methods.rs` to w3bstream, and you can send messages to prover successfully, then you can execute `ioctl ws message send --project-id 10000 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Stark\"}"` to obtain a stark proof, then put it in a file, like `start-proof.json`.

2. Execute the `docker run` command for local verification. Note that the directory where the proof is located needs to be mounted into the image.
   It's simple, just input the proof file and image-id. You can also use the help command to check how to use it.

```shell
docker run -v /host/stark-proof.json:/stark-proof.json iotexdev/zkverifier:v0.0.1 /verifier/risc0-circuit verify -p /stark-proof.json -i "520991199, 1491489009, 3725421922, 2701107036, 261900524, 710029518, 655219346, 3077599842"
```

## Verify with binary file
### Build Verifier
1. Install the `risc0` toolchain

``` shell
cargo install cargo-binstall
cargo binstall cargo-risczero
cargo risczero install
```

2. build release

``` shell
cargo build --release
```

After this command is successful, a `risc0-circuit` executable file will be generated in the `target/release` directory.

### Verify
You can execute the binary file in the `target/release` directory. It's simple, just input the proof file and image-id. You can also use the help command to check how to use it.

``` shell
target/release/risc0-circuit verify -p stark-proof.json -i "520991199, 1491489009, 3725421922, 2701107036, 261900524, 710029518, 655219346, 3077599842"
```

