package main

import (
	"bufio"
	tb "github.com/nsf/termbox-go"
	"os"
	"strings"
	//	"time"
)

// all UI features are defined in ui.go
func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

var (
	f     os.File
	fname string
	src   []byte
	err   error
)

func main() {
	if len(os.Args) > 1 {
		fname = os.Args[1]
		// boring file reading stuff
		f, err := os.Open(fname)
		check(err)
		defer f.Close()

		fstat, err := f.Stat()
		check(err)

		// TODO: possibly optimize this for bigger files
		buf := bufio.NewReader(f)
		src, err = buf.Peek(int(fstat.Size()))
		check(err)
	} else {
		src = []byte(`
welcome to gophet!

  - this project is in its earliest stage
  - report any bugs directly to me
  - everything you currently see on screen is subject to change
  - right now you can't open files from this menu
  - to open a file, run this from your command line:
      on windows:   gophet.exe [filename] 
      on linux/mac: ./gophet [filename]
  `)
	}

	// init ui
	check(tb.Init())
	defer tb.Close()
	width, height := tb.Size()

	// TODO: fix this mess
	ui := UI{
		FileName:        fname,
		FileContent:     string(src),
		Command:         "",
		FileLines:       strings.Split(string(src), "\n"),
		Width:           width,
		Height:          height,
		FileModified:    false,
		Exit:            false,
		InfoBarColor:    Color{tb.ColorWhite, tb.ColorBlack},
		TextFieldColor:  Color{tb.ColorBlack, tb.ColorWhite},
		RulerColor:      Color{tb.ColorBlack, tb.ColorLightGray},
		CommandBarColor: Color{tb.ColorWhite, tb.ColorBlack},
		CursorColor:     Color{tb.ColorWhite, tb.ColorDefault},
		Cursor:          []int{0, 0},
	}

	// init input
	input := Input{}
	tb.SetInputMode(tb.InputAlt)

	for {
		if ui.Exit {
			break
		}
		// resize
		ui.Width, ui.Height = tb.Size()
		err = tb.Clear(tb.ColorDefault, tb.ColorDefault)

		check(err)
		ui.DrawTextField()
		ui.DrawInfoBar()
		ui.DrawCommandBar()

		tb.Flush()
		input.GetKey(&ui)
		// TODO: lower this when idle
		// this is probably not needed though, since input.GetKey() is blocking
		// time.Sleep(16 * time.Millisecond)
	}
}
