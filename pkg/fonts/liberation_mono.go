package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type liberationMono struct{}

func (f *liberationMono) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/liberation-mono/LiberationMono-Regular.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/liberation-mono/LiberationMono-Bold.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilyLiberationMono, opts...)
}

func (f *liberationMono) VerticalPadModifier() float64   { return 6 }
func (f *liberationMono) HorizontalPadModifier() float64 { return 1.9 }
func (f *liberationMono) TextSize() int                  { return 9 }
func (f *liberationMono) HeaderSize() int                { return 26 }
