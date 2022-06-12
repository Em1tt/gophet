arch=`echo $GOARCH`
os=`echo $GOOS`

go build -ldflags="-s -w" -trimpath -o "gophet-$os-$arch"
