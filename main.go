package main

import (
	"os"
	"runtime"
	"strings"
	"time"

	tb "github.com/nsf/termbox-go"
)

// all UI features are defined in ui.go
var (
	f          os.File
	fname, src string
	err        error
)

// TODO: put this into a separate file
func main() {
	if len(os.Args) > 1 {
		fname, src = os.Args[1], fopen[string](os.Args[1])
	} else {
		src = `welcome to gophet!

  - this project is in its earliest stage
  - everything you currently see on screen is subject to change
  - right now you can't open files from this menu
  - to exit, press Esc`
	}

	// TODO: create some default config instead
	cfg := Config{
		color{
			[2][3]uint8{{0, 0, 0}, {255, 255, 255}},
			[2][3]uint8{{0, 0, 0}, {255, 255, 255}},
			[2][3]uint8{{255, 255, 255}, {0, 0, 0}},
			[2][3]uint8{{0, 0, 0}, {155, 155, 155}},
			[2][3]uint8{{0, 0, 0}, {255, 255, 255}},
		},
		delay{16, 16},
		4,
	}

	if _, err := os.Stat("config.json"); !os.IsNotExist(err) {
		readConfig("config.json", &cfg)
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
		TabSize:         cfg.TabSize,
		InputDelay:      time.Duration(cfg.Delay.Input),
		DrawDelay:       time.Duration(cfg.Delay.Draw),
		InfoBarColor:    RGBToTB(cfg.Color.InfoBar),
		TextFieldColor:  RGBToTB(cfg.Color.TextField),
		RulerColor:      RGBToTB(cfg.Color.Ruler),
		CommandBarColor: RGBToTB(cfg.Color.CommandBar),
		CursorColor:     RGBToTB(cfg.Color.Cursor),
		Cursor:          [2]int{0, 0},
	}
	tb.SetOutputMode(tb.OutputRGB)

	// init input
	input := Input{
		Ui: &ui,
	}
	tb.SetInputMode(tb.InputEsc)

	runtime.GC()
	go func() {
		for !ui.Exit {
			// resize
			ui.Width, ui.Height = tb.Size()
			tb.Clear(ui.TextFieldColor.FG, ui.TextFieldColor.BG)

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
