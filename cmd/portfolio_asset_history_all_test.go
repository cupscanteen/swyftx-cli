package cmd

import (
	"encoding/json"
	"github.com/carlmjohnson/requests"
	"testing"
)

const portfolioAssetHistoryAllTestdata = "testdata/portfolio/all"

func Test_PortfolioAssetHistoryAllCmd_Default(t *testing.T) {
	resetClient()
	resetPortfolioArgs()
	client.Transport = requests.Replay(portfolioAssetHistoryAllTestdata)
	RecordNewTestdata(portfolioAssetHistoryAllTestdata)
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
func Test_PortfolioAssetHistoryAllCmd_Args(t *testing.T) {

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
				"all",
			},
			expected: Metadata{
				CurrentPage:  1,
				PageSize:     20,
				FirstPage:    1,
				LastPage:     215,
				TotalRecords: 4291,
			},
		},
		{
			name: "set limit of 5 for BTC",
			args: []string{
				"portfolio",
				"asset",
				"history",
				"all",
				"--limit",
				"5",
			},
			expected: Metadata{
				CurrentPage:  1,
				PageSize:     5,
				FirstPage:    1,
				LastPage:     859,
				TotalRecords: 4291,
			},
		},
		{
			name: "only BUY orders",
			args: []string{
				"portfolio",
				"asset",
				"history",
				"all",
				"--order-type",
				"BUY",
			},
			expected: Metadata{
				CurrentPage:  1,
				PageSize:     20,
				FirstPage:    1,
				LastPage:     144,
				TotalRecords: 2876,
			},
		},
		{
			name: "read second page of SELL results which have COMPLETED",
			args: []string{
				"portfolio",
				"asset",
				"history",
				"all",
				"--page",
				"2",
				"--limit",
				"2",
				"--order-status",
				"COMPLETED",
			},
			expected: Metadata{
				CurrentPage:  2,
				PageSize:     2,
				FirstPage:    1,
				LastPage:     515,
				TotalRecords: 1030,
			},
		},
	}
	for _, tc := range cases {
		resetClient()
		resetPortfolioArgs()
		t.Run(tc.name, func(t *testing.T) {
			client.Transport = requests.Replay(portfolioAssetHistoryAllTestdata)
			RecordNewTestdata(portfolioAssetHistoryAllTestdata)
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
