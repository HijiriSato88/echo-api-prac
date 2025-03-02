package handlers

import (
    "echo_api_prac/services"
    "net/http"
    "strconv"
    "github.com/go-playground/validator/v10"

    "github.com/labstack/echo/v4"
)

var validate = validator.New()

type CreateUserRequest struct {
    Name  string `json:"name" validate:"required,min=2,max=100"`
    Email string `json:"email" validate:"required,email"`
}

func CreateUser(c echo.Context) error {
	var req CreateUserRequest

    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    if err := validate.Struct(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

	user, err := services.CreateUser(req.Name, req.Email)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
    users, err := services.GetUsers()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
    }

    return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
    }

    user, err := services.GetUser(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }

    return c.JSON(http.StatusOK, user)
}
