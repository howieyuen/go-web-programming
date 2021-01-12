package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	ID        int
	Content   string
	Author    string `sql:"not null"`
	PostID    int    `sql:"index"`
	CreatedAt time.Time
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(postgres.Open("database=gwp user=gwp password=gwp sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Post{}, &Comment{})
}

// https://gorm.io/docs/associations.html
func main()  {
	p := Post{Content: "xxx", Author: "zzz"}
	fmt.Println(p)
	db.Create(&p)
	fmt.Println(p)

	c := Comment{Content: "+1", Author: "vvv"}
	db.Model(&p).Association("Comments").Append(&c)
	fmt.Println(c)

	var readPost Post
	db.Where("id = ?", p.ID).Take(&readPost)
	fmt.Println(readPost)

	var cs []Comment
	db.Model(&readPost).Association("Comments").Find(&cs)
	fmt.Println(cs[0])
}