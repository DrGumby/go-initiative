package main

import "fmt"

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
