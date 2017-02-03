package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestStreamingAddition(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{Input: "123456", Output: "6925"},
		{Input: "1a2bh3kn4lkn5lkj6", Output: "6925"},
		{Input: "12", Output: ""},
		{Input: "Only Letters", Output: ""},
	}

	for _, test := range tests {
		var out bytes.Buffer

		streamAddition(strings.NewReader(test.Input), &out)
		actual := out.String()
		if actual != test.Output {
			t.Errorf("Got: %s, want: %s for input %s", actual, test.Output, test.Input)
		}
	}
}
