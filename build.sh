#!/bin/sh
arch=`echo $GOARCH`
os=`echo $GOOS`

go mod tidy
go build -ldflags="-s -w" -gcflags=all="-l -B -C" -trimpath -o "gophet-$os-$arch"
