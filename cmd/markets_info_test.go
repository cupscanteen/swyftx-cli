package cmd

import (
	"encoding/json"
	"github.com/carlmjohnson/requests"
	"testing"
)

const marketInfoTestdata = "testdata/marketinfo"

func Test_MarketInfoBasicCmd(t *testing.T) {
	resetClient()
	resetInfoArgs()
	client.Transport = requests.Replay(marketInfoTestdata)
	RecordNewTestdata(marketInfoTestdata)
	stdOut := CaptureStdout(func() {
		rootCmd.SetArgs([]string{
			"markets",
			"info",
		})
		rootCmd.Execute()
	})
	if len(stdOut) == 0 {
		t.Fatal("failed to read JSON response correctly")
	}
	var i MarketsInfoBasicDTO
	err := json.Unmarshal([]byte(stdOut), &i)
	if err != nil {
		t.Fatalf("json unmarshall error: %q", err)
	}
}
func Test_MarketsInfoCmds(t *testing.T) {
	cases := []struct {
		name     string
		args     []string
		expected MarketsInfoBasic
	}{
		{
			name: "market info basic for ETH",
			args: []string{
				"markets",
				"info",
				"-a",
				"5",
			},
			expected: MarketsInfoBasic{
				Name:      "Ethereum",
				AltName:   "Ethereum",
				Code:      "ETH",
				ID:        5,
				Rank:      2,
				Buy:       "3754.37969042",
				Sell:      "3737.15431289",
				Spread:    "0.45",
				Volume24H: 10000686345,
				MarketCap: 325271200727,
			},
		},
	}
	for _, tc := range cases {
		resetClient()
		resetInfoArgs()
		t.Run(tc.name, func(t *testing.T) {
			client.Transport = requests.Replay(marketInfoTestdata)
			//RecordNewTestdata(marketInfoTestdata)
			stdOut := CaptureStdout(func() {
				rootCmd.SetArgs(tc.args)
				rootCmd.Execute()
			})
			var i MarketsInfoBasicDTO
			err := json.Unmarshal([]byte(stdOut), &i)
			if err != nil {
				t.Fatalf("json unmarshall error: %q", err)
			}
			if i[0] != tc.expected {
				FatalfFormatter(t, i[0], tc.expected)
			}
		})
	}
}
func resetInfoArgs() {
	infoAssetId = ""
	infoPretty = false
	infoOutput = "csv"
}
