package main

import (
	"bufio"
	"os"
)

// Simple error checking to prettify code
func check(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// Opens a file with name fname and returns its content. Uses bufio
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
