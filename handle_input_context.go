package main

import "github.com/gdamore/tcell/v2"

type HandleInputContext interface {
	HandleEvent(ev *tcell.EventKey)
}
