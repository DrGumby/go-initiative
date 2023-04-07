package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

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

	initList := &InitiativeList{
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

	dialog := &Dialog{
		"", "", "", 0,
	}

	currentContextIdx := 0
	contexts := []HandleInputContext{
		initList,
		dialog,
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
			} else if ev.Key() == tcell.KeyTAB {
				currentContextIdx = (currentContextIdx + 1) % len(contexts)
			} else {
				contexts[currentContextIdx].HandleEvent(ev)
			}
		}
		s.Clear()
		initList.Draw(s, 1, 1)
		dialog.Draw(s, 1, 20)
	}
}
