#!/usr/bin/env bash

set -euo pipefail

# build binaries for multiple platforms
# copied from https://github.com/kubernetes/kubernetes/blob/v1.15.1/hack/lib/golang.sh#L42-L53
readonly KUBE_SUPPORTED_CLIENT_PLATFORMS=(
  linux_amd64
  linux_386
  linux_arm
  linux_arm64
  linux_s390x
  linux_ppc64le
  darwin_amd64
  darwin_386
  windows_amd64
  windows_386
)

for platform in "${KUBE_SUPPORTED_CLIENT_PLATFORMS[@]}"; do
  bazel build --platforms "@io_bazel_rules_go//go/toolchain:${platform}" //:tar
  mv -f bazel-bin/tar.tgz "bazel-bin/bin-${platform}.tgz"
done
