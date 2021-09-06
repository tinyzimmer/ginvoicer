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

type FontFamily string

func (f FontFamily) String() string { return string(f) }

const (
	FontFamilyHack           FontFamily = "Hack"
	FontFamilyUbuntuMono     FontFamily = "UbuntuMono"
	FontFamilyAnonymousPro   FontFamily = "AnonymousPro"
	FontFamilyGoMono         FontFamily = "GoMono"
	FontFamilySpaceMono      FontFamily = "SpaceMono"
	FontFamilyLiberationMono FontFamily = "LiberationMono"
	FontFamilyLuxiMono       FontFamily = "LuxiMono"
)

type BuildOutput string

const (
	BuildOutputPDF BuildOutput = "pdf"
)
