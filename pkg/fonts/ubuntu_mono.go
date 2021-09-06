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

type ubuntuMono struct{}

func (u *ubuntuMono) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/ubuntu-mono/UbuntuMono.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/ubuntu-mono/UbuntuMono_B.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilyUbuntuMono, opts...)
}

func (u *ubuntuMono) VerticalPadModifier() float64   { return 6 }
func (u *ubuntuMono) HorizontalPadModifier() float64 { return 1.8 }
func (u *ubuntuMono) TextSize() int                  { return 10 }
func (u *ubuntuMono) HeaderSize() int                { return 28 }
