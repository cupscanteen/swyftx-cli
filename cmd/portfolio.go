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
	"github.com/spf13/cobra"
)

var (
	// Query params
	portfolioAssetId       string
	portfolioPage          string
	portfolioLimit         string
	portfolioSortKey       string
	portfolioSortDirection string
	portfolioStartDate     string
	portfolioEndDate       string
	portfolioOrderType     string
	portfolioOrderStatus   string

	// pretty will format to Stdout with prettified tab spacing
	portfolioPretty bool
	// fileType determines the output file type when using the --output option
	portfolioOutput string
)

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Portfolio endpoints for Swyftx",
	Long: `These sub-commands will allow you to qeury the portfolio endpoints for Swyftx orders and price history by asset.
`,
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
}
