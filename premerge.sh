#!/usr/bin/env bash

set -e

echo
echo "Verifying: gofmt"
unformatted=$(gofmt -l ./pkg/.)
if [ -z "$unformatted" ]
then
  echo "gofmt successful"
else
  echo >&2 "Go files must be formatted with gofmt. Please run:"
  for fn in $unformatted; do
    echo >&2 "  gofmt -w $PWD/$fn"
  done
  exit 1
fi

## run tests with coverage
echo
echo "Running: tests"
for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f test_result.out ]; then
        cat test_result.out >> coverage.txt
        rm test_result.out
    fi
done