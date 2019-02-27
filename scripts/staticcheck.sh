#!/bin/bash

set -euo pipefail

echo "Running static checks"
staticcheck ./...