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
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Requests a new Access Token from Swyftx using the API key.",
	Long: `
Request a new Access Token from Swyftx using the API key provided during authentication.

This command will fail if the API key has been added to the configuration file. To set the API key run:

	'swyftx-cli authenticate --apiKey <key>'

The Access Token expires after one week so this command must be run periodically.
`,
	RunE: refresher,
}

func init() {
	authenticateCmd.AddCommand(refreshCmd)
}

func refresher(cmd *cobra.Command, args []string) error {
	refreshToken()
	return nil
}

func refreshToken() {
	apiKey := viper.GetString("apikey")

	body := struct {
		ApiKey string `json:"apiKey"`
	}{ApiKey: apiKey}
	type Result struct {
		AccessToken string `json:"accessToken,omitempty"`
		Scopes      string `json:"scopes,omitempty"`
	}
	var result Result

	err := requests.
		URL("/auth/refresh/").
		Host(SwyftxAPI).
		ContentType("application/json").
		BodyJSON(&body).
		ToJSON(&result).
		Fetch(context.Background())
	if err != nil {
		fmt.Println("Access Token request failed")
		cobra.CheckErr(err)
	}

	viper.Set("token", result.AccessToken)
	if err := viper.WriteConfig(); err != nil {
		cobra.CheckErr(err)
	}

	fmt.Println("Access Token successfully saved. You can now use this application to call Swyftx.")
	fmt.Println("The token will expire in one week. To renew the token run 'swyftx-cli authenticate refresh'.")
}
