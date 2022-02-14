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

// infoCmd represents the info command
var marketsInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Basic information about assets. If left blank will show **all** assets traded on Swyftx.",
	Long: `Retrieve basic information about a tradable asset on Swyftx. 

If no asset is provided, will return all tradable assets on Swyftx.
`,
}

func init() {
	marketsCmd.AddCommand(marketsInfoCmd)
	// Asset to query
	marketsInfoCmd.Flags().StringVarP(&assetId, "asset", "a", "3", "Asset by ID which should be queried. Must be in 'id' format, not name.")
	// Helpers
	marketsInfoCmd.Flags().BoolVarP(&pretty, "pretty", "", false, "Pretty print the response")
	marketsInfoCmd.Flags().StringVarP(&output, "output", "o", "csv", "Write the output to a file. Options: csv **coming soon**")
}
