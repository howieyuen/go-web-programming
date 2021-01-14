package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	post := Post{
		ID:      1,
		Content: "Hello World!",
		Author: Author{
			ID:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			{
				ID:      1,
				Content: "Have a great day!",
				Author: Author{
					ID:   2,
					Name: "Adam",
				},
			},
			{
				ID:      2,
				Content: "How are you today?",
				Author: Author{
					ID:   3,
					Name: "Betty",
				},
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	err = ioutil.WriteFile("post.json", output, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
