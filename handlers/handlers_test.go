package handlers

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

// Test Template rendering
func TestTemplateRenderer(t *testing.T) {
	w := httptest.NewRecorder()
	p := PageUploader{InputName: "uploadFile", Port: PORT}
	err := renderTemplate(w, "upload", "../public/upload.html", &p)
	if err == nil {
		fmt.Println("Template has been rendered")
	}

	// Output:
	// Template has been rendered
}
