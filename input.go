package main

import (
	tb "github.com/nsf/termbox-go"
)

type Input struct{}

func (i Input) GetKey(ui *UI) {
	switch event := tb.PollEvent(); event.Type {
	case tb.EventKey:
		switch event.Key {
		case tb.KeyArrowUp:
			if ui.Cursor[1] > 0 {
				ui.Cursor[1]--
			}
		case tb.KeyArrowDown:
			if ui.Cursor[1] < len(ui.FileLines)-1 {
				ui.Cursor[1]++
			}
		case tb.KeyArrowLeft:
			if ui.Cursor[0] > 0 {
				ui.Cursor[0]--
			}
		case tb.KeyArrowRight:
			if ui.Cursor[0] < len(ui.FileLines[ui.Cursor[1]])-1 {
				ui.Cursor[0]++
			}
		}
	}

	// fix cursor position
	if len(ui.FileLines[ui.Cursor[1]]) < ui.Cursor[0] {
		ui.Cursor[0] = len(ui.FileLines[ui.Cursor[1]]) - 1
	}
	if ui.Cursor[0] < 0 {
		ui.Cursor[0] = 0
	}
}
