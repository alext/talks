package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alext/talks/gds-go/01/exercises/dice/tim/dice"
)

func makeRequest(path string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(diceHandler{})
	handler.ServeHTTP(rr, req)

	return rr, nil
}

func TestDiceHandlerSuccess(t *testing.T) {
	var tests = []struct {
		qs     string
		status int
	}{
		{"", http.StatusOK},
		{"?dice=", http.StatusOK},
		{"?dice=d6", http.StatusOK},
		{"?dice=d6,d10", http.StatusOK},
		{"?dice=x", http.StatusBadRequest},
	}

	for _, test := range tests {
		res, err := makeRequest(test.qs)
		if err != nil {
			t.Fatal(err)
		}

		if status := res.Code; status != test.status {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, test.status)
		}
	}
}

func TestDiceHandlerFailure(t *testing.T) {
	res, err := makeRequest("?dice=x")
	if err != nil {
		t.Fatal(err)
	}

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "{\"error\":\"x is not a valid dice definition\"}\n"
	if res.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

type fixedRoll struct {
	dice []int
}

func (r fixedRoll) Roll() dice.RollResult {
	rr := dice.RollResult{Value: 0, Dice: []dice.DieResult{}}
	for _, die := range r.dice {
		rr.Value += die
		t := fmt.Sprintf("d%d", die)
		rr.Dice = append(rr.Dice, dice.DieResult{Type: t, Value: die})
	}
	return rr
}

func TestDoRoll(t *testing.T) {
	var tests = []struct {
		dice []int
		json string
	}{
		{[]int{}, `{"roll":0,"dice":[]}`},
		{[]int{6}, `{"roll":6,"dice":[{"type":"d6","roll":6}]}`},
		{[]int{6, 6}, `{"roll":12,"dice":[{"type":"d6","roll":6},{"type":"d6","roll":6}]}`},
	}

	for _, test := range tests {
		r := fixedRoll{dice: test.dice}
		res := httptest.NewRecorder()
		doRoll(r, res)

		if res.Body.String() != test.json {
			t.Errorf("doRoll returned unexpected body: got %v want %v",
				res.Body.String(), test.json)
		}
	}
}
