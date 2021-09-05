package invoice

import (
	"fmt"
	"strings"

	"github.com/tinyzimmer/ginvoicer/pkg/fonts"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

func (p *pdfBuilder) writeHeader(info *types.InvoiceDetails, font fonts.Font) (err error) {
	// Invoice in big text
	if err = p.setRegular(p.headerTextSize); err != nil {
		return
	}
	if err = p.Text("INVOICE"); err != nil {
		return
	}

	p.SetX(p.hpad + p.horizontalPadding(font)/2)
	p.SetY(p.vpad + float64(p.headerTextSize))

	// Write payer info below the heading

	if err = p.setRegular(p.textSize); err != nil {
		return
	}
	for _, line := range info.Payer.Strings() {
		if err = p.Text(strings.ToUpper(line)); err != nil {
			return
		}
		p.SetX(p.hpad + p.horizontalPadding(font)/2)
		p.SetY(p.GetY() + float64(p.textSize+4))
	}

	// Write invoice info

	p.SetY(p.vpad)
	p.SetX(p.pageWidth - p.hpad - p.pageWidth/3)
	if err = p.setBold(p.textSize); err != nil {
		return
	}
	if err = p.Text("Invoice Date"); err != nil {
		return
	}

	p.SetY(p.GetY() + float64(p.textSize) + p.verticalPadding(font)*2)
	p.SetX(p.pageWidth - p.hpad - p.pageWidth/3)
	if err = p.Text("Invoice Number"); err != nil {
		return
	}

	if err = p.setRegular(p.textSize); err != nil {
		return
	}

	p.SetY(p.vpad + float64(p.textSize) + 2)
	p.SetX(p.pageWidth - p.hpad - p.pageWidth/3)
	if err = p.Text(info.InvoiceDate.Format("2 Jan 2006")); err != nil {
		return
	}

	p.SetY(p.GetY() + float64(p.textSize) + p.verticalPadding(font)*2)
	p.SetX(p.pageWidth - p.hpad - p.pageWidth/3)
	if err = p.Text(info.InvoiceNumber); err != nil {
		return
	}

	// Write payee info to side of page

	p.SetY(p.vpad)
	p.SetX(p.pageWidth - p.hpad - info.Payee.TextWidth() - p.horizontalPadding(font)*2.5)

	for _, line := range info.Payee.Strings() {
		if err = p.Text(strings.ToUpper(line)); err != nil {
			return
		}
		p.SetX(p.pageWidth - p.hpad - info.Payee.TextWidth() - p.horizontalPadding(font)*2.5)
		p.SetY(p.GetY() + float64(p.textSize+4))
	}

	return
}

func (p *pdfBuilder) writeInvoiceTable(info *types.InvoiceDetails, font fonts.Font) (err error) {
	// Write the invoice table header
	if err = p.setBold(p.textSize); err != nil {
		return
	}
	if err = p.Text("Description"); err != nil {
		return
	}
	p.SetY(p.vpad + p.pageHeight/5)
	p.SetX(p.GetX() + p.horizontalPadding(font)*3.5)
	if err = p.Text("Quantity"); err != nil {
		return
	}
	p.SetY(p.vpad + p.pageHeight/5)
	p.SetX(p.GetX() + p.horizontalPadding(font))
	if err = p.Text("Unit Price"); err != nil {
		return
	}
	p.SetY(p.vpad + p.pageHeight/5)
	p.SetX(p.GetX() + p.horizontalPadding(font))
	if err = p.Text("Discount"); err != nil {
		return
	}

	p.SetY(p.vpad + p.pageHeight/5)
	p.SetX(p.GetX() + p.horizontalPadding(font))
	if err = p.Text("Tax"); err != nil {
		return
	}
	p.SetY(p.vpad + p.pageHeight/5)
	amountText := fmt.Sprintf("Amount %s", p.currency)
	p.SetX(p.pageWidth - p.hpad - p.horizontalPadding(font)*1.3)
	if err = p.Text(amountText); err != nil {
		return
	}

	// Place a divider under the header
	p.SetLineWidth(1)
	p.SetStrokeColor(96, 96, 96)
	p.Line(p.hpad, p.GetY()+7, p.pageWidth-p.horizontalPadding(font)*1.7+float64(p.textSize)*2, p.GetY()+7)

	// Write out the invoice items
	p.SetY(p.GetY() + 25)
	p.SetX(p.hpad)
	p.SetStrokeColor(184, 184, 184)
	if err = p.setRegular(p.textSize); err != nil {
		return
	}
	for _, item := range info.Items {
		baseY := p.GetY()
		if err = p.Text(item.Description); err != nil {
			return
		}

		var text string
		var max int

		p.SetY(baseY)
		text = item.FormattedQuantity()
		max = info.Items.MaxQuantitySize()
		p.SetX(p.hpad + (p.horizontalPadding(font) * float64(p.textSize) * 0.6) + float64(max-len(text))*(float64(p.textSize)/2))
		if err = p.Text(text); err != nil {
			return
		}

		p.SetY(baseY)
		text = item.FormattedUnitPrice()
		max = info.Items.MaxUnitPriceSize()
		p.SetX(p.hpad + (p.horizontalPadding(font) * float64(p.textSize) * 0.92) + float64(max-len(text))*(float64(p.textSize)/2))
		if len(text) > 6 {
			p.SetX(p.GetX() - p.horizontalPadding(font)/2)
		}
		if err = p.Text(text); err != nil {
			return
		}

		p.SetY(baseY)
		text = item.FormattedDiscount()
		max = info.Items.MaxDiscountSize()
		p.SetX(p.hpad + (p.horizontalPadding(font) * float64(p.textSize) * 1.15) + float64(max-len(text))*(float64(p.textSize)/2))
		if err = p.Text(text); err != nil {
			return
		}

		p.SetY(baseY)
		text = item.FormattedTax()
		max = info.Items.MaxTaxSize()
		p.SetX(p.hpad + (p.horizontalPadding(font) * float64(p.textSize) * 1.30) + float64(max-len(text))*(float64(p.textSize)/2))
		if err = p.Text(text); err != nil {
			return
		}

		p.SetY(baseY)
		text = item.FormattedSubtotal()
		max = info.Items.MaxSubtotalSize()
		p.SetX(p.hpad + (p.horizontalPadding(font) * float64(p.textSize) * 1.545) + float64(max-len(text))*(float64(p.textSize)/2))
		if err = p.Text(text); err != nil {
			return
		}

		p.Line(p.hpad, p.GetY()+10, p.pageWidth-p.horizontalPadding(font)*1.7+float64(p.textSize)*2, p.GetY()+10)
		p.SetY(p.GetY() + 25)
		p.SetX(p.hpad)
	}

	// Write out totals

	p.SetY(p.GetY() + 10)

	p.SetX(p.hpad + p.horizontalPadding(font)*11.5)
	if err = p.Text("Subtotal"); err != nil {
		return
	}
	p.SetX(p.GetX() + p.horizontalPadding(font))
	if err = p.Text(info.Items.FormattedSubtotal()); err != nil {
		return
	}

	p.SetStrokeColor(96, 96, 96)
	p.Line(p.pageWidth/2, p.GetY()+10, p.pageWidth-p.horizontalPadding(font)*1.75+float64(p.textSize)*2, p.GetY()+10)

	p.SetY(p.GetY() + 30)
	p.SetX(p.hpad + p.horizontalPadding(font)*11.5 - float64(p.textSize)/1.8)
	if err = p.setBold(p.textSize); err != nil {
		return
	}
	if err = p.Text(fmt.Sprintf("TOTAL %s", p.currency)); err != nil {
		return
	}
	p.SetX(p.GetX() + p.horizontalPadding(font))
	if err = p.Text(info.Items.FormattedTotal()); err != nil {
		return
	}

	return
}

func (p *pdfBuilder) writePaymentOptions(info *types.InvoiceDetails, font fonts.Font) (err error) {
	if err = p.setRegular(p.headerTextSize - 4); err != nil {
		return
	}
	if err = p.Text("PAYMENT ADVICE"); err != nil {
		return
	}

	p.SetX(p.hpad + p.horizontalPadding(font)/4)
	startY := p.GetY() + float64(p.headerTextSize)
	p.SetY(startY)

	// Write payee info again
	if err = p.setRegular(p.textSize); err != nil {
		return
	}

	if err = p.Text("To: "); err != nil {
		return
	}
	p.SetX(p.hpad + p.horizontalPadding(font))
	p.SetY(startY)

	for _, line := range info.Payee.Strings() {
		if err = p.Text(strings.ToUpper(line)); err != nil {
			return
		}
		p.SetX(p.hpad + p.horizontalPadding(font))
		p.SetY(p.GetY() + float64(p.textSize+4))
	}

	// Write out payer info and details

	p.SetY(p.pageHeight - (p.pageHeight / 3) + 20)
	p.SetX(p.pageWidth - p.hpad - p.pageWidth/3)
	if err = p.setBold(p.textSize); err != nil {
		return
	}
	if err = p.Text("Customer"); err != nil {
		return
	}

	p.SetX(p.pageWidth - p.hpad - p.pageWidth/3)
	p.SetY(p.GetY() + float64(p.textSize)*1.67)
	if err = p.Text("Invoice Number"); err != nil {
		return
	}

	p.SetLineType("")
	p.SetStrokeColor(156, 156, 156)
	p.Line(p.pageWidth-p.hpad-p.pageWidth/3, p.GetY()+7, p.pageWidth-p.horizontalPadding(font)*1.75+float64(p.textSize)*2, p.GetY()+7)

	p.SetX(p.pageWidth - p.hpad - p.pageWidth/3)
	p.SetY(p.GetY() + float64(p.textSize)*1.67 + 5)
	if err = p.Text("Amount Due"); err != nil {
		return
	}

	p.SetX(p.pageWidth - p.hpad - p.pageWidth/3)
	p.SetY(p.GetY() + float64(p.textSize)*1.67)
	if err = p.Text("Due Date"); err != nil {
		return
	}
	p.Line(p.pageWidth-p.hpad-p.pageWidth/3, p.GetY()+7, p.pageWidth-p.horizontalPadding(font)*1.75+float64(p.textSize)*2, p.GetY()+7)

	p.SetY(p.pageHeight - (p.pageHeight / 3) + 20)
	p.SetX(p.pageWidth - p.hpad - p.horizontalPadding(font)*3)
	if err = p.setRegular(p.textSize); err != nil {
		return
	}
	if err = p.Text(strings.Split(info.Payer.Name, "\n")[0]); err != nil {
		return
	}

	p.SetX(p.pageWidth - p.hpad - p.horizontalPadding(font)*3)
	p.SetY(p.GetY() + float64(p.textSize)*1.67)
	if err = p.Text(info.InvoiceNumber); err != nil {
		return
	}

	p.SetX(p.pageWidth - p.hpad - p.horizontalPadding(font)*3)
	p.SetY(p.GetY() + float64(p.textSize)*1.67 + 5)
	if err = p.Text(info.Items.FormattedTotal()); err != nil {
		return
	}

	p.SetX(p.pageWidth - p.hpad - p.horizontalPadding(font)*3)
	p.SetY(p.GetY() + float64(p.textSize)*1.67)
	if err = p.Text(info.DueDate.Format("2 Jan 2006")); err != nil {
		return
	}

	// Write out payment options

	p.SetX(p.pageWidth/2 - p.horizontalPadding(font)*2)
	p.SetY(p.GetY() + p.verticalPadding(font)*1.5)
	if err = p.setBold(p.textSize); err != nil {
		return
	}
	if err = p.Text("Payment Options"); err != nil {
		return
	}

	if err = p.setRegular(p.textSize); err != nil {
		return
	}

	p.SetY(p.GetY() + p.verticalPadding(font)*1.15)
	if info.Payee.BankAccount != nil {
		p.SetX(p.pageWidth/4 + p.horizontalPadding(font))
		p.Text("Bank Transfer")

		p.SetX(p.pageWidth/4 + p.horizontalPadding(font)*1.5)
		p.SetY(p.GetY() + p.verticalPadding(font))
		p.Text("Account: " + info.Payee.BankAccount.AccountNumber)

		p.SetX(p.pageWidth/4 + p.horizontalPadding(font)*1.5)
		p.SetY(p.GetY() + p.verticalPadding(font))
		p.Text("Routing: " + info.Payee.BankAccount.RoutingNumber)

		for _, line := range info.Payee.BankAccount.Strings() {
			p.SetX(p.pageWidth/4 + p.horizontalPadding(font)*2)
			p.SetY(p.GetY() + p.verticalPadding(font))
			p.Text(line)
		}

	}
	return
}

func (p *pdfBuilder) writeFooter(info *types.InvoiceDetails, font fonts.Font) error {
	p.SetX(p.hpad)
	p.SetY(p.pageHeight - p.verticalPadding(font)*1.5)
	if err := p.setRegular(p.textSize - 2); err != nil {
		return err
	}
	return p.Text(fmt.Sprintf("Company Registration No: %s. Registered Office: %s", info.Payee.RegistrationNo, info.Payee.Address.String()))
}

func (p *pdfBuilder) horizontalPadding(font fonts.Font) float64 {
	return p.pageWidth / (float64(p.textSize) * font.HorizontalPadModifier())
}

func (p *pdfBuilder) verticalPadding(font fonts.Font) float64 {
	return p.pageHeight / (float64(p.textSize) * font.VerticalPadModifier())
}

func (p *pdfBuilder) setRegular(size int) error {
	return p.SetFont(p.fontFamily.String(), "", size)
}

func (p *pdfBuilder) setBold(size int) error {
	return p.SetFont(p.fontFamily.String(), "Bold", size)
}
