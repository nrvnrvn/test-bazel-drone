#!/usr/bin/env bash

set -euo pipefail

# such status, much volatile
echo lol woot
# ref: https://docs.drone.io/reference/environ/
echo image_tag "${DRONE_TAG:-latest}"

# Build time variable to embed.
# Ref: https://github.com/bazelbuild/rules_go/blob/master/go/core.rst#defines-and-stamping
echo version "${DRONE_TAG:-latest}"
echo gitCommit "${DRONE_COMMIT_SHA:-$(git rev-parse HEAD)}"
echo buildDate "$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
