package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/duexcoast/nlp"
	"github.com/duexcoast/nlp/stemmer"
)

func main() {
	// routing
	// /health is an exact match
	// /health/ is a prefix match
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", tokenizeHandler).Methods(http.MethodPost)
	r.HandleFunc("/stem/{word}", stemHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// run server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	stem := stemmer.Stem(word)
	fmt.Fprintln(w, stem)
}

// Exercise: write a tokenizeHandler that will read the text from the
// request body and return JSON in the format {"tokens": ["who", "on", "first"]}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	// Step 1: Convert and validate error
	// We don't want to accept just whatever - this caps the data
	// we read from the req at 1 Mb
	// We're handling our validation in the next couple steps
	rdr := io.LimitReader(r.Body, 1_000_000)
	data, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, "cant read", http.StatusBadRequest)
		// we need to return after http.Error
		return
	}
	if len(data) == 0 {
		http.Error(w, "missing data", http.StatusBadRequest)
		return
	}

	// step 2: Work
	tokens := nlp.Tokenize(string(data))

	// step 3: Encode & emit output
	resp := map[string]any{
		"tokens": tokens,
	}
	data, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't encode", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	// fmt.Fprintln(w, "OK")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: run a health check
	fmt.Fprintln(w, "OK")
}
