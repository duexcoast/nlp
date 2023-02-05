package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/duexcoast/nlp"
)

func main() {
	// routing
	// /health is an exact match
	// /health/ is a prefix match
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)

	// run server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

// Exercise: write a tokenizeHandler that will read the text from the
// request body and return JSON in the format {"tokens": ["who", "on", "first"]}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	// Convert and validate error
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cant read", http.StatusBadRequest)
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
