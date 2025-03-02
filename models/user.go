package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  string `gorm:"not null" validate:"required,min=2,max=100"`
    Email string `gorm:"unique;not null" validate:"required,email"`
}