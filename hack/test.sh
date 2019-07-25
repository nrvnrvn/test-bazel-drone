#!/usr/bin/env bash

set -euo pipefail

# test testing step

# test go package
bazel test //...
