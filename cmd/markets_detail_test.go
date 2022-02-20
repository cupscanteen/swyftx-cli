package cmd

import (
	"encoding/json"
	"github.com/carlmjohnson/requests"
	"testing"
)

const marketDetailTestdata = "testdata/marketdetail"

func Test_RequestBasicInfo(t *testing.T) {
	resetClient()
	client.Transport = requests.Replay(marketInfoTestdata)
	RecordNewTestdata(marketInfoTestdata)
	result, err := requestInfoDetail(&client)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func Test_MarketsDetailCmdArgs(t *testing.T) {
	cases := []struct {
		name     string
		args     []string
		expected MarketsInfoDetail
	}{
		{
			name: "market info detail for ETH",
			args: []string{
				"markets",
				"detail",
				"-a",
				"5",
			},
			expected: MarketsInfoDetail{
				ID:          5,
				Description: "Ethereum is a <a href=\"https://www.coingecko.com/en?category_id=29&view=market\">smart contract platform</a> that enables developers to build tokens and decentralized applications (dapps). ETH is the native currency for the Ethereum platform and also works as the transaction fees to miners on the Ethereum network.\r\n\r\nEthereum is the pioneer for blockchain based smart contracts. Smart contract is essentially a computer code that runs exactly as programmed without any possibility of downtime, censorship, fraud or third-party interference. It can facilitate the exchange of money, content, property, shares, or anything of value. When running on the blockchain a smart contract becomes like a self-operating computer program that automatically executes when specific conditions are met.\r\n\r\nEthereum allows programmers to run complete-turing smart contracts that is capable of any customizations. Rather than giving a set of limited operations, Ethereum allows developers to have complete control over customization of their smart contract, giving developers the power to build unique and innovative applications.\r\n\r\nEthereum being the first blockchain based smart contract platform, they have gained much popularity, resulting in new competitors fighting for market share. The competitors includes: <a href=\"https://www.coingecko.com/en/coins/ethereum_classic\">Ethereum Classic</a> which is the oldchain of Ethereum, <a href=\"https://www.coingecko.com/en/coins/qtum\">Qtum</a>, <a href=\"https://www.coingecko.com/en/coins/eos\">EOS</a>, <a href=\"https://www.coingecko.com/en/coins/neo\">Neo</a>, <a href=\"https://www.coingecko.com/en/coins/icon\">Icon</a>, <a href=\"https://www.coingecko.com/en/coins/tron\">Tron</a> and <a href=\"https://www.coingecko.com/en/coins/cardano\">Cardano</a>.\r\n\r\nEthereum wallets are fairly simple to set up with multiple popular choices such as myetherwallet, <a href=\"https://www.coingecko.com/buzz/complete-beginners-guide-to-metamask?locale=en\">metamask</a>, and <a href=\"https://www.coingecko.com/buzz/trezor-model-t-wallet-review\">Trezor</a>. Read here for more guide on using ethereum wallet: <a href=\"https://www.coingecko.com/buzz/how-to-use-an-ethereum-wallet\">How to Use an Ethereum Wallet</a>",
				Name:        "Ethereum",
				Category:    "Smart Contract Platform",
				Mineable:    true,
				Spread:      "0.46",
				Rank:        2,
				RankSuffix:  "nd",
				Volume: struct {
					Two4H float64 `json:"24H"`
				}{
					Two4H: 1.1559285651e+10,
				},
				PriceChange: struct {
					Week  float64 `json:"week"`
					Month float64 `json:"month"`
				}{
					Week:  -8.85215,
					Month: -11.79925,
				},
				Urls: struct {
					Explorer string `json:"explorer"`
					Reddit   string `json:"reddit"`
					Twitter  string `json:"twitter"`
					Website  string `json:"website"`
				}{
					Explorer: "https://etherscan.io/",
					Reddit:   "https://www.reddit.com/r/ethereum",
					Twitter:  "https://twitter.com/ethereum",
					Website:  "https://www.ethereum.org/",
				},
				Supply: struct {
					Circulating float64 `json:"circulating"`
					Total       float64 `json:"total"`
					Max         float64 `json:"max"`
				}{
					Circulating: 1.196554476865e+08,
					Total:       0,
					Max:         0,
				},
			},
		},
	}
	for _, tc := range cases {
		resetClient()
		resetInfoArgs()
		t.Run(tc.name, func(t *testing.T) {
			client.Transport = requests.Replay(marketDetailTestdata)
			RecordNewTestdata(marketDetailTestdata)
			stdOut := CaptureStdout(func() {
				rootCmd.SetArgs(tc.args)
				rootCmd.Execute()
			})
			var i MarketsInfoDetailDTO
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
