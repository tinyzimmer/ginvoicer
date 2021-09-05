package types

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func init() {
	usr, err := user.Current()
	if err != nil {
		DefaultConfigPath = "config.yaml"
		return
	}
	DefaultConfigPath = filepath.Join(usr.HomeDir, ".ginvoicer", "config.yaml")
}

var DefaultConfigPath string

type Config struct {
	// The configuration for invoices
	Invoices InvoicesConfig `json:"invoices" yaml:"invoices"`
	// The entity issuing invoices
	Payee *Entity `json:"payee" yaml:"payee"`
	// Entities eligible for invoicing.
	Payers []*Entity `json:"payers" yaml:"payers"`
	// Items that can be added to invoices
	Billables BillableList `json:"billables" yaml:"billables"`
}

type InvoicesConfig struct {
	// A prefix appended to invoice numbers
	Prefix string `json:"prefix" yaml:"prefix"`
	// A directory to store invoices by default
	Directory string `json:"directory" yaml:"directory"`
	// The font family to use for invoices
	FontFamily FontFamily `json:"fontFamily" yaml:"fontFamily"`
	// The number of zeroes to pad invoice numbers with
	ZeroPadding int `json:"zeroPadding" yaml:"zeroPadding"`
}

func (c *Config) FormatInvoiceNumber(num int) string {
	numFmt := fmt.Sprintf("%%0%dd", c.Invoices.ZeroPadding+1)
	numStr := fmt.Sprintf(numFmt, num)
	if c.Invoices.Prefix != "" {
		return fmt.Sprintf("%s%s", c.Invoices.Prefix, numStr)
	}
	return numStr
}

func (c *Config) GetNextInvoiceNumber() int {
	path := c.Invoices.Directory
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return 0
		}
	}

	last := -1
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// Strip any prefix
		nameFull := filepath.Base(path)
		if c.Invoices.Prefix != "" {
			// We are using a prefix and this file doesn't have one.
			// No need for further checks.
			if !strings.HasPrefix(nameFull, c.Invoices.Prefix) {
				return nil
			}
			nameFull = strings.TrimPrefix(nameFull, c.Invoices.Prefix)
		}

		// Strip any extension
		name := strings.Split(nameFull, ".")[0]

		inum, err := strconv.Atoi(name)
		if err != nil {
			// Skip files that are not parseable numbers
			return nil
		}
		if inum > last {
			last = inum
		}

		return nil
	})

	if err != nil {
		return 0
	}

	return last + 1
}
