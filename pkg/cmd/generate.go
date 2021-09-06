package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/tinyzimmer/ginvoicer/pkg/invoice"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

var (
	outFile       string
	payer         string
	invoiceNumber int
	billables     map[string]int
	dueIn         time.Duration
)

func init() {
	flags := generateCommand.Flags()

	flags.StringVarP(&payer, "payer", "p", "", "payer (by name or alias) to generate the invoice for")
	flags.IntVarP(&invoiceNumber, "number", "n", -1, "invoice number (default: auto-increment)")
	flags.DurationVarP(&dueIn, "due-in", "d", time.Hour*24*7, "time until invoice is due")
	flags.StringVarP(&outFile, "output", "o", "", "path to write the invoice to (defaults to the format and locations in your configuration)")
	flags.StringToIntVarP(&billables, "item", "i", map[string]int{}, "items to add to the invoice in the format of alias=quantity")

	cobra.CheckErr(cobra.MarkFlagRequired(flags, "payer"))

	rootCommand.AddCommand(generateCommand)
}

var generateCommand = &cobra.Command{
	Use:     "generate",
	Short:   "Generate an invoice",
	Aliases: []string{"new", "new-invoice"},
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := loadConfig()
		if err != nil {
			return err
		}

		if config.Payee == nil {
			return fmt.Errorf("no payee information found in configuration")
		}

		var payerEntity *types.Entity
		for _, p := range config.Payers {
			if payer == p.Name || payer == p.Alias {
				payerEntity = p
			}
		}
		if payerEntity == nil {
			return fmt.Errorf("no payer found by the name %q", payer)
		}

		if invoiceNumber == -1 {
			invoiceNumber = config.GetNextInvoiceNumber()
		}

		invoiceItems := make(types.BillableList, 0)
		for name, quantity := range billables {
		GetItem:
			for _, item := range config.Billables {
				if name == item.Alias {
					invoiceItems = append(invoiceItems, &types.Billable{
						Description: item.Description,
						Quantity:    float64(quantity),
						UnitPrice:   item.UnitPrice,
						Discount:    item.Discount,
						Tax:         item.Tax,
					})
					break GetItem
				}
			}

		}

		builder, err := invoice.NewBuilder(types.BuildOutputPDF)
		if err != nil {
			return err
		}

		if config.Invoices.FontFamily != "" {
			builder.SetFontFamily(config.Invoices.FontFamily)
		}

		invoiceNumberStr := config.FormatInvoiceNumber(invoiceNumber)

		invoiceInfo := &types.InvoiceDetails{
			InvoiceNumber: invoiceNumberStr,
			InvoiceDate:   time.Now(),
			Payee:         config.Payee,
			Payer:         payerEntity,
			Items:         invoiceItems,
			DueDate:       time.Now().Add(dueIn),
		}

		output := outFile
		if output == "" {
			output = fmt.Sprintf("%s.pdf", invoiceNumberStr)
			if config.Invoices.Directory != "" {
				output = filepath.Join(config.Invoices.Directory, output)
			}
		}

		if err := builder.BuildAndWriteInvoice(invoiceInfo, output); err != nil {
			return err
		}

		fmt.Println("Invoice written to", output)
		return nil
	},
}
