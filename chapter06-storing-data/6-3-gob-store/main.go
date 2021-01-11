package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

// Post is
type Post struct {
	ID      int
	Content string
	Author  string
}

func store(data interface{}, fileName string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileName, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, fileName string) {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{ID: 1, Content: "zzz", Author: "xxx"}
	store(post, "post.gob")
	var postRead Post
	load(&postRead, "post.gob")
	fmt.Print(postRead)
}
