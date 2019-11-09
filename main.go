package main

import (
	"fmt"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	value, foundValue := os.LookupEnv(key)
	if foundValue {
		return value
	}
	return fallback
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	statusTemplate.Execute(w, fmt.Sprintf("Project status 2"))
}

func main() {
	// Say that when we receive a request for the '/' (or "root") URL
	// we want the function `indexHandler` to handle it.
	http.HandleFunc("/", statusHandler)
	// Start listening for HTTP requests.
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}
