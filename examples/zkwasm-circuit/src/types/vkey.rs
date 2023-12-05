use anyhow::{Context, Result};
use delphinus_zkwasm::circuits::config::CircuitConfigure;
use halo2_proofs::arithmetic::MultiMillerLoop;
use halo2_proofs::plonk::{Circuit, VerifyingKey};
use halo2_proofs::poly::commitment::Params;
use serde::{Deserialize, Serialize};

use super::BinarySerializer;

#[derive(Clone, Serialize, Deserialize)]
pub struct VkeyMetadata {
    vkey_data: Vec<u8>,
    circuit_cfg: CircuitConfigure,
}

impl VkeyMetadata {
    pub fn new<E: MultiMillerLoop>(
        vkey: VerifyingKey<E::G1Affine>,
        circuit_cfg: CircuitConfigure,
    ) -> Result<Self> {
        let mut ob = Self {
            vkey_data: vec![],
            circuit_cfg,
        };
        vkey.write(&mut ob.vkey_data)?;
        Ok(ob)
    }

    pub fn load_vkey_with<E: MultiMillerLoop, C: Circuit<E::Scalar>>(
        &self,
        params: &Params<E::G1Affine>,
    ) -> Result<VerifyingKey<E::G1Affine>> {
        self.circuit_cfg.clone().set_global_CIRCUIT_CONFIGURE();
        let vkey = VerifyingKey::read::<_, C>(&mut &self.vkey_data[..], params)
            .context("failed to load verifying key")?;
        Ok(vkey)
    }
}

//TODO: optimize the serialization
impl BinarySerializer for VkeyMetadata {
    type T = Self;

    fn from_binary(buf: &[u8]) -> Result<Self::T> {
        serde_json::from_slice(buf).context("failed to deserialize vkey metadata")
    }

    fn to_binary(&self) -> Result<Vec<u8>> {
        serde_json::to_vec(self).context("failed to serialize vkey metadata")
    }
}
