#/bin/bash

REPORT_DIR="$PWD/reports"

if [ ! -d "$REPORT_DIR" ]; then
    echo "Creating the $REPORT_DIR directory"
    mkdir "$REPORT_DIR"
fi

COVERAGE_REPORT_PATH="$REPORT_DIR/coverage.out"

PARENT_PATH=$( cd $(dirname "${BASH_SOURCE[0]}") ; pwd -P )

if [ ! -z "$GITHUB_WORKSPACE" ]; then
    echo "Running on github server. Setting package path accordingly"
    PACKAGE_PATH="$GITHUB_WORKSPACE/..."
else
    echo "Running on local. Setting package path accordingly"
    PACKAGE_PATH="$PARENT_PATH/.././..."
fi


go test -v --cover --covermode=count --coverprofile="$COVERAGE_REPORT_PATH" "$PACKAGE_PATH"

echo "--------------------------------------------------"
echo "                Test coverage report              "
echo "--------------------------------------------------"

go tool cover -func="$COVERAGE_REPORT_PATH"

