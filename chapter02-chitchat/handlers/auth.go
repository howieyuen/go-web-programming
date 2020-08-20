package handlers

import (
	"net/http"
	"time"
	
	"github.com/google/uuid"
	"k8s.io/klog"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/core"
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/utils"
)

// GET /login
// Show the login page
func Login(writer http.ResponseWriter, _ *http.Request) {
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	err := t.Execute(writer, nil)
	if err != nil {
		klog.Errorf("template execute error: %v", err)
	}
}

// GET /signup
// show the signup page
func Signup(writer http.ResponseWriter, _ *http.Request) {
	utils.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		klog.Errorf("Cannot parse form: %v", err)
	}
	user := core.User{
		Uuid:      uuid.New().String(),
		Name:      request.PostFormValue("name"),
		Email:     request.PostFormValue("email"),
		Password:  request.PostFormValue("password"),
		CreatedAt: time.Now(),
	}
	if err := user.Create(); err != nil {
		klog.Errorf("Cannot create user: %v", err)
		return
	}
	klog.Infof("create user: %s successfully", user.Name)
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := core.GetUserByEmail(request.PostFormValue("email"))
	if err != nil {
		klog.Errorf("Cannot find user: &v", err)
	}
	
	if user.Password != utils.Encrypt(request.PostFormValue("password")) {
		klog.Error("user %s try to login with wrong password!", user.Name)
		http.Redirect(writer, request, "/login", 302)
		return
	}
	
	session, err := user.CreateSession()
	if err != nil {
		klog.Errorf("Cannot create session: %v", err)
		return
	}
	
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(writer, &cookie)
	http.Redirect(writer, request, "/", 302)
}

// GET /logout
// Logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie && cookie != nil {
		klog.Warning("Failed to get cookie")
		session := core.Session{Uuid: cookie.Value}
		err := session.DeleteByUUID()
		if err != nil {
			klog.Errorf("delete session by uuid error: %v", err)
		}
	}
	http.Redirect(writer, request, "/", 302)
}
