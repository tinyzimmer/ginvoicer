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
	"embed"
	"fmt"

	"github.com/signintech/gopdf"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

//go:embed data/*
var fontFS embed.FS

type Font interface {
	Load(*gopdf.GoPdf) error
	TextSize() int
	HeaderSize() int
	VerticalPadModifier() float64
	HorizontalPadModifier() float64
}

func GetFont(family types.FontFamily) (Font, error) {
	switch family {
	case types.FontFamilyHack:
		return &hack{}, nil
	case types.FontFamilyAnonymousPro:
		return &anonymousPro{}, nil
	case types.FontFamilyUbuntuMono:
		return &ubuntuMono{}, nil
	case types.FontFamilyGoMono:
		return &goMono{}, nil
	case types.FontFamilySpaceMono:
		return &spaceMono{}, nil
	case types.FontFamilyLiberationMono:
		return &liberationMono{}, nil
	case types.FontFamilyLuxiMono:
		return &luxiMono{}, nil
	default:
		return nil, fmt.Errorf("unrecognized font family: %s", family)
	}
}
