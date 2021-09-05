package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type anonymousPro struct{}

func (a *anonymousPro) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/anonymous-pro/AnonymousPro.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/anonymous-pro/AnonymousPro_B.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilyAnonymousPro, opts...)
}

func (a *anonymousPro) VerticalPadModifier() float64   { return 6 }
func (a *anonymousPro) HorizontalPadModifier() float64 { return 1.7 }
func (a *anonymousPro) TextSize() int                  { return 10 }
func (a *anonymousPro) HeaderSize() int                { return 28 }
