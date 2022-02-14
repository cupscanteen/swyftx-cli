/*
Copyright © 2022 Cupscanteen Industries

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	// Global flags
	cfgFile string
	debug   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "excellerate",
	Short: "excellerate is a command line interface (CLI) for the Swyftx API.",
	Long:  `excellerate is a command line interface (CLI) for the Swyftx API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.excellerate.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "Debug verbose output")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	configName := ".excellerate"
	configType := "yaml"

	configPath := fmt.Sprintf("%s/%s.%s", home, configName, configType)

	_, err = os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	cobra.CheckErr(err)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".excellerate" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if debug {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}

}