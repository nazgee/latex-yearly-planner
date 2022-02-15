package cal

import (
	"strconv"
	"time"
)

type Actions []*Action
type Action struct {
	Time time.Time
}

func (a Action) Action() string {
	if a.Time.IsZero() {
		return ""
	}

	firstOfMonth := time.Date(a.Time.Year(), a.Time.Month(), 1, 0, 0, 0, 0, a.Time.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	dayNumber := strconv.FormatInt(int64(a.Time.Day()), 10)

	if a.Time.Weekday() == time.Sunday || a.Time.Weekday() == time.Saturday {
		dayNumber = `\textbf{` + dayNumber + `}`
	}

	ret := dayNumber + ` & $\Delta$\:\:$K$\:\:$\beta \hspace{0.7em}\square$`
	if a.Time.Day() != lastOfMonth.Day() {
		ret = ret + `\\`
	}

	return ret
}
