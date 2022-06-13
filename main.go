package main

import (
	tb "github.com/nsf/termbox-go"
	"os"
	"runtime"
	"strings"
	//	"time"
)

// all UI features are defined in ui.go
var (
	f          os.File
	fname, src string
	err        error
)

// TODO: put this into a separate file
var (
	ibc = Color{tb.ColorWhite, tb.ColorBlack}
	tfc = Color{tb.ColorBlack, tb.ColorWhite}
	rc  = Color{tb.ColorBlack, tb.ColorLightGray}
	cbc = Color{tb.ColorWhite, tb.ColorBlack}
	cc  = Color{tb.ColorWhite, tb.ColorDefault}
)

func main() {
	if len(os.Args) > 1 {
		fname, src = os.Args[1], fopen(os.Args[1])
	} else {
		src = `welcome to gophet!

  - this project is in its earliest stage
  - everything you currently see on screen is subject to change
  - right now you can't open files from this menu
  - to exit, press Ctrl + Q`
	}

	// init ui
	check(tb.Init())
	defer tb.Close()
	width, height := tb.Size()

	// TODO: fix this mess
	ui := UI{
		FileName:        fname,
		FileContent:     strings.SplitAfter(src, "\n"),
		Width:           width,
		Height:          height,
		TabSize:         4,
		InfoBarColor:    ibc,
		TextFieldColor:  tfc,
		RulerColor:      rc,
		CommandBarColor: cbc,
		CursorColor:     cc,
		Cursor:          []int{0, 0},
	}

	// init input
	input := Input{&ui}
	tb.SetInputMode(tb.InputAlt)

	runtime.GC()
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
		input.GetKey()
		// TODO: lower this when idle
		// this is probably not needed though, since input.GetKey() is blocking
		// time.Sleep(16 * time.Millisecond)
	}
}
