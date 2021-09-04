package invoice

import (
	"fmt"

	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

func NewBuilder(output types.BuildOutput) (types.Builder, error) {
	switch output {
	case types.BuildOutputPDF:
		return newPDFBuilder()
	default:
		return nil, fmt.Errorf("unrecognized output format: %s", output)
	}
}

type builder struct {
	currency   string
	fontFamily types.FontFamily
}

func (b *builder) SetFontFamily(family types.FontFamily) { b.fontFamily = family }
func (b *builder) SetCurrency(currency string)           { b.currency = currency }
