#!/bin/bash

set -euo pipefail

echo "Checking for Go simplifications"
gosimple ./...