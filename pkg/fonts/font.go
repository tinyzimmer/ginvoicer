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
	default:
		return nil, fmt.Errorf("unrecognized font family: %s", family)
	}
}
