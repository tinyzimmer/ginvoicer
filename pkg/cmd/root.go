package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/tinyzimmer/ginvoicer/pkg/invoice"
	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

var (
	outFile       string
	configFile    string
	payer         string
	invoiceNumber int
	billables     map[string]int
	dueIn         time.Duration
)

func init() {
	flags := generateCommand.Flags()

	flags.StringVarP(&configFile, "config", "c", types.DefaultConfigPath, "path to a configuration file")
	flags.StringVarP(&payer, "payer", "p", "", "payer (by name or alias) to generate the invoice for")
	flags.IntVarP(&invoiceNumber, "number", "n", -1, "invoice number (default: auto-increment)")
	flags.DurationVarP(&dueIn, "due-in", "d", time.Hour*24*7, "time until invoice is due")
	flags.StringVarP(&outFile, "output", "o", "", "path to write the invoice to")
	flags.StringToIntVarP(&billables, "item", "i", map[string]int{}, "items to add to the invoice in the format of alias=quantity")

	cobra.MarkFlagRequired(flags, "payer")

	rootCommand.AddCommand(generateCommand)
}

func Execute() error { return rootCommand.Execute() }

var rootCommand = &cobra.Command{
	Use:          "invoicer",
	Short:        "Invoice generator for freelancers",
	SilenceUsage: true,
}

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generate an invoice",
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

		invoiceNumberStr := config.FormatInvoiceNumber(invoiceNumber)

		invoiceInfo := &types.InvoiceDetails{
			InvoiceNumber: invoiceNumberStr,
			InvoiceDate:   time.Now(),
			Payee:         config.Payee,
			Payer:         payerEntity,
			Items:         invoiceItems,
			DueDate:       time.Now().Add(dueIn),
		}

		if err := builder.BuildInvoice(invoiceInfo); err != nil {
			return err
		}

		output := outFile
		if output == "" {
			output = fmt.Sprintf("%s.pdf", invoiceNumberStr)
		}
		if config.InvoiceDirectory != "" {
			output = filepath.Join(config.InvoiceDirectory, output)
		}

		if err := builder.WriteFile(output); err != nil {
			return err
		}

		fmt.Println("Invoice written to", output)
		return nil
	},
}

func loadConfig() (*types.Config, error) {
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	config := &types.Config{}
	return config, yaml.Unmarshal(configBytes, config)
}
