package main

import "github.com/gdamore/tcell/v2"

var (
	defStyle        = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	highligtedStyle = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
)
