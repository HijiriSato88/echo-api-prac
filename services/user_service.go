package services

import (
    "echo_api_prac/models"
    "echo_api_prac/repositories"
)

func CreateUser(name, email string) (*models.User, error) {
    user := &models.User{Name: name, Email: email}
    err := repositories.CreateUser(user)
    return user, err
}

func GetUsers() ([]models.User, error) {
    return repositories.GetAllUsers()
}

func GetUser(id uint) (*models.User, error) {
    return repositories.GetUserByID(id)
}
