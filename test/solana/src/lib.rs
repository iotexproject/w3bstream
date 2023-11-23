use solana_program::{
    account_info::{next_account_info, AccountInfo},
    entrypoint,
    entrypoint::ProgramResult,
    msg,
    pubkey::Pubkey,
};

entrypoint!(process_instruction);

pub fn process_instruction(
    program_id: &Pubkey,
    accounts: &[AccountInfo],
    input: &[u8],
) -> ProgramResult {
    let accounts_iter = &mut accounts.iter();
    let account = next_account_info(accounts_iter)?;

    match input[0] {
        // 0 means submit proof
        0 => {
            // proof is in the rest of the input
            msg!("Received proof: {:?}", &input[1..]);
            // save proof to account
            account.realloc(input[1..].len(), true)?;
            let mut data = account.try_borrow_mut_data()?;
            data.copy_from_slice(&input[1..]);
            msg!("updated data: {:?}", data);
        }
        _ => {}
    }

    Ok(())
}
