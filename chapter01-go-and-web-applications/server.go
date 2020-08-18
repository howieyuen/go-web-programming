package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello world, %s", request.URL.Path[1:])
}
