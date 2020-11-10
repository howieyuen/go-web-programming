package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello")
}

type worldHandler struct{}

func (w *worldHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(" world!"))
}

var _ http.Handler = &helloHandler{}
var _ http.Handler = &worldHandler{}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.Handle("/hello", &helloHandler{})
	http.Handle("/world", &worldHandler{})
	server.ListenAndServe()
}
