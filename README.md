# tonicpow-go
**tonicpow-go** is the official golang implementation for the TonicPow API

[![Build Status](https://travis-ci.org/tonicpow/tonicpow-go.svg?branch=master&v=1)](https://travis-ci.org/tonicpow/tonicpow-go)
[![Report](https://goreportcard.com/badge/github.com/tonicpow/tonicpow-go?style=flat&v=1)](https://goreportcard.com/report/github.com/tonicpow/tonicpow-go)
[![Release](https://img.shields.io/github/release-pre/tonicpow/tonicpow-go.svg?style=flat)](https://github.com/tonicpow/tonicpow-go/releases)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/tonicpow/tonicpow-go?status.svg&style=flat)](https://godoc.org/github.com/tonicpow/tonicpow-go)

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

**tonicpow-go** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy) and [dep](https://github.com/golang/dep).
```bash
$ go get -u github.com/tonicpow/tonicpow-go
```

Updating dependencies in **tonicpow-go**:
```bash
$ cd ../tonicpow-go
$ dep ensure -update -v
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/tonicpow/tonicpow-go).

### Features
- Complete coverage for the [TonicPow.com](https://tonicpow.com/) API
- Client is completely configurable
- Customize API Key and User Agent per request
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more

## Examples & Tests
All unit tests and [examples](tonicpow_test.go) run via [Travis CI](https://travis-ci.org/tonicpow/tonicpow-go) and uses [Go version 1.13.x](https://golang.org/doc/go1.13). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```bash
$ cd ../tonicpow-go
$ go test ./... -v
```

Run tests (excluding integration tests)
```bash
$ cd ../tonicpow-go
$ go test ./... -v -test.short
```

## Benchmarks
Run the Go [benchmarks](tonicpow_test.go):
```bash
$ cd ../tonicpow-go
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
	"github.com/tonicpow/tonicpow-go"
)

func main() {
    client, _ := NewClient(privateGUID)
    resp, _ = client.ConvertGoal("signup-goal", "f773c231ee9.....", 0, "")
}
```

## Maintainers

[@MrZ1836](https://github.com/mrz1836)

## Contributing

If you're looking for a JS version, checkout [this package](https://github.com/tonicpow/tonicpow-js)

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=tonicpow-go)

## License

![License](https://img.shields.io/github/license/tonicpow/tonicpow-go.svg?style=flat)
