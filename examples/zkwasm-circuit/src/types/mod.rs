pub mod params;
pub mod proof;
pub mod vkey;
mod instance;

use anyhow::Result;

pub trait JSONSerializer {
    type T;
    fn from_json(ob: serde_json::Value) -> Result<Self::T>;
    fn to_json(&self) -> Result<serde_json::Value>;
}

pub trait BinarySerializer {
    type T;
    fn from_binary(b: &[u8]) -> Result<Self::T>;
    fn to_binary(&self) -> Result<Vec<u8>>;
}