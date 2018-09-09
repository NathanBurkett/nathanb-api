#!/usr/bin/env bash
# Generate test coverage statistics for Go packages.
#
# Works around the fact that `go test -coverprofile` currently does not work
# with multiple packages, see https://code.google.com/p/go/issues/detail?id=6909
#
# Usage: test-coverage/run [--html]
#
#     --html      Additionally create HTML report and open it in browser
#

set -e

_workdir=test-coverage/output
_profile="$_workdir/cover.out"
_mode=count

generate_cover_data() {
    rm -rf "$_workdir"
    mkdir "$_workdir"

    for pkg in "$@"; do
        f="$_workdir/$(echo $pkg | tr / -).cover"
        go test -v -covermode="$_mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $_mode" >"$_profile"
    grep -h -v "^mode:" "$_workdir"/*.cover >>"$_profile"
}

show_cover_report() {
    go tool cover -${1}="$_profile"
}

generate_cover_data $(go list ./...)
show_cover_report func
case "$1" in
"")
    ;;
--html)
    show_cover_report html ;;
*)
    echo >&2 "error: invalid option: $1"; exit 1 ;;
esac
