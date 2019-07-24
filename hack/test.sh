#!/usr/bin/env bash

set -euo pipefail

# test go package
bazel test //...
