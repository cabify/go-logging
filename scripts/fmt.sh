#!/bin/bash

set -euo pipefail

echo "Checking Go file formatting"

# shellcheck disable=SC2044
GOIMP=$(for f in $(find . -name '*.go' ! -path "./.cache/*" ! -path "./vendor/*" ! -name "bindata.go") ; do goimports -l "$f" ; done) && echo "$GOIMP" && test -z "$GOIMP"