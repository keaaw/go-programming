package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!") // Write "Hello, World!" to the response
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Listen and serve requests
}
