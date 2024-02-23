package handlers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a user"))
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a project"))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a task"))
}

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a document"))
}

func CreateVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a version"))
}
