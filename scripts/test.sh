#!/usr/bin/env bash

set -e

mkdir -p coverage
go test -coverprofile coverage/cover.out -v $(go list ./...)
go tool cover -html=coverage/cover.out -o coverage/index.html
