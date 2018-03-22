use futures::Future;
use futures::future::ok;

use hyper::{StatusCode, Error};
use hyper::server::{Http, Request, Response, Service};

use core::obtain_ticket;

struct Server;

impl Service for Server {
    type Request = Request;
    type Response = Response;
    type Error = Error;
    type Future = Box<Future<Item = Self::Response, Error = Self::Error>>;

    fn call(&self, req: Request) -> Self::Future {
        match req.path() {
            "/" => {
                return Box::new(ok(Response::new().with_body("it works")));
            }
            "/rpc/ping.action" => {
                return Box::new(ok(Response::new().with_body("pong")));
            }
            "/rpc/releaseTicket.action" => {
                return Box::new(ok(Response::new().with_body("release it")));
            }
            "/rpc/obtainTicket.action" => {
                return Box::new(ok(
                    Response::new().with_body("give you ticket"),
                ));
            }
            _ => {
                return Box::new(ok(Response::new()
                    .with_body("not found")
                    .with_status(StatusCode::NotFound)));
            }
        }
    }
}

pub fn start_server() {
    obtain_ticket("D:/General Project/jbls/tests/test.der");
    let addr = "127.0.0.1:3000".parse().unwrap();
    let server = Http::new().bind(&addr, || Ok(Server)).unwrap();
    println!("Starting server!");
    server.run().unwrap();
}
