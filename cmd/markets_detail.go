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

// detailCmd represents the detail command
var detailCmd = &cobra.Command{
	Use:   "detail",
	Short: "Detailed information about assets. If parameter left blank, returns all traded assets on Swyftx.",
	Long: `Retrieve detailed information about a tradable asset on Swyftx.

If no asset is provided, all assets are returned.`,
	RunE: marketInfoDetail,
}

func init() {
	marketsCmd.AddCommand(detailCmd)
	// Asset to query
	detailCmd.Flags().StringVarP(&infoAssetId, "asset", "a", "", "Asset by ID which should be queried. Must be in 'id' format, not name.")
	// Helpers
	detailCmd.Flags().BoolVarP(&infoPretty, "pretty", "", false, "Pretty print the response")
	detailCmd.Flags().StringVarP(&infoOutput, "output", "o", "csv", "Write the output to a file. Options: csv **coming soon**")

}
func marketInfoDetail(cmd *cobra.Command, args []string) error {
	result, err := requestInfoDetail(&client)
	cobra.CheckErr(err)

	stdout := GenericPrinter(result, infoPretty)
	fmt.Println(stdout)
	return nil
}

func requestInfoDetail(c *http.Client) (MarketsInfoDetailDTO, error) {
	var result MarketsInfoDetailDTO
	url := fmt.Sprintf("/markets/info/detail/%s/", infoAssetId)
	if infoAssetId == "" {
		url = "/markets/info/detail/"
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
		return MarketsInfoDetailDTO{}, err
	}
	return result, nil
}
