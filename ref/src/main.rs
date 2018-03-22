extern crate futures;
extern crate hyper;
extern crate ring;
extern crate untrusted;

mod core;
mod server;

use server::start_server;

fn main() {
    start_server();
}
