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
	"net/http"
)

var marketsInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Basic information about assets. If left blank will show **all** assets traded on Swyftx.",
	Long: `Retrieve basic information about a tradable asset on Swyftx. 

If no asset is provided, all assets are returned.
`,
	RunE: marketInfoBasic,
}

func init() {
	marketsCmd.AddCommand(marketsInfoCmd)
	// Asset to query
	marketsInfoCmd.Flags().StringVarP(&infoAssetId, "asset", "a", "", "Asset by ID which should be queried. Must be in 'id' format, not name.")
	// Helpers
	marketsInfoCmd.Flags().BoolVarP(&infoPretty, "pretty", "", false, "Pretty print the response")
	marketsInfoCmd.Flags().StringVarP(&infoOutput, "output", "o", "csv", "Write the output to a file. Options: csv **coming soon**")
}
func marketInfoBasic(cmd *cobra.Command, args []string) error {
	result, err := requestBasicInfo(&client)
	cobra.CheckErr(err)

	stdout := GenericPrinter(result, infoPretty)
	fmt.Println(stdout)
	return nil
}

func requestBasicInfo(c *http.Client) (MarketsInfoBasicDTO, error) {
	var result MarketsInfoBasicDTO
	url := fmt.Sprintf("/markets/info/basic/%s/", infoAssetId)
	if infoAssetId == "" {
		url = "/markets/info/basic/"
	}
	err := requests.
		URL(url).
		Client(c).
		Host(SwyftxAPI).
		ContentType("application/json").
		ToJSON(&result).
		AddValidator(StatusChecker).
		Fetch(context.Background())
	if err != nil {
		return MarketsInfoBasicDTO{}, err
	}
	return result, nil
}
