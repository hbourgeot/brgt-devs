package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	s := NewServer()
	s.MountHandlers()
	http.ListenAndServe(":4666", s.Router)
}
