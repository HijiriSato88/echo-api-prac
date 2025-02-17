package main

import (
	"gin_api_prac/infra"
	"gin_api_prac/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}