package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// var url1 string = "http://albaab.free.fr/The%20Rules%20of%20Life.pdf"
var url string = "http://localhost:8080/pdf"

func downloadfilefromurl(w http.ResponseWriter, r *http.Request) {
	timeout := time.Duration(5) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Disposition", "attachment;filename=book-from-link.pdf")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
}
