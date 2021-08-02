package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

// PORT is the port used to serve the app
var PORT string

// UploadPath uploades files to this folder
const UploadPath = "./static"

// PageUploader is a structure that will be populate data to the upload form template
type PageUploader struct {
	Port      string
	InputName string
}

// templates is a upload file template
var templates = template.Must(template.ParseFiles("public/upload.html"))

// Rendering form template
func renderTemplate(w http.ResponseWriter, tmpl string, p *PageUploader) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UploadPageHandler Index page with the upload form, peerhaps could be renamed
func UploadPageHandler(w http.ResponseWriter, r *http.Request) {
	p := PageUploader{InputName: "uploadFile", Port: PORT}
	renderTemplate(w, "upload", &p)
}

// UploadFile uploades files into static folder
func UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get handler for filename, size and headers
	file, h, err := r.FormFile("uploadFile")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tmpfile, err := os.Create(UploadPath + "/" + h.Filename)
	defer tmpfile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/static/", http.StatusSeeOther)
	return
}
