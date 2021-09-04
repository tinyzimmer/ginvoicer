package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type fontOptions struct {
	path  string
	style int
}

func loadToPDF(pdf *gopdf.GoPdf, family types.FontFamily, opts ...fontOptions) error {
	for _, o := range opts {
		f, err := fontFS.Open(o.path)
		if err != nil {
			return err
		}
		err = pdf.AddTTFFontByReaderWithOption(family.String(), f, gopdf.TtfOption{
			Style: o.style,
		})
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	return nil
}
