use std::io::Read;
use std::fs::File;
use std::io::Error;
use std::sync::Arc;

use ring;
use ring::{rand, signature};

use untrusted;

pub fn obtain_ticket(filepath: &str) {
    sign("a", filepath).unwrap();
}

fn sign(data: &str, filepath: &str) -> Result<(), SignError> {
    let private_key_der = read_file(filepath)?;
    let private_key_der = untrusted::Input::from(&private_key_der);
    let key_pair = signature::RSAKeyPair::from_der(private_key_der)
        .map_err(|ring::error::Unspecified| SignError::BadPrivateKey)?;
    let key_pair = Arc::new(key_pair);
    let mut signing_state =
        signature::RSASigningState::new(key_pair).map_err(|ring::error::Unspecified| SignError::OOM)?;
    let rng = rand::SystemRandom::new();
    let mut signature = vec![0; signing_state.key_pair().public_modulus_len()];
    signing_state
        .sign(&signature::RSA_PKCS1_SHA256, &rng, data.as_bytes(), &mut signature)
        .map_err(|ring::error::Unspecified| SignError::OOM)?;

    println!("{:?}", signature);
    Ok(())
}

#[derive(Debug)]
enum SignError {
    IO(Error),
    BadPrivateKey,
    OOM,
}

fn read_file(path: &str) -> Result<Vec<u8>, SignError> {
    let mut f = File::open(path).map_err(|e| SignError::IO(e))?;
    let mut buffer: Vec<u8> = Vec::new();
    f.read_to_end(&mut buffer).map_err(|e| SignError::IO(e))?;;
    Ok(buffer)
}
