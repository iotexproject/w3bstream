Risc0 local verify proof
==================

## Verify
1. Get start proof and image id
If you have successfully built `methods.rs` in the `method` path, you can find `xx_ID`(this is `RANGE_ID` in example `methods.rs`), the corresponding [u32; 8] array for `xx_ID` without `[]` is `image-id` string. 
If you have successfully deployed `methods.rs` to w3bstream, and you can send messages to znode successfully, then you can execute `ioctl ws message send --project-id 10000 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Stark\"}"` to obtain a stark proof, then put it in a file, like `start-proof.json`.

2. You can directly execute the binary file in the `bin` directory. It's simple, just input the proof file and image-id. You can also use the help command to check how to use it.

``` shell
bin/risc0-circuit verify -p stark-proof.json -i "520991199, 1491489009, 3725421922, 2701107036, 261900524, 710029518, 655219346, 3077599842"
```

## Build Verifier
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