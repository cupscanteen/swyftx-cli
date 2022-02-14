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

// marketsCmd represents the markets command
var marketsCmd = &cobra.Command{
	Use:   "markets",
	Short: "The subcommands of the markets Swyftx endpoints",
	Long: `Markets endpoints for Swyftx's API endpoints. 

This includes basic and detailed asset information, all assets and the market live rates per asset.
`,
}

func init() {
	rootCmd.AddCommand(marketsCmd)
}
