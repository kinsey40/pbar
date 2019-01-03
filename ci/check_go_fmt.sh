#!/bin/bash
# This script is used by the CI to check if the code is gofmt formatted.

set -euo pipefail

if [ -z "$(gofmt -l $( find . -type f -name '*.go' ) )" ]; then
    echo "Gofmt correct"
else
    echo "The go source files aren't formatted correctly."
    exit 1
fi

