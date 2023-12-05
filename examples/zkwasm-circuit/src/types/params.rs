use anyhow::{Ok, Result};
use halo2_proofs::arithmetic::CurveAffine;
use halo2_proofs::poly::commitment::Params;

use crate::types::BinarySerializer;

impl<C: CurveAffine> BinarySerializer for Params<C> {
    type T = Params<C>;

    fn from_binary(b: &[u8]) -> Result<Self::T> {
        Ok(Params::<C>::read(&b[..])?)
    }

    fn to_binary(&self) -> Result<Vec<u8>> {
        let mut ret = vec![];
        self.write(&mut ret)?;
        Ok(ret)
    }
}
