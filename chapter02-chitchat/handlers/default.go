package handlers

import (
	"net/http"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/core"
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/utils"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := core.GetAllThreads()
	if err != nil {
		utils.ToErrorPage(writer, request, "Cannot get threads")
		return
	}
	_, err = CheckSession(request)
	fileName := "private.navbar"
	if err != nil {
		fileName = "public.navbar"
	}
	utils.GenerateHTML(writer, threads, "layout", fileName, "Index")
}

// GET /Err?msg=
// shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	_, err := CheckSession(request)
	fileName := "private.navbar"
	if err != nil {
		fileName = "public.navbar"
	}
	utils.GenerateHTML(writer, values.Get("msg"), "layout", fileName, "error")
}
