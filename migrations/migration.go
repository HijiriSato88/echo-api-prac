package main

import (
	"echo_api_prac/infra"
	"echo_api_prac/models"
)

func main() {
	infra.Initialize()
	infra.SetupDB()

	if err := infra.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		panic("Failed to migrate database")
	}
}