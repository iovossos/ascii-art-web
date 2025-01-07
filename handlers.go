package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// HandleRoutes handles all incoming requests and routes them appropriately
func HandleRoutes(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		IndexHandler(w, r)
	case "/download":
		DownloadHandler(w, r) // Renders the download page with redirection logic
	case "/download-file":
		ServeFileHandler(w, r) // Directly serves the file
	case "/ascii-art":
		AsciiArtHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 Not Found: %s %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // Set header before WriteHeader
	w.WriteHeader(http.StatusNotFound)
	RenderTemplate(w, "404", nil)
}

func outputExists() bool {
	info, err := os.Stat("output.txt")
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	// Set no-cache headers to prevent caching behavior
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")

	// Check if output.txt exists
	if !outputExists() {
		log.Println("output.txt not found. Rendering notexist.html.")
		RenderTemplate(w, "notexist", nil)
		return
	}

	// Get file info for Content-Length
	fileInfo, err := os.Stat("output.txt")
	if err != nil {
		log.Printf("Error getting file info: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set headers for download
	w.Header().Set("Content-Disposition", "attachment; filename=output.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Serve output.txt
	http.ServeFile(w, r, "output.txt")

	// Clear the ASCII result after serving the file
	asciiResult = ""

	defer func() {
		err := os.Remove("output.txt")
		if err != nil {
			log.Printf("Error deleting output.txt: %v", err)
		}
	}()
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Check if output.txt exists
	if !outputExists() {
		// Render the notexist.html page if file doesn't exist
		RenderTemplate(w, "notexist", nil)
		return
	}

	// Render the download page if output.txt exists
	RenderTemplate(w, "download", nil)
}

// POST handler to process the ASCII art generation
var asciiResult string // Global variable to temporarily store the result

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	if r.Method != http.MethodPost {
		log.Printf("Method Not Allowed: %s", r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		log.Printf("400 Bad Request: Missing page or banner")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		RenderTemplate(w, "400", nil)
		return
	}

	if !CheckValidity(text) {
		log.Printf("Bad Request: Invalid input")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		RenderTemplate(w, "400", nil)
		return
	}

	asciiMap, asciiHeight, err := LoadBanner("assets/" + banner + ".txt")
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		RenderTemplate(w, "500", nil)
		return
	}

	// Generate the ASCII art and store it in the global variable
	asciiResult = ProcessString(text, asciiMap, asciiHeight)
	log.Printf("Rendering ASCII art result for input: %s", text)

	// Render the index page with the ASCII result
	RenderTemplate(w, "index", asciiResult)
}

// GET handler for the home page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("200 OK - Serving index page for request: %s %s", r.Method, r.URL.Path)
	RenderTemplate(w, "index", asciiResult)
	asciiResult = "" // Clear the result after rendering
}

// RenderTemplate renders the specified HTML template with the provided data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Parse the specified template file
	t, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", tmpl))
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal Server Error: Unable to load template.", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error: Unable to render template.", http.StatusInternalServerError)
		return
	}
}
