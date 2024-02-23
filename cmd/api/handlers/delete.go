package handlers

import "net/http"

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a user"))
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a project"))
}
