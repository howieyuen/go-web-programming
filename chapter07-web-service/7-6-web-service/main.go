package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w, r)
	case "POST":
		handlePost(w, r)
	case "PUT":
		handlePut(w, r)
	case "DELETE":
		handleDelete(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post := getPost(id)
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	post.create()
	w.WriteHeader(http.StatusCreated)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post := getPost(id)
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	post.update()
	w.WriteHeader(http.StatusOK)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post := getPost(id)
	post.delete()
	w.WriteHeader(http.StatusOK)
	return
}
