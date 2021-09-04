package cmd

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/tinyzimmer/ginvoicer/pkg/types"
)

var configFile string

func init() {
	rootCommand.PersistentFlags().StringVarP(&configFile, "config", "c", types.DefaultConfigPath, "path to a configuration file")
}

func Execute() error { return rootCommand.Execute() }

var rootCommand = &cobra.Command{
	Use:          "ginvoicer",
	Short:        "Invoice generator for independent contractors",
	SilenceUsage: true,
}

func loadConfig() (*types.Config, error) {
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	config := &types.Config{}
	return config, yaml.Unmarshal(configBytes, config)
}
