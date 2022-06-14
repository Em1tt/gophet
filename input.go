package main

import (
	tb "github.com/nsf/termbox-go"
)

type Input struct {
	Ui    *UI
	Event tb.Event
}

// Parses termbox.Event bound to UI
func (i Input) GetKey() {
	switch i.Event.Type {
	case tb.EventKey:
		switch i.Event.Key {
		case tb.KeyArrowUp:
			if i.Ui.Cursor[1] > 0 {
				i.Ui.Cursor[1]--
			}
		case tb.KeyArrowDown:
			if i.Ui.Cursor[1] < len(i.Ui.FileContent)-1 {
				i.Ui.Cursor[1]++
			}
		case tb.KeyArrowLeft:
			if i.Ui.Cursor[0] > 0 {
				i.Ui.Cursor[0]--
			}
		case tb.KeyArrowRight:
			if i.Ui.Cursor[0] < len(i.Ui.FileContent[i.Ui.Cursor[1]])-1 {
				i.Ui.Cursor[0]++
			}
		case tb.KeyEsc:
			i.Ui.Exit = true
		}
	}

	// fix cursor position
	if len(i.Ui.FileContent[i.Ui.Cursor[1]]) < i.Ui.Cursor[0] {
		i.Ui.Cursor[0] = len(i.Ui.FileContent[i.Ui.Cursor[1]]) - 1
	}
	if i.Ui.Cursor[0] < 0 {
		i.Ui.Cursor[0] = 0
	}
}
