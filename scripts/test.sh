#!/bin/bash

set -euo pipefail

echo "Running unit tests"
go test -race ./...