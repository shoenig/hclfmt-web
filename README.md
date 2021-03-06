hclfmt-web
==========

The `hclfmt` command as a service.

[Online](https://sethops1.net/hclfmt)

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/cmds/hclfmt-web)](https://goreportcard.com/report/gophers.dev/cmds/hclfmt-web)
[![Build Status](https://travis-ci.com/shoenig/hclfmt-web.svg?branch=master)](https://travis-ci.com/shoenig/hclfmt-web)
[![GoDoc](https://godoc.org/gophers.dev/cmds/hclfmt-web?status.svg)](https://godoc.org/gophers.dev/cmds/hclfmt-web)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/hclfmt-web.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/hclfmt-web.svg)](LICENSE)

# Project Overview

Module `gophers.dev/cmds/hclfmt-web` provides a web server for applying
the `hclfmt` command on input and returning the results.

# Getting Started

The `hclfmt-web` package can be installed by running
```bash
$ go install gophers.dev/cmds/hclfmt-web@latest
```

# Example Usage

```bash
$ hclfmt-web -config hack/example.hcl
```

# Configuration

See [example.hcl](hack/example.hcl) for an example configuration file.

# Contributing

The `gophers.dev/cmds/hclfmt-web` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `gophers.dev/cmds/hclfmt-web` module is open source under the [BSD-3-Clause](LICENSE) license.
