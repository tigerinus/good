# 👍 good

[![Go Reference](https://pkg.go.dev/badge/github.com/tigerinus/good.svg)](https://pkg.go.dev/github.com/tigerinus/good) [![Go Report Card](https://goreportcard.com/badge/github.com/tigerinus/good)](https://goreportcard.com/report/github.com/tigerinus/good) [![goreleaser](https://github.com/tigerinus/good/actions/workflows/release.yml/badge.svg)](https://github.com/tigerinus/good/actions/workflows/release.yml)

A tool for installing a Go package to an isolated path, to keep the global GOPATH
clean. Because of isolated path, uninstalling is also possible.

This tool is greatly inspired by [`pipx`](https://github.com/pypa/pipx).

![a logo of gopher with thumb up without any text](logo.png)
> Thanks [DALL-e](https://openai.com/dall-e-2/) for the logo.

```text
A tool for installing a Go package to an isolated path, to keep the global GOPATH
clean. Because of isolated path, uninstalling is also possible.

Usage:
  good [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     Install a package named by its import path
  list        List all installed packages
  uninstall   Uninstall a package named by its import path
  version     Show version

Flags:
  -d, --debug   debug mode
  -h, --help    help for good

Use "good [command] --help" for more information about a command.
```

## Installation

```bash
go install github.com/tigerinus/good
```

## Contributing

This tool is still at its very early stage. Issues and Pull Requests are definitely welcome!
