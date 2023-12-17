use crate::eigentrust::{
    params::rns::{compose_big, decompose_big, RnsParams},
    utils::{fe_to_big, field_to_bits},
    FieldExt,
};
use num_bigint::BigUint;
use num_traits::{One, Zero};
use std::marker::PhantomData;

/// Enum for the two different type of Quotient.
#[derive(Clone, Debug)]
pub enum Quotient<W: FieldExt, N: FieldExt, const NUM_LIMBS: usize, const NUM_BITS: usize, P>
where
    P: RnsParams<W, N, NUM_LIMBS, NUM_BITS>,
{
    /// Quotient type for the addition and subtraction.
    Short(N),
    /// Quotient type for the multiplication and division.
    Long(Integer<W, N, NUM_LIMBS, NUM_BITS, P>),
}

impl<W: FieldExt, N: FieldExt, const NUM_LIMBS: usize, const NUM_BITS: usize, P>
    Quotient<W, N, NUM_LIMBS, NUM_BITS, P>
where
    P: RnsParams<W, N, NUM_LIMBS, NUM_BITS>,
{
    /// Returns Quotient type for the addition or the subtraction.
    pub fn short(self) -> Option<N> {
        match self {
            Quotient::Short(res) => Some(res),
            _ => None,
        }
    }

    /// Returns Quotient type for the multiplication or the division.
    pub fn long(self) -> Option<Integer<W, N, NUM_LIMBS, NUM_BITS, P>> {
        match self {
            Quotient::Long(res) => Some(res),
            _ => None,
        }
    }
}

/// Structure for the ReductionWitness.
#[derive(Debug, Clone)]
pub struct ReductionWitness<
    W: FieldExt,
    N: FieldExt,
    const NUM_LIMBS: usize,
    const NUM_BITS: usize,
    P,
> where
    P: RnsParams<W, N, NUM_LIMBS, NUM_BITS>,
{
    /// Result from the operation.
    pub(crate) result: Integer<W, N, NUM_LIMBS, NUM_BITS, P>,
    /// Quotient from the operation.
    pub(crate) quotient: Quotient<W, N, NUM_LIMBS, NUM_BITS, P>,
    /// Intermediate values from the operation.
    #[allow(dead_code)]
    pub(crate) intermediate: [N; NUM_LIMBS],
    /// Residue values from the operation.
    #[allow(dead_code)]
    pub(crate) residues: Vec<N>,
}

/// Structure for the Integer.
// TODO: Add ReductionWitness as a part of Integer
// TODO: Add `add_add`, `mul2`, `mul3`, `sub_sub`
#[derive(Clone, Debug, PartialEq)]
pub struct Integer<W: FieldExt, N: FieldExt, const NUM_LIMBS: usize, const NUM_BITS: usize, P>
where
    P: RnsParams<W, N, NUM_LIMBS, NUM_BITS>,
{
    /// Integer limbs for the non-native operations.
    pub(crate) limbs: [N; NUM_LIMBS],
    /// Phantom data for the Wrong Field.
    _wrong_field: PhantomData<W>,
    /// Phantom data for the RnsParams.
    _rns: PhantomData<P>,
}

impl<W: FieldExt, N: FieldExt, const NUM_LIMBS: usize, const NUM_BITS: usize, P> Default
    for Integer<W, N, NUM_LIMBS, NUM_BITS, P>
where
    P: RnsParams<W, N, NUM_LIMBS, NUM_BITS>,
{
    fn default() -> Self {
        Self::zero()
    }
}

impl<W: FieldExt, N: FieldExt, const NUM_LIMBS: usize, const NUM_BITS: usize, P>
    Integer<W, N, NUM_LIMBS, NUM_BITS, P>
where
    P: RnsParams<W, N, NUM_LIMBS, NUM_BITS>,
{
    /// Creates a new object by decomposing given `limbs`.
    pub fn new(num: BigUint) -> Self {
        let limbs = decompose_big::<N, NUM_LIMBS, NUM_BITS>(num);
        Self::from_limbs(limbs)
    }

    /// Creates and new Integer from wrong field
    pub fn from_w(num: W) -> Self {
        let num_bn = fe_to_big(num);
        Self::new(num_bn)
    }

    /// Creates and new Integer from native field
    pub fn from_n(num: N) -> Self {
        let num_bn = fe_to_big(num);
        Self::new(num_bn)
    }

    /// Construct an Integer from given `limbs`.
    pub fn from_limbs(limbs: [N; NUM_LIMBS]) -> Self {
        Self {
            limbs,
            _wrong_field: PhantomData,
            _rns: PhantomData,
        }
    }

    /// Construct an Integer from given `limbs` vector.
    pub fn from_slice(limbs_vec: &[N]) -> Self {
        assert_eq!(limbs_vec.len(), NUM_LIMBS);
        let limbs: [N; NUM_LIMBS] = limbs_vec.try_into().unwrap();
        Self::from_limbs(limbs)
    }

    /// Convert the value to bits representation
    pub fn to_bits(&self) -> Vec<bool> {
        let mut bits = Vec::new();
        for i in 0..NUM_LIMBS {
            let limb_bits = field_to_bits::<_, NUM_BITS>(self.limbs[i]).map(|x| x == N::ONE);
            bits.extend(&limb_bits);
        }

        bits
    }

    /// Returns integer with value zero
    pub fn zero() -> Self {
        Self::new(BigUint::zero())
    }

    /// Returns integer with value one
    pub fn one() -> Self {
        Self::new(BigUint::one())
    }

    /// Returns [`BigUint`] representation from the given `limbs`.
    pub fn value(&self) -> BigUint {
        let limb_values = self.limbs.map(|limb| fe_to_big(limb));
        compose_big::<NUM_LIMBS, NUM_BITS>(limb_values)
    }

    /// Reduce function for the [`Integer`].
    pub fn reduce(&self) -> ReductionWitness<W, N, NUM_LIMBS, NUM_BITS, P> {
        let p_prime = P::negative_wrong_modulus_decomposed();
        let p_in_n = P::wrong_modulus_in_native_modulus();
        let a = self.value();
        let (q, res) = P::construct_reduce_qr(a);

        // Calculate the intermediate values for the ReductionWitness.
        let mut t = [N::ZERO; NUM_LIMBS];
        for i in 0..NUM_LIMBS {
            t[i] = self.limbs[i] + p_prime[i] * q;
        }

        // Calculate the residue values for the ReductionWitness.
        let residues = P::residues(&res, &t);

        let satisfied = P::constrain_binary_crt(t, res, residues.clone());
        assert!(satisfied);

        let native_constraint = P::compose(self.limbs) - q * p_in_n - P::compose(res);
        assert!(native_constraint == N::ZERO);

        // Construct correct type for the ReductionWitness
        let result_int = Integer::from_limbs(res);
        let quotient_n = Quotient::Short(q);
        ReductionWitness {
            result: result_int,
            quotient: quotient_n,
            intermediate: t,
            residues,
        }
    }

    /// Non-native addition for given two [`Integer`].
    pub fn add(
        &self,
        other: &Integer<W, N, NUM_LIMBS, NUM_BITS, P>,
    ) -> ReductionWitness<W, N, NUM_LIMBS, NUM_BITS, P> {
        let p_prime = P::negative_wrong_modulus_decomposed();
        let p_in_n = P::wrong_modulus_in_native_modulus();
        let a = self.value();
        let b = other.value();
        let (q, res) = P::construct_add_qr(a, b);

        // Calculate the intermediate values for the ReductionWitness.
        let mut t = [N::ZERO; NUM_LIMBS];
        for i in 0..NUM_LIMBS {
            t[i] = self.limbs[i] + other.limbs[i] + p_prime[i] * q;
        }

        // Calculate the residue values for the ReductionWitness.
        let residues = P::residues(&res, &t);

        let satisfied = P::constrain_binary_crt(t, res, residues.clone());
        assert!(satisfied);

        let native_constraint =
            P::compose(self.limbs) + P::compose(other.limbs) - q * p_in_n - P::compose(res);
        assert!(native_constraint == N::ZERO);

        // Construct correct type for the ReductionWitness
        let result_int = Integer::from_limbs(res);
        let quotient_n = Quotient::Short(q);
        ReductionWitness {
            result: result_int,
            quotient: quotient_n,
            intermediate: t,
            residues,
        }
    }

    /// Non-native subtraction for given two [`Integer`].
    pub fn sub(
        &self,
        other: &Integer<W, N, NUM_LIMBS, NUM_BITS, P>,
    ) -> ReductionWitness<W, N, NUM_LIMBS, NUM_BITS, P> {
        let p_prime = P::negative_wrong_modulus_decomposed();
        let p_in_n = P::wrong_modulus_in_native_modulus();
        let a = self.value();
        let b = other.value();
        let (q, res) = P::construct_sub_qr(a, b);

        // Calculate the intermediate values for the ReductionWitness.
        let mut t = [N::ZERO; NUM_LIMBS];
        for i in 0..NUM_LIMBS {
            t[i] = self.limbs[i] - other.limbs[i] + p_prime[i] * q;
        }

        // Calculate the residue values for the ReductionWitness.
        let residues = P::residues(&res, &t);

        let satisfied = P::constrain_binary_crt(t, res, residues.clone());
        assert!(satisfied);

        let native_constraint =
            P::compose(self.limbs) - P::compose(other.limbs) + q * p_in_n - P::compose(res);
        assert!(native_constraint == N::ZERO);

        // Construct correct type for the ReductionWitness
        let result_int = Integer::from_limbs(res);
        let quotient_n = Quotient::Short(q);

        ReductionWitness {
            result: result_int,
            quotient: quotient_n,
            intermediate: t,
            residues,
        }
    }

    /// Non-native multiplication for given two [`Integer`].
    pub fn mul(
        &self,
        other: &Integer<W, N, NUM_LIMBS, NUM_BITS, P>,
    ) -> ReductionWitness<W, N, NUM_LIMBS, NUM_BITS, P> {
        let p_prime = P::negative_wrong_modulus_decomposed();
        let p_in_n = P::wrong_modulus_in_native_modulus();
        let a = self.value();
        let b = other.value();
        let (q, res) = P::construct_mul_qr(a, b);

        // Calculate the intermediate values for the ReductionWitness.
        let mut t: [N; NUM_LIMBS] = [N::ZERO; NUM_LIMBS];
        for k in 0..NUM_LIMBS {
            for i in 0..=k {
                let j = k - i;
                t[i + j] = t[i + j] + self.limbs[i] * other.limbs[j] + p_prime[i] * q[j];
            }
        }

        // Calculate the residue values for the ReductionWitness.
        let residues = P::residues(&res, &t);

        let satisfied = P::constrain_binary_crt(t, res, residues.clone());
        assert!(satisfied);

        let native_constraint = P::compose(self.limbs) * P::compose(other.limbs)
            - P::compose(q) * p_in_n
            - P::compose(res);
        assert!(native_constraint == N::ZERO);

        // Construct correct type for the ReductionWitness.
        let result_int = Integer::from_limbs(res);
        let quotient_int = Quotient::Long(Integer::from_limbs(q));
        ReductionWitness {
            result: result_int,
            quotient: quotient_int,
            intermediate: t,
            residues,
        }
    }

    /// Non-native division for given two [`Integer`].
    pub fn div(
        &self,
        other: &Integer<W, N, NUM_LIMBS, NUM_BITS, P>,
    ) -> ReductionWitness<W, N, NUM_LIMBS, NUM_BITS, P> {
        let p_prime = P::negative_wrong_modulus_decomposed();
        let p_in_n = P::wrong_modulus_in_native_modulus();
        let a = self.value();
        let b = other.value();
        let (q, res) = P::construct_div_qr(a, b);

        // Calculate the intermediate values for the ReductionWitness.
        let mut t: [N; NUM_LIMBS] = [N::ZERO; NUM_LIMBS];
        for k in 0..NUM_LIMBS {
            for i in 0..=k {
                let j = k - i;
                t[i + j] = t[i + j] + res[i] * other.limbs[j] + p_prime[i] * q[j];
            }
        }

        // Calculate the residue values for the ReductionWitness.
        let residues = P::residues(&res, &t);

        let satisfied = P::constrain_binary_crt(t, res, residues.clone());
        assert!(satisfied);

        let native_constraints = P::compose(other.limbs) * P::compose(res)
            - P::compose(self.limbs)
            - P::compose(q) * p_in_n;
        assert!(native_constraints == N::ZERO);

        // Construct correct type for the ReductionWitness.
        let result_int = Integer::from_limbs(res);
        let quotient_int = Quotient::Long(Integer::from_limbs(q));

        ReductionWitness {
            result: result_int,
            quotient: quotient_int,
            intermediate: t,
            residues,
        }
    }

    /// Check if two integers are equal
    pub fn is_eq(&self, other: &Self) -> bool {
        let self_native = P::compose(self.limbs);
        let other_native = P::compose(other.limbs);
        self_native == other_native
    }
}

#[cfg(test)]
mod test {
    use super::*;
    use crate::eigentrust::params::rns::bn256::Bn256_4_68;
    use halo2_curves::bn256::{Fq, Fr};
    use num_integer::Integer as NumInteger;
    use num_traits::{FromPrimitive, One, Zero};
    use std::str::FromStr;

    #[test]
    fn should_mul_two_numbers() {
        // Testing mul with two elements.
        let a_big = BigUint::from_str("2188824282428718582428782428718558718582").unwrap();
        let b_big = Bn256_4_68::wrong_modulus() - BigUint::from_u8(1).unwrap();
        let big_answer = a_big.clone() * b_big.clone();
        let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big);
        let b = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(b_big);
        let c = a.mul(&b);

        assert!(Bn256_4_68::constrain_binary_crt(
            c.intermediate,
            c.result.limbs,
            c.residues
        ));
        assert_eq!(
            c.result.value(),
            big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
        );
    }

    #[test]
    fn should_mul_zero() {
        // Testing mul with input zero.
        let a_big = BigUint::from_str("0").unwrap();
        let b_big = Bn256_4_68::wrong_modulus() - BigUint::from_u8(1).unwrap();
        let big_answer = a_big.clone() * b_big.clone();
        let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big);
        let b = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(b_big);
        let c = a.mul(&b);

        assert!(Bn256_4_68::constrain_binary_crt(
            c.intermediate,
            c.result.limbs,
            c.residues
        ));
        assert_eq!(
            c.result.value(),
            big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
        );
    }

    #[test]
    fn should_mul_accumulate_array_of_small_numbers() {
        // Testing mul with array of 8 small elements.
        let a_big = BigUint::from_str("2188824286654430").unwrap();
        let a_big_array = [
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big,
        ];
        let carry = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(BigUint::one());
        let mut acc = carry.mul(&carry);
        let mut big_answer = BigUint::one();
        for i in 0..8 {
            big_answer *= a_big_array[i].clone();
            let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big_array[i].clone());
            acc = acc.result.mul(&a);

            assert!(Bn256_4_68::constrain_binary_crt(
                acc.intermediate,
                acc.result.limbs,
                acc.residues
            ));
            assert_eq!(
                acc.result.value(),
                big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
            );
        }
    }

    #[test]
    fn should_mul_accumulate_array_of_big_numbers() {
        // Testing mul with array of 8 big elements.
        let a_big = BigUint::from_str(
            "21888242871839275222246405745257275088696311157297823662689037894645226208582",
        )
        .unwrap();
        let a_big_array = [
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big,
        ];
        let carry = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(BigUint::one());
        let mut acc = carry.mul(&carry);
        let mut big_answer = BigUint::one();
        for i in 0..8 {
            big_answer *= a_big_array[i].clone();
            let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big_array[i].clone());
            acc = acc.result.mul(&a);

            assert!(Bn256_4_68::constrain_binary_crt(
                acc.intermediate,
                acc.result.limbs,
                acc.residues
            ));
            assert_eq!(
                acc.result.value(),
                big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
            );
        }
    }

    #[test]
    fn should_add_two_numbers() {
        // Testing add with two elements.
        let a_big = BigUint::from_str(
            "2188824287183927522224640574525727508869631115729782366268903789426208582",
        )
        .unwrap();
        let b_big = BigUint::from_str(
            "21888242871839275222246405745257275088696311157297823662689037894645226208582",
        )
        .unwrap();
        let big_answer = a_big.clone() + b_big.clone();
        let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big);
        let b = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(b_big);
        let c = a.add(&b);

        assert!(Bn256_4_68::constrain_binary_crt(
            c.intermediate,
            c.result.limbs,
            c.residues
        ));
        assert_eq!(
            c.result.value(),
            big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
        );
    }

    #[test]
    fn should_add_biggest_number_plus() {
        // Testing add with biggest field value + 1.
        let a_big = BigUint::from_str("1").unwrap();
        let b_big = BigUint::from_str(
            "21888242871839275222246405745257275088696311157297823662689037894645226208582",
        )
        .unwrap();
        let big_answer = a_big.clone() + b_big.clone();
        let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big);
        let b = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(b_big);
        let c = a.add(&b);

        assert!(Bn256_4_68::constrain_binary_crt(
            c.intermediate,
            c.result.limbs,
            c.residues
        ));
        assert_eq!(
            c.result.value(),
            big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
        );
    }

    #[test]
    fn should_add_accumulate_array_of_small_numbers() {
        // Testing add with array of 8 small elements.
        let a_big = BigUint::from_str("4057452572750886963137894").unwrap();
        let a_big_array = [
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big,
        ];
        let carry = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(BigUint::zero());
        let mut acc = carry.add(&carry);
        let mut big_answer = BigUint::zero();
        for i in 0..8 {
            big_answer += a_big_array[i].clone();
            let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big_array[i].clone());
            acc = acc.result.add(&a);

            assert!(Bn256_4_68::constrain_binary_crt(
                acc.intermediate,
                acc.result.limbs,
                acc.residues
            ));
            assert_eq!(
                acc.result.value(),
                big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
            );
        }
    }

    #[test]
    fn should_add_accumulate_array_of_big_numbers() {
        // Testing add with array of 8 big elements.
        let a_big = BigUint::from_str(
            "21888242871839275222246405745257275088696311157297823662689037894645226208582",
        )
        .unwrap();
        let a_big_array = [
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big.clone(),
            a_big,
        ];
        let carry = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(BigUint::one());
        let mut acc = carry.mul(&carry);
        let mut big_answer = BigUint::one();
        for i in 0..8 {
            big_answer += a_big_array[i].clone();
            let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big_array[i].clone());
            acc = acc.result.add(&a);

            assert!(Bn256_4_68::constrain_binary_crt(
                acc.intermediate,
                acc.result.limbs,
                acc.residues
            ));
            assert_eq!(
                acc.result.value(),
                big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
            );
        }
    }

    #[test]
    fn should_sub_two_numbers_positive() {
        // Testing sub with two elements bigger - smaller.
        let a_big = BigUint::from_str("2904853095839045839045839045738478394657834658737465873645")
            .unwrap();
        let b_big = BigUint::from_str("1345345345345").unwrap();
        let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big.clone());
        let b = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(b_big.clone());
        let c = a.sub(&b);

        assert!(Bn256_4_68::constrain_binary_crt(
            c.intermediate,
            c.result.limbs,
            c.residues
        ));
        assert_eq!(
            (c.result.value() + b_big).mod_floor(&Bn256_4_68::wrong_modulus()),
            a_big
        );
    }

    #[test]
    fn should_sub_two_numbers_negative() {
        // Testing sub with two elements smaller - bigger.
        let a_big = BigUint::from_str("1345345345345").unwrap();
        let b_big = BigUint::from_str("2904853095839045839045839045738478394657834658737465873645")
            .unwrap();
        let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big.clone());
        let b = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(b_big.clone());
        let c = a.sub(&b);

        assert!(Bn256_4_68::constrain_binary_crt(
            c.intermediate,
            c.result.limbs,
            c.residues
        ));
        assert_eq!(
            (c.result.value() + b_big).mod_floor(&Bn256_4_68::wrong_modulus()),
            a_big
        );
    }

    #[test]
    fn should_div_two_numbers() {
        // Testing div with two elements.
        let a_big = BigUint::from_str(
            "21888242871839275222246405745257275088696311157297823662689037894645226208582",
        )
        .unwrap();
        let b_big = BigUint::from_str("2").unwrap();
        let big_answer = a_big.clone() / b_big.clone();
        let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big);
        let b = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(b_big);
        let c = a.div(&b);
        assert!(Bn256_4_68::constrain_binary_crt(
            c.intermediate,
            c.result.limbs,
            c.residues
        ));
        assert_eq!(
            c.result.value(),
            big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
        );
    }

    #[test]
    fn should_div_two_numbers_zero() {
        // Testing div with dividing 0 to a another value.
        let a_big = BigUint::from_str("0").unwrap();
        let b_big = BigUint::from_str(
            "21888242871839275222246405745257275088696311157297823662689037894645226208582",
        )
        .unwrap();
        let big_answer = a_big.clone() / b_big.clone();
        let a = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(a_big);
        let b = Integer::<Fq, Fr, 4, 68, Bn256_4_68>::new(b_big);
        let c = a.div(&b);
        assert!(Bn256_4_68::constrain_binary_crt(
            c.intermediate,
            c.result.limbs,
            c.residues
        ));
        assert_eq!(
            c.result.value(),
            big_answer.mod_floor(&Bn256_4_68::wrong_modulus())
        );
    }
}
