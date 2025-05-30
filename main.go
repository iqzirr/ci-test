// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// helloHandler responds to HTTP requests with "Hello, World!".
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	greetings := "Haloooo ci cd github coy"
	// Print a message to the response writer.
	// This message will be sent to the client (web browser).
	fmt.Fprintf(w, "%s \nServed by container: %s\n", greetings, hostname)
	// Log the request to the server console
	log.Printf("Received request for %s from %s\n", r.URL.Path, r.RemoteAddr)
}

func helloHandler_coba(w http.ResponseWriter, r *http.Request) {
	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	greetings := "Hai, belajar github action lagi cuy"

	// Print a message to the response writer.
	// This message will be sent to the client (web browser).
	fmt.Fprintf(w, "%s \nServed by container: %s\n", greetings, hostname)
	// Log the request to the server console
	log.Printf("Received request for %s from %s\n", r.URL.Path, r.RemoteAddr)
}

func main() {
	// Define the port the server will listen on.
	// You can use an environment variable or a default.
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// Register the helloHandler function to handle all requests to the root path ("/").
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/coba", helloHandler_coba)

	// Start the HTTP server.
	log.Printf("Server starting on port %s...\n", port)
	// http.ListenAndServe will block until the server is stopped or an error occurs.
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		// If ListenAndServe returns an error, log it and exit.
		log.Fatalf("Could not start server: %s\n", err)
	}
}
