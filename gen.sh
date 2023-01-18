#!/usr/bin/env bash

set -euo pipefail

if [[ "v${1:-}" != v[0-9][0-9]* ]]; then
	echo "usage: $0 <cldr-version>" >&2
	exit 1
fi

V="${1}"
V0="$V"

if [[ "v$V" = v[0-9][0-9] ]]; then
	V0=$V.0
elif [[ "v$V" = v[0-9][0-9].0 ]]; then
	V=${V%%.*}
fi

if [[ ! -f "cldr-common-$V0.zip" ]]; then
	./fetch-cldr.sh "$V"
fi

echo '--[ currency ]--'
cd currency
# go run github.com/blueboardio/cldr/internal/gen-currencies ../cldr-common-$V0.zip
go run -modfile ../gen/go.mod ../gen/gen-currencies/main.go "../cldr-common-$V0.zip"

echo '--[ country ]--'
cd ../country
go run -modfile ../gen/go.mod ../gen/gen-countries/main.go "../cldr-common-$V0.zip"


