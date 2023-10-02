package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// serve the pdf file
func servePDF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/pdf")
	f, err := os.Open("./static/pdf-2.pdf")
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}
