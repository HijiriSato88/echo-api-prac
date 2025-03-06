package main

import (
	"echo_api_prac/infra"
	"echo_api_prac/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	infra.Initialize()
	infra.SetupDB()
	infra.InitRedis()
	infra.InitCache()

	e := echo.New()

	routes.SetupRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}
