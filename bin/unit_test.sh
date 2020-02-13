#!/usr/bin/env bash

echo "Running tests in $PWD"
echo ""
go clean -testcache
go test ./...
