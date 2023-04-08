package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

const (
	NAMELINE    = 0
	HPLINE      = 1
	INITLINE    = 2
	CONFIRMLINE = 3
	END         = 4
)

type Dialog struct {
	name          TextInput
	hpStr         TextInput
	initiativeStr TextInput
	selected      int
	initlist      *InitiativeList
}

func (d *Dialog) HandleEvent(ev *tcell.EventKey) {
	if ev.Key() == tcell.KeyUp {
		d.MoveUp()
	} else if ev.Key() == tcell.KeyDown {
		d.MoveDown()
	} else if ev.Key() == tcell.KeyEnter {
		if d.selected == CONFIRMLINE {
			init, err := d.ToInitiativeLine()
			if err != nil {
				return
			}
			d.initlist.Append(*init)
			d.name = TextInput{" ", 0}
			d.initiativeStr = TextInput{" ", 0}
			d.hpStr = TextInput{" ", 0}
		}
	} else {
		switch d.selected {
		case NAMELINE:
			d.name.HandleEvent(ev)
		case HPLINE:
			d.hpStr.HandleEvent(ev)
		case INITLINE:
			d.initiativeStr.HandleEvent(ev)
		}
	}
}

func (d Dialog) ToInitiativeLine() (*InitiativeLine, error) {
	hp, err := strconv.Atoi(strings.TrimSpace(d.hpStr.GetValue()))
	if err != nil {
		return nil, err
	}

	initiative, err := strconv.Atoi(strings.TrimSpace(d.initiativeStr.GetValue()))
	if err != nil {
		return nil, err
	}

	return &InitiativeLine{
		name:       strings.TrimSpace(d.name.GetValue()),
		hp:         hp,
		maxhp:      hp,
		initiative: initiative,
	}, nil
}

func (d *Dialog) Draw(s tcell.Screen, startx int, starty int) {
	currenty := starty
	lines := []struct {
		s string
		t *TextInput
	}{
		NAMELINE:    {fmt.Sprintf("Name:       "), &d.name},
		HPLINE:      {fmt.Sprintf("Max HP:     "), &d.hpStr},
		INITLINE:    {fmt.Sprintf("Initiative: "), &d.initiativeStr},
		CONFIRMLINE: {"Create new  ", nil},
	}

	for i, row := range lines {
		var style tcell.Style
		if i == d.selected {
			style = highligtedStyle
		} else {
			style = defStyle
		}

		emitStr(s, startx, currenty, style, row.s)
		if row.t != nil {
			row.t.Draw(s, startx+len(row.s), currenty)
		}
		currenty += 1
	}
}

func (d *Dialog) MoveDown() {
	if d.selected < END-1 {
		d.selected += 1
	}
}

func (d *Dialog) MoveUp() {
	if d.selected > 0 {
		d.selected -= 1
	}
}
