package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primary_key"`
	Name            string    `gorm:"unique_index" json:"name"`
	Email           string    `json:"email" validate:"email,required"`
	Password        string    `gorm:"type:varchar(255);not null" json:"password" validate:"required,min=5"`
	PasswordConfirm string    `gorm:"type:varchar(255);not null" json:"passwordConfirm" validate:"required"`
	Token           string    `json:"token"`
	Provider        string    `gorm:"default:'local';"`
	Role            string    `gorm:"type:varchar(20);default:'user';"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoCreateTime"`
	LoginHistory    []Login
}

// type User struct {
// 	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	Name      string    `gorm:"type:varchar(255);not null"`
// 	Email     string    `gorm:"uniqueIndex;not null"`
// 	Password  string    `gorm:"not null"`
// 	Role      string    `gorm:"type:varchar(255);not null"`
// 	Provider  string    `gorm:"not null"`
// 	Photo     string    `gorm:"not null"`
// 	Verified  bool      `gorm:"not null"`
// 	CreatedAt time.Time `gorm:"not null"`
// 	UpdatedAt time.Time `gorm:"not null"`
// }
