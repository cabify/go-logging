#!/bin/bash

set -euo pipefail

echo "Fixing Go formatting"
# shellcheck disable=SC2044
for f in $(find . -name '*.go' ! -path "./vendor/*" ! -path "./.cache/*") ; do
   	if [ -n "$(goimports -d "$f")" ]; then
		echo "Fixing $f";
		goimports -w "$f";
	fi;
done
