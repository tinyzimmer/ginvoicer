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

import (
	"io"
	"time"
)

// Builder represents an invoice builder saving to various formats.
// Currently only PDF is implemented.
type Builder interface {
	// Builders can be used as readers after they have generated the invoice.
	io.ReadCloser

	// Set the font family to use when BuildInvoice is called.
	SetFontFamily(FontFamily)

	// Build an invoice with the given details
	BuildInvoice(*InvoiceDetails) error

	// Write the generated invoice to the given path
	WriteFile(path string) error

	// A convenience method for building an invoice and automatically
	// writing it to an output file.
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
