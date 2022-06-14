package main

import (
	"bufio"
	"encoding/json"
	"os"

	tb "github.com/nsf/termbox-go"
)

type color struct {
	InfoBar    [2][3]uint8
	TextField  [2][3]uint8
  Cursor     [2][3]uint8
	Ruler      [2][3]uint8
	CommandBar [2][3]uint8
}

type delay struct {
	Input int
	Draw  int
}

// Boilerplate for configuration
type Config struct {
	Color   color
	Delay   delay
  TabSize int
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

// Reads configuration from fname to config
func readConfig(fname string, config *Config) {
	src := fopen[[]byte](fname)
	check(json.Unmarshal(src, &config))
}

// Converts RGB color used in config.json to Color for use with termbox
func RGBToTB (col [2][3]uint8) Color {
  result := make([]tb.Attribute, 2)
  for i, layer := range col {
    result[i] = tb.RGBToAttribute(layer[0], layer[1], layer[2])
  }
  return Color{result[0], result[1]}
}
