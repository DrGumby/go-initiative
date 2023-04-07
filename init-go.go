package main

import (
	"fmt"
	"log"
  "sort"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

var (
	defStyle        = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	highligtedStyle = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
)

type InitiativeLine struct {
	name       string
	initiative int
	hp         int
	maxhp      int
}

func (l *InitiativeLine) ModifyHP(delta int) {
	l.hp += delta
}

func (l InitiativeLine) ToString() string {
  hpstr := fmt.Sprintf("%d/%d", l.hp, l.maxhp)
  initstr := fmt.Sprintf("%d", l.initiative)

	return fmt.Sprintf("%-*.*s%-*.*s%-*.*s", 10, 10, l.name, 8, 8, hpstr, 4, 4, initstr)
}

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

func (l InitiativeList) Draw(s tcell.Screen, startx int, starty int) {
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

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func main() {
	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	initList := InitiativeList{
		[]InitiativeLine{
			{
				"Hello", 20, 3, 40,
			},
			{
				"World", 43, 1, 5,
			},
			{
				"Test", 21, 3, 40,
			},
		},
		0,
	}

	// Event loop
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'q' {
				return
			} else if ev.Key() == tcell.KeyUp {
				initList.MoveUp()
			} else if ev.Key() == tcell.KeyDown {
				initList.MoveDown()
			} else if ev.Rune() == '-' {
        initList.GetSelected().ModifyHP(-1)
			} else if ev.Rune() == '+' {
        initList.GetSelected().ModifyHP(+1)
      }
		}
		initList.Draw(s, 1, 1)
	}
}
