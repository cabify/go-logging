#!/bin/bash

set -euo pipefail

echo "Checking for ignored Go errors"
errcheck -ignoretests ./...