package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hbourgeot/brgt-devs/internal/database"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	s := NewServer()
	s.MountHandlers()

	db, err := NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	_ = &database.Brgt{DB: db}

	http.ListenAndServe(":4666", s.Router)
}

func NewConnection() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&database.User{}, &database.Project{}, &database.Task{}, &database.Document{}, &database.Version{}, &database.UsersProjects{})
	return db, nil
}
