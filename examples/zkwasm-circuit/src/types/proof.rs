use anyhow::{Context, Ok, Result};
use data_encoding::BASE64;
use halo2_proofs::pairing::bn256::Bn256;
use serde_json::{json, Value};

use crate::types::instance::Instances;
use crate::types::{BinarySerializer, JSONSerializer};

pub struct ProofData {
    pub proof: Vec<u8>,
    pub instances: Instances<Bn256>,
    pub vk: Vec<u8>,
}

impl BinarySerializer for ProofData {
    type T = Self;

    fn from_binary(buf: &[u8]) -> Result<Self::T> {
        let ob: Value = serde_json::from_slice(buf)?;
        let instances_ob = ob.get("instances").context("no instances")?.to_owned();
        let proof = ob["proof"].as_str().context("no proof")?;
        let vk = ob["vk"].as_str().context("no proof")?;

        Ok(ProofData {
            proof: BASE64.decode(proof.as_bytes())?,
            instances: Instances::from_json(instances_ob)?,
            vk: BASE64.decode(vk.as_bytes())?,
        })
    }

    fn to_binary(&self) -> Result<Vec<u8>> {
        let ob: Value = json!({
            "proof": BASE64.encode(&self.proof),
            "instances": self.instances.to_json()?,
            "vk": BASE64.encode(&self.vk),
        });
        Ok(serde_json::to_vec(&ob)?)
    }
}
