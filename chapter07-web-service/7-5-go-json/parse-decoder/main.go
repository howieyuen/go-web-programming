package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(post)
	}
}
