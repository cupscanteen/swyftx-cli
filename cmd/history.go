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

var (
	// Query params
	page          string
	limit         string
	sortKey       string
	sortDirection string
	startDate     string
	endDate       string
	orderType     string
	orderStatus   string

	// Output options
	// pretty will format to Stdout with prettified tab spacing
	pretty bool
	// fileType determines the output file type when using the --output option
	fileType string
)

func init() {
	ordersCmd.AddCommand(historyCmd)
	// Query params
	historyCmd.Flags().StringVarP(&limit, "limit", "l", "20", "Number of orders to display in a single page")
	historyCmd.Flags().StringVarP(&page, "page", "p", "1", "Page to display of paginated orders")
	historyCmd.Flags().StringVarP(&sortKey, "sort-key", "k", "date", "Return orders by {date,amount,userCountryValue}. Defaults to date")
	historyCmd.Flags().StringVarP(&sortDirection, "sort-direction", "d", "ASC", "Sort direction, default is ASC. Options: {ASC|DSC}")
	historyCmd.Flags().StringVarP(&orderType, "order-type", "t", "", "Returns only a single order type. Defaults to all. Options: {BUY,SELL,WITHDRAWAL,DEPOSIT}")
	historyCmd.Flags().StringVarP(&orderStatus, "order-status", "s", "", "Returns only a single order by status. Defaults to all status types. Options: {PENDING, COMPLETED, FAILED}")
	historyCmd.Flags().StringVarP(&startDate, "start-date", "", "", "Only return results from after a timestamp. Expects unix time. Defaults to none.")
	historyCmd.Flags().StringVarP(&endDate, "end-date", "", "", "Only return results from after a timestamp. Expects unix time. Defaults to none.")
	// Output options
	historyCmd.Flags().BoolVarP(&pretty, "pretty", "", false, "Pretty print the response")
	historyCmd.Flags().StringVarP(&fileType, "filetype", "f", "csv", "The filetype to generate")
}

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Get all assets history for Swyftx",
	Long: `
Get the buy/sell/withdraw/deposit history of all assets. This command takes many optional parameters
`,
	RunE: orderHistory,
}

func orderHistory(cmd *cobra.Command, args []string) error {
	token := viper.GetString("token")

	result, err := requester(token)
	cobra.CheckErr(err)

	// Print to stdout
	stdout := assetPrinter(result, pretty)
	fmt.Println(stdout)

	return nil
}

func assetPrinter(result AssetHistoryAll, prettify bool) string {
	if prettify {
		return prettyPrint(result)
	}
	return printer(result)
}

func requester(token string) (AssetHistoryAll, error) {
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
		os.Exit(-1)
	}
	return result, nil
}

type AssetHistoryAll struct {
	Items []struct {
		Date             int64  `json:"date"`
		Amount           string `json:"amount"`
		Movement         string `json:"movement"`
		ApproxMovement   string `json:"approxMovement"`
		UserCountryValue string `json:"userCountryValue"`
		UUID             string `json:"uuid"`
		Type             string `json:"type"`
		Status           string `json:"status"`
		StatusRaw        int    `json:"statusRaw"`
		OrderType        int    `json:"orderType"`
		SecondaryAsset   int    `json:"secondaryAsset"`
		PrimaryAsset     int    `json:"primaryAsset"`
		SecondaryAmount  string `json:"secondaryAmount"`
		PrimaryAmount    string `json:"primaryAmount"`
	} `json:"items"`
	RecordCount int `json:"recordCount"`
}
