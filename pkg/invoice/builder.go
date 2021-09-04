package invoice

import (
	"fmt"

	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

func NewBuilder(output types.BuildOutput) (types.Builder, error) {
	switch output {
	case types.BuildOutputPDF:
		return newPDFBuilder()
	default:
		return nil, fmt.Errorf("unrecognized output format: %s", output)
	}
}
