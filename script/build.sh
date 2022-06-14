#!/bin/sh
os="$1"
arch="$2"

# if no arguments specified, use host defaults
[ -z "$1" ] && os=`go env GOHOSTOS`
[ -z "$2" ] && arch=`go env GOHOSTARCH`

# construct filename
fname="gophet-$os-$arch"
[ $os = "windows" ] && fname="$fname.exe"

# make needed directories if they don't exist
[ ! -d build ] && mkdir build

# verify everything is in place
go mod vendor
go mod verify
go mod tidy

# build with needed optimization
env GOOS=$os GORCH=$arch go build -ldflags="-s -w" -gcflags=all="-l -B -C" -trimpath -o "build/$fname"
