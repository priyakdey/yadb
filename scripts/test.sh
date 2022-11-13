#!/bin/bash

PARENT_PATH=$( cd $(dirname "${BASH_SOURCE[0]}") ; pwd -P)

echo "Running all unit test cases.."
go test -v "$PARENT_PATH/../..."

