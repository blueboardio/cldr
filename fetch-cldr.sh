#!/usr/bin/env bash

# Usage:
#   ./fetch-cldr.sh <version>
#
# Examples:
#   ./fetch-cldr.sh 37
#   ./fetch-cldr.sh 37.0

set -eou pipefail

V=${1:-37}

if [[ "$V" != *.* ]]; then
	V=$V.0
fi

major="$(expr "x$V" : 'x\([0-9]*\)')"

f="cldr-common-$V.zip"

curl -o "$f" "http://unicode.org/Public/cldr/$major/$f"
