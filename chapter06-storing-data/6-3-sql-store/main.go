package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Post is
type Post struct {
	ID      int
	Content string
	Author  string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func list(limit int) (posts []Post, err error) {
	rows, err := db.Query("select * from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		var post = Post{}
		err = rows.Scan(&post.ID, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func getPost(id int) (*Post, error) {
	var post = new(Post)
	err := db.QueryRow("select * from posts where id = $1", id).
		Scan(&post.ID, &post.Content, &post.Author)
	return post, err
}

func (p *Post) create() error {
	sql := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	err = stmt.QueryRow(p.Content, p.Author).Scan(&p.ID)
	return err
}

func (p *Post) update() error {
	_, err := db.Exec("update posts set content=$1, author=$2 where id=$3", p.Content, p.Author, p.ID)
	return err
}

func (p *Post) delete() error {
	_, err := db.Exec("delete from posts where id = $1", p.ID)
	return err
}

func main() {
	post := Post{Content: "Hello", Author: "zhangsan"}

	fmt.Println(post)
	err := post.create()
	if err != nil {
		panic(err)
	}
	fmt.Println(post)

	readPost, _ := getPost(post.ID)
	fmt.Println(readPost)

	readPost.Content = "world"
	readPost.Author = "lisi"
	readPost.update()

	posts, _ := list(5)
	fmt.Println(posts)

	readPost.delete()

	posts, _ = list(5)
	fmt.Println(posts)
}
