use crate::eigentrust::{
    poseidon::{native::Poseidon, RoundParams},
    FieldExt, SpongeHasher,
};
use std::marker::PhantomData;

/// Constructs objects.
#[derive(Clone)]
pub struct PoseidonSponge<F: FieldExt, const WIDTH: usize, P>
where
    P: RoundParams<F, WIDTH>,
{
    /// Constructs a vector for the inputs.
    inputs: Vec<F>,
    /// Internal state
    state: [F; WIDTH],
    /// Constructs a phantom data for the parameters.
    _params: PhantomData<P>,
}

impl<F: FieldExt, const WIDTH: usize, P> PoseidonSponge<F, WIDTH, P>
where
    P: RoundParams<F, WIDTH>,
{
    /// Create objects.
    pub fn new() -> Self {
        Self {
            inputs: Vec::new(),
            state: [F::ZERO; WIDTH],
            _params: PhantomData,
        }
    }

    /// Clones and appends all elements from a slice to the vec.
    pub fn update(&mut self, inputs: &[F]) {
        self.inputs.extend_from_slice(inputs);
    }

    /// Absorb the data in and split it into
    /// chunks of size WIDTH.
    fn load_state(chunk: &[F]) -> [F; WIDTH] {
        assert!(chunk.len() <= WIDTH);
        let mut fixed_chunk = [F::ZERO; WIDTH];
        fixed_chunk[..chunk.len()].copy_from_slice(chunk);
        fixed_chunk
    }

    /// Squeeze the data out by
    /// permuting until no more chunks are left.
    pub fn squeeze(&mut self) -> F {
        if self.inputs.is_empty() {
            self.inputs.push(F::ZERO);
        }

        for chunk in self.inputs.chunks(WIDTH) {
            let mut input = [F::ZERO; WIDTH];

            // Absorb
            let loaded_state = Self::load_state(chunk);
            for i in 0..WIDTH {
                input[i] = loaded_state[i] + self.state[i];
            }

            // Permute
            let pos = Poseidon::<_, WIDTH, P>::new(input);
            self.state = pos.permute();
        }

        // Clear the inputs, and return the result
        self.inputs.clear();
        self.state[0]
    }
}

impl<F: FieldExt, const WIDTH: usize, P> Default for PoseidonSponge<F, WIDTH, P>
where
    P: RoundParams<F, WIDTH>,
{
    fn default() -> Self {
        Self::new()
    }
}

impl<F: FieldExt, const WIDTH: usize, P> SpongeHasher<F> for PoseidonSponge<F, WIDTH, P>
where
    P: RoundParams<F, WIDTH>,
{
    fn new() -> Self {
        Self::new()
    }

    fn update(&mut self, inputs: &[F]) {
        Self::update(self, inputs)
    }

    fn squeeze(&mut self) -> F {
        PoseidonSponge::squeeze(self)
    }
}
