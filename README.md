# go-tonicpow
**go-tonicpow** is the official golang implementation for interacting with the TonicPow API

[![Build Status](https://travis-ci.com/tonicpow/go-tonicpow.svg?branch=master)](https://travis-ci.com/tonicpow/go-tonicpow)
[![Report](https://goreportcard.com/badge/github.com/tonicpow/go-tonicpow?style=flat)](https://goreportcard.com/report/github.com/tonicpow/go-tonicpow)
[![Release](https://img.shields.io/github/release-pre/tonicpow/go-tonicpow.svg?style=flat)](https://github.com/tonicpow/go-tonicpow/releases)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/tonicpow/go-tonicpow?status.svg&style=flat)](https://godoc.org/github.com/tonicpow/go-tonicpow)

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Installation

**go-tonicpow** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/tonicpow/go-tonicpow
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/tonicpow/go-tonicpow).

### Features
- Client is completely configurable
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Coverage for the [TonicPow.com API](https://docs.tonicpow.com/)
    - [x] Authentication
    - [x] Users
    - [x] Advertiser Profiles
    - [x] Campaigns
    - [x] Goals
    - [ ] Links

## Examples & Tests
All unit tests and [examples](tonicpow_test.go) run via [Travis CI](https://travis-ci.org/tonicpow/go-tonicpow) and uses [Go version 1.13.x](https://golang.org/doc/go1.13). View the [deployment configuration file](.travis.yml).

View a [full example application](examples/examples.go).

Run all tests (including integration tests)
```bash
$ cd ../go-tonicpow
$ go test ./... -v
```

Run tests (excluding integration tests)
```bash
$ cd ../go-tonicpow
$ go test ./... -v -test.short
```

## Benchmarks
Run the Go [benchmarks](tonicpow_test.go):
```bash
$ cd ../go-tonicpow
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [examples & benchmarks](tonicpow_test.go)

Basic implementation:
```golang
package main

import (
	"os"
	"github.com/tonicpow/go-tonicpow"
)

func main() {
    api, _ := tonicpow.NewClient(os.Getenv("TONICPOW_API_KEY"), tonicpow.LiveEnvironment, nil)
    _ = api.ProlongSession("")
}
```

## Maintainers

[@MrZ](https://github.com/mrz1836)

## Contributing

If you're looking for a JS version, checkout [this package](https://github.com/tonicpow/tonicpow-js)

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-tonicpow)

## License

![License](https://img.shields.io/github/license/tonicpow/go-tonicpow.svg?style=flat)
