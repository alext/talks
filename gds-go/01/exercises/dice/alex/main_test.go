package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testCycles = 10

func TestRoll(t *testing.T) {
	tests := []struct {
		Path         string
		ResponseCode int
		Faces        int
	}{
		{
			Path:         "/roll",
			ResponseCode: http.StatusOK,
			Faces:        6,
		},
		{
			Path:         "/roll?die=D9",
			ResponseCode: http.StatusOK,
			Faces:        9,
		},
		{
			Path:         "/roll?die=wibble",
			ResponseCode: http.StatusBadRequest,
		},
		{
			Path:         "/roll?die=D-4",
			ResponseCode: http.StatusBadRequest,
		},
		{
			Path:         "/roll?die=D1",
			ResponseCode: http.StatusOK,
			Faces:        1,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.Path, nil)

		for i := 1; i <= testCycles; i += 1 {
			w := httptest.NewRecorder()
			rollHandler(w, req)
			if w.Code != test.ResponseCode {
				t.Errorf("GET %s (run %d), want status %d. Got %d", test.Path, i, test.ResponseCode, w.Code)
				continue
			}
			if test.Faces > 0 {
				var data map[string]int
				err := json.NewDecoder(w.Body).Decode(&data)
				if err != nil {
					t.Errorf("GET %s (run %d), error decoding JSON: %s", test.Path, i, err.Error())
					continue
				}
				if data["faces"] != test.Faces {
					t.Errorf("GET %s (run %d), want %d faces, got %d", test.Path, i, test.Faces, data["faces"])
					continue
				}
				if data["roll"] > test.Faces {
					t.Errorf("GET %s (run %d), want roll <= %d, got %d", test.Path, i, test.Faces, data["roll"])
					continue
				}
			}
		}
	}
}
