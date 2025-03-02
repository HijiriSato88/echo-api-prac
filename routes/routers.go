package routes

import (
    "echo_api_prac/handlers"

    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
    e.POST("/users", handlers.CreateUser)
    e.GET("/users", handlers.GetUsers)
    e.GET("/users/:id", handlers.GetUser)
}
