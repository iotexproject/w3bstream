@external("env", "wasm_input")
declare function wasm_input(x: i32): i64

@external("env", "require")
declare function require(x: i32): void

export function read_public_input(): i64 {
    return wasm_input(1);
}

export function read_private_input(): i64 {
    return wasm_input(0);
}

export function zkmain(): void {
  let a = read_private_input();
  let b = read_private_input();
  let expect = read_public_input();

  require(a+b == expect);
}
