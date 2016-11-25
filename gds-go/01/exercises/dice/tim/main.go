package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alext/talks/gds-go/01/exercises/dice/tim/dice"
)

// - GET `/roll` generates a random roll from a D6 die.
// - GET `/roll?die=D<N>` (where N is a positive integer)
// - Both calls return JSON including the die type, and
//   the random roll from the die.
// - Invalid parameters return a 400 and suitable JSON
//   encoded error message.

type diceHandler struct{}

func (d diceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	roll, err := dice.NewRoll(r.FormValue("dice"))
	if err != nil {
		outputError(err, w)
		return
	}

	doRoll(roll, w)
}

func doRoll(r dice.Roller, w http.ResponseWriter) {
	res := r.Roll()

	json, err := json.Marshal(res)
	if err != nil {
		outputError(err, w)
		return
	}

	fmt.Fprint(w, string(json))
}

type handlerError struct {
	Err string `json:"error"`
}

func outputError(err error, w http.ResponseWriter) {
	s, err := json.Marshal(handlerError{err.Error()})
	if err != nil {
		s = []byte(`{"error":"Error handling an error condition"}`)
	}
	http.Error(w, string(s), http.StatusBadRequest)
}

func main() {
	fmt.Println("Listening on :8080")
	http.Handle("/roll", diceHandler{})
	http.ListenAndServe(":8080", nil)
}
