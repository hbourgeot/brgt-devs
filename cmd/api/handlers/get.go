package handlers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a user"))
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all projects"))
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a project"))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all tasks"))
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a task"))
}

func GetDocuments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all documents"))
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a document"))
}

func GetVersions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all versions"))
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a version"))
}
