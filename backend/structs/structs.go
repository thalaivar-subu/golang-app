package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type InsertBody struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	ID    int    `json:"id"`
}
