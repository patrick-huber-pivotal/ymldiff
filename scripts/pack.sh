#!/bin/bash

# get script directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
BINDIR="$(realpath $DIR/../bin)"

pushd $BINDIR    
    tar czvf ymldiff-windows-amd64.tgz ymldiff-windows-amd64.exe
    tar czvf ymldiff-linux-amd64.tgz ymldiff-linux-amd64
    zip -9 ymldiff-windows-amd64.zip ymldiff-windows-amd64.exe
popd