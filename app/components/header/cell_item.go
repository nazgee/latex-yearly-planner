package header

import "log"

type CellItem struct {
	Text     string
	Ref      string
	selected bool
}

func NewCellItem(text string) CellItem {
	return CellItem{Text: text}
}

func (c CellItem) Select() CellItem {
	c.selected = true

	return c
}

func (c CellItem) Selected(selected bool) CellItem {
	c.selected = selected

	return c
}

func (c CellItem) Refer(ref string) CellItem {
	c.Ref = ref

	return c
}

func (c CellItem) Display() string {
	if len(c.Ref) == 0 {
		c.Ref = c.Text
		// TODO this would break when translation is used
		log.Panic("text used as a ref, would cause translation issues")
	}

	link := `\hyperlink{` + c.Ref + `}{` + c.Text + `}`

	if c.selected {
		return `\cellcolor{black}{\textcolor{white}{` + link + `}}`
	}

	return link
}
