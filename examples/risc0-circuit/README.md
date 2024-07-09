Risc0 local verify proof
==================

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

