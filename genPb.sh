#!/usr/bin/env sh
set -eu

goBin="$(go env GOPATH)/bin"
PATH="$goBin:$PATH"
export PATH

cd "$(dirname "$0")/proto"
exec sh ./protoc.sh
