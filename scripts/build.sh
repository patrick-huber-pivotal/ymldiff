#!/bin/bash

# get script directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
BINDIR="$(realpath $DIR/../bin)"

# clean the binaries
rm -rf $BINDIR/*

# build windows
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=0

go build -o "$BINDIR/ymldiff-windows-amd64.exe" "$DIR/../main.go"

# build linux
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

go build -o "$BINDIR/ymldiff-linux-amd64" "$DIR/../main.go"