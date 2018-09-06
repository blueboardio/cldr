# cldr - Unicode CLDR data exposed as Go packages

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/blueboardio/cldr)
[![Travis-CI](https://api.travis-ci.org/blueboardio/cldr.svg?branch=master)](https://travis-ci.org/dolmen-go/jsonptr)
[![Go Report Card](https://goreportcard.com/badge/github.com/blueboardio/cldr)](https://goreportcard.com/report/github.com/blueboardio/cldr)

## Packages

* [github.com/blueboardio/cldr/country](https://godoc.org/github.com/blueboardio/cldr/country)
* [github.com/blueboardio/cldr/currency](https://godoc.org/github.com/blueboardio/cldr/currency)

Each package exposes a string-based type called `code` (`country.Code`, `currency.Code`) corresponding to an ISO code and which is the entrypoint of the API.

Data comes from the [Unicode Common Locale Data Repository](http://cldr.unicode.org/index). Code generators are bundled to update the data from the latest CLDR release.

## Status

This is business code used in production at [BlueBoard.io](https://blueboard.io).
Features will be added to the existing packages only if they are useful to our business. General use, completeness and bloat are not our aim.

## License

Copyright (c) 2018 BlueBoard SAS.

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
