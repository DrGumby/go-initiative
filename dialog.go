package main

import (
	"fmt"
	"strconv"

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
	name          string
	hpStr         string
	initiativeStr string
	selected      int
	initlist      *InitiativeList
}

func (d *Dialog) HandleEvent(ev *tcell.EventKey) {
	if ev.Key() == tcell.KeyUp {
		d.MoveUp()
	} else if ev.Key() == tcell.KeyDown {
		d.MoveDown()
	} else if c := ev.Rune(); strconv.IsPrint(c) {
		switch d.selected {
		case NAMELINE:
			d.name += string(c)
		case HPLINE:
			d.hpStr += string(c)
		case INITLINE:
			d.initiativeStr += string(c)
		}
	} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
		switch d.selected {
		case NAMELINE:
			if len(d.name) > 0 {
				d.name = d.name[:len(d.name)-1]
			}
		case HPLINE:
			if len(d.hpStr) > 0 {
				d.hpStr = d.hpStr[:len(d.hpStr)-1]
			}
		case INITLINE:
			if len(d.initiativeStr) > 0 {
				d.initiativeStr = d.initiativeStr[:len(d.initiativeStr)-1]
			}
		}
	}
}

func (d Dialog) ToInitiativeLine() (*InitiativeLine, error) {
	hp, err := strconv.Atoi(d.hpStr)
	if err != nil {
		return nil, err
	}

	initiative, err := strconv.Atoi(d.initiativeStr)
	if err != nil {
		return nil, err
	}

	return &InitiativeLine{
		name:       d.name,
		hp:         hp,
		maxhp:      hp,
		initiative: initiative,
	}, nil
}

func (d *Dialog) Draw(s tcell.Screen, startx int, starty int) {
	currenty := starty
	lines := []string{
		NAMELINE:    fmt.Sprintf("Name:       %s", d.name),
		HPLINE:      fmt.Sprintf("Max HP:     %s", d.hpStr),
		INITLINE:    fmt.Sprintf("Initiative: %s", d.initiativeStr),
		CONFIRMLINE: "Create new  ",
	}

	for i, row := range lines {
		var style tcell.Style
		if i == d.selected {
			style = highligtedStyle
		} else {
			style = defStyle
		}

		emitStr(s, startx, currenty, style, row)
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
