use halo2_curves::ff::Field;
use halo2_proofs::{
    circuit::{Layouter, SimpleFloorPlanner, Value},
    plonk::{Advice, Circuit, Column, ConstraintSystem, Instance, Selector},
};

#[derive(Clone, Debug)]
pub struct Config {
    pub advice: Column<Advice>,
    pub instance: Column<Instance>,
    pub selector: Selector,
}

#[derive(Default, Clone)]
pub struct EmptyCircuit<F: Field> {
    pub constant: F,
    pub a: Value<F>,
}

impl<F: Field> Circuit<F> for EmptyCircuit<F> {
    type Config = Config;
    type FloorPlanner = SimpleFloorPlanner;

    fn without_witnesses(&self) -> Self {
        Self::default()
    }

    fn configure(_meta: &mut ConstraintSystem<F>) -> Self::Config {
        todo!()
    }

    fn synthesize(
        &self,
        _config: Self::Config,
        mut _layouter: impl Layouter<F>,
    ) -> Result<(), halo2_proofs::plonk::Error> {
        todo!()
    }
}
