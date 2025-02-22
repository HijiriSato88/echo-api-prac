package main

import (
	"echo_api_prac/controllers"
	"echo_api_prac/infra"
	"echo_api_prac/repositories"
	"echo_api_prac/services"

	"github.com/labstack/echo/v4"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	e := echo.New()
	e.POST("/users", userController.Create)
	e.Logger.Fatal(e.Start(":8080"))
}
