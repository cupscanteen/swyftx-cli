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
	"os"
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
}
func marketInfoDetail(cmd *cobra.Command, args []string) error {

	result, err := requestInfoDetail()
	cobra.CheckErr(err)

	stdout := detailInfoPrinter(result, infoPretty)
	fmt.Println(stdout)
	return nil
}

func requestInfoDetail() (MarketsInfoDetail, error) {
	var result MarketsInfoDetail
	url := fmt.Sprintf("/markets/info/detail/%s/", infoAssetId)
	if infoAssetId == "" {
		url = "/markets/info/detail/"
	}
	err := requests.
		URL(url).
		Host(SwyftxAPI).
		ContentType("application/json").
		ToJSON(&result).
		CheckStatus(200).
		Fetch(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	return result, nil
}
