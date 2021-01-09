package main

import (
	"net/http"
	"text/template"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	// layout := "Mon Jan 2 15:04:05 MST 2006"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
