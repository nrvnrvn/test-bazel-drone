#!/usr/bin/env bash

set -euo pipefail

# test testing step

# test go package
bazel test //...

# test building steps

# build the binary
bazel build //:test-bazel-drone

# build the image
bazel build //:test-bazel-drone-image.tar
