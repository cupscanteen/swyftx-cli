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
)

var liveRatesAssetId string

var liveRatesCmd = &cobra.Command{
	Use:   "live",
	Short: "Current prices for all assets against asset. Currently only supported for AUD, USD and BTC",
	Long: `Current prices for an asset. Only AUD, USD and BTC comparison rates are supported.

Only accepts the following as valid ID's
	1:  AUD
	3:  BTC
	36: USD(t)
`,
	RunE: marketLiveRates,
}

func init() {
	marketsCmd.AddCommand(liveRatesCmd)

	// Asset to query
	liveRatesCmd.Flags().StringVarP(&liveRatesAssetId, "asset", "a", "1", "Asset by ID which should be queried. Must be in 'id' format, not name. Accepts 1:AUD, 3:BTC or 36:USD(t)")
	liveRatesCmd.MarkFlagRequired("asset")
	// Helpers
	liveRatesCmd.Flags().BoolVarP(&infoPretty, "pretty", "", false, "Pretty print the response")
	liveRatesCmd.Flags().StringVarP(&infoOutput, "output", "o", "csv", "Write the output to a file. Options: csv **coming soon**")
}

func marketLiveRates(cmd *cobra.Command, args []string) error {

	result, err := requestLiveRates()
	cobra.CheckErr(err)

	stdout := GenericPrinter(result, infoPretty)
	fmt.Println(stdout)
	return nil
}

func requestLiveRates() (LiveRates, error) {
	var result LiveRates
	url := fmt.Sprintf("/live-rates/%s/", liveRatesAssetId)
	err := requests.
		URL(url).
		Host(SwyftxAPI).
		ContentType("application/json").
		ToJSON(&result).
		AddValidator(StatusChecker).
		Fetch(context.Background())
	if err != nil {
		return LiveRates{}, err
	}
	return result, nil
}
