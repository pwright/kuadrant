# Limitador

[![Limitador GH Workflow](https://github.com/Kuadrant/limitador/actions/workflows/rust.yml/badge.svg)](https://github.com/Kuadrant/limitador/actions/workflows/rust.yml)
[![docs.rs](https://docs.rs/limitador/badge.svg)](https://docs.rs/limitador)
[![Crates.io](https://img.shields.io/crates/v/limitador)](https://crates.io/crates/limitador)
[![Docker Repository on Quay](https://quay.io/repository/3scale/limitador/status
"Docker Repository on Quay")](https://quay.io/repository/3scale/limitador)

Limitador is a generic rate-limiter written in Rust. It can be used as a
library, or as a service. The service exposes HTTP endpoints to apply and manage
limits. Limitador can be used with Envoy because it also exposes a grpc service, on a different
port, that implements the Envoy Rate Limit protocol (v3).

- [**Getting started**](#getting-started)
- [**Limits storage**](#limits-storage)
- [**Development**](#development)
- [**Kubernetes**](limitador-server/kubernetes/)
- [**License**](#license)

Limitador is under active development, and its API has not been stabilized yet.

## Getting started

- [Rust library](#rust-library)
- [Server](#server)

### Rust library

Add this to your `Cargo.toml`:
```toml
[dependencies]
limitador = { version = "0.2.0" }
```

To use Limitador in a project that compiles to WASM, there are some features
that need to be disabled. Add this to your `Cargo.toml` instead:
```toml
[dependencies]
limitador = { version = "0.2.0", default-features = false }
```

### Server

Run with Docker (replace `latest` with the version you want):
```bash
docker run --rm --net=host -it quay.io/3scale/limitador:latest
```

Run locally:
```bash
cargo run --release --bin limitador-server
```

To use Redis, specify the URL with `REDIS_URL`:
```bash
REDIS_URL=redis://127.0.0.1:6379 cargo run --release --bin limitador-server
```

By default, Limitador starts the HTTP server in `localhost:8080` and the grpc
service that implements the Envoy Rate Limit protocol in `localhost:8081`. That
can be configured with these ENVs: `ENVOY_RLS_HOST`, `ENVOY_RLS_PORT`,
`HTTP_API_HOST`, and `HTTP_API_PORT`.

The OpenAPI spec of the HTTP service is
[here](limitador-server/docs/http_server_spec.json).

Limitador can be started with a YAML file that has some limits predefined. Keep
in mind that they can be modified using the HTTP API. There's an [example
file](limitador-server/examples/limits.yaml) that allows 10 requests per minute
and per user_id when the HTTP method is "GET" and 5 when it is a "POST". You can
run it with Docker (replace `latest` with the version you want):
```bash
docker run -e LIMITS_FILE=/home/limitador/my_limits.yaml --rm --net=host -it -v $(pwd)/limitador-server/examples/limits.yaml:/home/limitador/my_limits.yaml:ro quay.io/3scale/limitador:latest
```

You can also use the YAML file when running locally:
```bash
LIMITS_FILE=./limitador-server/examples/limits.yaml cargo run --release --bin limitador-server
```

If you want to use Limitador with Envoy, there's a minimal Envoy config for
testing purposes [here](limitador-server/examples/envoy.yaml). The config
forwards the "userid" header and the request method to Limitador. It assumes
that there's an upstream API deployed on port 1323. You can use
[echo](https://github.com/labstack/echo), for example.

Limitador has several options that can be configured via ENV. This
[doc](limitador-server/docs/configuration.md) specifies them.

## Limits storage

Limitador can store its limits and counters in-memory or in Redis. In-memory is
faster, but the limits are applied per instance. When using Redis, multiple
instances of Limitador can share the same limits, but it's slower.


## Development

### Build

```bash
cargo build
```

### Run the tests

Some tests need a redis deployed in `localhost:6379`. You can run it in Docker with:
```bash
docker run --rm -p 6379:6379 -it redis
```

Then, run the tests:

```bash
cargo test
```

or you can run tests disabling the "redis storage" feature:
```bash
cd limitador; cargo test --no-default-features
```

Limitador also offers experimental support for Infinispan as a storage for the
limits and counters. It's under a feature not enabled by default. To build with it and test it:
```bash
cargo build --features=infinispan_storage
cargo test --features=infinispan_storage
```

## License

[Apache 2.0 License](LICENSE)
