Risc0 guest template
==================

## Build risc0 circuits program
1. get template 

``` shell
git clone git@github.com:machinefi/sprout.git && cd examples/risc0-circuits
```

2. build

``` shell
cargo build --release
```

you will find `methods.rs` in the `target/release/build/risc0-circuits-xxx/out/methods.rs` 

## Advanced
You can also develop your own risc0 guest program.

1. Edit `guest/Cargo.toml`, changing the line `name = "method_name"` to instead read `name = "your_method_name"`.
2. Edit `guest/src/main.rs`, changing the `main` func.
3. Build wasm with `cargo build --release`, and you will find `methods.rs` in the `target/release/build/your_method_name-xxx/out/methods.rs`  .

[more risc0 guest examples](https://github.com/risc0/risc0/tree/main/examples)