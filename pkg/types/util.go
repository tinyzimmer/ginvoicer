package types

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func formatVal(val float64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%.02f", val)
}
