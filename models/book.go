package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	User   []User    `json:"user" gorm:"primary_key"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
	Price  int       `json:"price"`
	Sales  int       `json:"sales"`
	Stock  int       `json:"stock"`
}
