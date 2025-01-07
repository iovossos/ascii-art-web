package main

import (
	"log"
	"net/http"
)

func main() {
	// Route for the main page
	http.HandleFunc("/", HandleRoutes)

	// Route to handle the ASCII art generation
	http.HandleFunc("/ascii-art", AsciiArtHandler)
	http.HandleFunc("/download", DownloadHandler)

	// Serve favicon
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	// Serve static files (CSS & images) from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server on port 8080
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
