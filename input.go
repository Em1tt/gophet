package main

import (
	tb "github.com/nsf/termbox-go"
)

type Input struct {
	ui *UI
}

func (i Input) GetKey() {
	switch event := tb.PollEvent(); event.Type {
	case tb.EventKey:
		switch event.Key {
		case tb.KeyArrowUp:
			if i.ui.Cursor[1] > 0 {
				i.ui.Cursor[1]--
			}
		case tb.KeyArrowDown:
			if i.ui.Cursor[1] < len(i.ui.FileContent)-1 {
				i.ui.Cursor[1]++
			}
		case tb.KeyArrowLeft:
			if i.ui.Cursor[0] > 0 {
				i.ui.Cursor[0]--
			}
		case tb.KeyArrowRight:
			if i.ui.Cursor[0] < len(i.ui.FileContent[i.ui.Cursor[1]])-1 {
				i.ui.Cursor[0]++
			}
		case tb.KeyCtrlQ:
			i.ui.Exit = true
		}
	}

	// fix cursor position
	if len(i.ui.FileContent[i.ui.Cursor[1]]) < i.ui.Cursor[0] {
		i.ui.Cursor[0] = len(i.ui.FileContent[i.ui.Cursor[1]]) - 1
	}
	if i.ui.Cursor[0] < 0 {
		i.ui.Cursor[0] = 0
	}
}
