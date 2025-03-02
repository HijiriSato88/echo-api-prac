package models

import "time"

type User struct {
    ID           uint       `gorm:"primary_key"`
    Name         string
    Email        string     `gorm:"type:varchar(100);unique"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}