package main

import (
	"fmt"
	"net/http"
)

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("first")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, c)
	
	cookies := r.Cookies()
	fmt.Fprintln(w, cookies)
	
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
