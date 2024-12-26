package main

import (
	asciiart "asciiart/Utils"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./templates"))

	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	//POST REUEST
	http.HandleFunc("/generate", asciiart.AsciiHandler)

	//GET REQUEST
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			//200
			tmpl, err := template.ParseFiles("./templates/index.html")
			if err != nil {
				//500
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println("Oops! 500 error 🤦‍♀️Something went haywire 🌀 Please stand by as we restore order 👨‍💻.")
				return
			}
			tmpl.Execute(w, nil)
			return
		}

		if _, err := os.Stat("./templates" + r.URL.Path); os.IsNotExist(err) {
			// Not Found (404): Missing template or invalid path
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "./templates/404.html")
			log.Println("404! 🐱 This page is a no-show!")
			return
		}
		// OK (200): Serve existing files via the file server
		fs.ServeHTTP(w, r)
	})

	log.Println("🔧 Server's working hard at http://localhost:8080! 💪")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
