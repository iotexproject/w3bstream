use std::fmt;

use anyhow::{Ok, Result};
use data_encoding::BASE64;
use halo2_proofs::{
    arithmetic::{BaseExt, MultiMillerLoop},
    pairing::group::ff::PrimeField,
};

use super::JSONSerializer;

#[derive(Clone)]
pub struct Instances<E: MultiMillerLoop>(pub Vec<E::Scalar>);

impl<E: MultiMillerLoop> JSONSerializer for Instances<E> {
    type T = Self;

    fn from_json(ob: serde_json::Value) -> Result<Self::T> {
        let arr = ob.as_array().expect("the input is not an array");
        let ins: Vec<E::Scalar> = arr
            .iter()
            .map(|val| {
                let str = val.as_str().expect("the input is not a string");
                let bytes = BASE64
                    .decode(str.as_bytes())
                    .expect("the input is not a base64 string");
                E::Scalar::read(&mut bytes.as_slice())
                    .expect("the input is not a valid field element")
            })
            .collect();
        Ok(Instances(ins))
    }

    fn to_json(&self) -> Result<serde_json::Value> {
        let ob: Vec<String> = self
            .0
            .iter()
            .map(|instance| {
                let mut bytes = vec![];
                instance
                    .write(&mut bytes)
                    .expect("the input is not a valid field element");
                BASE64.encode(&bytes)
            })
            .collect();
        Ok(serde_json::to_value(ob)?)
    }
}

impl<E: MultiMillerLoop> fmt::Display for Instances<E> {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let arr: Vec<String> = self
            .0
            .iter()
            .map(|ins| format!("{:x?}", ins.to_repr().as_ref()))
            .collect();
        write!(f, "{:?}", arr)
    }
}