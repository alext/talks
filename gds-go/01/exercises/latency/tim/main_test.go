package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeRequest(path string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(latencyHandler{})
	handler.ServeHTTP(rr, req)

	return rr, nil
}

func TestLatencyHandler(t *testing.T) {
	var tests = []struct {
		qs     string
		status int
		body   string
		wait   time.Duration
	}{
		{"", http.StatusOK, "OK", defaultWait},
		{"?duration=", http.StatusOK, "OK", defaultWait},
		{"?duration=1ms", http.StatusOK, "OK", time.Millisecond},
		{"?duration=100ms", http.StatusOK, "OK", time.Millisecond * 100},
		{"?duration=1s", http.StatusOK, "OK", time.Second},
		{"?duration=x", http.StatusBadRequest, "time: invalid duration x", 0},
	}

	for _, test := range tests {
		start := time.Now()

		res, err := makeRequest(test.qs)
		if err != nil {
			t.Fatal(err)
		}

		slept := time.Now().Sub(start)
		if slept < test.wait {
			t.Errorf("Handler only slept for %v; expected %v",
				slept, test.wait)
		}

		if status := res.Code; status != test.status {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, test.status)
		}

		if res.Body.String() != test.body {
			t.Errorf("Handler returned unexpected body: got %v want %v",
				res.Body.String(), test.body)
		}
	}
}

func TestGetWaitTime(t *testing.T) {
	h := latencyHandler{}

	var tests = []struct {
		input    string
		duration time.Duration
		isErr    bool
	}{
		{"", defaultWait, false},
		{"1ms", time.Millisecond, false},
		{"100ms", time.Millisecond * 100, false},
		{"1s", time.Second, false},
		{"x", 0, true},
	}

	for _, test := range tests {
		d, err := h.getWaitTime(test.input)

		if (err != nil) != test.isErr {
			t.Errorf("getWaitTime returned unexpected error status: got %v want %v",
				(err != nil), test.isErr)
		}

		if d != test.duration {
			t.Errorf("getWaitTime returned unexpected duration: got %v want %v",
				d, test.duration)
		}
	}
}
