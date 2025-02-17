package controllers

import (
	"gin_api_prac/dto"
	"gin_api_prac/services"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type UUserController interface {
	Create(ctx *gin.Context)
}

type UserController struct {
	service services.UUserService
}

func NewUserController(service services.UUserService) UUserController {
	return &UserController{service: service}
}

func (c *UserController) Create(ctx *gin.Context) {
	var input dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newUser})
}
