#!/usr/bin/env bash

set -e
rm -f coverage.txt

for pkg in $(go list ./... | grep -v vendor); do
	rm -f .cover.out
	go test -race -coverprofile=.cover.out -covermode=atomic $pkg
	if [[ -f .cover.out ]]; then
		cat .cover.out >> coverage.txt
	fi
done
