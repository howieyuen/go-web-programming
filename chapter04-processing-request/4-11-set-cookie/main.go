package main

import (
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first",
		Value:    "zz",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second",
		Value:    "yy",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	http.SetCookie(w, &c2)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/cookie", setCookie)
	server.ListenAndServe()
}
