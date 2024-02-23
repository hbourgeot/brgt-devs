package main_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hbourgeot/brgt-devs/internal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewConnection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}))

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	gormDB.AutoMigrate(&database.User{}, &database.Project{}, &database.Task{}, &database.Document{}, &database.Version{}, &database.UsersProjects{})
}
