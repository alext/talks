package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRot13(t *testing.T) {
	tests := []struct {
		Input  rune
		Output rune
	}{
		{Input: 'a', Output: 'n'},
		{Input: 'A', Output: 'N'},
		{Input: 'm', Output: 'z'},
		{Input: 'n', Output: 'a'},
		{Input: 'z', Output: 'm'},
		{Input: ' ', Output: ' '},
		{Input: 'é', Output: 'é'},
	}

	for _, test := range tests {
		actual := rot13(test.Input)
		if actual != test.Output {
			t.Errorf("Got: %c, want: %c for input %c", actual, test.Output, test.Input)
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{Input: "Hello World", Output: "Uryyb Jbeyq"},
	}

	for _, test := range tests {
		var out bytes.Buffer

		rotate(strings.NewReader(test.Input), &out)
		actual := out.String()
		if actual != test.Output {
			t.Errorf("Got: %s, want: %s for input %s", actual, test.Output, test.Input)
		}
	}
}
