package main

import (
	"os"

	"github.com/tinyzimmer/ginvoicer/pkg/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
