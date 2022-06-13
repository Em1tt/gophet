package main

import (
	"bufio"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func fopen(fname string) []byte {
	f, err := os.Open(fname)
	check(err)
	defer f.Close()

	fstat, err := f.Stat()
	check(err)
	size := int(fstat.Size())

	// TODO: make this better, if possible
	buf := bufio.NewReaderSize(f, size)
	src, err = buf.Peek(size)
	check(err)

	return src
}
