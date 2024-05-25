#!/usr/bin/env bash

set -e

export LC_ALL="en_US.UTF-8"
export LANG="en_US.UTF-8"

export CYAN="\033[0;36m"

clear() {
    printf '\33c\e[3J'
}

log() {
    echo -e "${CYAN}${1}${NC}"
}

clear

log "Deleting generated BUILD.bazel build files ..."
find . -type f -name "BUILD.bazel" -mindepth 2 -delete

log "Creating deps.bzl file from go.mod files ..."
bazelisk run //:gazelle -- update-repos -from_file="go.mod" -to_macro="deps.bzl%go_dependencies"

log "Updating BUILD.bazel files with fix option ..."
bazelisk run //:gazelle -- fix

log "Formatting and lint Bazel files ..."
bazelisk run //:buildifier

log "Installing goimports utility ..."
go install golang.org/x/tools/cmd/goimports@latest && which goimports

log "Formatting Go files ..."
goimports -l -w .

log "Building all targets ..."
bazelisk build --disk_cache="${HOME}/.cache/bazel" //...

log "Testing all targets ..."
bazelisk test //... --test_output=all

log "Querying all targets ..."
bazelisk query //...

log "Running main target ..."
./bazel-bin/src/src_/src
