package core

import (
	"time"
	
	"github.com/google/uuid"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/storage"
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/utils"
)

type User struct {
	Id        int64
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Get a single user given the email
func GetUserByEmail(email string) (*User, error) {
	user := User{}
	err := storage.Db.QueryRow("SELECT * FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// create a user
func (user *User) Create() error {
	query := "insert into users (uuid, name, email, password, created_at) values (?,?,?,?,?)"
	stmt, err := storage.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	result, err := stmt.Exec(user.Uuid, user.Name, user.Email, utils.Encrypt(user.Password), user.CreatedAt)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}

// create a session for a existing user
func (user *User) CreateSession() (*Session, error) {
	query := "insert into sessions (uuid, email, user_id, created_at) values (?,?,?,?)"
	stmt, err := storage.Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	session := &Session{
		Uuid:      uuid.New().String(),
		Email:     user.Email,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	
	result, err := stmt.Exec(session.Uuid, session.Email, session.UserId, session.CreatedAt)
	if err != nil {
		return nil, err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	session.Id = id
	return session, nil
}

// delete user from database
func (user *User) Delete() error {
	statement := "delete from users where id = ?"
	stmt, err := storage.Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(user.Id)
	return nil
}

// Update user information in the database
func (user *User) Update() error {
	statement := "update users set name = ?, email = ? where id = ?"
	stmt, err := storage.Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return nil
}

// Get the session for an existing user
func (user *User) Session() (*Session, error) {
	session := Session{}
	err := storage.Db.QueryRow("SELECT * FROM sessions WHERE user_id = ?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// create a new thread
func (user *User) CreateThread(topic string) (*Thread, error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
	stmt, err := storage.Db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	thread := Thread{
		Uuid:      uuid.New().String(),
		Topic:     topic,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	// use QueryRow to return a row and scan the returned id into the Session struct
	result, err := stmt.Exec(thread.Uuid, thread.Topic, thread.UserId, thread.CreatedAt)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	thread.Id = id
	return &thread, nil
}

// create a new post to a thread
func (user *User) CreatePost(thread *Thread, body string) (*Post, error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values (?,?,?,?,?)"
	stmt, err := storage.Db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	post := Post{
		Uuid:      uuid.New().String(),
		Body:      body,
		UserId:    user.Id,
		ThreadId:  thread.Id,
		CreatedAt: time.Now(),
	}
	result, err := stmt.Exec(post.Uuid, post.Body, post.UserId, post.ThreadId, post.CreatedAt)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	post.Id = id
	return &post, nil
}
