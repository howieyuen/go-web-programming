package core

import (
	"time"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/storage"
)

type Session struct {
	Id        int64
	Uuid      string
	Email     string
	UserId    int64
	CreatedAt time.Time
}

// Check if session is valid in the database
func (session *Session) Check() (bool, error) {
	err := storage.Db.QueryRow("SELECT * FROM sessions WHERE uuid = ?", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Delete session from database
func (session *Session) DeleteByUUID() error {
	statement := "delete from sessions where uuid = ?"
	stmt, err := storage.Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(session.Uuid)
	return nil
}

// Get the user from the session
func (session *Session) User() (*User, error) {
	var user = User{}
	err := storage.Db.QueryRow("select * from users where id = ?", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
