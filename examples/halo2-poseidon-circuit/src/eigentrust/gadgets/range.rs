use crate::eigentrust::{params::rns::decompose_big, utils::fe_to_big, FieldExt, RegionCtx};
use halo2_proofs::{
    circuit::{AssignedCell, Layouter, Region, Value},
    plonk::{ConstraintSystem, Constraints, Error, Selector},
    poly::Rotation,
};

use super::lt_eq_lookup::{MockChip, MockChipset, MockCommonConfig};

/// Short length bits check chip using lookup table.
///
/// Checks if the target value is `S` bits long.
#[derive(Debug, Clone)]
pub struct LookupShortWordCheckChip<F: FieldExt, const K: usize, const S: usize> {
    x: AssignedCell<F, F>,
}

impl<F: FieldExt, const K: usize, const S: usize> LookupShortWordCheckChip<F, K, S> {
    /// Construct new instance
    pub fn new(x: AssignedCell<F, F>) -> Self {
        Self { x }
    }
}

impl<F: FieldExt, const K: usize, const S: usize> MockChip<F>
    for LookupShortWordCheckChip<F, K, S>
{
    type Output = AssignedCell<F, F>;

    fn configure(mock_common: &MockCommonConfig, meta: &mut ConstraintSystem<F>) -> Selector {
        //
        // IMPORTANT: For the maximal usage of CommonConfig columns(20 advice + 10 fixed),
        //			  we use the advice column 16 - 17. (17th ~ 18th)
        //

        assert!(0 < S && S < K, "Word bits should be less than target bits.");

        let bitshift_selector = meta.complex_selector();

        let word_column = mock_common.common.advice[16];
        let shifted_word_column = mock_common.common.advice[17];

        meta.lookup("Check K bits limit", |meta| {
            let bitshift_selector = meta.query_selector(bitshift_selector);
            let shifted_word = meta.query_advice(shifted_word_column, Rotation::cur());

            vec![(bitshift_selector * shifted_word, mock_common.table)]
        });

        // For short lookups, check that the word has been shifted by the correct number
        // of bits. https://p.z.cash/halo2-0.1:decompose-short-lookup
        meta.create_gate("Short word S bits limit", |meta| {
            let bitshift_selector = meta.query_selector(bitshift_selector);
            let word = meta.query_advice(word_column, Rotation::cur());
            let shifted_word = meta.query_advice(shifted_word_column, Rotation::cur());

            let two_pow_k_min_s = F::from(1 << (K - S));

            // shifted_word = word * 2^{K-S}
            Constraints::with_selector(
                bitshift_selector,
                Some(word * two_pow_k_min_s - shifted_word),
            )
        });

        bitshift_selector
    }

    fn synthesize(
        self,
        mock_common: &MockCommonConfig,
        selector: &Selector,
        mut layouter: impl Layouter<F>,
    ) -> Result<Self::Output, Error> {
        layouter.assign_region(
            || "short word check chip",
            |region| {
                let mut ctx = RegionCtx::new(region, 0);
                ctx.enable(*selector)?;

                // Assign original value
                let assigned_x = ctx.copy_assign(mock_common.common.advice[16], self.x.clone())?;

                // Assign shifted value
                let shifted_word = self.x.value().cloned() * Value::known(F::from(1 << (K - S)));
                ctx.assign_advice(mock_common.common.advice[17], shifted_word)?;

                Ok(assigned_x)
            },
        )
    }
}

/// Range check chip using lookup table.
///
/// Check if the target value is `K * N` bits long.
#[derive(Debug, Clone)]
pub struct LookupRangeCheckChip<F: FieldExt, const K: usize, const N: usize> {
    x: AssignedCell<F, F>,
}

impl<F: FieldExt, const K: usize, const N: usize> LookupRangeCheckChip<F, K, N> {
    /// Construct a new instance
    pub fn new(x: AssignedCell<F, F>) -> Self {
        Self { x }
    }
}

impl<F: FieldExt, const K: usize, const N: usize> MockChip<F> for LookupRangeCheckChip<F, K, N> {
    type Output = (AssignedCell<F, F>, AssignedCell<F, F>);

    fn configure(mock_common: &MockCommonConfig, meta: &mut ConstraintSystem<F>) -> Selector {
        //
        // IMPORTANT: For the maximal usage of CommonConfig columns(20 advice + 10 fixed),
        //			  we use the advice column 16 - 17. (17th ~ 18th)
        //

        let running_sum_selector = meta.complex_selector();

        let running_sum = mock_common.common.advice[16];

        meta.lookup("running_sum check", |meta| {
            let running_sum_selector = meta.query_selector(running_sum_selector);
            let z_cur = meta.query_advice(running_sum, Rotation::cur());

            // Example of 16-bit number:
            // K = 8
            // N = 2

            // z_0 1111111100000001 (little-endian)
            // z_1 00000001
            // z_2 00000000

            // a_0 = z_0 - z_1 * (1 << 8) = 11111111
            // a_1 = z_1 - z_2 * (1 << 8) = 00000001

            // In the case of a running sum decomposition, we recover the word from
            // the difference of the running sums:
            // z_i = 2^{K}⋅z_{i + 1} + a_i
            // => a_i = z_i - 2^{K}⋅z_{i + 1}

            let z_next = meta.query_advice(running_sum, Rotation::next());
            let running_sum_word = z_cur - z_next * F::from(1 << K);
            let running_sum_lookup = running_sum_selector * running_sum_word;

            // Combine the running sum and short lookups:
            vec![(running_sum_lookup, mock_common.table)]
        });

        running_sum_selector
    }

    fn synthesize(
        self,
        mock_common: &MockCommonConfig,
        selector: &Selector,
        mut layouter: impl Layouter<F>,
    ) -> Result<Self::Output, Error> {
        layouter.assign_region(
            || "range check chip",
            |region: Region<'_, F>| {
                let mut ctx = RegionCtx::new(region, 0);
                let x = ctx.copy_assign(mock_common.common.advice[16], self.x.clone())?;

                // Take first "num_bits" bits of `element`.
                let words = x
                    .value()
                    .map(|&element| {
                        let fe_big = fe_to_big(element);
                        let limbs: [F; N] = decompose_big::<_, N, K>(fe_big);
                        limbs
                    })
                    .transpose_array();

                // Example of 16-bit number:
                // K = 8
                // N = 2
                // x:   1111111100000001
                // a_0: 11111111
                // a_1: 00000001
                //
                // z_0 = x
                // z_1 = (z_0 - a_0) / (2 ^ K)
                // z_2 = (z_1 - a_1) / (2 ^ K)
                //
                // [1111111100000001, 00000001, 0]
                // Last element should be 0

                // Assign cumulative sum such that
                // z_i = 2^{K}⋅z_{i + 1} + a_i =>
                // z_i - a_i = 2^{K}⋅z_{i + 1} =>
                // z_{i + 1} = (z_i - a_i) / (2^K)

                let mut z = x.clone();
                let mut last_word_cell = x;
                let words_len = words.len();
                let inv_two_pow_k = F::from(1u64 << K).invert().unwrap();
                for (id, word) in words.into_iter().enumerate() {
                    // Enable q_lookup on this row
                    ctx.enable(*selector)?;

                    if id == words_len - 1 {
                        last_word_cell = z.clone();
                    }

                    // z_next = (z_cur - m_cur) / 2^K
                    let z_next = z
                        .value()
                        .zip(word)
                        .map(|(z, word)| (*z - word) * inv_two_pow_k);

                    // Assign z_next
                    ctx.next();
                    z = ctx.assign_advice(mock_common.common.advice[16], z_next)?;
                }

                ctx.constrain_to_constant(z, F::ZERO)?;

                Ok((self.x.clone(), last_word_cell))
            },
        )
    }
}

/// RangeChipsetConfig
#[derive(Debug, Clone)]
pub struct RangeChipsetConfig {
    running_sum_selector: Selector,
    short_word_selecor: Selector,
}

impl RangeChipsetConfig {
    /// Construct a new instance
    pub fn new(running_sum_selector: Selector, short_word_selecor: Selector) -> Self {
        Self {
            running_sum_selector,
            short_word_selecor,
        }
    }
}

/// RangeChipset
#[derive(Debug, Clone)]
pub struct RangeChipset<F: FieldExt, const K: usize, const N: usize, const S: usize> {
    x: AssignedCell<F, F>,
}

impl<F: FieldExt, const K: usize, const N: usize, const S: usize> RangeChipset<F, K, N, S> {
    /// Constructs new chipset
    pub fn new(x: AssignedCell<F, F>) -> Self {
        Self { x }
    }
}

impl<F: FieldExt, const K: usize, const N: usize, const S: usize> MockChipset<F>
    for RangeChipset<F, K, N, S>
{
    type Config = RangeChipsetConfig;
    type Output = AssignedCell<F, F>;

    fn synthesize(
        self,
        mock_common: &MockCommonConfig,
        config: &RangeChipsetConfig,
        mut layouter: impl Layouter<F>,
    ) -> Result<Self::Output, Error> {
        // First, check if x is less than 256 bits
        let range_chip = LookupRangeCheckChip::<F, K, N>::new(self.x);
        let (x, last_word_cell) = range_chip.synthesize(
            mock_common,
            &config.running_sum_selector,
            layouter.namespace(|| "range check"),
        )?;

        let short_word_chip = LookupShortWordCheckChip::<F, K, S>::new(last_word_cell);
        let _ = short_word_chip.synthesize(
            mock_common,
            &config.short_word_selecor,
            layouter.namespace(|| "last word check"),
        )?;

        Ok(x)
    }
}

#[cfg(test)]
mod tests {
    use MockCommonConfig;

    use super::*;

    use halo2_proofs::{
        circuit::{Region, SimpleFloorPlanner, Table, Value},
        dev::MockProver,
        halo2curves::bn256::Fr,
        plonk::{Circuit, ConstraintSystem, Error},
    };

    const K: usize = 8;
    const N: usize = 2;
    const S: usize = 3;

    #[derive(Clone)]
    struct TestConfig {
        mock_common: MockCommonConfig,
        short_word_selector: Selector,
        running_sum_selector: Selector,
        lookup_range_check: RangeChipsetConfig,
    }

    impl TestConfig {
        fn new(meta: &mut ConstraintSystem<Fr>) -> Self {
            let mock_common = MockCommonConfig::new(meta);
            let short_word_selector =
                LookupShortWordCheckChip::<Fr, K, S>::configure(&mock_common, meta);
            let running_sum_selector =
                LookupRangeCheckChip::<Fr, K, N>::configure(&mock_common, meta);
            let lookup_range_check =
                RangeChipsetConfig::new(running_sum_selector, short_word_selector);

            Self {
                mock_common,
                short_word_selector,
                running_sum_selector,
                lookup_range_check,
            }
        }
    }

    struct TableAssigner;

    impl MockChipset<Fr> for TableAssigner {
        type Config = ();
        type Output = ();

        fn synthesize(
            self,
            common: &MockCommonConfig,
            _: &Self::Config,
            mut layouter: impl Layouter<Fr>,
        ) -> Result<Self::Output, Error> {
            // Loads the values [0..2^K) into `common.table`.
            layouter.assign_table(
                || "table_column",
                |mut table: Table<Fr>| {
                    // We generate the row values lazily (we only need them during keygen).
                    for index in 0..(1 << K) {
                        table.assign_cell(
                            || "table_column",
                            common.table,
                            index,
                            || Value::known(Fr::from(index as u64)),
                        )?;
                    }
                    Ok(())
                },
            )
        }
    }

    #[derive(Clone)]
    struct SWTestCircuit {
        x: Value<Fr>,
    }

    impl SWTestCircuit {
        fn new(x: Fr) -> Self {
            Self { x: Value::known(x) }
        }
    }

    impl Circuit<Fr> for SWTestCircuit {
        type Config = TestConfig;
        type FloorPlanner = SimpleFloorPlanner;

        fn without_witnesses(&self) -> Self {
            Self {
                x: Value::unknown(),
            }
        }

        fn configure(meta: &mut ConstraintSystem<Fr>) -> TestConfig {
            TestConfig::new(meta)
        }

        fn synthesize(
            &self,
            config: TestConfig,
            mut layouter: impl Layouter<Fr>,
        ) -> Result<(), Error> {
            let table_assigner = TableAssigner;
            table_assigner.synthesize(
                &config.mock_common,
                &(),
                layouter.namespace(|| "table setup"),
            )?;

            let x = layouter.assign_region(
                || "temp",
                |region: Region<'_, Fr>| {
                    let mut ctx = RegionCtx::new(region, 0);
                    let x = ctx.assign_advice(config.mock_common.common.advice[0], self.x)?;
                    Ok(x)
                },
            )?;

            let short_word_check_chip = LookupShortWordCheckChip::<Fr, K, S>::new(x);
            let _ = short_word_check_chip.synthesize(
                &config.mock_common,
                &config.short_word_selector,
                layouter.namespace(|| "short word check"),
            )?;
            Ok(())
        }
    }

    #[test]
    fn test_short_word_case() {
        let k = 9;
        let pub_ins = vec![];

        // Testing x is 3 bits
        let x = Fr::from(0x07); // 0b111

        let test_chip = SWTestCircuit::new(x);
        let prover = MockProver::run(k, &test_chip, vec![pub_ins.clone()]).unwrap();
        assert!(prover.verify().is_ok());

        // Should fail since x is 4 bits
        let x = Fr::from(0x09); // 0b1001

        let test_chip = SWTestCircuit::new(x);
        let prover = MockProver::run(k, &test_chip, vec![pub_ins]).unwrap();
        assert!(prover.verify().is_err());
    }

    #[derive(Clone)]
    struct LookupRangeTestCircuit {
        x: Value<Fr>,
    }

    impl LookupRangeTestCircuit {
        fn new(x: Fr) -> Self {
            Self { x: Value::known(x) }
        }
    }

    impl Circuit<Fr> for LookupRangeTestCircuit {
        type Config = TestConfig;
        type FloorPlanner = SimpleFloorPlanner;

        fn without_witnesses(&self) -> Self {
            Self {
                x: Value::unknown(),
            }
        }

        fn configure(meta: &mut ConstraintSystem<Fr>) -> TestConfig {
            TestConfig::new(meta)
        }

        fn synthesize(
            &self,
            config: TestConfig,
            mut layouter: impl Layouter<Fr>,
        ) -> Result<(), Error> {
            let table_assigner = TableAssigner;
            table_assigner.synthesize(
                &config.mock_common,
                &(),
                layouter.namespace(|| "table setup"),
            )?;

            let x = layouter.assign_region(
                || "temp",
                |region: Region<'_, Fr>| {
                    let mut ctx = RegionCtx::new(region, 0);
                    let x = ctx.assign_advice(config.mock_common.common.advice[0], self.x)?;
                    Ok(x)
                },
            )?;

            let range_check_chip = LookupRangeCheckChip::<Fr, K, N>::new(x);
            let _ = range_check_chip.synthesize(
                &config.mock_common,
                &config.running_sum_selector,
                layouter.namespace(|| "range check"),
            )?;
            Ok(())
        }
    }

    #[test]
    fn test_lookup_range_check() {
        let k = 9;
        let pub_ins = vec![];

        // Testing x is 16 bits
        let x = Fr::from(0xffff); // 0b1111111111111111

        let test_chip = LookupRangeTestCircuit::new(x);
        let prover = MockProver::run(k, &test_chip, vec![pub_ins.clone()]).unwrap();
        assert!(prover.verify().is_ok());

        // Should fail since x is 17 bits
        let x = Fr::from(0x10000); // 0b10000000000000000

        let test_chip = LookupRangeTestCircuit::new(x);
        let prover = MockProver::run(k, &test_chip, vec![pub_ins]).unwrap();
        assert!(prover.verify().is_err());
    }

    #[derive(Clone)]
    struct RangeTestCircuit {
        x: Value<Fr>,
    }

    impl RangeTestCircuit {
        fn new(x: Fr) -> Self {
            Self { x: Value::known(x) }
        }
    }

    impl Circuit<Fr> for RangeTestCircuit {
        type Config = TestConfig;
        type FloorPlanner = SimpleFloorPlanner;

        fn without_witnesses(&self) -> Self {
            Self {
                x: Value::unknown(),
            }
        }

        fn configure(meta: &mut ConstraintSystem<Fr>) -> TestConfig {
            TestConfig::new(meta)
        }

        fn synthesize(
            &self,
            config: TestConfig,
            mut layouter: impl Layouter<Fr>,
        ) -> Result<(), Error> {
            let table_assigner = TableAssigner;
            table_assigner.synthesize(
                &config.mock_common,
                &(),
                layouter.namespace(|| "table setup"),
            )?;

            let x = layouter.assign_region(
                || "temp",
                |region: Region<'_, Fr>| {
                    let mut ctx = RegionCtx::new(region, 0);
                    let x = ctx.assign_advice(config.mock_common.common.advice[0], self.x)?;
                    Ok(x)
                },
            )?;

            let lookup_range_check_chipset = RangeChipset::<Fr, K, N, S>::new(x);
            let _ = lookup_range_check_chipset.synthesize(
                &config.mock_common,
                &config.lookup_range_check,
                layouter.namespace(|| "lookup range check chipset"),
            )?;
            Ok(())
        }
    }

    #[test]
    fn test_range_check_chipset() {
        let k = 9;
        let pub_ins = vec![];

        // Testing x is 11 bits
        let x = Fr::from(0x7ff); // 0b11111111111

        let test_chip = RangeTestCircuit::new(x);
        let prover = MockProver::run(k, &test_chip, vec![pub_ins.clone()]).unwrap();
        assert!(prover.verify().is_ok());

        // Should fail since x is 12 bits
        let x = Fr::from(0xfff); // 0b111111111111

        let test_chip = RangeTestCircuit::new(x);
        let prover = MockProver::run(k, &test_chip, vec![pub_ins]).unwrap();
        assert!(prover.verify().is_err());
    }
}
