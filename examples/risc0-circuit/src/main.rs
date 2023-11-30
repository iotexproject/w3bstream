use std::fs;

use clap::Parser;
use risc0_circuit::opts::Opts;
use risc0_zkvm::{serde::from_slice, Receipt};
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize)]
pub enum RiscReceipt {
    /// The [Receipt].
    Stark(Receipt),
}

fn main() {
    let opts = Opts::parse();
    match opts.sub {
        risc0_circuit::opts::Subcommands::Verfiy { proof, image_id } => {
            let mut id: [u32; 8] = [0; 8];
            let vec_u32: Result<Vec<u32>, _> = image_id
                .split(",")
                .into_iter()
                .map(|s| s.trim().parse::<u32>())
                .collect();

            match vec_u32 {
                Ok(v) => {
                    if v.len() == id.len() {
                        id.copy_from_slice(&v);
                    } else {
                        println!("The length of image_id is not 8.");
                        return;
                    }
                }
                Err(e) => println!("image_id parse error: {}", e),
            }

            let proof_raw = fs::read(proof).expect("read proof file error");
            let proof_raw = String::from_utf8(proof_raw).unwrap();
            let risc_receipt: RiscReceipt = serde_json::from_str(&proof_raw).unwrap();
            match risc_receipt {
                RiscReceipt::Stark(receipt) => {
                    let verify_result = receipt.verify(id);
                    let result = match verify_result {
                Ok(_) => from_slice(&receipt.journal.bytes).expect(
                    "Journal output should deserialize into the same types (& order) that it was written",
                ),
                Err(e) => format!("{}", e),
            };
                    println!("{:?}", result);
                }
            }
        }
    }
}
