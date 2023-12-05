use clap::Parser;
use halo2_proofs::{poly::commitment::Params, pairing::bn256::Bn256, arithmetic::Engine};
use zkwasm_circuit::{types::BinarySerializer, load_binary_from_file, verify, opts::Opts};


const PARAMS_FILE_PATH: &str = "./setup/K18.params";

fn main() {
    // let mut params_file = std::fs::File::open(PARAMS_FILE_PATH).unwrap();
    // let params = Params::<<Bn256 as Engine>::G1Affine>::from_binary(&load_binary_from_file(&mut params_file)).unwrap();

    // let mut proof_tmp_file = std::fs::File::open("/Volumes/HIKVISION/project_rust/zkwasm-server/proof.json").unwrap();
    // verify(&params, &mut proof_tmp_file);

    let opts = Opts::parse();
    match opts.sub {
        zkwasm_circuit::opts::Subcommands::Verfiy { proof } => {
            let mut params_file = std::fs::File::open(PARAMS_FILE_PATH).unwrap();
            let params = Params::<<Bn256 as Engine>::G1Affine>::from_binary(&load_binary_from_file(&mut params_file)).unwrap();
        
            let mut proof_tmp_file = std::fs::File::open(proof).unwrap();
            let result = verify(&params, &mut proof_tmp_file);
            assert!(result.is_ok());
            println!("proof is locally verified!");
        },
    }
}