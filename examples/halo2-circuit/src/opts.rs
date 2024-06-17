use clap::{Parser, Subcommand};

#[derive(Debug, Parser)]
#[clap(name = "halo2-circuit", version = "0.1.0")]
pub struct Opts {
    #[clap(subcommand)]
    pub sub: Subcommands,
}

#[derive(Debug, Subcommand)]
pub enum Subcommands {
    #[clap(name = "solidity")]
    #[clap(about = "Generate verifier solidity contract.")]
    Solidity {
        #[clap(
            long,
            short,
            value_name = "file",
            default_value = "Verifier.sol"
        )]
        file: String,
    },

    #[clap(name = "proof")]
    #[clap(about = "Generate proof.")]
    Proof {
        #[clap(
            long,
            value_name = "private_a",
            default_value = "3"
        )]
        private_a: u64,
        #[clap(
            long,
            value_name = "private_b",
            default_value = "4"
        )]
        private_b: u64,
        #[clap(
            long,
            value_name = "project_id",
        )]
        project_id: u64,
        #[clap(
            long,
            value_name = "task_id",
        )]
        task_id: u64,
    },

    #[clap(name = "verify")]
    #[clap(about = "Local verify proof.")]
    Verfiy {
        #[clap(long, value_name = "proof-file")]
        proof: String,
        #[clap(long, value_name = "public-input")]
        public: u64,
        #[clap(long, value_name = "project-input")]
        project: u64,
        #[clap(long, value_name = "task-input")]
        task: u64,
    },
}
