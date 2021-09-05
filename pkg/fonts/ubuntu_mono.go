package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type ubuntuMono struct{}

func (u *ubuntuMono) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/ubuntu-mono/UbuntuMono.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/ubuntu-mono/UbuntuMono_B.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilyUbuntuMono, opts...)
}

func (u *ubuntuMono) VerticalPadModifier() float64   { return 6 }
func (u *ubuntuMono) HorizontalPadModifier() float64 { return 1.8 }
func (u *ubuntuMono) TextSize() int                  { return 10 }
func (u *ubuntuMono) HeaderSize() int                { return 28 }
