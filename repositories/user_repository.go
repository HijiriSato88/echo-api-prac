package repositories

import (
    "echo_api_prac/infra"
    "echo_api_prac/models"
)

func CreateUser(user *models.User) error {
    return infra.DB.Create(user).Error
}

func GetAllUsers() ([]models.User, error) {
    var users []models.User
    err := infra.DB.Find(&users).Error
    return users, err
}

func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    err := infra.DB.First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}