package cmd

import (
	"bytes"
	"github.com/carlmjohnson/requests"
	"io"
	"net/http"
	"os"
	"testing"
)

// resetClient is used to return the http.Client to its default state. This is
// required because in testing we use requests.Replay() to replace the global
// client variable
func resetClient() {
	client = http.Client{}
}

// CaptureStdout captures prints so that we can inspect and remove them from logs
func CaptureStdout(f func()) string {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	defer r.Close()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdOut

	var b bytes.Buffer
	io.Copy(&b, r)

	return b.String()
}

// RecordNewTestdata is used refresh out testdata with up-to-date data
func RecordNewTestdata(testdata string) {
	if os.Getenv("SWYFTX_TESTDATA_RECORD") != "" {
		println("RECORDING NEW TESTDATA")
		client.Transport = requests.Record(nil, testdata)
	}
}

// FatalfFormatter is a helper that prints test comparison asserts nicely
func FatalfFormatter(t *testing.T, got, want interface{}) {
	t.Fatalf("\nwant %q\ngot  %#v", want, got)
}
