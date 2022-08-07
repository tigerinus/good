# Contributing

This project looks forward for your contribution in following areas.

## Coding

### To add a new command

This project uses [Cobra](https://github.com/spf13/cobra) for command implementations. To add a new command, check out [its CLI tool](https://github.com/spf13/cobra-cli).

### Conventions

Any logic that's shared across commands should be coded under `/common`.

### Static Code Analysis

This project uses [golangci-lint](https://golangci-lint.run/) and [gofumpt](https://github.com/mvdan/gofumpt).

If you use Visual Studio Code, make sure your `settings.json` includes

```json
    "go.lintFlags": [
        "-D", "staticcheck", "-E", "gocyclo,gosec,makezero,prealloc,revive"
    ],
    "go.lintTool": "golangci-lint",
    "gopls": {
        "formatting.gofumpt": true
    },
```
