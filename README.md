# cldr - Unicode CLDR data exposed as Go packages

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/blueboardio/cldr/v2?tab=subdirectories)
[![Travis-CI](https://api.travis-ci.org/blueboardio/cldr.svg?branch=master)](https://travis-ci.org/blueboardio/cldr)
[![Go Report Card](https://goreportcard.com/badge/github.com/blueboardio/cldr)](https://goreportcard.com/report/github.com/blueboardio/cldr)

Author: [@dolmen](https://github.com/dolmen)  (Olivier Mengué).

## Packages

* [github.com/blueboardio/cldr/v2/country](https://pkg.go.dev/github.com/blueboardio/cldr/v2/country?tab=doc)
* [github.com/blueboardio/cldr/v2/currency](https://pkg.go.dev/github.com/blueboardio/cldr/v2/currency?tab=doc)

Each package exposes a string-based type called `Code` (`country.Code`, `currency.Code`) corresponding to an ISO code and which is the entrypoint of the API. Those types implement interfaces used for input and output from/to
JSON ([json.Marshaler](https://golang.org/pkg/encoding/json/#Marshaler), [json.Unmarshaler](https://golang.org/pkg/encoding/json/#Unmarshaler)),
command line flags ([flag.Value](https://golang.org/pkg/flag/#Value)) and SQL databases
([sql.Scanner](https://golang.org/pkg/database/sql/#Scanner), [driver.Valuer](https://golang.org/pkg/database/sql/driver/#Valuer)).


Data (country names, currencies per country) comes from the [Unicode Common Locale Data Repository](http://cldr.unicode.org/index). Code generators are bundled to update the data from the latest CLDR release.

## Status

This is business code used in production at [BlueBoard.io](https://blueboard.io).

## License

Copyright © 2018-2019 BlueBoard SAS.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
