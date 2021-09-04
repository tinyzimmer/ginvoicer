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
	// A prefix appended to invoice numbers
	InvoicePrefix string `json:"invoicePrefix" yaml:"invoicePrefix"`
	// A directory to store invoices by default
	InvoiceDirectory string `json:"invoiceDirectory" yaml:"invoiceDirectory"`
	// The number of zeroes to pad invoice numbers with
	ZeroPadding int `json:"zeroPadding" yaml:"zeroPadding"`
	// The entity issuing invoices
	Payee *Entity `json:"payee" yaml:"payee"`
	// Entities eligible for invoicing.
	Payers []*Entity `json:"payers" yaml:"payers"`
	// Items that can be added to invoices
	Billables BillableList `json:"billables" yaml:"billables"`
}

func (c *Config) FormatInvoiceNumber(num int) string {
	numFmt := fmt.Sprintf("%%0%dd", c.ZeroPadding+1)
	numStr := fmt.Sprintf(numFmt, num)
	if c.InvoicePrefix != "" {
		return fmt.Sprintf("%s%s", c.InvoicePrefix, numStr)
	}
	return numStr
}

func (c *Config) GetNextInvoiceNumber() int {
	path := c.InvoiceDirectory
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
		if c.InvoicePrefix != "" {
			// We are using a prefix and this file doesn't have one.
			// No need for further checks.
			if !strings.HasPrefix(nameFull, c.InvoicePrefix) {
				return nil
			}
			nameFull = strings.TrimPrefix(nameFull, c.InvoicePrefix)
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
