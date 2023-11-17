use std::marker::PhantomData;

use halo2_proofs::{
    circuit::{Layouter, SimpleFloorPlanner, Value},
    plonk::{Advice, Circuit, Column, ConstraintSystem, Error, Expression, Instance, Selector},
    poly::Rotation,
};

use super::gadgets::{
    is_zero::{IsZeroChip, IsZeroConfig, IsZeroInstructin},
    Field,
};

#[derive(Clone, Debug)]
pub struct FunctionConfig<F: Field> {
    pub selector: Selector,
    pub a: Column<Advice>,
    pub b: Column<Advice>,
    pub c: Column<Advice>,
    pub a_equals_b: IsZeroConfig<F>,
    pub instance: Column<Instance>,
}

/// prove: f(a, b, c) = if a == b {c} else {a - b}
#[derive(Clone, Default)]
pub struct FunctionCircuit<F: Field> {
    a: u64,
    b: u64,
    c: u64,
    _marker: PhantomData<F>,
}

impl<F: Field> Circuit<F> for FunctionCircuit<F> {
    type Config = FunctionConfig<F>;
    type FloorPlanner = SimpleFloorPlanner;

    fn without_witnesses(&self) -> Self {
        FunctionCircuit::default()
    }

    fn configure(meta: &mut ConstraintSystem<F>) -> Self::Config {
        let selector = meta.selector();
        let a = meta.advice_column();
        let b = meta.advice_column();
        let c = meta.advice_column();
        let instance = meta.instance_column();

        meta.enable_equality(a);
        meta.enable_equality(instance);

        let is_zero_value_inv = meta.advice_column();

        let a_equals_b = IsZeroChip::configure(
            meta,
            |meta| meta.query_selector(selector),
            |meta| meta.query_advice(a, Rotation::cur()) - meta.query_advice(b, Rotation::cur()),
            is_zero_value_inv,
        );

        meta.create_gate("function gate", |meta| {
            let s = meta.query_selector(selector);
            let ac = meta.query_advice(a, Rotation::cur());
            let b = meta.query_advice(b, Rotation::cur());
            let c = meta.query_advice(c, Rotation::cur());
            let output = meta.query_advice(a, Rotation::next());

            vec![
                s.clone() * a_equals_b.clone().is_zero_expression.clone() * (output.clone() - c),
                s * (Expression::Constant(F::ONE) - a_equals_b.clone().is_zero_expression)
                    * (output - ac + b),
            ]
        });

        FunctionConfig {
            selector,
            a,
            b,
            c,
            a_equals_b,
            instance,
        }
    }

    fn synthesize(
        &self,
        config: Self::Config,
        mut layouter: impl Layouter<F>,
    ) -> Result<(), Error> {
        let is_zero_chip = IsZeroChip::construct(config.a_equals_b);

        let output = layouter.assign_region(
            || "function region",
            |mut region| {
                config.selector.enable(&mut region, 0)?;

                region.assign_advice(|| "load a", config.a, 0, || Value::known(F::from(self.a)))?;
                region.assign_advice(|| "load b", config.b, 0, || Value::known(F::from(self.b)))?;
                region.assign_advice(|| "load c", config.c, 0, || Value::known(F::from(self.c)))?;
                let output = if self.a == self.b {
                    F::from(self.c)
                } else {
                    F::from(self.a) - F::from(self.b)
                };
                let output =
                    region.assign_advice(|| "output", config.a, 1, || Value::known(output))?;

                is_zero_chip.assign(
                    &mut region,
                    0,
                    Value::known(F::from(self.a) - F::from(self.b)),
                )?;

                Ok(output)
            },
        )?;

        layouter.constrain_instance(output.cell(), config.instance, 0)
    }
}

#[cfg(test)]
mod tests {
    use halo2_curves::bn256::Fr;
    use halo2_proofs::dev::MockProver;
    use std::marker::PhantomData;

    use super::FunctionCircuit;

    #[test]
    fn verify() {
        let a = 2;
        let b = 10;
        let c = 3;
        let circuit = FunctionCircuit::<Fr> {
            a,
            b,
            c,
            _marker: PhantomData,
        };

        let out = if a == b {
            Fr::from(c)
        } else {
            Fr::from(a) - Fr::from(b)
        };
        let out = vec![out];

        let prover = MockProver::run(4, &circuit, vec![out]).unwrap();
        prover.assert_satisfied();
    }
}
