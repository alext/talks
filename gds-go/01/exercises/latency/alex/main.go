package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const defaultDuration = 500 * time.Millisecond

func parseDuration(durationParam string) (time.Duration, error) {
	if durationParam == "" {
		return defaultDuration, nil
	}
	return time.ParseDuration(durationParam)
}

func latencyHandler(w http.ResponseWriter, req *http.Request) {
	duration, err := parseDuration(req.FormValue("duration"))
	if err != nil {
		http.Error(w, "Invalid duration: "+err.Error(), http.StatusBadRequest)
		return
	}
	time.Sleep(duration)
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/latency", latencyHandler)
	err := http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
