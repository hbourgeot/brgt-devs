package main

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hbourgeot/brgt-devs/cmd/api/handlers"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (s *Server) MountHandlers() {
	// Mount middlewares
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.RedirectSlashes)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.StripSlashes)
	s.Router.Use(middleware.Timeout(60 * time.Second))

	// Mount the handlers
	s.Router.Get("/", handlers.Home)
}
