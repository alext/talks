package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Die struct {
	Faces int `json:"faces"`
	Roll  int `json:"roll"`
}

func NewDie(faces int) Die {
	d := Die{
		Faces: faces,
		Roll:  rand.Intn(faces) + 1,
	}
	return d
}

func parseDieParam(dieParam string) (Die, error) {
	if dieParam == "" {
		return NewDie(6), nil
	}
	var faces int
	_, err := fmt.Sscanf(dieParam, "D%d", &faces)
	if err != nil {
		return Die{}, err
	}
	if faces <= 0 {
		return Die{}, errors.New("Die must have a positive integer number of faces")
	}
	return NewDie(faces), nil
}

func rollHandler(w http.ResponseWriter, req *http.Request) {
	d, err := parseDieParam(req.FormValue("die"))
	if err != nil {
		http.Error(w, "Invalid die param: "+err.Error(), http.StatusBadRequest)
		return
	}
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/roll", rollHandler)
	err := http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
