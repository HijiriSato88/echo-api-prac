package models

import "time"

type Comment struct {
    ID        uint      `gorm:"primaryKey"`
    Content   string    `gorm:"type:text;not null"`
    PostID    uint      `gorm:"not null"`
    UserID    uint      `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time

    Post      Post      `gorm:"foreignKey:PostID"`
    User      User      `gorm:"foreignKey:UserID"`
}