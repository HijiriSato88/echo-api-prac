package main

import (
	"Echo-CRUD-demo/infra"
	"Echo-CRUD-demo/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}