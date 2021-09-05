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
