package main

import (
	"bufio"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func fopen(fname string) []byte {
	f, err := os.Open(fname)
	check(err)
	defer f.Close()

	fstat, err := f.Stat()
	check(err)

	// TODO: fix buffer full
	buf := bufio.NewReader(f)
	src, err = buf.Peek(int(fstat.Size()))
	check(err)

	return src
}
