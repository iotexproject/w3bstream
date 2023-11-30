use halo2_proofs::halo2curves::{
    bn256::{Fq, Fr},
    ff::{Field as Halo2Field, FromUniformBytes, PrimeField},
};

pub mod is_zero;

pub trait Field: Halo2Field + PrimeField<Repr = [u8; 32]> + FromUniformBytes<64> + Ord {}
impl Field for Fr {}
impl Field for Fq {}
