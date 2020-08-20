package core

import (
	"time"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/storage"
)

type Post struct {
	Id        int64
	Uuid      string
	Body      string
	UserId    int64
	ThreadId  int64
	CreatedAt time.Time
}

// Get the user who wrote the post
func (post *Post) User() (*User, error) {
	user := User{}
	err := storage.Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("2006-01-02 15:04:05")
}
