extern crate futures;
extern crate hyper;

mod server;

use server::start_server;

fn main() {
    start_server();
}
