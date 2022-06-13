package main

import (
	tb "github.com/nsf/termbox-go"
	"os"
	"strings"
	"time"
)

// all UI features are defined in ui.go
var (
	f          os.File
	fname, src string
	err        error
)

// TODO: put this into a separate file
var (
	ibc        = Color{tb.ColorWhite, tb.ColorBlack}
	tfc        = Color{tb.ColorBlack, tb.ColorWhite}
	rc         = Color{tb.ColorBlack, tb.ColorLightGray}
	cbc        = Color{tb.ColorWhite, tb.ColorBlack}
	cc         = Color{tb.ColorWhite, tb.ColorDefault}
	inputDelay = 16
	drawDelay  = 16
)

func main() {
	if len(os.Args) > 1 {
		fname, src = os.Args[1], fopen(os.Args[1])
	} else {
		src = `welcome to gophet!

  - this project is in its earliest stage
  - everything you currently see on screen is subject to change
  - right now you can't open files from this menu
  - to exit, press Esc`
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
		InputDelay:      time.Duration(inputDelay),
		DrawDelay:       time.Duration(drawDelay),
		InfoBarColor:    ibc,
		TextFieldColor:  tfc,
		RulerColor:      rc,
		CommandBarColor: cbc,
		CursorColor:     cc,
		Cursor:          []int{0, 0},
	}

	// init input
	input := Input{
		Ui: &ui,
	}
	tb.SetInputMode(tb.InputEsc)

	clearColFG, clearColBG := ui.SplitColor(tfc)
	go func() {
		for !ui.Exit {
			// resize
			ui.Width, ui.Height = tb.Size()
			tb.Clear(clearColFG, clearColBG)

			ui.DrawTextField()
			ui.DrawInfoBar()
			ui.DrawCommandBar()

			go tb.Flush()
			time.Sleep(ui.DrawDelay * time.Millisecond)
		}
	}()

	for !ui.Exit {
		input.Event = tb.PollEvent()
		input.GetKey()
		time.Sleep(ui.InputDelay * time.Millisecond)
	}
}
