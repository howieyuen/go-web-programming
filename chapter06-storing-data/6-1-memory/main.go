package main

import "fmt"

// Post is
type Post struct {
	ID      int
	Content string
	Author  string
}

var postByID map[int]*Post
var postByAuthor map[string][]*Post

func store(posts ...Post) {
	for i := range posts {
		postByID[posts[i].ID] = &posts[i]
		postByAuthor[posts[i].Author] = append(postByAuthor[posts[i].Author], &posts[i])
	}
}

func main() {
	postByID = make(map[int]*Post, 4)
	postByAuthor = make(map[string][]*Post)

	post1 := Post{ID: 1, Content: "Hello zhangsan", Author: "zhangsan"}
	post2 := Post{ID: 2, Content: "Hello lisi", Author: "lisi"}
	post3 := Post{ID: 3, Content: "zhangsan World", Author: "zhangsan"}
	post4 := Post{ID: 4, Content: "lisi World", Author: "lisi"}

	store(post1, post2, post3, post4)
	// store([]Post{post1, post2, post3, post4}...)

	fmt.Println(postByID[1])
	fmt.Println(postByID[2])

	for _, post := range postByAuthor["zhangsan"] {
		fmt.Println(post)
	}

	for _, post := range postByAuthor["lisi"] {
		fmt.Println(post)
	}
}
