package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Post is
type Post struct {
	ID      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Open("posts.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	// reader := csv.NewReader(bufio.NewReader(csvFile))
	reader := csv.NewReader(csvFile)
	var posts []Post
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		id, _ := strconv.Atoi(line[0])
		post := Post{ID: id, Content: line[1], Author: line[2]}
		posts = append(posts, post)
	}

	fileName := "posts.csv.bak"
	_, err = os.Lstat(fileName)
	if !os.IsNotExist(err) {
		os.Remove(fileName)
	}

	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)

	for _, post := range posts {
		fmt.Println(post)
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		writer.Write(line)
	}
	writer.Flush()
}
