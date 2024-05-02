// Package main provides a simple HTTP server that serves static files from a specified directory.
// The server's settings are configurable via command-line flags.
//
// Usage:
//
//	-p="8080": Port to listen on. Default is "8080".
//	-d=".": Directory of static files to serve. Default is the current directory.
//	-b="/": Base URL path. Default is "/".
//
// Example:
//
//	To serve files out of "./static" on "localhost:8100" with a base path of "/files",
//	start the server with: `go run . -d="./static" -p="8100" -b="/files"`
//	Then, navigate to http://localhost:8100/files to access the served files.
package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

const (
	defaultPort      = "8100" // defaultPort defines the default port to listen on.
	defaultDirectory = "."    // defaultDirectory is the default directory to serve files from.
	defaultBasePath  = "/"    // defaultBasePath is the default base URL path.
)

func main() {
	// Command-line flags
	dirFlag := flag.String("d", defaultDirectory, "Directory to serve static files from")
	portFlag := flag.String("p", defaultPort, "Port to listen on")
	basePathFlag := flag.String("b", defaultBasePath, "Base URL path for the server")

	flag.Parse()

	// Resolve the absolute path of the directory to serve.
	absDir, err := filepath.Abs(*dirFlag)
	if err != nil {
		log.Fatalf("Error getting absolute directory path: %s", err)
	}

	// Setup and start the HTTP server.
	log.Printf("Serving %s at http://localhost:%s%s", absDir, *portFlag, *basePathFlag)
	setupServer(absDir, *portFlag, *basePathFlag)
}

// setupServer configures and starts an HTTP server.
// It serves static files from the specified directory at the given port and base path.
func setupServer(directory, port, basePath string) {
	fileServer := http.FileServer(http.Dir(directory))
	http.Handle(basePath, http.StripPrefix(basePath, fileServer))
	server := &http.Server{
		Addr:              ":" + port,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
