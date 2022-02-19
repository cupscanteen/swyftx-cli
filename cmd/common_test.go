package cmd

import (
	"testing"
)

func TestErrCheck401(t *testing.T) {
	testCases := []struct {
		name       string
		stringTest string
		expected   bool
	}{
		{
			"401 error check true",
			"status: 401",
			true,
		},
		{
			"401 error check false",
			"should fail",
			false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			CaptureStdout(func() {
				result := errCheck401(tt.stringTest)
				if result != tt.expected {
					t.Fatalf("expected %v but got %v", tt.expected, result)
				}
			})
		})
	}
}
