package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello")
}

func world(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(" world!"))
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	server.ListenAndServe()
}
