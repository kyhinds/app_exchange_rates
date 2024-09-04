package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/internal/api" // Ensure this import path correctly points to your API package
)

func main() {
	// Serve React static files from the client/build directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get working directory: %v", err)
	}
	reactBuildPath := filepath.Join(workingDir, "..", "client", "build")
	fs := http.FileServer(http.Dir(reactBuildPath))

	// Serve static files on root and all other unmatched routes so the React router can handle them.
	// Ensures that users can refresh on any page without getting a 404 error.
	http.Handle("/", http.StripPrefix("/", fs))

	// API Endpoints
	http.HandleFunc("/rates", api.RatesHandler)
	http.HandleFunc("/convert", api.ConvertHandler)

	// Start the server on port 8080
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
