package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type hack struct{}

func (h *hack) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/hack/Hack.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/hack/Hack_B.ttf",
			style: gopdf.Bold,
		},
		{
			path:  "data/hack/Hack_I.ttf",
			style: gopdf.Italic,
		},
	}
	return loadToPDF(pdf, types.FontFamilyHack, opts...)
}

func (h *hack) VerticalPadModifier() float64   { return 6 }
func (h *hack) HorizontalPadModifier() float64 { return 2 }
