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
	"fmt"
	"strings"
)

type Entity struct {
	*Address `json:",inline" yaml:",inline"`

	Alias          string       `json:"alias" yaml:"alias"`
	BankAccount    *BankAccount `json:"bankAccount" yaml:"bankAccount"`
	RegistrationNo string       `json:"registrationNo" yaml:"registrationNo"`
	Email          string       `json:"email" yaml:"email"`
	Currency       string       `json:"currency" yaml:"currency"`
}

type BankAccount struct {
	*Address `json:",inline" yaml:",inline"`

	AccountNumber string `json:"accountNumber" yaml:"accountNumber"`
	RoutingNumber string `json:"routingNumber" yaml:"routingNumber"`
}

type Address struct {
	Name     string `json:"name" yaml:"name"`
	Address  string `json:"address" yaml:"address"`
	Address2 string `json:"address2" yaml:"address2"`
	City     string `json:"city" yaml:"city"`
	State    string `json:"state" yaml:"state"`
	ZipCode  string `json:"zipCode" yaml:"zipCode"`
	Country  string `json:"country" yaml:"country"`
}

func (a *Address) String() string {
	var sb strings.Builder
	if a.Address != "" {
		sb.WriteString(a.Address)
	}
	if a.Address2 != "" {
		sb.WriteString(", " + a.Address2)
	}
	if a.City != "" {
		sb.WriteString(", " + a.City)
	}
	if a.State != "" {
		sb.WriteString(", " + a.State)
	}
	if a.ZipCode != "" {
		sb.WriteString(", " + a.ZipCode)
	}
	if a.Country != "" {
		sb.WriteString(", " + a.Country)
	}
	return sb.String()
}

func (a *Address) Strings() []string {
	out := make([]string, 0)
	if a.Name != "" {
		out = append(out, strings.Split(a.Name, "\n")...)
	}
	if a.Address != "" {
		out = append(out, a.Address)
	}
	if a.Address2 != "" {
		out = append(out, a.Address2)
	}
	if a.City != "" || a.State != "" || a.ZipCode != "" {
		out = append(out, strings.Replace(fmt.Sprintf("%s %s %s", a.City, a.State, a.ZipCode), "  ", " ", -1))
	}
	if a.Country != "" {
		out = append(out, a.Country)
	}
	return out
}

func (e *Entity) TextWidth() float64 {
	var max float64
	for _, s := range e.Address.Strings() {
		if float64(len(s)) > max {
			max = float64(len(s))
		}
	}
	return max
}
