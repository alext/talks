package main

import (
	"fmt"
	"net/http"
	"time"
)

// - GET `/latency` - returns the string "OK" after 500ms
// - GET `/latency?duration=100ms` - override the default
//   delay duration. The value can be any string accepted
//   by `time.ParseDuration`.
// - An invalid duration parameter should return a 400
//   along with an error message.

const defaultWait = time.Millisecond * 500

type latencyHandler struct{}

func (l latencyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wait, err := l.getWaitTime(r.FormValue("duration"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	time.Sleep(wait)
	fmt.Fprint(w, "OK")
}

func (latencyHandler) getWaitTime(duration string) (time.Duration, error) {
	if duration != "" {
		return time.ParseDuration(duration)
	}
	return defaultWait, nil
}

func main() {
	fmt.Println("Listening on :8080")
	http.Handle("/latency", latencyHandler{})
	http.ListenAndServe(":8080", nil)
}
