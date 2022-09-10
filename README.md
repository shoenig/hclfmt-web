hclfmt-web
==========

The `hclfmt` command as a service.

[Online](https://sethops1.net/hclfmt)

[![GitHub](https://img.shields.io/github/license/shoenig/hclfmt-web.svg)](LICENSE)

# Project Overview

Module `github.com/shoenig/hclfmt-web` provides a web server for applying
the `hclfmt` command on input and returning the results.

# Getting Started

The `hclfmt-web` package can be installed by running
```bash
$ go install github.com/shoenig/hclfmt-web@latest
```

# Example Usage

```bash
$ hclfmt-web -config hack/example.hcl
```

# Configuration

See [example.hcl](hack/example.hcl) for an example configuration file.

# Contributing

The `github.com/shoenig/hclfmt-web` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `github.com/shoenig/hclfmt-web` module is open source under the [BSD-3-Clause](LICENSE) license.
