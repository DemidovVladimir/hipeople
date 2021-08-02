package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DemidovVladimir/publishfiles/handlers"
)

func main() {
	handlers.PORT = "8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number: ", handlers.PORT)
	} else {
		handlers.PORT = arguments[1]
		fmt.Println("Using port number: ", handlers.PORT)
	}
	fileServer := http.FileServer(http.Dir(handlers.UploadPath))
	r := http.NewServeMux()
	r.HandleFunc("/", handlers.UploadPageHandler)
	r.HandleFunc("/upload", handlers.UploadFile)
	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Print("Server started on localhost:8080, use / for uploading files and /static/ for checking the files.")
	err := http.ListenAndServe(":"+handlers.PORT, r)
	if err != nil {
		log.Println(err)
		return
	}
}
