#!/bin/bash

set -euo pipefail

echo "Checking shell scripts for errors"
find . -type f -name "*.sh" ! -path "./vendor/*" ! -path "./.cache/*" -print0 | xargs -0 shellcheck