package common

import (
	"time"
)

type Menu struct {
	Location string `json:"location"`
	Date     string `json:"date"`
	Pizza    string `json:"pizza"`
}

type Menus struct {
	Pizzas []*Menu `json:"pizzas"`
}

type Location int

const (
	Telegraph Location = iota
	Shattuck
	Broadway
)

func (l Location) String() string {
	return [...]string{"telegraph", "shattuck", "broadway"}[l]
}

func (menu *Menu) SetTime(now time.Time) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	currentTime := now.In(loc)
	menu.Date = currentTime.Format("2006-01-02")
}

func (menu *Menu) SetLocation(loc Location) {
	menu.Location = loc.String()
}
