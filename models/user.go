package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:100;not null"`
    Email     string    `gorm:"size:255;not null;unique"`
    CreatedAt time.Time
    UpdatedAt time.Time

    Posts     []Post    `gorm:"foreignKey:UserID"`
}
