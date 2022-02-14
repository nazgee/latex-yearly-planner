package header

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
	"github.com/kudrykv/latex-yearly-planner/app/components/simpletranslate"
	"time"
)

type MonthItem struct {
	Val     time.Month
	ref     bool
	shorten bool
}

func (m MonthItem) Display() string {
	ref := m.Val.String()
	text := simpletranslate.TranslateMonth(m.Val)

	if m.shorten {
		text = simpletranslate.TranslateMonthShort(m.Val)
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
