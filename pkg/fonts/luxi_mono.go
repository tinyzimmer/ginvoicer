package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type luxiMono struct{}

func (f *luxiMono) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/luxi-mono/luximr.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/luxi-mono/luximb.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilyLuxiMono, opts...)
}

func (f *luxiMono) VerticalPadModifier() float64   { return 6 }
func (f *luxiMono) HorizontalPadModifier() float64 { return 1.9 }
func (f *luxiMono) TextSize() int                  { return 9 }
func (f *luxiMono) HeaderSize() int                { return 26 }
