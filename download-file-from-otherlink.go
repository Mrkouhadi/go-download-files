package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func downloadPDFOtherLink(w http.ResponseWriter, r *http.Request) {
	url := "http://albaab.free.fr/The%20Rules%20of%20Life.pdf"
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
