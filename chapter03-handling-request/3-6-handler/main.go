package main

import (
	"fmt"
	"net/http"
)

var _ http.Handler = handler{}

type handler struct{}

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello")
	writer.Write([]byte(" world!"))
}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler{},
	}
	server.ListenAndServe()
}
