package main

import (
	"fmt"
	"net/http"
)

type handler struct{}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello~~"))
}

var _ http.Handler = &handler{}

func log(f http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("Handler called, %T\n", f)
		f.ServeHTTP(writer, request)
	})
}

func protect(f http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("login success\n"))
		f.ServeHTTP(writer, request)
	})
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.Handle("/hello", protect(log(&handler{})))
	server.ListenAndServe()
}
