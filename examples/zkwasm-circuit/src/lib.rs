use std::{fs::File, io::{Seek, Read}};

use anyhow::Result;
use delphinus_zkwasm::circuits::TestCircuit;
use halo2_proofs::{poly::commitment::Params, arithmetic::Engine, pairing::bn256::Bn256};

use crate::{types::{proof::ProofData, vkey::VkeyMetadata, BinarySerializer}, prover::verifier};

pub mod types;
mod prover;
pub mod opts;

pub fn verify(
    params: &Params<<Bn256 as Engine>::G1Affine>,
    proof_tmp_file: &mut File,
) -> Result<()> {

    // 1. load proof and instance
    let pd = ProofData::from_binary(&load_binary_from_file(proof_tmp_file)).unwrap();
    let (proof, instances, vk) = (pd.proof, pd.instances, pd.vk);

    // 2. load vkey
    let vk = VkeyMetadata::from_binary(&vk).unwrap();
    let vkey = vk.load_vkey_with::<Bn256, TestCircuit<_>>(&params).unwrap();

    // 3. build verifier
    let mut verifier = verifier::ZKVerifier::<Bn256>::new();
    verifier.load_params(&params);
    verifier.load_vkey(&vkey);
    return verifier.verify_proof(&instances.0, &proof);
}

pub fn load_binary_from_file(fd: &mut File) -> Vec<u8> {
    fd.seek(std::io::SeekFrom::Start(0)).unwrap();
    let mut ret = vec![];
    fd.read_to_end(&mut ret).unwrap();
    ret
}