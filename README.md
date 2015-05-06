# ok
Toolkit for high productivity web development in Go language (WIP).

[![GoDoc](https://godoc.org/github.com/anonx/ok?status.svg)](https://godoc.org/github.com/anonx/ok)
[![Build Status](https://travis-ci.org/anonx/ok.svg?branch=master)](https://travis-ci.org/anonx/ok)
[![Coverage Status](https://coveralls.io/repos/anonx/ok/badge.svg)](https://coveralls.io/r/anonx/ok)
[![Go Report Card](http://goreportcard.com/badge/anonx/ok?t=3)](http:/goreportcard.com/report/anonx/ok)

## Commands
* `ok run path/to/app` - Run a tool that watches the app files and rebuilds if necessary. It can be used as a task runner, too.
* `ok generate handler -i path/to/app/controllers -o path/to/app/result/handlers` - Util that is used by `go:generate` to transform controllers into `handlers` package. Read more about [controllers](https://github.com/anonx/concept/blob/master/basics.md#basics).

## License
Distributed under the BSD 2-clause "Simplified" License.