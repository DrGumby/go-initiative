package main

import (
	"sort"

	"github.com/gdamore/tcell/v2"
)

type InitiativeList struct {
	initlist []InitiativeLine
	selected int
}

func (l *InitiativeList) Append(item InitiativeLine) {
	l.initlist = append(l.initlist, item)
}

func (l InitiativeList) GetSelected() *InitiativeLine {
	return &l.initlist[l.selected]
}

func (l *InitiativeList) DeleteSelected() {
	l.initlist = append(l.initlist[:l.selected], l.initlist[l.selected+1:]...)
	l.correctSelected()
}

func (l *InitiativeList) correctSelected() {
	if l.selected >= len(l.initlist) {
		l.selected = len(l.initlist)
	}

	if l.selected < 0 {
		l.selected = 0
	}
}

func (l *InitiativeList) MoveDown() {
	if l.selected < len(l.initlist)-1 {
		l.selected += 1
	}
}

func (l *InitiativeList) MoveUp() {
	if l.selected > 0 {
		l.selected -= 1
	}
}

func (l *InitiativeList) Sort() {
	sort.SliceStable(l.initlist, func(i, j int) bool {
		return l.initlist[i].initiative < l.initlist[j].initiative
	})
}

func (l *InitiativeList) Draw(s tcell.Screen, startx int, starty int) {
	l.Sort()
	currenty := starty

	for i, row := range l.initlist {
		var style tcell.Style
		if i == l.selected {
			style = highligtedStyle
		} else {
			style = defStyle
		}

		emitStr(s, startx, currenty, style, row.ToString())
		currenty += 1
	}
}

func (l *InitiativeList) HandleEvent(ev *tcell.EventKey) {
	// Process event
	if ev.Key() == tcell.KeyUp {
		l.MoveUp()
	} else if ev.Key() == tcell.KeyDown {
		l.MoveDown()
	} else if ev.Rune() == '-' {
		l.GetSelected().ModifyHP(-1)
	} else if ev.Rune() == '+' {
		l.GetSelected().ModifyHP(+1)
	} else if ev.Rune() == 'd' {
		l.DeleteSelected()
	}
}
