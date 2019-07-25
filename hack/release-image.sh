#!/usr/bin/env bash

set -euo pipefail

# build and push the image
DOCKER_CONFIG=/tmp
export DOCKER_CONFIG
echo "${DOCKER_AUTH}" > "${DOCKER_CONFIG}/config.json"
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:image
