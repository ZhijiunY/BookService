package models

import (
	"time"
)

type Register struct {
	Name            string    `json:"name" validate:"required"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `gorm:"type:varchar(255);not null" json:"password" validate:"required,min=5"`
	PasswordConfirm string    `gorm:"type:varchar(255);not null" json:"passwordConfirm" validate:"required"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoCreateTime"`
}

type Login struct {
	Email     string `json:"email" validate:"required"`
	Password  string `gorm:"type:varchar(255);not null" json:"password" validate:"required,min=5"`
	LogID     uint   `gorm:"primary_key"`
	UserID    uint   `gorm:"index"`
	LoginTime time.Time
	IPAddress string
}
