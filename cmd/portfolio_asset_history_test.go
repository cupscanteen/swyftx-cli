package cmd

import (
	"encoding/json"
	"github.com/carlmjohnson/requests"
	"testing"
)

const portfolioAssetHistoryTestdata = "testdata/portfolio"

func Test_PortfolioAssetHistoryCmd_Default(t *testing.T) {
	resetClient()
	resetPortfolioArgs()
	client.Transport = requests.Replay(portfolioAssetHistoryTestdata)
	RecordNewTestdata(portfolioAssetHistoryTestdata)
	stdOut := CaptureStdout(func() {
		rootCmd.SetArgs([]string{
			"portfolio",
			"asset",
			"history",
		})
		rootCmd.Execute()
	})
	if len(stdOut) == 0 {
		t.Fatal("failed to read JSON response correctly")
	}
	var i AssetHistoryAllDTO
	err := json.Unmarshal([]byte(stdOut), &i)
	if err != nil {
		t.Fatalf("json unmarshall error: %q", err)
	}
	btcAssetId := 3
	for _, x := range i.Items {
		if x.SecondaryAsset != btcAssetId {
			FatalfFormatter(t, x.SecondaryAsset, btcAssetId)
		}
	}
}
func Test_PortfolioAssetHistoryCmd_Args(t *testing.T) {
	cases := []struct {
		name     string
		args     []string
		expected Metadata
	}{
		{
			name: "default limit 20 of BTC",
			args: []string{
				"portfolio",
				"asset",
				"history",
			},
			expected: Metadata{
				CurrentPage:  1,
				PageSize:     20,
				FirstPage:    1,
				LastPage:     1,
				TotalRecords: 10,
			},
		},
		{
			name: "set limit of 5 for BTC",
			args: []string{
				"portfolio",
				"asset",
				"history",
				"--limit",
				"5",
			},
			expected: Metadata{
				CurrentPage:  1,
				PageSize:     5,
				FirstPage:    1,
				LastPage:     2,
				TotalRecords: 10,
			},
		},
		{
			name: "only BUY orders",
			args: []string{
				"portfolio",
				"asset",
				"history",
				"--order-type",
				"BUY",
			},
			expected: Metadata{
				CurrentPage:  1,
				PageSize:     20,
				FirstPage:    1,
				LastPage:     1,
				TotalRecords: 4,
			},
		},
		{
			name: "read second page of SELL results for ETH",
			args: []string{
				"portfolio",
				"asset",
				"history",
				"--page",
				"2",
				"--limit",
				"2",
				"--asset",
				"5",
			},
			expected: Metadata{
				CurrentPage:  2,
				PageSize:     2,
				FirstPage:    1,
				LastPage:     4,
				TotalRecords: 8,
			},
		},
	}
	for _, tc := range cases {
		resetClient()
		resetPortfolioArgs()
		t.Run(tc.name, func(t *testing.T) {
			client.Transport = requests.Replay(portfolioAssetHistoryTestdata)
			RecordNewTestdata(portfolioAssetHistoryTestdata)
			stdOut := CaptureStdout(func() {
				rootCmd.SetArgs(tc.args)
				rootCmd.Execute()
			})
			var i AssetHistoryAllDTO
			err := json.Unmarshal([]byte(stdOut), &i)
			if err != nil {
				t.Fatalf("json unmarshall error: %q", err)
			}
			if i.Metadata != tc.expected {
				FatalfFormatter(t, i.Metadata, tc.expected)
			}
		})
	}
}

func resetPortfolioArgs() {
	portfolioAssetId = "3"
	portfolioLimit = "20"
	portfolioPage = "1"
	portfolioSortKey = "date"
	portfolioSortDirection = "ASC"
	portfolioOrderType = ""
	portfolioOrderStatus = ""
	portfolioStartDate = ""
	portfolioEndDate = ""
	portfolioPretty = false
	portfolioOutput = "csv"
}
