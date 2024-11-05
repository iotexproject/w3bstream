// Copyright 2023 RISC Zero, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#![no_main]
// #![no_std]

use risc0_zkvm::guest::env;
use serde_json::Value as JsonValue;

risc0_zkvm::guest::entry!(main);

pub fn main() {
    let data: Vec<String> = env::read();
    env::log(&format!("data {:?}", data));

    let v: JsonValue = serde_json::from_str(&data[0]).unwrap();

    // Parse private input directly as u64
    let pri_a = v["private_input"].as_str().unwrap().parse::<u64>().unwrap();

    // Split public input string and parse both numbers
    let public_input = v["public_input"].as_str().unwrap();
    let pub_nums: Vec<u64> = public_input
        .split(',')
        .map(|s| s.trim().parse::<u64>().unwrap())
        .collect();

    let pub_b = pub_nums[0];
    let pub_c = pub_nums[1];

    if pri_a > pub_b && pri_a < pub_c {
        let s = format!(
            "I know your private input is greater than {} and less than {}, and I can prove it!",
            pub_b, pub_c
        );
        env::commit(&s);
    } else {
        let s = format!(
            "I know your private input is not greater than {} or less than {}, and I can not prove it!",
            pub_b, pub_c
        );
        env::commit(&s);
    }
}
