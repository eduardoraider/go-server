package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function to handle incoming requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}

	// Register the handler function for the "/" URL pattern
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
