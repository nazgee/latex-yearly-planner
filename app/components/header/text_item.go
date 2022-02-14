package header

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
	"log"
	"runtime/debug"
	"strconv"
	"unicode/utf8"
)

type TextItem struct {
	Name      string
	bold      bool
	ref       bool
	refPrefix string
	refText   string
}

func NewTextItem(name string) TextItem {
	return TextItem{
		Name: name,
	}
}

func (t TextItem) Display() string {
	var (
		out string
		ref string
	)
	if t.bold {
		out = "\\textbf{" + t.Name + "}"
	} else {
		out = t.Name
	}

	if len(t.refText) > 0 {
		ref = t.refText
	} else {
		ref = t.refPrefix + t.Name
		// TODO this would break when translation is used
		debug.PrintStack()
		log.Panic("text used as a ref, would cause translation issues, refPrefix=" + t.refPrefix + " (len=" + strconv.FormatInt(int64(utf8.RuneCountInString(t.refPrefix)), 10) + ") name=" + t.Name)
	}

	if t.ref {
		return hyper.Target(ref, out)
	}

	return hyper.Link(ref, out)
}

func (t TextItem) Ref(ref bool) TextItem {
	t.ref = ref

	return t
}

func (t TextItem) Bold(f bool) TextItem {
	t.bold = f

	return t
}

func (t TextItem) RefPrefix(refPrefix string) TextItem {
	t.refPrefix = refPrefix

	return t
}

func (t TextItem) RefText(refText string) TextItem {
	t.refText = refText

	return t
}
