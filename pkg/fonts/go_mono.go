package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type goMono struct{}

func (f *goMono) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/go-mono/GoMono.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/go-mono/GoMono_B.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilyGoMono, opts...)
}

func (f *goMono) VerticalPadModifier() float64   { return 6 }
func (f *goMono) HorizontalPadModifier() float64 { return 1.9 }
func (f *goMono) TextSize() int                  { return 9 }
func (f *goMono) HeaderSize() int                { return 26 }
