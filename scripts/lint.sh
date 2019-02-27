#!/bin/bash

set -euo pipefail

echo "Checking for style errors"
golint ./...