#!/bin/sh
arch=`echo $GOARCH`
os=`echo $GOOS`

go build -ldflags="-s -w" -gcflags=all="-l -B -C" -trimpath -o "gophet-$os-$arch"
