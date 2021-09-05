package invoice

import (
	"github.com/signintech/gopdf"

	"github.com/tinyzimmer/ginvoicer/pkg/fonts"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type pdfBuilder struct {
	*gopdf.GoPdf

	currency              string
	fontFamily            types.FontFamily
	hpad, vpad            float64
	pageWidth, pageHeight float64
}

func newPDFBuilder() (*pdfBuilder, error) {
	pdf := &gopdf.GoPdf{}
	cfg := gopdf.Config{PageSize: *gopdf.PageSizeLetter}
	pdf.Start(cfg)
	pdf.AddPage()
	return &pdfBuilder{
		GoPdf:      pdf,
		currency:   "USD",
		fontFamily: types.FontFamilyHack,
		vpad:       70,
		hpad:       30,
		pageWidth:  cfg.PageSize.W,
		pageHeight: cfg.PageSize.H,
	}, nil
}

func (p *pdfBuilder) SetFontFamily(family types.FontFamily) { p.fontFamily = family }

func (p *pdfBuilder) BuildInvoice(info *types.InvoiceDetails) (err error) {
	p.SetX(p.hpad)
	p.SetY(p.vpad)

	// Change currency if necessary
	if info.Payer.Currency != "" {
		p.currency = info.Payer.Currency
	}

	// Load the font family
	font, err := fonts.GetFont(p.fontFamily)
	if err != nil {
		return
	}
	if err = font.Load(p.GoPdf); err != nil {
		return
	}

	// Write the heading

	if err = p.writeHeader(info, font); err != nil {
		return
	}

	// Jump a fifth of the way down to begin invoice items
	p.SetY(p.vpad + p.pageHeight/5)
	p.SetX(p.hpad)

	if err = p.writeInvoiceTable(info, font); err != nil {
		return
	}

	// Jump down and write the due date

	p.SetX(p.hpad)
	p.SetY(p.GetY() + 50)

	if err = p.Text("Due Date: "); err != nil {
		return
	}
	if err = p.Text(info.DueDate.Format("2 Jan 2006")); err != nil {
		return
	}

	// Jump down and write a dotted line before payment advice
	p.SetLineType("dashed")
	p.SetStrokeColor(0, 0, 0)
	p.SetY(p.pageHeight - (p.pageHeight / 3))
	p.Line(p.hpad-p.horizontalPadding(font)/4, p.GetY(), p.pageWidth-p.horizontalPadding(font), p.GetY())

	p.SetX(p.hpad)
	p.SetY(p.GetY() + 30)

	if err = p.writePaymentOptions(info, font); err != nil {
		return
	}

	// Write out footer
	return p.writeFooter(info, font)
}

func (p *pdfBuilder) WriteFile(path string) error { return p.WritePdf(path) }

func (p *pdfBuilder) BuildAndWriteInvoice(info *types.InvoiceDetails, outpath string) error {
	if err := p.BuildInvoice(info); err != nil {
		return err
	}
	return p.WriteFile(outpath)
}
