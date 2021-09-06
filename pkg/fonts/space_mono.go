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

type spaceMono struct{}

func (f *spaceMono) Load(pdf *gopdf.GoPdf) (err error) {
	opts := []fontOptions{
		{
			path:  "data/space-mono/SpaceMono-Regular.ttf",
			style: gopdf.Regular,
		},
		{
			path:  "data/space-mono/SpaceMono-Bold.ttf",
			style: gopdf.Bold,
		},
	}
	return loadToPDF(pdf, types.FontFamilySpaceMono, opts...)
}

func (f *spaceMono) VerticalPadModifier() float64   { return 6 }
func (f *spaceMono) HorizontalPadModifier() float64 { return 1.9 }
func (f *spaceMono) TextSize() int                  { return 9 }
func (f *spaceMono) HeaderSize() int                { return 26 }
