use halo2_proofs::{
    circuit::{Chip, Region, Value},
    plonk::{Advice, Column, ConstraintSystem, Error, Expression, VirtualCells},
    poly::Rotation,
};

use super::Field;

pub trait IsZeroInstructin<F: Field> {
    fn assign(
        &self,
        region: &mut Region<'_, F>,
        offset: usize,
        value: Value<F>,
    ) -> Result<(), Error>;
}

#[derive(Clone, Debug)]
pub struct IsZeroConfig<F> {
    pub value_inv: Column<Advice>,
    pub is_zero_expression: Expression<F>,
}

pub struct IsZeroChip<F: Field> {
    config: IsZeroConfig<F>,
}

impl<F: Field> IsZeroChip<F> {
    pub fn construct(config: IsZeroConfig<F>) -> Self {
        IsZeroChip { config }
    }
}

impl<F: Field> IsZeroChip<F> {
    pub fn configure(
        meta: &mut ConstraintSystem<F>,
        q_enable: impl FnOnce(&mut VirtualCells<'_, F>) -> Expression<F>,
        value: impl FnOnce(&mut VirtualCells<'_, F>) -> Expression<F>,
        value_inv: Column<Advice>,
    ) -> IsZeroConfig<F> {
        // dummy
        let mut is_zero_expression = Expression::Constant(F::ZERO);

        meta.create_gate("is zero gate", |meta| {
            let q_enable = q_enable(meta);

            let value_inv = meta.query_advice(value_inv, Rotation::cur());
            let value = value(meta);

            is_zero_expression = Expression::Constant(F::ONE) - value.clone() * value_inv;

            [q_enable * value * is_zero_expression.clone()]
        });

        IsZeroConfig::<F> {
            value_inv,
            is_zero_expression,
        }
    }
}

impl<F: Field> IsZeroInstructin<F> for IsZeroChip<F> {
    fn assign(
        &self,
        region: &mut Region<'_, F>,
        offset: usize,
        value: Value<F>,
    ) -> Result<(), Error> {
        let config = self.config();

        let value_inv = value.into_field().invert();
        region.assign_advice(
            || "witness inverse of value",
            config.value_inv,
            offset,
            || value_inv,
        )?;

        Ok(())
    }
}

impl<F: Field> Chip<F> for IsZeroChip<F> {
    type Config = IsZeroConfig<F>;
    type Loaded = ();

    fn config(&self) -> &Self::Config {
        &self.config
    }

    fn loaded(&self) -> &Self::Loaded {
        &()
    }
}

#[cfg(test)]
mod test {
    use std::marker::PhantomData;

    use halo2_proofs::{
        circuit::{Layouter, SimpleFloorPlanner, Value},
        dev::MockProver,
        halo2curves::bn256::Fr,
        plonk::{Advice, Circuit, Column, ConstraintSystem, Error, Selector},
        poly::Rotation,
    };

    use super::{Field, IsZeroChip, IsZeroConfig, IsZeroInstructin};

    macro_rules! try_test_circuit {
        ($values:expr, $checks:expr) => {{
            let k = usize::BITS - $values.len().leading_zeros() + 2;
            let circuit = TestCircuit::<Fr> {
                values: Some($values),
                checks: Some($checks),
                _marker: PhantomData,
            };
            let prover = MockProver::<Fr>::run(k, &circuit, vec![]).unwrap();
            prover.assert_satisfied()
        }};
    }

    macro_rules! try_test_circuit_error {
        ($values:expr, $checks:expr) => {{
            let k = usize::BITS - $values.len().leading_zeros() + 2;
            let circuit = TestCircuit::<Fr> {
                values: Some($values),
                checks: Some($checks),
                _marker: PhantomData,
            };
            let prover = MockProver::<Fr>::run(k, &circuit, vec![]).unwrap();
            assert!(prover.verify_par().is_err());
        }};
    }

    #[derive(Clone, Debug)]
    struct TestCircuitConfig<F> {
        q_enable: Selector,
        value: Column<Advice>,
        check: Column<Advice>,
        is_zero: IsZeroConfig<F>,
    }

    #[derive(Default)]
    struct TestCircuit<F: Field> {
        values: Option<Vec<u64>>,
        checks: Option<Vec<bool>>,
        _marker: PhantomData<F>,
    }

    impl<F: Field> Circuit<F> for TestCircuit<F> {
        type Config = TestCircuitConfig<F>;
        type FloorPlanner = SimpleFloorPlanner;

        fn without_witnesses(&self) -> Self {
            Self::default()
        }

        fn configure(meta: &mut ConstraintSystem<F>) -> Self::Config {
            let q_enable = meta.complex_selector();
            let value = meta.advice_column();
            let check = meta.advice_column();
            let value_inv = meta.advice_column();

            let is_zero = IsZeroChip::configure(
                meta,
                |meta| meta.query_selector(q_enable),
                |meta| {
                    let value_prev = meta.query_advice(value, Rotation::prev());
                    let value_cur = meta.query_advice(value, Rotation::cur());
                    value_cur - value_prev
                },
                value_inv,
            );

            let config = Self::Config {
                q_enable,
                value,
                check,
                is_zero,
            };

            meta.create_gate("check is zero", |meta| {
                let q_enable = meta.query_selector(q_enable);

                let check = meta.query_advice(config.check, Rotation::cur());

                vec![q_enable * (config.is_zero.is_zero_expression.clone() - check)]
            });

            config
        }

        fn synthesize(
            &self,
            config: Self::Config,
            mut layouter: impl Layouter<F>,
        ) -> Result<(), Error> {
            let chip = IsZeroChip::construct(config.is_zero.clone());

            let values: Vec<_> = self
                .values
                .as_ref()
                .map(|values| values.iter().map(|value| F::from(*value)).collect())
                .ok_or(Error::Synthesis)?;
            let checks = self.checks.as_ref().ok_or(Error::Synthesis)?;
            let (first_value, values) = values.split_at(1);
            let first_value = first_value[0];

            layouter.assign_region(
                || "witness",
                |mut region| {
                    region.assign_advice(
                        || "first row value",
                        config.value,
                        0,
                        || Value::known(first_value),
                    )?;

                    let mut value_prev = first_value;
                    for (idx, (value, check)) in values.iter().zip(checks).enumerate() {
                        region.assign_advice(
                            || "value",
                            config.value,
                            idx + 1,
                            || Value::known(*value),
                        )?;
                        region.assign_advice(
                            || "check",
                            config.check,
                            idx + 1,
                            || Value::known(F::from(*check as u64)),
                        )?;

                        config.q_enable.enable(&mut region, idx + 1)?;
                        chip.assign(&mut region, idx + 1, Value::known(*value - value_prev))?;

                        value_prev = *value;
                    }

                    Ok(())
                },
            )
        }
    }

    #[test]
    fn test_circuit() {
        try_test_circuit!(vec![1, 2, 3, 4, 5], vec![false, false, false, false]);
        try_test_circuit!(vec![1, 2, 3, 5, 5], vec![false, false, false, true]);
        try_test_circuit_error!(vec![1, 2], vec![true]);
        try_test_circuit_error!(vec![1, 1], vec![false]);
    }
}
