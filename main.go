package main

import (
	"net/http"
)

func main() {
	// serving a list files
	handleTest := http.StripPrefix("/allfiles/", http.FileServer(http.Dir("static")))
	http.Handle("/allfiles/", handleTest)

	// ability to download that pdf file from: "http://localhost:8080/files"
	http.HandleFunc("/downloadfile", downloadPDFFromFS)

	// aility to download a pdf file from different link
	http.HandleFunc("/downloadfilefromotherlink", downloadPDFOtherLink)

	// server on port:8080
	panic(http.ListenAndServe(":8080", nil))
}
