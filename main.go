package main

import (
	"gin_api_prac/controllers"
	"gin_api_prac/infra"
	"gin_api_prac/repositories"
	"gin_api_prac/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	r := gin.Default()
	r.POST("/users", userController.Create)
	r.Run(":8080")
}
