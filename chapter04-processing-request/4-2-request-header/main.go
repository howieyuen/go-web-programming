package main

import (
	"fmt"
	"net/http"
)

func header(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(write, request.Header)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/header", header)
	server.ListenAndServe()
}
