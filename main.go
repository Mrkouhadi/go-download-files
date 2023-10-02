package main

import (
	"net/http"
)

func main() {
	// serving a list files in static directory
	serveAllfilesInstaticDir := http.StripPrefix("/", http.FileServer(http.Dir("static")))
	http.Handle("/", serveAllfilesInstaticDir)

	// serve a single pdf file : ./static/pdf-2.pdf
	http.HandleFunc("/pdf", servePDF)

	// ability to download that pdf file from static directory
	http.HandleFunc("/download-file-from-fs", downloadPDFFromFS)

	// aility to download a pdf file from a url
	http.HandleFunc("/download-file-from-url", downloadfilefromurl)

	// server on port:8080
	panic(http.ListenAndServe(":8080", nil))
}
