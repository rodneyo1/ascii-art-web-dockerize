package server

import (
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func renderTemplate(w http.ResponseWriter, data *PageData) {
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

func TestRenderTemplate(t *testing.T) {
	// Prepare a mock template in memory for testing
	Tmpl = template.Must(template.New("test").Parse(`{{if .Error}}{{.Error}}{{else}}{{.Art}}{{end}}`))

	tests := []struct {
		data       *PageData
		expected   string
		statusCode int
	}{
		{&PageData{Art: "Art Content"}, "Art Content", http.StatusOK},
		{&PageData{Error: "Error Content"}, "Error Content", http.StatusOK},
	}

	for _, tt := range tests {
		rr := httptest.NewRecorder()
		renderTemplate(rr, tt.data)

		if status := rr.Code; status != tt.statusCode {
			t.Errorf("renderTemplate returned wrong status code: got %v want %v", status, tt.statusCode)
		}

		if !strings.Contains(rr.Body.String(), tt.expected) {
			t.Errorf("renderTemplate returned unexpected body: got %v want %v", rr.Body.String(), tt.expected)
		}
	}
}
