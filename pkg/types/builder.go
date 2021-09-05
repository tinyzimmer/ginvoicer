package types

import "time"

// Builder represents an invoice builder saving to various formats.
// Currently only PDF is implemented.
type Builder interface {
	SetFontFamily(FontFamily)

	BuildInvoice(*InvoiceDetails) error
	WriteFile(path string) error
	BuildAndWriteInvoice(info *InvoiceDetails, outpath string) error
}

// InvoiceDetails represents the details required for
// building an invoice.
type InvoiceDetails struct {
	// The invoice number
	InvoiceNumber string
	// The date the invoice is being issued
	InvoiceDate time.Time
	// The date the invoice is due
	DueDate time.Time

	// The entity sending the invoice
	Payee *Entity
	// The entity responsible for paying the invoice
	Payer *Entity

	// The billable items in the invoice
	Items BillableList
}
