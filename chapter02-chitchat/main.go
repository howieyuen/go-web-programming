package main

import (
	"net/http"
	"time"
	
	"k8s.io/klog"
	
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/handlers"
	"github.com/howieyuen/go-web-programming/chapter02-chitchat/options"
)

var config *options.Configuration

func init() {
	config = options.LoadConfig()
}

func main() {
	klog.Infof("ChitChat %s started at %s", options.Version(), config.Address)
	
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	
	// index
	mux.HandleFunc("/", handlers.Index)
	// error
	mux.HandleFunc("/err", handlers.Err)
	
	// defined in route_auth.go
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.HandleFunc("/signup", handlers.Signup)
	mux.HandleFunc("/signup_account", handlers.SignupAccount)
	mux.HandleFunc("/authenticate", handlers.Authenticate)
	
	// defined in route_thread.go
	mux.HandleFunc("/thread/new", handlers.NewThread)
	mux.HandleFunc("/thread/create", handlers.CreateThread)
	mux.HandleFunc("/thread/post", handlers.PostThread)
	mux.HandleFunc("/thread/read", handlers.ReadThread)
	
	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		klog.Fatal("starting Mux Server error: %v", err)
	}
}
