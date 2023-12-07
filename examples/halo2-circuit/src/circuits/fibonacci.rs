use halo2_proofs::{
    arithmetic::Field,
    circuit::{Layouter, SimpleFloorPlanner},
    plonk::{Advice, Circuit, Column, ConstraintSystem, Constraints, Error, Instance, Selector},
    poly::Rotation,
};
use std::marker::PhantomData;

#[derive(Clone, Debug)]
pub struct FibonacciConfig {
    advice: [Column<Advice>; 2],

    instance: Column<Instance>,

    selector: Selector,
}

#[derive(Debug, Clone)]
struct FibonacciChip<F: Field> {
    config: FibonacciConfig,
    _marker: PhantomData<F>,
}

impl<F: Field> FibonacciChip<F> {
    pub fn construct(config: FibonacciConfig) -> Self {
        Self {
            config,
            _marker: PhantomData,
        }
    }

    fn configure(
        meta: &mut ConstraintSystem<F>,
        advice: [Column<Advice>; 2],
        instance: Column<Instance>,
    ) -> FibonacciConfig {
        meta.enable_equality(instance);
        for column in &advice {
            meta.enable_equality(*column);
        }
        let selector = meta.selector();

        meta.create_gate("fibonacci", |meta| {
            // | ins   | a0     |   a1   | seletor|
            // |-------|------- |------- |------- |
            // |   a   | f(0)=a | f(1)=b |    1   |
            // |   b   | f(2)=b | f(3)   |    1   |
            // |  out  | f(4)   | f(5)   |    0   |
            let s = meta.query_selector(selector);
            let lc = meta.query_advice(advice[0], Rotation::cur());
            let rc = meta.query_advice(advice[1], Rotation::cur());
            let ln = meta.query_advice(advice[0], Rotation::next());
            let rn = meta.query_advice(advice[1], Rotation::next());
            Constraints::with_selector(s, vec![(lc + rc.clone() - ln.clone()), (rc + ln - rn)])
        });

        FibonacciConfig {
            advice,
            selector,
            instance,
        }
    }
}

#[derive(Debug, Default, Clone)]
pub struct FibonacciCircuit<F: Field> {
    pub n: usize,
    pub _marker: PhantomData<F>,
}

impl<F: Field> Circuit<F> for FibonacciCircuit<F> {
    type Config = FibonacciConfig;

    type FloorPlanner = SimpleFloorPlanner;

    fn without_witnesses(&self) -> Self {
        Self::default()
    }

    fn configure(meta: &mut ConstraintSystem<F>) -> Self::Config {
        let advice = [meta.advice_column(), meta.advice_column()];
        let instance = meta.instance_column();

        FibonacciChip::configure(meta, advice, instance)
    }

    fn synthesize(
        &self,
        config: Self::Config,
        mut layouter: impl halo2_proofs::circuit::Layouter<F>,
    ) -> Result<(), Error> {
        let field_chip = FibonacciChip::<F>::construct(config);

        let out = layouter.assign_region(
            || "fibo region",
            |mut region| {
                let lhs = field_chip.config.advice[0];
                let rhs = field_chip.config.advice[1];
                let instance = field_chip.config.instance;
                let s = field_chip.config.selector;

                let mut prev_left =
                    region.assign_advice_from_instance(|| "f0", instance, 0, lhs, 0)?;
                let mut prev_right =
                    region.assign_advice_from_instance(|| "f1", instance, 1, rhs, 0)?;

                for i in 1..=self.n / 2 {
                    s.enable(&mut region, i - 1)?;
                    let value = prev_left.value().copied() + prev_right.value().copied();

                    let cur_left = region.assign_advice(|| "f left", lhs, i, || value)?;
                    let value = prev_right.value().copied() + cur_left.value().copied();
                    let cur_right = region.assign_advice(|| "f right", rhs, i, || value)?;
                    prev_left = cur_left;
                    prev_right = cur_right;
                }

                if self.n % 2 == 0 {
                    return Ok(prev_left);
                }
                Ok(prev_right)
            },
        )?;

        layouter
            .namespace(|| "out")
            .constrain_instance(out.cell(), field_chip.config.instance, 2)
    }
}

pub fn fib(n: u64) -> u64 {
    match n {
        0 => 1,
        1 => 1,
        _ => fib(n - 1) + fib(n - 2),
    }
}

#[cfg(test)]
mod tests {
    use std::marker::PhantomData;

    use super::{fib, FibonacciCircuit};
    use halo2_curves::bn256::Fr;
    use halo2_proofs::dev::MockProver;

    #[test]
    fn verify() {
        let k = 4;

        let f0 = Fr::from(1);
        let f1 = Fr::from(1);
        let n = 11;
        let out = Fr::from(fib(n));

        let circuit = FibonacciCircuit {
            n: n as usize,
            _marker: PhantomData,
        };

        let public_inputs = vec![f0, f1, out];

        let prover = MockProver::run(k, &circuit, vec![public_inputs.clone()]).unwrap();
        assert_eq!(prover.verify(), Ok(()));
    }
}
