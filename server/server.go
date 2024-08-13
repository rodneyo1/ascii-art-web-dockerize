package server

import (
	"asciiartserver/asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var Tmpl *template.Template

type PageData struct {
	Art   string
	Error string
}

// AsciiArtHandler generates ascii art as data and calls the rendeer function with "index.html"
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// If not a POST request, just render the form
		data := &PageData{}
		RenderTemplate(w, "templates/index.html", data)
		return
	}

	input := r.FormValue("input")
	banner := "asciiart/banners/" + r.FormValue("banner")

	data := &PageData{}
	if input == "" || banner == "" {
		handleError(w, data, http.StatusBadRequest, "Both text input and banner selection are required.", "Error: Missing input or banner selection")
		return
	}

	art, err := asciiart.GenerateASCIIArt(input, banner)
	if err != nil {
		switch err {
		case asciiart.ErrNotFound:
			handleError(w, data, http.StatusNotFound, "The specified banner was not found.", fmt.Sprintf("Error: %v", err))
		case asciiart.ErrBadRequest:
			handleError(w, data, http.StatusBadRequest, "The request was incorrect. Please check your input.", fmt.Sprintf("Error: %v", err))
		default:
			handleError(w, data, http.StatusInternalServerError, "An internal error occurred. Please try again later.", fmt.Sprintf("Internal error: %v", err))
		}
		return
	}

	data.Art = art
	RenderTemplate(w, "templates/index.html", data)
}

// RenderTemplate renders and executes templates
func RenderTemplate(w http.ResponseWriter, templateFile string, data *PageData) {
	var err error
	// Parse the template file
	Tmpl, err = template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
	}
	if Tmpl == nil {
		log.Println("Template file not found")
		http.Error(w, "Template file not found", http.StatusNotFound)
		return
	}
	if err := Tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		log.Println("Template executed successfully")
	}
}

// handleError logs errors and sends appropriate status codes
func handleError(w http.ResponseWriter, data *PageData, statusCode int, errMsg string, logMsg string) {
	data.Error = errMsg
	log.Println(logMsg)
	// Set the status code here
	w.WriteHeader(statusCode)
	// Render the template after setting the status code
	RenderTemplate(w, "templates/index.html", data)
}

// DownloadHandler downloads the generated ascii art into a .txt file
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handleError(w, &PageData{}, http.StatusMethodNotAllowed, "Invalid request method", "Invalid request method")
		return
	}

	// Extract the ASCII art from the query parameters (if available)
	art := r.URL.Query().Get("art")
	if art == "" {
		handleError(w, &PageData{}, http.StatusBadRequest, "No ASCII art provided", "No ASCII art provided")
		return
	}

	// Convert the ASCII art to bytes
	artBytes := []byte(art)
	contentLength := len(artBytes)

	// Set the headers to trigger a file download
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(contentLength))

	// Write the ASCII art to the response
	if _, err := w.Write(artBytes); err != nil {
		handleError(w, &PageData{}, http.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("Error writing file: %v", err))
	}
}
