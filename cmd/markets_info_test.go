package cmd

import (
	"encoding/json"
	"github.com/carlmjohnson/requests"
	"testing"
)

const marketInfoTestdata = "testdata/marketinfo"

func Test_RequestBasicInfo(t *testing.T) {
	t.Cleanup(resetClient)
	client.Transport = requests.Replay(marketInfoTestdata)
	RecordNewTestdata(marketInfoTestdata)
	result, err := requestBasicInfo(&client)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func Test_MarketInfoBasicCmd(t *testing.T) {
	t.Cleanup(resetClient)
	client.Transport = requests.Replay(marketInfoTestdata)
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
