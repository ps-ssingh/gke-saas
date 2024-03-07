package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
	"fmt"
)

type Quote struct {
	Quote      string   `json:"Quote"`
	Author     string   `json:"Author"`
	Popularity float64  `json:"Popularity"`
	Category   string   `json:"Category"`
}

func main() {
	// Read quotes from JSON file
	quotesFile, err := ioutil.ReadFile("quotes.json")
	if err != nil {
		log.Fatal(err)
	}

	var quotes []Quote
	err = json.Unmarshal(quotesFile, &quotes)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		// Generate random quote
		randomQuote := quotes[rand.Intn(len(quotes))]

		// Marshal quote to JSON
		jsonQuote, err := json.Marshal(randomQuote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set Content-Type header and write response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonQuote)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
}

