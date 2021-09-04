package types

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type FontFamily string

func (f FontFamily) String() string { return string(f) }

const (
	FontFamilyHack FontFamily = "Hack"
)

type BuildOutput string

const (
	BuildOutputPDF BuildOutput = "pdf"
)

// Builder represents an invoice builder saving to various formats.
type Builder interface {
	SetFontFamily(FontFamily)
	SetCurrency(string)

	BuildInvoice(*InvoiceDetails) error
	WriteFile(path string) error
}

type InvoiceDetails struct {
	InvoiceNumber string
	InvoiceDate   time.Time
	Payee         *Entity
	Payer         *Entity

	Items BillableList

	DueDate time.Time
}

type Config struct {
	Payee     *Entity      `json:"payee" yaml:"payee"`
	Payers    []*Entity    `json:"payers" yaml:"payers"`
	Billables BillableList `json:"billables" yaml:"billables"`
}

type Entity struct {
	Alias          string       `json:"alias" yaml:"alias"`
	Name           string       `json:"name" yaml:"name"`
	Address        string       `json:"address" yaml:"address"`
	Address2       string       `json:"address2" yaml:"address2"`
	City           string       `json:"city" yaml:"city"`
	State          string       `json:"state" yaml:"state"`
	ZipCode        string       `json:"zipCode" yaml:"zipCode"`
	Country        string       `json:"country" yaml:"country"`
	RegistrationNo string       `json:"registrationNo" yaml:"registrationNo"`
	BankAccount    *BankAccount `json:"bankAccount" yaml:"bankAccount"`
	Email          string       `json:"email" yaml:"email"`
}

type BankAccount struct {
	AccountNumber string `json:"accountNumber" yaml:"accountNumber"`
	RoutingNumber string `json:"routingNumber" yaml:"routingNumber"`
	Address       string `json:"address" yaml:"address"`
}

func (e *Entity) AddressString() string {
	var sb strings.Builder
	if e.Address != "" {
		sb.WriteString(e.Address)
	}
	if e.Address2 != "" {
		sb.WriteString(", " + e.Address2)
	}
	if e.City != "" {
		sb.WriteString(", " + e.City)
	}
	if e.State != "" {
		sb.WriteString(", " + e.State)
	}
	if e.ZipCode != "" {
		sb.WriteString(", " + e.ZipCode)
	}
	if e.Country != "" {
		sb.WriteString(", " + e.Country)
	}
	return sb.String()
}

func (e *Entity) Strings() []string {
	out := make([]string, 0)
	if e.Name != "" {
		out = append(out, strings.Split(e.Name, "\n")...)
	}
	if e.Address != "" {
		out = append(out, e.Address)
	}
	if e.Address2 != "" {
		out = append(out, e.Address2)
	}
	if e.City != "" || e.State != "" || e.ZipCode != "" {
		out = append(out, strings.Replace(fmt.Sprintf("%s %s %s", e.City, e.State, e.ZipCode), "  ", " ", -1))
	}
	if e.Country != "" {
		out = append(out, e.Country)
	}
	return out
}

func (e *Entity) TextWidth() float64 {
	var max float64
	for _, s := range e.Strings() {
		if float64(len(s)) > max {
			max = float64(len(s))
		}
	}
	return max
}

type BillableList []*Billable

func (b BillableList) MaxQuantitySize() int {
	var max int
	for _, i := range b {
		if len(i.FormattedQuantity()) > max {
			max = len(i.FormattedQuantity())
		}
	}
	return max
}

func (b BillableList) MaxUnitPriceSize() int {
	var max int
	for _, i := range b {
		if len(i.FormattedUnitPrice()) > max {
			max = len(i.FormattedUnitPrice())
		}
	}
	return max
}

func (b BillableList) MaxDiscountSize() int {
	var max int
	for _, i := range b {
		if len(i.FormattedDiscount()) > max {
			max = len(i.FormattedDiscount())
		}
	}
	return max
}

func (b BillableList) MaxTaxSize() int {
	var max int
	for _, i := range b {
		if len(i.FormattedTax()) > max {
			max = len(i.FormattedTax())
		}
	}
	return max
}

func (b BillableList) MaxSubtotalSize() int {
	var max int
	for _, i := range b {
		if len(i.FormattedSubtotal()) > max {
			max = len(i.FormattedSubtotal())
		}
	}
	return max
}

func (b BillableList) Subtotal() float64 {
	var total float64
	for _, i := range b {
		total += i.Subtotal()
	}
	return total
}

func (b BillableList) FormattedSubtotal() string {
	return formatVal(b.Subtotal())
}

func (b BillableList) Total() float64 {
	var total float64
	for _, i := range b {
		total += i.Total()
	}
	return total
}

func (b BillableList) FormattedTotal() string {
	return formatVal(b.Total())
}

type Billable struct {
	Alias       string  `json:"alias" yaml:"alias"`
	Description string  `json:"description" yaml:"description"`
	Quantity    float64 `json:"quantity" yaml:"quantity"`
	UnitPrice   float64 `json:"unitPrice" yaml:"unitPrice"`
	Discount    float64 `json:"discount" yaml:"discount"`
	Tax         float64 `json:"tax" yaml:"tax"`
}

func (b *Billable) Subtotal() float64 {
	return b.Quantity * b.UnitPrice
}

func (b *Billable) Total() float64 {
	return b.Subtotal() - (b.Subtotal() * b.Discount) + (b.Subtotal() * b.Tax)
}

func (b *Billable) FormattedQuantity() string {
	return formatVal(b.Quantity)
}

func (b *Billable) FormattedUnitPrice() string {
	return formatVal(b.UnitPrice)
}

func (b *Billable) FormattedDiscount() string {
	return formatVal(b.Discount*100) + "%"
}

func (b *Billable) FormattedTax() string {
	if b.Tax == 0 {
		return "Tax Exempt"
	}
	return formatVal(b.Tax*100) + "%"
}

func (b *Billable) FormattedSubtotal() string {
	return formatVal(b.Subtotal())
}

func formatVal(val float64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%.02f", val)
}
