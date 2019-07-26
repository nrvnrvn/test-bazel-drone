#!/usr/bin/env bash

set -euo pipefail

# test testing step

# test go package
bazel test //...

# test building steps

# try build all
bazel build //...
