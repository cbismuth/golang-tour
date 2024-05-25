#!/usr/bin/env bash

set -e

export LC_ALL="en_US.UTF-8"
export LANG="en_US.UTF-8"

bazelisk test //... --test_output=all
