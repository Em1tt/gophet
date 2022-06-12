package main

import (
	rw "github.com/mattn/go-runewidth"
	tb "github.com/nsf/termbox-go"
	"strconv"
)

type Color struct {
	BG, FG tb.Attribute
}

type UI struct {
	FileName, FileContent, Command                                         string
	FileLines                                                              []string
	Width, Height, RulerPadding                                            int
	FileModified, Exit                                                     bool
	InfoBarColor, TextFieldColor, RulerColor, CommandBarColor, CursorColor Color
	Cursor                                                                 []int
}

func splitColor(col []Color) (tb.Attribute, tb.Attribute) {
	var (
		bg, fg tb.Attribute
	)

	bg = tb.ColorDefault
	if len(col) == 1 {
		bg, fg = col[0].BG, col[0].FG
	}
	return bg, fg
}

func tbprint(x, y int, text string, col ...Color) {
	bg, fg := splitColor(col)
	cx, cy := x, y

	for _, char := range text {
		if char == '\n' {
			cx, cy = x, cy+1
		}
		tb.SetCell(cx, cy, char, fg, bg)
		cx += rw.RuneWidth(char)
	}
}

func tbputchar(x, y int, char rune, col ...Color) {
	bg, fg := splitColor(col)
	tb.SetCell(x, y, char, fg, bg)
}

func (ui UI) DrawInfoBar() {
	for i := 0; i < ui.Width; i++ {
		tbputchar(i, 0, ' ', ui.InfoBarColor)
	}

	if ui.FileModified {
		tbprint(0, 0, "[*]", ui.InfoBarColor)
	}
	tbprint(4, 0, ui.FileName, ui.InfoBarColor)
}

func (ui *UI) DrawTextField() {
	ui.RulerPadding = len(strconv.Itoa(len(ui.FileLines))) + 1

	// print ruler
	for l := 1; l < len(ui.FileLines); l++ {
		tbprint(0, l, strconv.Itoa(l), ui.RulerColor)
	}

	// colour background
	for x := ui.RulerPadding; x < ui.Width; x++ {
		for y := 1; y < ui.Height; y++ {
			tbputchar(x, y, ' ', ui.TextFieldColor)
		}
	}

	tbprint(ui.RulerPadding, 1, ui.FileContent, ui.TextFieldColor)
	x, y := ui.Cursor[0], ui.Cursor[1]
	if ui.FileLines[y] == "" {
		return
	}
	tbputchar(ui.RulerPadding+x, y+1, rune(ui.FileLines[y][x]), ui.CursorColor)
}

func (ui UI) DrawCommandBar() {
	// clear space
	for i := 0; i < ui.Width; i++ {
		tbputchar(i, ui.Height-1, ' ', ui.CommandBarColor)
	}
	tbprint(0, ui.Height-1, ui.Command, ui.CommandBarColor)
}
