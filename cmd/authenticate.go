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
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var authApiKey string

func init() {
	rootCmd.AddCommand(authenticateCmd)

	authenticateCmd.Flags().StringVarP(&authApiKey, "apikey", "a", "", "Swyftx API Key used to fetch Access Tokens")
	authenticateCmd.MarkFlagRequired("apikey")
}

// authenticateCmd represents the authenticate command
var authenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Registers a Swyftx API Key for use throughout the application.",
	Long: `
This command allows you to register a Swyftx API key for use in this CLI. You can get an API key from the Swyftx Dashboard.

The API key will be saved locally to your device and used to fetch refresh tokens from Swyftx for use in other endpoints.

Excellerate must have an API Key registered before it can access any Swyftx endpoints.
`,
	Example: "excellerate authenticate --apikey abc123",
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("authenticate called with " + authApiKey)
	//},
	RunE: authenticate,
}

func authenticate(cmd *cobra.Command, args []string) error {
	err := writeApiKeyToConfig(authApiKey)
	cobra.CheckErr(err)

	return nil
}

func writeApiKeyToConfig(apiKey string) error {
	viper.Set("apikey", apiKey)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	refreshToken()
	return nil
}
