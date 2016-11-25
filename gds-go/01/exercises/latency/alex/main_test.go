package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLatency(t *testing.T) {
	tests := []struct {
		Path                 string
		ResponseCode         int
		ResponseDuration     time.Duration
		ResponseBodyContains string
	}{
		{
			Path:             "/latency",
			ResponseCode:     http.StatusOK,
			ResponseDuration: 500 * time.Millisecond,
		},
		{
			Path:             "/latency?duration=50ms",
			ResponseCode:     http.StatusOK,
			ResponseDuration: 50 * time.Millisecond,
		},
		{
			Path:                 "/latency?duration=wibble",
			ResponseCode:         http.StatusBadRequest,
			ResponseBodyContains: "invalid duration wibble",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.Path, nil)
		w := httptest.NewRecorder()
		startTime := time.Now()
		latencyHandler(w, req)
		actualDuration := time.Since(startTime)
		if w.Code != test.ResponseCode {
			t.Errorf("GET %s, want status %d. Got %d", test.Path, test.ResponseCode, w.Code)
		}
		if test.ResponseDuration != 0 && absDuration(actualDuration-test.ResponseDuration) > 10*time.Millisecond {
			t.Errorf("GET %s, want duration %s. Got %s", test.Path, test.ResponseDuration, actualDuration)
		}
		if test.ResponseBodyContains != "" && !strings.Contains(w.Body.String(), test.ResponseBodyContains) {
			t.Errorf("GET %s, expected to find '%s' in response body: '%s'", test.Path, test.ResponseBodyContains, w.Body.String())
		}
	}
}

func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
