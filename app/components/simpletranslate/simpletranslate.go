package simpletranslate

import (
	"log"
	"strings"
	"time"
)

var monthsPolish = [...]string{
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

var daysPolish = [...]string{
	"Poniedziałek",
	"Wtorek",
	"Środa",
	"Czwartek",
	"Piątek",
	"Sobota",
	"Niedziela",
}

var remap = map[string]string{
	"Notes":       "Notatki",
	"Notes Index": "Indeks notatek",
	"Week":        "Tydzień",
	"Monday":      "Poniedziałek",
	"Tuesday":     "Wtorek",
	"Wednesday":   "Środa",
	"Thursday":    "Czwartek",
	"Friday":      "Piątek",
	"Saturday":    "Sobota",
	"Sunday":      "Niedziela",
	"W":           "T",
	"Notebook":    "Notatnik",
}

var remapPrefixes = map[string]string{
	"Mon,":      "Pon,",
	"Tue,":      "Wto,",
	"Wed,":      "Śro,",
	"Thu,":      "Czw,",
	"Fri,":      "Ptk,",
	"Sat,":      "Sob,",
	"Sun,":      "Nie,",
	"Monday":    "Poniedziałek",
	"Tuesday":   "Wtorek",
	"Wednesday": "Środa",
	"Thursday":  "Czwartek",
	"Friday":    "Piątek",
	"Saturday":  "Sobota",
	"Sunday":    "Niedziela",
}

func Translate(word string) string {
	return remap[word]
}

func TranslatePrefix(wordWithPrefix string) string {
	for k := range remapPrefixes {
		var translated string
		if strings.HasPrefix(wordWithPrefix, k) {
			translated = strings.Replace(wordWithPrefix, k, remapPrefixes[k], 1)
			//println("translated to '" + translated + "'")
			return translated
		}
	}

	log.Panic("Failed to translate '" + wordWithPrefix + "'")
	return ""
}

func TranslateShort(word string) string {
	return string([]rune(Translate(word))[:3])
}

func TranslateDay(d time.Weekday) string {
	return Translate(d.String())
}

func TranslateDayShort(d time.Weekday) string {
	return string([]rune(TranslateDay(d))[:3])
}

func TranslateMonth(m time.Month) string {
	return monthsPolish[m-1]
}

func TranslateMonthShort(m time.Month) string {
	return string([]rune(TranslateMonth(m))[:3])
}
