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
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var portfolioAssetHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Retrieve the buy/sell/withdraw/deposit history of a *single* asset",
	Long: `Retrieve portfolio asset history for a **single** asset. The asset must be in its numerical 'id' state.

To ascertain the 'id' of a given asset from its human-readable name you can run the following search:
	swyftx-cli markets info basic BTC

Defaults to getting the asset history for BTC if no '--asset' flag is supplied
`,
	RunE: portfolioAssetHistorySingle,
}

func init() {
	portfolioAssetCmd.AddCommand(portfolioAssetHistoryCmd)
	// Asset to query
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioAssetId, "asset", "a", "3", "Asset by ID which should be queried. Must be in 'id' format, not name.")
	// Query params
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioLimit, "limit", "l", "20", "Number of orders to display in a single page")
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioPage, "page", "p", "1", "Page to display of paginated orders")
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioSortKey, "sort-key", "k", "date", "Return orders by {date,amount,userCountryValue}. Defaults to date")
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioSortDirection, "sort-direction", "d", "ASC", "Sort direction, default is ASC. Options: {ASC|DSC}")
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioOrderType, "order-type", "t", "", "Returns only a single order type. Defaults to all. Options: {BUY,SELL,WITHDRAWAL,DEPOSIT}")
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioOrderStatus, "order-status", "s", "", "Returns only a single order by status. Defaults to all status types. Options: {PENDING, COMPLETED, FAILED}")
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioStartDate, "start-date", "", "", "Only return results from after a timestamp. Expects unix time. Defaults to none.")
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioEndDate, "end-date", "", "", "Only return results from after a timestamp. Expects unix time. Defaults to none.")
	// Output options
	portfolioAssetHistoryCmd.Flags().BoolVarP(&portfolioPretty, "pretty", "", false, "Pretty print the response")
	portfolioAssetHistoryCmd.Flags().StringVarP(&portfolioOutput, "output", "o", "csv", "Write the output to a file. Options: csv **coming soon**")

}

func portfolioAssetHistorySingle(cmd *cobra.Command, args []string) error {
	token, err := AccessTokenGetter()
	cobra.CheckErr(err)

	result, err := requestSingleAsset(token, &client)
	cobra.CheckErr(err)

	stdout := GenericPrinter(result, portfolioPretty)
	fmt.Println(stdout)

	return nil
}

func requestSingleAsset(token string, c *http.Client) (AssetHistoryAllDTO, error) {
	var result AssetHistoryAllDTO
	err := requests.
		URL(fmt.Sprintf("/portfolio/assethistory/%s/", portfolioAssetId)).
		Client(c).
		Param("limit", portfolioLimit).
		Param("page", portfolioPage).
		Param("type", portfolioOrderType).
		Param("status", portfolioOrderStatus).
		Param("sortDirection", portfolioSortDirection).
		Param("sortKey", portfolioSortKey).
		Param("startDate", portfolioStartDate).
		Param("endDate", portfolioEndDate).
		Host(SwyftxAPI).
		ContentType("application/json").
		Header("Authorization", fmt.Sprintf("Bearer %s", token)).
		ToJSON(&result).
		CheckStatus(200).
		Fetch(context.Background())
	if err != nil {
		if !errCheck401(err.Error()) {
			fmt.Println(err.Error())
		}
		os.Exit(-1)
	}
	page, _ := strconv.Atoi(portfolioPage)
	pageSize, _ := strconv.Atoi(portfolioLimit)
	metadata := CalculateMetadata(result.RecordCount, page, pageSize)
	result.Metadata = metadata
	// remove RecordCount as we now expost Metadata
	result.RecordCount = 0
	log.Printf("%#v", result.Metadata)
	return result, nil
}
