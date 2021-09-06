/*

Copyright (C) 2021 Avi Zimmerman

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program; if not, write to the Free Software Foundation,
Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

*/

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
