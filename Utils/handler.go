package asciiart

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Result string
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
			// Internal Server Error (500): Failed to parse template
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		// our server receives this data through r.FormValue()
		text := r.FormValue("text")
		font := r.FormValue("font")

		if text == "" || !IsValidASCII(text) {
		// Bad Request (400): Invalid or missing ASCII input
			 
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "./templates/400.html")
			log.Println("Bad request: Missing text or invalid ASCII")
			log.Println("Whoops! Looks like the user hit a little bump in the road 🚧. 400 error! 💥")

			return
		}

		if font == "" {
			font = "standard"
		}

		result, err := GenerateAsciiArt(text, font)
		if err != nil {
		// Internal Server Error (500): Error in generating ASCII art
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./templates/500.html")
			log.Println("Oops! 500 error 🤦‍♀️Something went haywire 🌀 Please stand by as we restore order 👨‍💻.")

			return
		}

		// OK (200): Successfully generated ASCII art
		tmpl.Execute(w, TemplateData{Result: result})
		return
	}

	// After processing, it sends back the result using Go templates:
	tmpl.Execute(w, TemplateData{Result: ""})
}

func IsValidASCII(text string) bool {
	for _, char := range text {
		if char > 127 {
			return false
		}
	}
	return true
}
