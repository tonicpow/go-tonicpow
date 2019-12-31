<img src="https://repository-images.githubusercontent.com/215606155/bf4a1180-2b4d-11ea-96a0-3b4020e72e35" alt="TonicPow & Go">

**go-tonicpow** is the official golang implementation for interacting with the [TonicPow API](https://docs.tonicpow.com)

[![Go](https://img.shields.io/github/go-mod/go-version/tonicpow/go-tonicpow)](https://golang.org/)
[![Build Status](https://travis-ci.com/tonicpow/go-tonicpow.svg?branch=master)](https://travis-ci.com/tonicpow/go-tonicpow)
[![Report](https://goreportcard.com/badge/github.com/tonicpow/go-tonicpow?style=flat)](https://goreportcard.com/report/github.com/tonicpow/go-tonicpow)
[![Release](https://img.shields.io/github/release-pre/tonicpow/go-tonicpow.svg?style=flat)](https://github.com/tonicpow/go-tonicpow/releases)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme)
[![Slack](https://img.shields.io/badge/slack-tonicpow-orange.svg?style=flat)](https://atlantistic.slack.com/app_redirect?channel=tonicpow)
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
- [Client](client.go) is completely configurable
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Coverage for the [TonicPow.com API](https://docs.tonicpow.com/)
    - [x] [Authentication](https://docs.tonicpow.com/#632ed94a-3afd-4323-af91-bdf307a399d2)
    - [x] [Users](https://docs.tonicpow.com/#50b3c130-7254-4a05-b312-b14647736e38)
    - [x] [Advertiser Profiles](https://docs.tonicpow.com/#2f9ec542-0f88-4671-b47c-d0ee390af5ea)
    - [x] [Campaigns](https://docs.tonicpow.com/#5aca2fc7-b3c8-445b-aa88-f62a681f8e0c)
    - [x] [Goals](https://docs.tonicpow.com/#316b77ab-4900-4f3d-96a7-e67c00af10ca)
    - [x] [Links](https://docs.tonicpow.com/#ee74c3ce-b4df-4d57-abf2-ccf3a80e4e1e)

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
    _ = api.ConvertGoal("new-lead-goal", "s358wef983283...", "", "")
}
```

## Maintainers

[@MrZ](https://github.com/mrz1836)

## Contributing

If you're looking for a library using JavaScript, checkout [tonicpow-js](https://github.com/tonicpow/tonicpow-js)

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://tonicpow.com/?af=go-tonicpow)

## License

[![license](https://img.shields.io/badge/license-Open%20BSV-brightgreen.svg?style=flat)](/LICENSE)
