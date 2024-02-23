package handlers

import "net/http"

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a user"))
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a project"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a task"))
}

func UpdateDocument(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a document"))
}

func UpdateVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a version"))
}
