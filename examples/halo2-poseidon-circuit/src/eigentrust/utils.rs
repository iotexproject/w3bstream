//! Helper functions for generating params, pk/vk pairs, creating and verifying
//! proofs, etc.
use crate::eigentrust::{params::rns::decompose_big_decimal, FieldExt};
use halo2_proofs::{
    circuit::{AssignedCell, Value},
    halo2curves::{
        ff::{PrimeField, WithSmallOrderMulGroup},
        pairing::{Engine, MultiMillerLoop},
        serde::SerdeObject,
    },
    plonk::{
        create_proof, keygen_pk, keygen_vk, verify_proof, Circuit, Error, ProvingKey, VerifyingKey,
    },
    poly::{
        commitment::{CommitmentScheme, Params, ParamsProver},
        kzg::{
            commitment::{KZGCommitmentScheme, ParamsKZG},
            multiopen::{ProverGWC, VerifierGWC},
            strategy::AccumulatorStrategy,
        },
        VerificationStrategy,
    },
    transcript::{
        Blake2bRead, Blake2bWrite, Challenge255, TranscriptReadBuffer, TranscriptWriterBuffer,
    },
};
use num_bigint::{BigInt, BigUint};
use num_rational::BigRational;
use num_traits::{Num, One};
use rand::Rng;
use std::{
    env::current_dir,
    fmt::Debug,
    fs::{write, File},
    io::{BufReader, Error as IoError, Read},
    path::Path,
    time::Instant,
};

/// Reads raw bytes from the file
pub fn read_bytes(path: impl AsRef<Path>) -> Vec<u8> {
    let f = File::open(path).unwrap();
    let mut reader = BufReader::new(f);
    let mut buffer = Vec::new();

    // Read file into vector.
    reader.read_to_end(&mut buffer).unwrap();

    buffer
}

/// Reads raw bytes from the file
pub fn read_bytes_data(file_name: &str) -> Vec<u8> {
    let current_dir = current_dir().unwrap();
    let bin_path = current_dir.join(format!("../data/{}.bin", file_name));
    read_bytes(bin_path)
}

/// Write bytes to a file
pub fn write_bytes(bytes: Vec<u8>, path: impl AsRef<Path>) -> Result<(), IoError> {
    write(path, bytes)
}

/// Write bytes to data directory
pub fn write_bytes_data(bytes: Vec<u8>, file_name: &str) -> Result<(), IoError> {
    let current_dir = current_dir().unwrap();
    let bin_path = current_dir.join(format!("../data/{}.bin", file_name));
    write(bin_path, bytes)
}

/// Returns boolean value from the assigned cell value
pub fn assigned_as_bool<F: FieldExt>(bit: AssignedCell<F, F>) -> bool {
    let bit_value = bit.value();
    let mut is_one = false;
    bit_value.map(|f| {
        is_one = F::ONE == *f;
        f
    });
    is_one
}

/// Returns field value from the assigned cell value
pub fn assigned_to_field<F: FieldExt>(cell: AssignedCell<F, F>) -> Option<F> {
    let cell_value = cell.value();
    let mut arr = None;
    cell_value.map(|f| {
        arr = Some(*f);
    });
    arr
}

/// Converts given bytes to the bits.
pub fn to_bits(num: &[u8]) -> Vec<bool> {
    let len = num.len() * 8;
    let mut bits = Vec::new();
    for i in 0..len {
        let bit = num[i / 8] & (1 << (i % 8)) != 0;
        bits.push(bit);
    }
    bits
}

/// Converts given field element to the bits.
pub fn field_to_bits_vec<F: FieldExt>(num: F) -> Vec<F> {
    let bits = to_bits(num.to_repr().as_ref());
    let sliced_bits = bits[..F::NUM_BITS as usize].to_vec();
    sliced_bits.iter().map(|&x| F::from(u64::from(x))).collect()
}

/// Converts given field element to the bits.
pub fn field_to_bits<F: FieldExt, const B: usize>(num: F) -> [F; B] {
    let bits = to_bits(num.to_repr().as_ref());
    let sliced_bits = bits[..B].to_vec();
    let vec: Vec<F> = sliced_bits.iter().map(|&x| F::from(u64::from(x))).collect();
    vec.try_into().unwrap()
}

/// Convert bytes array to a wide representation of 64 bytes
pub fn to_wide(b: &[u8]) -> [u8; 64] {
    let mut bytes = [0u8; 64];
    bytes[..b.len()].copy_from_slice(b);
    bytes
}

/// Convert bytes array to a short representation of 32 bytes
pub fn to_short(b: &[u8]) -> [u8; 32] {
    let mut bytes = [0u8; 32];
    bytes[..b.len()].copy_from_slice(b);
    bytes
}

/// Converts field element to string
pub fn field_to_string<F: FieldExt>(f: &F) -> String {
    let bytes = f.to_repr();
    let bn_f = BigUint::from_bytes_le(bytes.as_ref());
    bn_f.to_string()
}

/// Generate parameters with polynomial degree = `k`.
pub fn generate_params<E: Engine + Debug>(k: u32) -> ParamsKZG<E>
where
    E::G1Affine: SerdeObject,
    E::G2Affine: SerdeObject,
    E::Scalar: PrimeField,
{
    ParamsKZG::<E>::new(k)
}

/// Write parameters to a file.
pub fn write_params<E: Engine + Debug>(params: &ParamsKZG<E>)
where
    E::G1Affine: SerdeObject,
    E::G2Affine: SerdeObject,
    E::Scalar: PrimeField,
{
    let mut buffer: Vec<u8> = Vec::new();
    params.write(&mut buffer).unwrap();
    let name = format!("params-{}", params.k());
    write_bytes_data(buffer, &name).unwrap();
}

/// Read parameters from a file.
pub fn read_params<E: Engine + Debug>(k: u32) -> ParamsKZG<E>
where
    E::G1Affine: SerdeObject,
    E::G2Affine: SerdeObject,
    E::Scalar: PrimeField,
{
    let buffer: Vec<u8> = read_bytes_data(&format!("params-{}", k));
    ParamsKZG::<E>::read(&mut &buffer[..]).unwrap()
}

/// Proving/verifying key generation.
pub fn keygen<E: Engine + Debug, C: Circuit<E::Scalar>>(
    params: &ParamsKZG<E>,
    circuit: C,
) -> Result<ProvingKey<<E as Engine>::G1Affine>, Error>
where
    E::G1Affine: SerdeObject,
    E::G2Affine: SerdeObject,
    E::Scalar: FieldExt,
{
    let vk = keygen_vk::<<E as Engine>::G1Affine, ParamsKZG<E>, _>(params, &circuit)?;
    let pk = keygen_pk::<<E as Engine>::G1Affine, ParamsKZG<E>, _>(params, vk, &circuit)?;

    Ok(pk)
}

/// Helper function for finalizing verification
// Rust compiler can't infer the type, so we need to make a helper function
pub fn finalize_verify<
    'a,
    E: MultiMillerLoop + Debug,
    V: VerificationStrategy<'a, KZGCommitmentScheme<E>, VerifierGWC<'a, E>>,
>(
    v: V,
) -> bool
where
    E::G1Affine: SerdeObject,
    E::G2Affine: SerdeObject,
    E::Scalar: PrimeField,
{
    v.finalize()
}

/// Make a proof for generic circuit.
pub fn prove<E: Engine + Debug, C: Circuit<E::Scalar>, R: Rng + Clone>(
    params: &ParamsKZG<E>,
    circuit: C,
    pub_inps: &[&[<KZGCommitmentScheme<E> as CommitmentScheme>::Scalar]],
    pk: &ProvingKey<E::G1Affine>,
    rng: &mut R,
) -> Result<Vec<u8>, Error>
where
    E::G1Affine: SerdeObject,
    E::G2Affine: SerdeObject,
    E::Scalar: FieldExt + WithSmallOrderMulGroup<3>,
{
    let mut transcript = Blake2bWrite::<_, E::G1Affine, Challenge255<_>>::init(vec![]);
    create_proof::<KZGCommitmentScheme<E>, ProverGWC<_>, _, _, _, _>(
        params,
        pk,
        &[circuit],
        &[pub_inps],
        rng.clone(),
        &mut transcript,
    )?;

    let proof = transcript.finalize();
    Ok(proof)
}

/// Verify a proof for generic circuit.
pub fn verify<E: MultiMillerLoop + Debug>(
    params: &ParamsKZG<E>,
    pub_inps: &[&[<KZGCommitmentScheme<E> as CommitmentScheme>::Scalar]],
    proof: &[u8],
    vk: &VerifyingKey<E::G1Affine>,
) -> Result<bool, Error>
where
    E::G1Affine: SerdeObject,
    E::G2Affine: SerdeObject,
    E::Scalar: FieldExt + WithSmallOrderMulGroup<3>,
{
    let strategy = AccumulatorStrategy::<E>::new(params);
    let mut transcript = Blake2bRead::<_, E::G1Affine, Challenge255<_>>::init(proof);
    let output = verify_proof::<KZGCommitmentScheme<E>, VerifierGWC<E>, _, _, _>(
        params,
        vk,
        strategy,
        &[pub_inps],
        &mut transcript,
    )?;

    Ok(finalize_verify(output))
}

/// Helper function for doing proof and verification at the same time.
pub fn prove_and_verify<E: MultiMillerLoop + Debug, C: Circuit<E::Scalar> + Clone, R: Rng + Clone>(
    params: ParamsKZG<E>,
    circuit: C,
    pub_inps: &[&[<KZGCommitmentScheme<E> as CommitmentScheme>::Scalar]],
    rng: &mut R,
) -> Result<bool, Error>
where
    E::G1Affine: SerdeObject,
    E::G2Affine: SerdeObject,
    E::Scalar: FieldExt + WithSmallOrderMulGroup<3>,
{
    let pk = keygen(&params, circuit.clone())?;
    let start = Instant::now();
    let proof = prove(&params, circuit, pub_inps, &pk, rng)?;
    let end = start.elapsed();
    println!("Proving time: {:?}", end);
    let res = verify(&params, pub_inps, &proof[..], pk.get_vk())?;

    Ok(res)
}

/// The u64 integer represented by an L-bit little-endian bitstring.
///
/// # Panics
///
/// Panics if the bitstring is longer than 64 bits.
pub fn le_bits_to_u64<const L: usize>(bits: &[bool; L]) -> u64 {
    assert!(L <= 64);
    bits.iter()
        .enumerate()
        .fold(0u64, |acc, (i, b)| acc + if *b { 1 << i } else { 0 })
}

/// Convert big endian bits to usize
pub fn be_bits_to_usize(bits: &[bool]) -> usize {
    bits.iter()
        .rev()
        .enumerate()
        .fold(0usize, |acc, (i, b)| acc + if *b { 1 << i } else { 0 })
}

/// Convert big endian bits to usize
pub fn be_assigned_bits_to_usize<F: FieldExt>(bits: &[AssignedCell<F, F>]) -> usize {
    let mut bool_bits: Vec<bool> = Vec::new();
    for bit in bits {
        let _ = bit.value().and_then(|a| {
            bool_bits.push(!bool::from(a.clone().is_zero()));
            Value::known(a)
        });
    }
    be_bits_to_usize(bool_bits.as_slice())
}

/// Get the field element of `2 ^ n`
pub fn power_of_two<F: FieldExt>(n: usize) -> F {
    big_to_fe(BigUint::one() << n)
}

/// Get the little-endian bits array of [`Field`] element
pub fn fe_to_le_bits<F: FieldExt>(e: F) -> Vec<bool> {
    let le_bytes = fe_to_big(e).to_bytes_le();
    to_bits(&le_bytes)
}

/// Returns modulus of the [`FieldExt`] as [`BigUint`].
pub fn modulus<F: FieldExt>() -> BigUint {
    BigUint::from_str_radix(&F::MODULUS[2..], 16).unwrap()
}

/// Returns [`FieldExt`] for the given [`BigUint`].
pub fn big_to_fe<F: FieldExt>(e: BigUint) -> F {
    let modulus = modulus::<F>();
    let e = e % modulus;
    F::from_str_vartime(&e.to_str_radix(10)[..]).unwrap()
}

/// Returns [`BigUint`] representation for the given [`FieldExt`].
pub fn fe_to_big<F: FieldExt>(fe: F) -> BigUint {
    BigUint::from_bytes_le(fe.to_repr().as_ref())
}

/// Returns [`BigUint`] representation for the given [`FieldExt`].
pub fn fe_to_big_val<F: FieldExt>(fe: Value<F>) -> Value<BigUint> {
    fe.map(|fe| fe_to_big(fe))
}

/// Converts a `BigRational` into scaled, decomposed numerator and denominator arrays of field elements.
pub fn big_to_fe_rat<F: FieldExt, const NUM_DECIMAL_LIMBS: usize, const POWER_OF_TEN: usize>(
    ratio: BigRational,
) -> ([F; NUM_DECIMAL_LIMBS], [F; NUM_DECIMAL_LIMBS]) {
    let num = ratio.numer();
    let den = ratio.denom();
    let max_len = NUM_DECIMAL_LIMBS * POWER_OF_TEN;
    let bigger = num.max(den);
    let dig_len = bigger.to_string().len();
    let diff = max_len - dig_len;

    let scale = BigInt::from(10_u32).pow(diff as u32);
    let num_scaled = num * scale.clone();
    let den_scaled = den * scale;
    let num_scaled_uint = num_scaled.to_biguint().unwrap();
    let den_scaled_uint = den_scaled.to_biguint().unwrap();

    let num_decomposed =
        decompose_big_decimal::<F, NUM_DECIMAL_LIMBS, POWER_OF_TEN>(num_scaled_uint);
    let den_decomposed =
        decompose_big_decimal::<F, NUM_DECIMAL_LIMBS, POWER_OF_TEN>(den_scaled_uint);

    (num_decomposed, den_decomposed)
}
