package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello~~"))
}

func log(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Println("Handler func called", name)
		f(writer, request)
	}
}

func main()  {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}