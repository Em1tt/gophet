package main

import (
	"bufio"
	"os"
)

func check(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func fopen(fname string) string {
	f, err := os.Open(fname)
	check(err)
	defer f.Close()

	fstat, err := f.Stat()
	check(err)
	size := int(fstat.Size())

	// TODO: make this better, if possible
	buf := bufio.NewReaderSize(f, size)
	content, err := buf.Peek(size)
	check(err)

	return string(content)
}
