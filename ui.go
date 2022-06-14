package main

import (
	rw "github.com/mattn/go-runewidth"
	tb "github.com/nsf/termbox-go"
	"strconv"
	"time"
)

// Stores background and foreground as termbox.Attribute
type Color struct {
	BG, FG tb.Attribute
}

// TODO: split this
type UI struct {
	FileName, Command string
	// TODO: convert to strings.Builder
	FileContent                                                            []string
	Width, Height, RulerPadding, TabSize                                   int
	InputDelay, DrawDelay                                                  time.Duration
	FileModified, Exit                                                     bool
	InfoBarColor, TextFieldColor, RulerColor, CommandBarColor, CursorColor Color
	Cursor                                                                 [2]int
}

// Prints char to the console at (x, y)
func (ui UI) PutChar(x, y int, char rune, col Color) {
	tb.SetCell(x, y, char, col.FG, col.BG)
}

// Prints text to the console at (x, y), obeying tabs and newlines
func (ui UI) Print(x, y int, text string, col Color) {
	cx, cy := x, y
	for _, char := range text {
		switch char {
		case '\n':
			cx, cy = x, cy+1
		case '\t':
			cx += ui.TabSize
		}
		ui.PutChar(cx, cy, char, col)
		cx += rw.RuneWidth(char)
	}
}

// Draws the info bar with UI.FileModified and UI.FileName
func (ui UI) DrawInfoBar() {
	for i := 0; i < ui.Width; i++ {
		ui.PutChar(i, 0, ' ', ui.InfoBarColor)
	}

	if ui.FileModified {
		ui.Print(0, 0, "[*]", ui.InfoBarColor)
	}
	ui.Print(4, 0, ui.FileName, ui.InfoBarColor)
}

// Draws the text field with ruler and UI.FileContent
func (ui *UI) DrawTextField() {
	ui.RulerPadding = len(strconv.Itoa(len(ui.FileContent))) + 1

	// print ruler
	for l := 1; l < len(ui.FileContent)+1; l++ {
		ui.Print(0, l, strconv.Itoa(l), ui.RulerColor)
	}

	// colour background
	for x := ui.RulerPadding; x < ui.Width; x++ {
		for y := 1; y < ui.Height; y++ {
			ui.PutChar(x, y, ' ', ui.TextFieldColor)
		}
	}

	for y, line := range ui.FileContent {
		ui.Print(ui.RulerPadding, y+1, line, ui.TextFieldColor)
	}

	x, y := ui.Cursor[0], ui.Cursor[1]
	if ui.FileContent[y] == "" {
		return
	}
	x, y = ui.RulerPadding+x, y+1
	ui.PutChar(x, y, tb.GetCell(x, y).Ch, ui.CursorColor)
}

// Draws the command bar with UI.Command
func (ui UI) DrawCommandBar() {
	// clear space
	for i := 0; i < ui.Width; i++ {
		ui.PutChar(i, ui.Height-1, ' ', ui.CommandBarColor)
	}
	ui.Print(0, ui.Height-1, ui.Command, ui.CommandBarColor)
}
