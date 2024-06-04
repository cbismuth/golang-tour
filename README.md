# Go :: Template

[![Go](https://github.com/cbismuth/golang-tour/actions/workflows/go.yaml/badge.svg)](https://github.com/cbismuth/golang-tour/actions/workflows/go.yaml)

Here is my pretty own repository to hack and learn from the Go [Tour](https://go.dev/tour).

## Unit tests

* [Welcome](src/tour_01_welcome_test.go) (see playground [here](https://go.dev/tour/welcome/1))
* [Basics](src/tour_02_basics_test.go) (see playground [here](https://go.dev/tour/basics/1))
* [Flow Control](src/tour_03_flowcontrol_test.go) (see playground [here](https://go.dev/tour/flowcontrol/1))
* [More Types](src/tour_04_moretypes_test.go) (see playground [here](https://go.dev/tour/moretypes/1))
* [Methods](src/tour_05_methods_test.go) (see playground [here](https://go.dev/tour/methods/1))
* [Generics](src/tour_06_generics_test.go) (see playground [here](https://go.dev/tour/generics/1))
* [Concurrency](src/tour_07_concurrency_test.go) (see playground [here](https://go.dev/tour/concurrency/1))

### Utilities

See [tour_utils.go](src/tour_utils.go) for miscellaneous utilities: types, random generators, floating-point, deep copying ...

## Go local setup

Here is a sample Apple macOS setup with [Homebrew](https://brew.sh):

```bash
# Install Homebrew Apple macOS package manager
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install Go utilities with Homebrew
brew install bazelisk bazel go gopls delve

# Export Go environment variables
export GOPATH="${HOME}/.go" # Default is "${HOME}/go"
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="${PATH}:${GOPATH}/bin:${GOROOT}/bin"

# Create Go source directory
mkdir -p "${GOPATH}/src"
```

## How to build

* A sample [build.sh](build.sh) Bash scripts is available to ease local build,
* GitHub action is available in [go.yaml](.github/workflows/go.yaml) to ease automated remote Continuous Integration.

In short:

* [Bazelisk](https://github.com/bazelbuild/bazelisk) is a wrapper for Bazel written in Go,
* [Bazel](https://bazel.build) is a build tool for multi-language and multi-platform projects,
* [Gazelle](https://github.com/bazelbuild/bazel-gazelle) is a Go build file generator.

## Credits

Written by Christophe Bismuth, licensed under the [MIT](LICENSE) license.
