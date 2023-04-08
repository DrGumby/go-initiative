package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

type TextInput struct {
	value  string
	cursor int
}

func (t *TextInput) Draw(s tcell.Screen, startx int, starty int) {
	before := t.value[:t.cursor]
	at := t.value[t.cursor]
	after := t.value[t.cursor+1:]

	emitStr(s, startx, starty, defStyle, before)
	emitStr(s, startx+len(before), starty, highligtedStyle, string(at))
	emitStr(s, startx+len(before)+1, starty, defStyle, after)
}

func (t *TextInput) MoveLeft() {
	if t.cursor > 0 {
		t.cursor -= 1
	}
}

func (t *TextInput) MoveRight() {
	if t.cursor < len(t.value) {
		t.cursor += 1
	}
}

func (t *TextInput) MoveHome() {
	t.cursor = 0
}

func (t *TextInput) MoveEnd() {
	t.cursor = 0
}

func (t *TextInput) DeleteBeforeCursor() {
	if t.cursor > 0 {
		t.value = t.value[:t.cursor-1] + t.value[t.cursor:]
	}
}

func (t *TextInput) DeleteAtCursor() {
	t.value = t.value[:t.cursor] + t.value[t.cursor+1:]
	if t.cursor >= len(t.value) {
		t.cursor = len(t.value) - 1
	}
}

func (t *TextInput) WriteAtCursor(c rune) {
	t.value = t.value[:t.cursor] + string(c) + t.value[t.cursor:]
	t.cursor += 1
}

func (t *TextInput) HandleEvent(ev *tcell.EventKey) {
	if ev.Key() == tcell.KeyLeft {
		t.MoveLeft()
	} else if ev.Key() == tcell.KeyRight {
		t.MoveRight()
	} else if ev.Key() == tcell.KeyHome {
		t.MoveHome()
	} else if ev.Key() == tcell.KeyEnd {
		t.MoveEnd()
	} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
		t.DeleteBeforeCursor()
	} else if ev.Key() == tcell.KeyDelete {
		t.DeleteAtCursor()
	} else if c := ev.Rune(); strconv.IsPrint(c) {
		t.WriteAtCursor(c)
	}
}

func (t TextInput) GetValue() string {
	return t.value
}
