package main

import "net/http"

func main() {
	s := NewServer()
	s.MountHandlers()
	http.ListenAndServe(":4666", s.Router)
}
