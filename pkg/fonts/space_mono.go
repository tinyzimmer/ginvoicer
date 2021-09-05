package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type spaceMono struct{}

func (f *spaceMono) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/space-mono/SpaceMono-Regular.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/space-mono/SpaceMono-Bold.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilySpaceMono, opts...)
}

func (f *spaceMono) VerticalPadModifier() float64   { return 6 }
func (f *spaceMono) HorizontalPadModifier() float64 { return 1.9 }
func (f *spaceMono) TextSize() int                  { return 9 }
func (f *spaceMono) HeaderSize() int                { return 26 }
