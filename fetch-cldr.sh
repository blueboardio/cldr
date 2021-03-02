#!/usr/bin/env bash

# Usage:
#   ./fetch-cldr.sh <version>
#
# Examples:
#   ./fetch-cldr.sh 37
#   ./fetch-cldr.sh 37.0

set -eou pipefail

V="${1:-38.1}"
V0="$V"

if [[ "$V" != *.* ]]; then
	V0=$V.0
fi

major="$(expr "x$V" : 'x\([0-9]*\)')"

f="cldr-common-$V0.zip"
url="http://unicode.org/Public/cldr/$V/$f"
echo "$url"

curl -o "$f" "$url"
