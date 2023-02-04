package main

import (
	"fmt"
	"net/http"
)

func main() {
	// routing
	// /health is an exact match
	// /health/ is a prefix match
	http.HandleFunc("/health", healthHandler)

	// run server
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Run a health check/
	fmt.Println(w, "OK")
}
