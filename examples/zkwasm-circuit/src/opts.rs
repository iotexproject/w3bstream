use clap::{Parser, Subcommand};

#[derive(Debug, Parser)]
#[clap(name = "zkwasm-circuit", version = "0.1.0")]
pub struct Opts {
    #[clap(subcommand)]
    pub sub: Subcommands,
}

#[derive(Debug, Subcommand)]

pub enum Subcommands {
    #[clap(name = "verify")]
    #[clap(about = "Locally verify proof.")]
    Verfiy {
        #[clap(long, short, value_name = "proof-file")]
        proof: String,
    },
}