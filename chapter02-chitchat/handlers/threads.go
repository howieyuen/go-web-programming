package handlers

import (
	"fmt"
	"net/http"
	
	"k8s.io/klog"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/core"
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/utils"
)

// GET /threads/new
// Show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := CheckSession(request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// create the user account
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := CheckSession(request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	
	err = request.ParseForm()
	if err != nil {
		klog.Errorf("Cannot parse form: %v", err)
		return
	}
	user, err := sess.User()
	if err != nil {
		klog.Errorf("Cannot get user from session: %v", err)
		return
	}
	topic := request.PostFormValue("topic")
	if _, err := user.CreateThread(topic); err != nil {
		klog.Errorf("Cannot create thread: %v", err)
		return
	}
	klog.Infof("create a new thread: %s success", topic)
	http.Redirect(writer, request, "/", 302)
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	uuid := values.Get("id")
	thread, err := core.GetThreadByUUID(uuid)
	if err != nil {
		utils.ToErrorPage(writer, request, "Cannot read thread")
		return
	}
	_, err = CheckSession(request)
	fileName := "private.navbar"
	if err != nil {
		fileName = "public.navbar"
	}
	utils.GenerateHTML(writer, thread, "layout", fileName, "private.thread")
}

// POST /thread/post
// create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
	session, err := CheckSession(request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	
	err = request.ParseForm()
	if err != nil {
		klog.Errorf("Cannot parse form: %v", err)
		return
	}
	user, err := session.User()
	if err != nil {
		klog.Errorf("Cannot get user from session: %v", err)
		return
	}
	body := request.PostFormValue("body")
	uuid := request.PostFormValue("uuid")
	thread, err := core.GetThreadByUUID(uuid)
	if err != nil {
		utils.ToErrorPage(writer, request, "Cannot read thread")
		return
	}
	if _, err := user.CreatePost(thread, body); err != nil {
		klog.Errorf("Cannot create post: %v", err)
		return
	}
	klog.Infof("user: %s reply topic %s", user.Name, thread.Topic)
	url := fmt.Sprint("/thread/read?id=", uuid)
	http.Redirect(writer, request, url, 302)
}
