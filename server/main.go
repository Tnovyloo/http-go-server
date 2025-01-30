package main

import (
	"encoding/json"
	"http-server/api"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Config holds the allowed and denied file extensions
type Config struct {
	AllowedExtensions []string `json:"allowed_extensions"`
	DeniedExtensions  []string `json:"denied_extensions"`
}

func LoadConfig(filename string) (*Config, error) {
	config := &Config{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, config); err != nil {
		return nil, err
	}
	return config, nil
}

func FileExtensionMiddleware(config *Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve "index.html" when root path `/` is requested
		requestedPath := r.URL.Path
		if requestedPath == "/" {
			requestedPath = "/index.html"
		}

		ext := strings.ToLower(filepath.Ext(requestedPath))
		// log.Printf("Requested path: %s, resolved extension: %s", requestedPath, ext)

		// Check if the extension is denied
		for _, deniedExt := range config.DeniedExtensions {
			if ext == deniedExt {
				log.Printf("extension is in denied extensions list!: extensions: %s, request extension:%s", config.DeniedExtensions, ext)
				http.Error(w, "Access to this file type is denied", http.StatusForbidden)
				return
			}
		}

		// Check if the extension is allowed (if allowed_extensions is not empty)
		if len(config.AllowedExtensions) > 0 {
			allowed := false
			for _, allowedExt := range config.AllowedExtensions {
				if ext == allowedExt {
					log.Printf("extension is allowed! allowed extensions: %s, request extension:%s", config.AllowedExtensions, ext)
					allowed = true
					break
				}
			}
			if !allowed {
				log.Printf("extension is not allowed! allowed extensions: %s, request extension:%s", config.AllowedExtensions, ext)
				http.Error(w, "Access to this file type is not allowed", http.StatusForbidden)
				return
			}
		}

		// Serve the file if it passes the checks
		if requestedPath == "/index.html" {
			http.ServeFile(w, r, "./static/index.html")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	// Load the configuration file
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", FileExtensionMiddleware(config, fs))

	// Serve API
	http.HandleFunc("/api", api.HandleIndex)

	// Start the server on port 8080
	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
