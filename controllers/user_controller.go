package controllers

import (
	"echo_api_prac/dto"
	"echo_api_prac/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UUserController interface {
	Create(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}

type UserController struct {
	service services.UUserService
}

func NewUserController(service services.UUserService) UUserController {
	return &UserController{service: service}
}

func (c *UserController) Create(ctx echo.Context) error {
	var input dto.CreateUserInput
	if err := ctx.Bind(&input); err != nil { // Bind：リクエストボディのデータをuser_dtoの構造体に詰める
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newUser, err := c.service.Create(input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{"data": newUser})
}

func (c UserController) FindAll(ctx echo.Context) error {
	users, err := c.service.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{"data": users})
}