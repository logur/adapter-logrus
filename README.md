# Template Logur Adapter

[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-orange.svg?style=flat-square)](https://github.com/logur/adapter-template)
[![CircleCI](https://circleci.com/gh/logur/adapter-template.svg?style=svg)](https://circleci.com/gh/logur/adapter-template)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/adapter/template?style=flat-square)](https://goreportcard.com/report/logur.dev/adapter/template)
[![GolangCI](https://golangci.com/badges/github.com/logur/adapter-template.svg)](https://golangci.com/r/github.com/logur/adapter-template)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/adapter/template)

**Logur logger adapter for TEMPLATE.**


## Installation

```bash
go get logur.dev/adapter/template
```


## Usage

```go
package main

import (
	templateadapter "logur.dev/adapter/template"
)

func main() {
	logger := templateadapter.New(/*logger*/)
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
