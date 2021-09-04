package types

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
