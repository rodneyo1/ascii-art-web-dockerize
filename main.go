package main

import (
	"fmt"
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

	// Define the handler functions
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			server.AsciiArtHandler(w, r)
		case "/about":
			data := &server.PageData{}
			server.RenderTemplate(w, "templates/about.html", data)
		case "/download":
			server.DownloadHandler(w, r)
		default:
			// Handle 404 for unregistered paths
			if !strings.HasPrefix(r.URL.Path, "/static/") {
				data := &server.PageData{
					Error: "Page Not Found",
				}
				w.WriteHeader(http.StatusNotFound)
				server.RenderTemplate(w, "templates/error.html", data)
				return
			}
		}
	})

	// Serve static files using FileServer
	staticDir := http.Dir("static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	log.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
