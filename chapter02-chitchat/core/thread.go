package core

import (
	"time"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/storage"
)

type Thread struct {
	Id        int64
	Uuid      string
	Topic     string
	UserId    int64
	CreatedAt time.Time
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() int {
	rows, err := storage.Db.Query("SELECT count(*) FROM posts where thread_id = ?", thread.Id)
	if err != nil {
		return 0
	}
	var count int
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return -1
		}
	}
	_ = rows.Close()
	return count
}

// get posts to a thread
func (thread *Thread) Posts() ([]*Post, error) {
	rows, err := storage.Db.Query("SELECT * FROM posts where thread_id = ?",
		thread.Id)
	if err != nil {
		return nil, err
	}
	var posts []*Post
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	_ = rows.Close()
	return posts, nil
}

// Get the user who started this thread
func (thread *Thread) User() (*User, error) {
	user := User{}
	err := storage.Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Get all threads in the database and returns it
func GetAllThreads() ([]*Thread, error) {
	rows, err := storage.Db.Query("SELECT * FROM threads ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	
	var threads []*Thread
	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt); err != nil {
			return nil, err
		}
		threads = append(threads, &thread)
	}
	_ = rows.Close()
	return threads, nil
}

func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("2006-01-02 15:04:05")
}

// Get a thread by the UUID
func GetThreadByUUID(uuid string) (*Thread, error) {
	thread := Thread{}
	err := storage.Db.QueryRow("SELECT * FROM threads WHERE uuid = ?", uuid).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &thread, nil
}
