#!/usr/bin/env bash

set -euo pipefail

# such status, much volatile
echo lol woot
# ref: https://docs.drone.io/reference/environ/
echo image_tag ${DRONE_TAG:-latest}
