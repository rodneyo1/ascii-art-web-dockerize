package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	server "asciiartserver/server"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}

	var err error

	// Parse the template file
	server.Tmpl, err = template.ParseFiles("templates/index.html", "templates/about.html", "templates/error.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
	}

	// Define the handler function for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			server.AsciiArtHandler(w, r)
		case "/about":
			// Handle the /about path
			data := &server.PageData{}
			if err := server.Tmpl.ExecuteTemplate(w, "about.html", data); err != nil {
				log.Printf("Error executing template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		case "/download":
			// Handle file download
			server.DownloadHandler(w, r)
		default:
			// Handle 404 for unregistered paths
			if !strings.HasPrefix(r.URL.Path, "/static/") {
				data := &server.PageData{
					Error: "Page Not Found",
				}
				w.WriteHeader(http.StatusNotFound)
				if err := server.Tmpl.ExecuteTemplate(w, "error.html", data); err != nil {
					log.Printf("Error executing template: %v", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
				return
			}
		}
	})

	// Serve other static files (e.g., CSS, JS) using FileServer
	staticDir := http.Dir("static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
