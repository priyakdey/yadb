#/bin/bash

REPORT_DIR="$PWD/reports"
COVERAGE_REPORT_PATH="$REPORT_DIR/coverage.out"

PARENT_PATH=$(cd $(dirname "${BASH_SOURCE[0]}") ; pwd -P)

go test -v --cover --covermode=count --coverprofile="$COVERAGE_REPORT_PATH" "$PARENT_PATH/../..."

echo "--------------------------------------------------"
echo "                Test coverage report              "
echo "--------------------------------------------------"

go tool cover -func="$COVERAGE_REPORT_PATH"

