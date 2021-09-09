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

type fontOptions struct {
	path  string
	style int
}

func loadToPDF(pdf *gopdf.GoPdf, family types.FontFamily, opts ...fontOptions) error {
	for _, o := range opts {
		f, err := fontFS.Open(o.path)
		if err != nil {
			return err
		}
		err = pdf.AddTTFFontByReaderWithOption(family.String(), f, gopdf.TtfOption{
			Style:      o.style,
			UseKerning: true,
		})
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	return nil
}
