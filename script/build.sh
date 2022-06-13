#!/bin/sh
os="$1"
arch="$2"

# if no arguments specified, use host defaults
[ -z "$1" ] && os=`go env GOHOSTOS`
[ -z "$2" ] && arch=`go env GOHOSTARCH`

# mahe needed directories if they don't exist
[ ! -d build ] && mkdir build

# build!
go mod tidy
env GOOS=$os GORCH=$arch go build -ldflags="-s -w" -gcflags=all="-l -B -C" -trimpath -o "build/gophet-$os-$arch"
