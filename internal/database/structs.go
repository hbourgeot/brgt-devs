package database

import "gorm.io/gorm"

type UsersProjects struct {
	UserID    int     `json:"-" gorm:"primaryKey"`
	ProjectID int     `json:"-" gorm:"primaryKey"`
	Project   Project `json:"project"`
	User      User    `json:"user"`
	*gorm.Model
}

type User struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"password"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email" gorm:"unique"`
	Projects []Project `json:"projects" gorm:"many2many:users_projects"`
	Tasks    []Task    `json:"tasks"`
	*gorm.Model
}

type Project struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LeaderID    int    `json:"-"`
	Leader      User   `json:"leader" gorm:"foreignKey:LeaderID"`
	Users       []User `json:"users" gorm:"many2many:users_projects"`
	Tasks       []Task `json:"tasks"`
	*gorm.Model
}

type Task struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ProjectID   int        `json:"-"` // Clave foránea para Project
	UserID      int        `json:"-"` // Clave foránea para User
	Project     Project    `json:"project"`
	User        User       `json:"user"`
	Documents   []Document `json:"documents" gorm:"foreignKey:TaskID"`
	*gorm.Model
}

type Document struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Path        string    `json:"path"`
	TaskID      int       `json:"-"` // Clave foránea para Task
	Task        Task      `json:"task"`
	Versions    []Version `json:"versions" gorm:"foreignKey:DocumentID"`
	*gorm.Model
}

type Version struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Path        string   `json:"path"`
	Changes     string   `json:"changes"`
	DocumentID  int      `json:"-"` // Clave foránea para Document
	Document    Document `json:"document"`
	*gorm.Model
}
