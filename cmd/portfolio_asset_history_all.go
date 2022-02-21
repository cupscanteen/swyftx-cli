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
	"strconv"
)

// allCmd represents the all command
var portfolioAssetHistoryAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Retrieve buy/sell/withdraw/deposit history of *all* assets",
	Long: `Retrieve portfolio asset history for all assets. This **does not** filter by asset. To filter by a single asset run this command without the 'all' at the end.
`,
	RunE: portfolioAssetsHistoryAll,
}

func init() {
	portfolioAssetHistoryCmd.AddCommand(portfolioAssetHistoryAllCmd)
	// Query params
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioLimit, "limit", "l", "20", "Number of orders to display in a single page")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioPage, "page", "p", "1", "Page to display of paginated orders")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioSortKey, "sort-key", "k", "date", "Return orders by {date,amount,userCountryValue}. Defaults to date")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioSortDirection, "sort-direction", "d", "ASC", "Sort direction, default is ASC. Options: {ASC|DSC}")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioOrderType, "order-type", "t", "", "Returns only a single order type. Defaults to all. Options: {BUY,SELL,WITHDRAWAL,DEPOSIT}")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioOrderStatus, "order-status", "s", "", "Returns only a single order by status. Defaults to all status types. Options: {PENDING, COMPLETED, FAILED}")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioStartDate, "start-date", "", "", "Only return results from after a timestamp. Expects unix time. Defaults to none.")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioEndDate, "end-date", "", "", "Only return results from after a timestamp. Expects unix time. Defaults to none.")
	// Output options
	portfolioAssetHistoryAllCmd.Flags().BoolVarP(&portfolioPretty, "pretty", "", false, "Pretty print the response")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&portfolioOutput, "output", "o", "csv", "Write the output to a file. Options: csv **coming soon**")
}

func portfolioAssetsHistoryAll(cmd *cobra.Command, args []string) error {
	token, err := AccessTokenGetter()
	cobra.CheckErr(err)

	result, err := requestAllAssets(token, &client)
	cobra.CheckErr(err)

	stdout := GenericPrinter(result, portfolioPretty)
	fmt.Println(stdout)

	return nil
}

func requestAllAssets(token string, c *http.Client) (AssetHistoryAllDTO, error) {
	var result AssetHistoryAllDTO
	err := requests.
		URL("/portfolio/assethistory/all/").
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
		AddValidator(StatusChecker).
		Fetch(context.Background())
	if err != nil {
		return AssetHistoryAllDTO{}, nil
	}
	page, _ := strconv.Atoi(portfolioPage)
	pageSize, _ := strconv.Atoi(portfolioLimit)
	metadata := CalculateMetadata(result.RecordCount, page, pageSize)
	result.Metadata = metadata
	result.RecordCount = 0
	return result, nil
}
