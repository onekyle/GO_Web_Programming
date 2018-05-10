package main

import (
	"net/http"
	"html/template"
)

func main() {
	mux := http.NewServeMux()
	file := http.FileServer(http.Dir("/public"))

	mux.Handle("/static/", http.StripPrefix("/static/",file))

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createdThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:"0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads(); if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}