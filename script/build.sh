#!/bin/sh
os="$1"
arch="$2"

if [ -z "$1" ]; then
  os=`go env GOHOSTOS`
fi
if [ -z "$2" ]; then
  arch=`go env GOHOSTARCH`
fi

go mod tidy
env GOOS=$os GORCH=$arch go build -ldflags="-s -w" -gcflags=all="-l -B -C" -trimpath -o "gophet-$os-$arch"
