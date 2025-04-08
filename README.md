# Go Repository Template

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/matt-e/otelfx.svg)](https://pkg.go.dev/github.com/matt-e/otelfx)
[![go.mod](https://img.shields.io/github/go-mod/go-version/matt-e/otelfx)](go.mod)
[![LICENSE](https://img.shields.io/github/license/matt-e/otelfx)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/matt-e/otelfx)](https://goreportcard.com/report/github.com/matt-e/otelfx)
[![Codecov](https://codecov.io/gh/matt-e/otelfx/branch/main/graph/badge.svg)](https://codecov.io/gh/matt-e/otelfx)

⭐ `Star` this repository if you find it valuable and worth maintaining.

👁 `Watch` this repository to get notified about new releases, issues, etc.

## Description

TODO


## Build

### Terminal

- `make` - execute the build pipeline.
- `make help` - print help for the [Make targets](Makefile).

### Visual Studio Code

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)` to execute the build pipeline.

## Release

The release workflow is triggered each time a tag with `v` prefix is pushed.

_CAUTION_: Make sure to understand the consequences before you bump the major version.
More info: [Go Wiki](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher),
[Go Blog](https://blog.golang.org/v2-go-modules).

## Maintenance

Notable files:

- [.github/workflows](.github/workflows) - GitHub Actions workflows,
- [.github/dependabot.yml](.github/dependabot.yml) - Dependabot configuration,
- [.vscode](.vscode) - Visual Studio Code configuration files,
- [.golangci.yml](.golangci.yml) - golangci-lint configuration,
- [Makefile](Makefile) - Make targets used for development, [CI build](.github/workflows) and [.vscode/tasks.json](.vscode/tasks.json),

## Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
