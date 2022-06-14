package main

import (
	"bufio"
	"encoding/json"
	"os"
)

// Boilerplate for configuration
type Config struct {
	Colorscheme map[string]struct {
		InfoBar    Color `json:"infobar"`
		TextField  Color `json:"textfield"`
		Ruler      Color `json:"ruler"`
		CommandBar Color `json:"commandbar"`
	} `json:"colorscheme"`

	Delay map[string]struct {
		Input int `json:"input"`
		Draw  int `json:"draw"`
	} `json:"delay"`
}

// Simple error checking to prettify code
func check(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// Opens a file with name fname and returns its content. Uses bufio
func fopen[T []byte | string](fname string) T {
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

	return T(content)
}

func readConfig(fname string) Config {
  src, config := fopen[[]byte](fname), Config{}
  check(json.Unmarshal(src, &config))
  return config
}
