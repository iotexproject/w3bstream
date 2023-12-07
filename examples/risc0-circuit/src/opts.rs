use clap::{Parser, Subcommand};

#[derive(Debug, Parser)]
#[clap(name = "risc0-circuit", version = "0.1.0")]
pub struct Opts {
    #[clap(subcommand)]
    pub sub: Subcommands,
}

#[derive(Debug, Subcommand)]

pub enum Subcommands {
    #[clap(name = "verify")]
    #[clap(about = "Local verify proof.")]
    Verfiy {
        #[clap(long, short, value_name = "proof-file")]
        proof: String,
        #[clap(long, short, value_name = "image-id")]
        image_id: String,
    },
}