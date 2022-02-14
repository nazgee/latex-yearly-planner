package header

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
)

var month2pl = map[string]string {
	"January":"Styczeń",
	"February":"Luty",
	"March":"Marzec",
	"April":"Kwiecień",
	"May":"Maj",
	"June":"Czerwiec",
	"July":"Lipiec",
	"August":"Sierpień",
	"September":"Wrzesień",
	"October":"Październik",
	"November":"Listopad",
	"December":"Grudzień",
}

var monthspl = [...]string{
        "Styczeń",
        "Luty",
        "Marzec",
        "Kwiecień",
        "Maj",
        "Czerwiec",
        "Lipiec",
        "Sierpień",
        "Wrzesień",
        "Październik",
        "Listopad",
        "Grudzień",
}

type MonthItem struct {
	Val     time.Month
	ref     bool
	shorten bool
}

func (m MonthItem) Display() string {
	//ref := month2pl[m.Val.String()]
	ref := m.Val.String()
	text := month2pl[m.Val.String()]
	//text := ref

	if m.shorten {
		chars := []rune(text)
		text = string(chars[:3])
	}

	if m.ref {
		return hyper.Target(ref, text)
	}

	return hyper.Link(ref, text)
}

func (m MonthItem) Ref() MonthItem {
	m.ref = true

	return m
}

func (m MonthItem) Shorten(f bool) MonthItem {
	m.shorten = f

	return m
}

func NewMonthItem(m time.Month) MonthItem {
	return MonthItem{Val: m}
}
