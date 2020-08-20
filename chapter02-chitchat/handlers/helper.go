package handlers

import (
	"errors"
	"net/http"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/core"
)

// Checks if the user is logged in and has a session, if not err is not nil
func CheckSession(request *http.Request) (*core.Session, error) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		return nil, err
	}
	session := core.Session{Uuid: cookie.Value}
	if ok, _ := session.Check(); !ok {
		err = errors.New("invalid session")
	}
	return &session, nil
}
