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
	s.Router.Get("/users", handlers.GetUsers)
	s.Router.Get("/users/{id}", handlers.GetUser)
	s.Router.Post("/users", handlers.CreateUser)
	s.Router.Delete("/users/{id}", handlers.DeleteUser)
	s.Router.Put("/users/{id}", handlers.UpdateUser)

	s.Router.Get("/projects", handlers.GetProjects)
	s.Router.Get("/projects/{id}", handlers.GetProject)
	s.Router.Post("/projects", handlers.CreateProject)
	s.Router.Delete("/projects/{id}", handlers.DeleteProject)
	s.Router.Put("/projects/{id}", handlers.UpdateProject)

	s.Router.Get("/tasks", handlers.GetTasks)
	s.Router.Get("/tasks/{id}", handlers.GetTask)
	s.Router.Post("/tasks", handlers.CreateTask)
	s.Router.Put("/tasks/{id}", handlers.UpdateTask)

	s.Router.Get("/documents", handlers.GetDocuments)
	s.Router.Get("/documents/{id}", handlers.GetDocument)
	s.Router.Post("/documents", handlers.CreateDocument)
	s.Router.Put("/documents/{id}", handlers.UpdateDocument)

	s.Router.Get("/versions", handlers.GetVersions)
	s.Router.Get("/versions/{id}", handlers.GetVersion)
	s.Router.Post("/versions", handlers.CreateVersion)
	s.Router.Put("/versions/{id}", handlers.UpdateVersion)
}
