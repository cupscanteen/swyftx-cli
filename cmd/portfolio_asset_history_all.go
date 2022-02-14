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
	"os"
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
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&limit, "limit", "l", "20", "Number of orders to display in a single page")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&page, "page", "p", "1", "Page to display of paginated orders")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&sortKey, "sort-key", "k", "date", "Return orders by {date,amount,userCountryValue}. Defaults to date")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&sortDirection, "sort-direction", "d", "ASC", "Sort direction, default is ASC. Options: {ASC|DSC}")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&orderType, "order-type", "t", "", "Returns only a single order type. Defaults to all. Options: {BUY,SELL,WITHDRAWAL,DEPOSIT}")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&orderStatus, "order-status", "s", "", "Returns only a single order by status. Defaults to all status types. Options: {PENDING, COMPLETED, FAILED}")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&startDate, "start-date", "", "", "Only return results from after a timestamp. Expects unix time. Defaults to none.")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&endDate, "end-date", "", "", "Only return results from after a timestamp. Expects unix time. Defaults to none.")
	// Output options
	portfolioAssetHistoryAllCmd.Flags().BoolVarP(&pretty, "pretty", "", false, "Pretty print the response")
	portfolioAssetHistoryAllCmd.Flags().StringVarP(&output, "output", "o", "csv", "Write the output to a file. Options: csv **coming soon**")
}

func portfolioAssetsHistoryAll(cmd *cobra.Command, args []string) error {
	token := viper.GetString("token")

	result, err := requestAllAssets(token)
	cobra.CheckErr(err)

	// Print to stdout
	stdout := assetPrinter(result, pretty)
	fmt.Println(stdout)

	return nil
}

func requestAllAssets(token string) (AssetHistoryAll, error) {
	var result AssetHistoryAll
	err := requests.
		URL("/portfolio/assethistory/all/").
		Param("limit", limit).
		Param("page", page).
		Param("type", orderType).
		Param("status", orderStatus).
		Param("sortDirection", sortDirection).
		Param("sortKey", sortKey).
		Param("startDate", startDate).
		Param("endDate", endDate).
		Host(SwyftxAPI).
		ContentType("application/json").
		Header("Authorization", fmt.Sprintf("Bearer %s", token)).
		ToJSON(&result).
		CheckStatus(200).
		Fetch(context.Background())
	if err != nil {
		errCheck401(err.Error())
		if !errCheck401(err.Error()) {
			fmt.Println(err.Error())
		}
		os.Exit(-1)
	}
	return result, nil
}
