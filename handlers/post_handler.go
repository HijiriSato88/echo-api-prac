package handlers

import (
    "echo_api_prac/services"
    "net/http"
    "echo_api_prac/validation"

    "github.com/labstack/echo/v4"
)

type CreatePostRequest struct {
    Title   string `json:"title" validate:"required"`
    Content string `json:"content" validate:"required"`
    UserID  uint   `json:"user_id" validate:"required"`
}

func CreatePost(c echo.Context) error {
    var req CreatePostRequest

    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    if err := validation.Validator.Struct(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    post, err := services.CreatePost(req.UserID, req.Title, req.Content)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, post)
}