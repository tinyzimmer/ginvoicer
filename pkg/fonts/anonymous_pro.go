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

package fonts

import (
	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

type anonymousPro struct{}

func (a *anonymousPro) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/anonymous-pro/AnonymousPro.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/anonymous-pro/AnonymousPro_B.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilyAnonymousPro, opts...)
}

func (a *anonymousPro) VerticalPadModifier() float64   { return 6 }
func (a *anonymousPro) HorizontalPadModifier() float64 { return 1.7 }
func (a *anonymousPro) TextSize() int                  { return 10 }
func (a *anonymousPro) HeaderSize() int                { return 28 }
