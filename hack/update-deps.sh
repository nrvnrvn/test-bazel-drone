#!/usr/bin/env bash

set -euo pipefail

# clear dependencies
> go.mod.bzl
> go.sum

# verify and vendor all dependencies
go mod verify
go mod tidy -v
go mod vendor

# remove all files that were fetched with previous commands in order to avoid conflicts with proto rules
find vendor -iname BUILD.bazel -delete

# update repositories
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=go.mod.bzl%go_repositories

# beautify bazel files
buildifier $(find . -type f -iname BUILD.bazel -o -name WORKSPACE)

# create and update build files
bazel run //:gazelle