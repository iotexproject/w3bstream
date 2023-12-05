use anyhow::{Context, Result};
use halo2_proofs::arithmetic::MultiMillerLoop;
use halo2_proofs::plonk::{verify_proof, SingleVerifier, VerifyingKey};
use halo2_proofs::poly::commitment::{Params, ParamsVerifier};
use halo2aggregator_s::transcript::poseidon::PoseidonRead;

pub struct ZKVerifier<'a, E: MultiMillerLoop> {
    params: Option<&'a Params<E::G1Affine>>,
    vkey: Option<VerifyingKey<E::G1Affine>>,
}

impl<'a, E: MultiMillerLoop> ZKVerifier<'a, E> {
    pub fn new() -> Self {
        Self {
            params: None,
            vkey: None,
        }
    }

    pub fn load_params(&mut self, params: &'a Params<E::G1Affine>) {
        self.params = Some(params);
    }

    pub fn load_vkey(&mut self, vkey: &VerifyingKey<E::G1Affine>) {
        self.vkey = Some(vkey.clone());
    }

    pub fn verify_proof(&self, instances: &Vec<E::Scalar>, proof: &Vec<u8>) -> Result<()> {
        let params_verifier: ParamsVerifier<E> = self
            .params
            .context("params not loaded")?
            .verifier(instances.len())?;
        let res = verify_proof(
            &params_verifier,
            self.vkey.as_ref().context("vkey not loaded")?,
            SingleVerifier::new(&params_verifier),
            &[&[instances]],
            &mut PoseidonRead::init(&proof[..]),
        )?;
        Ok(res)
    }
}
