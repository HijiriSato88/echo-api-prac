package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname 	string `gorm:"not null"`
	Lastname 	string `gorm:"not null"`
	Email 		string `gorm:"not null"`
	Phonenumber string `gorm:"not null"`
	Age 		uint `gorm:"not null"`
	Is_adult 	bool `gorm:"not null;default:false"`
}