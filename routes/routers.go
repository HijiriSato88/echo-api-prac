package routes

import (
    "echo_api_prac/handlers"

    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
    e.POST("/users", handlers.CreateUser)       // ユーザ作成
    e.GET("/users", handlers.GetUsers)          // 全ユーザ取得
    e.GET("/users/:id", handlers.GetUser)       // 特定のユーザ取得
    e.POST("/posts", handlers.CreatePost)       // ポスト投稿
    e.GET("/posts", handlers.GetPostsByTag )    // tagをキーとしてポスト取得
}
