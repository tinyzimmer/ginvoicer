package types

type FontFamily string

func (f FontFamily) String() string { return string(f) }

const (
	FontFamilyHack FontFamily = "Hack"
)

type BuildOutput string

const (
	BuildOutputPDF BuildOutput = "pdf"
)
