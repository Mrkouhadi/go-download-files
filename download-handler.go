package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadPDFFromFS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment;filename=book.pdf")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	f, err := os.Open("./static/pdf-2.pdf")
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(w, f)
}
