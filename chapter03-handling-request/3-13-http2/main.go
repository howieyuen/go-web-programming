package main

import (
	"net/http"
	
	"golang.org/x/net/http2"
)

type handler struct {
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello"))
}

var _ http.Handler = &handler{}

func main() {
	sever := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler{},
	}
	http2.ConfigureServer(&sever, &http2.Server{})
	sever.ListenAndServe()
}
