# 👍 good

[![Go Reference](https://pkg.go.dev/badge/github.com/tigerinus/good.svg)](https://pkg.go.dev/github.com/tigerinus/good) [![Go Report Card](https://goreportcard.com/badge/github.com/tigerinus/good)](https://goreportcard.com/report/github.com/tigerinus/good) [![goreleaser](https://github.com/tigerinus/good/actions/workflows/release.yml/badge.svg)](https://github.com/tigerinus/good/actions/workflows/release.yml)
[![codecov](https://codecov.io/gh/tigerinus/good/branch/main/graph/badge.svg?token=PX4PGVR3QC)](https://codecov.io/gh/tigerinus/good)

A tool for installing a Go package to an isolated path, to keep the global `GOPATH/pkg`
clean. Because of isolated path, uninstalling is also possible.

This tool is greatly inspired by [`pipx`](https://github.com/pypa/pipx).

![a logo of gopher with thumb up without any text](logo.png)
> Thanks [DALL-e](https://openai.com/dall-e-2/) for the logo.

## Motivations

- Command `go clean -i -r ...` [does not really clean everything](https://www.reddit.com/r/golang/comments/pzeunz/proper_package_management_commands_in_go_117/).
- Sometime people just want to use `go install` to install a CLI for non-development purpose, but doing it under a Go project, it changes your `go.mod` and `go.sum` files.
- People could do `ls $GOPATH/bin` to see all the binaries but have to go thru extra steps to figure out the package name before uninstalling it.
- Simply removing the binary under `$GOPATH/bin` leaves a bunch of dependencies in `$GOPATH/pkg/mod`.

## Features

- Install CLI Go apps in an isolated path, without changing anything under `$GOPATH` or `go.mod`/`go.sum` files (if under a Go project).
- Uninstall the app in one command without leaving any files behind.
- List all CLI go apps installed
- *(UPCOMING)* Check for outdated CLI go apps among the installed.
- *(UPCOMING)* Strip the binary after as part of installation.
- *(UPCOMING)* Show information about a CLI go app.
- *(UPCOMING)* Run a CLI app once and uninstall.
- *(UPCOMING)* Search for CLI go apps from sources like [awesome-go](https://awesome-go.com/).

## Usage

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
