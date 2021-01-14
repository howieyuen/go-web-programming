package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(postgres.Open("database=gwp user=gwp password=gwp sslmode=disable"), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&Post{})
}

func getPost(id int) Post {
	post := Post{}
	db.Where("id=?", id).Take(&post)
	return post
}

func (p *Post) create()  {
	db.Create(p)
}

func (p *Post) update()  {
	db.Updates(p)
}

func (p *Post) delete()  {
	db.Delete(p)
}
