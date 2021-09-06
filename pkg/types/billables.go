/*

Copyright (C) 2021 Avi Zimmerman

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program; if not, write to the Free Software Foundation,
Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

*/

package types

// Billable represents a billable item in an invoice.
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

func (b BillableList) LongestQuantity() string {
	var max int
	var longest string
	for _, i := range b {
		this := i.FormattedQuantity()
		if len(this) > max {
			longest = this
			max = len(this)
		}
	}
	return longest
}

func (b BillableList) LongestUnitPrice() string {
	var max int
	var longest string
	for _, i := range b {
		this := i.FormattedUnitPrice()
		if len(this) > max {
			longest = this
			max = len(this)
		}
	}
	return longest
}

func (b BillableList) LongestDiscount() string {
	var max int
	var longest string
	for _, i := range b {
		this := i.FormattedDiscount()
		if len(this) > max {
			longest = this
			max = len(this)
		}
	}
	return longest
}

func (b BillableList) LongestTax() string {
	var max int
	var longest string
	for _, i := range b {
		this := i.FormattedTax()
		if len(this) > max {
			longest = this
			max = len(this)
		}
	}
	return longest
}

func (b BillableList) LongestSubtotal() string {
	var max int
	var longest string
	for _, i := range b {
		this := i.FormattedSubtotal()
		if len(this) > max {
			longest = this
			max = len(this)
		}
	}
	return longest
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
