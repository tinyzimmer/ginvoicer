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
	"strings"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func formatVal(val float64, cur string) string {
	p := message.NewPrinter(language.English)
	text := p.Sprintf("%.02f", val)
	if cur == "" {
		return text
	}
	unit, err := currency.ParseISO(strings.ToUpper(cur))
	if err != nil {
		return text
	}
	return p.Sprintf("%v%s", currency.Symbol(unit), text)
}
