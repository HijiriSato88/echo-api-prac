package repositories

import (
    "echo_api_prac/infra"
    "echo_api_prac/models"
)

func CreatePost(post *models.Post) error {
    return infra.DB.Create(post).Error
}