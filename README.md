# docker-ref

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sagikazarmark/docker-ref/ci.yaml?style=flat-square)](https://github.com/sagikazarmark/docker-ref/actions/workflows/ci.yaml)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/sagikazarmark/docker-ref)
[![built with nix](https://img.shields.io/badge/builtwith-nix-7d81f7?style=flat-square)](https://builtwithnix.org)

**A library for parsing Docker image references.**

This library is extracted from the [distribution](https://github.com/distribution/distribution) project (from [this](https://github.com/distribution/distribution/tree/7b502560cad43970472964166dcb095b1f883ae4/reference) commit).

It's sole purpose is to allow parsing Docker image references without having to import the entire distribution project with all of its dependencies.

## Installation

```shell
go get github.com/sagikazarmark/docker-ref
```

## Development

**For an optimal developer experience, it is recommended to install [Nix](https://nixos.org/download.html) and [direnv](https://direnv.net/docs/installation.html).**

Run the test suite:

```shell
just test
```

## Prior art

There are some existing solutions out there, but I decided it's easier if I just fork the repo myself.

- https://github.com/novln/docker-parser
- https://github.com/containers/image/tree/8c387a14f4ac95fb4aa0569e8e31439f988bc941/docker/reference

## License

The project is licensed under the [Apache 2.0 License](LICENSE).
