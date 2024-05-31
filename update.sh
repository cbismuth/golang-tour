#!/usr/bin/env bash

set -e

export LC_ALL="en_US.UTF-8"
export LANG="en_US.UTF-8"

find . -type f -name "BUILD.bazel" -mindepth 2 -delete

bazelisk run //:gazelle -- update-repos -from_file="go.mod" -to_macro="deps.bzl%go_dependencies"
bazelisk run //:gazelle -- fix
bazelisk run //:buildifier
