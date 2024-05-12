# Go :: Template

[![Go](https://github.com/cbismuth/golang-tour/actions/workflows/go.yaml/badge.svg)](https://github.com/cbismuth/golang-tour/actions/workflows/go.yaml)

Here is my pretty own template to hack and learn on the Go programming language.

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
