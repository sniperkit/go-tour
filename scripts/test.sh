#!/usr/bin/env bash
# The coverage tool cannot handle multiple packages see (https://github.com/golang/go/issues/6909)
# This script is a workaround that runs coverage for each package and then unifies it

set -e

WORKDIR=coverage
PROFILE="$WORKDIR/cover.out"
MODE=count

generate_cover_data() {
    rm -rf "$WORKDIR"
    mkdir -p "$WORKDIR"

    for pkg in "$@"; do
        f="$WORKDIR/$(echo $pkg | tr / -).cover"
        go test -covermode="$MODE" -coverprofile="$f" "$pkg"
    done

    echo "mode: $MODE" >"$PROFILE"
    grep -h -v "^mode:" "$WORKDIR"/*.cover >>"$PROFILE"
}

run_tests() {
    go test -v $(go list ./...)
}

show_cover_report() {
    go tool cover -html="$PROFILE" -o "$WORKDIR"/index.html
}

run_tests
generate_cover_data $(go list ./...)
show_cover_report