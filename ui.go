package main

import (
	rw "github.com/mattn/go-runewidth"
	tb "github.com/nsf/termbox-go"
	"strconv"
	"time"
)

type Color struct {
	BG, FG tb.Attribute
}

type UI struct {
	FileName, Command string
	// TODO: convert to strings.Builder
	FileContent                                                            []string
	Width, Height, RulerPadding, TabSize                                   int
	InputDelay, DrawDelay                                                  time.Duration
	FileModified, Exit                                                     bool
	InfoBarColor, TextFieldColor, RulerColor, CommandBarColor, CursorColor Color
	Cursor                                                                 []int
}

func (ui UI) SplitColor(col Color) (tb.Attribute, tb.Attribute) {
	return col.FG, col.BG
}

func (ui UI) TBPrint(x, y int, text string, col Color) {
	var (
		fg, bg = ui.SplitColor(col)
		cx, cy = x, y
	)

	for _, char := range text {
		switch char {
		case '\n':
			cx, cy = x, cy+1
		case '\t':
			cx += ui.TabSize
		}
		tb.SetCell(cx, cy, char, fg, bg)
		cx += rw.RuneWidth(char)
	}
}

func (ui UI) TBPutChar(x, y int, char rune, col Color) {
	fg, bg := ui.SplitColor(col)
	tb.SetCell(x, y, char, fg, bg)
}

func (ui UI) DrawInfoBar() {
	for i := 0; i < ui.Width; i++ {
		ui.TBPutChar(i, 0, ' ', ui.InfoBarColor)
	}

	if ui.FileModified {
		ui.TBPrint(0, 0, "[*]", ui.InfoBarColor)
	}
	ui.TBPrint(4, 0, ui.FileName, ui.InfoBarColor)
}

func (ui *UI) DrawTextField() {
	ui.RulerPadding = len(strconv.Itoa(len(ui.FileContent))) + 1

	// print ruler
	for l := 1; l < len(ui.FileContent)+1; l++ {
		ui.TBPrint(0, l, strconv.Itoa(l), ui.RulerColor)
	}

	// colour background
	for x := ui.RulerPadding; x < ui.Width; x++ {
		for y := 1; y < ui.Height; y++ {
			ui.TBPutChar(x, y, ' ', ui.TextFieldColor)
		}
	}

	for y, line := range ui.FileContent {
		ui.TBPrint(ui.RulerPadding, y+1, line, ui.TextFieldColor)
	}

	x, y := ui.Cursor[0], ui.Cursor[1]
	if ui.FileContent[y] == "" {
		return
	}
	x, y = ui.RulerPadding+x, y+1
	ui.TBPutChar(x, y, tb.GetCell(x, y).Ch, ui.CursorColor)
}

func (ui UI) DrawCommandBar() {
	// clear space
	for i := 0; i < ui.Width; i++ {
		ui.TBPutChar(i, ui.Height-1, ' ', ui.CommandBarColor)
	}
	ui.TBPrint(0, ui.Height-1, ui.Command, ui.CommandBarColor)
}
