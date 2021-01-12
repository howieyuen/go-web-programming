package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	ID       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	ID      int
	Content string
	Author  string
	Post    *Post
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "database=gwp user=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

type operator interface {
	create() error
}

func (p *Post) create() error {
	return db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", p.Content, p.Author).Scan(&p.ID)
}

func (c *Comment) create() error {
	if c.Post == nil {
		return errors.New("post not found")
	}
	return db.QueryRow("insert into comments(content, author, post_id) values ($1, $2, $3)", c.Content, c.Author, c.Post.ID).Scan(&c.ID)
}

func getPost(id int) (*Post, error) {
	var post = &Post{ID: id}
	err := db.QueryRow("select content, author from posts where id = $1", id).
		Scan(&post.Content, &post.Author)
	rows, err := db.Query("select id, content, author from comments where post_id = $1", id)
	if rows != nil {
		defer rows.Close()
	}
	for rows != nil && rows.Next() {
		comment := Comment{Post: post}
		err := rows.Scan(&comment.ID, &comment.Content, &comment.Author)
		if err != nil {
			return post, err
		}
		post.Comments = append(post.Comments, comment)
	}
	return post, err
}

func main() {
	var err error
	handle := func(o operator) {
		if err != nil {
			log.Fatal(err)
			return
		}
		err = o.create()
	}

	post := &Post{Content: "zzz", Author: "xxx"}
	handle(post)

	comment := &Comment{Content: "+1", Author: "www", Post: post}
	handle(comment)

	readPost, err := getPost(post.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%v\n", readPost)
	fmt.Printf("%v\n", readPost.Comments)
	for _, c := range readPost.Comments {
		fmt.Printf("%v\n", c.Post)
	}
}
