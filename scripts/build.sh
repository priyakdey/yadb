#!/bin/bash

PROJECT_DIR="$PWD"

BIN_DIR="$PROJECT_DIR/bin"
EXE_PATH="$BIN_DIR/yadb"

PACKAGE="github.com/priyakdey/yadb"

# https://stackoverflow.com/a/24112741/10368507
PARENT_PATH=$( cd "$(dirname ${BASH_SOURCE[0]})" ; pwd -P )

VERSION_FILE="$PARENT_PATH/conf/version.properties"
VERSION=`grep "${1}" "$VERSION_FILE" | cut -d "=" -f 2`


if [ -d "$BIN_DIR/" ]; then
    echo "Deleting $BIN_DIR directory.."
    rm -r "$BIN_DIR"
fi

echo "Building yadb-$VERSION.."
go build -o "$EXE_PATH" "$PACKAGE"

