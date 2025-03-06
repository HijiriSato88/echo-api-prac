package repositories

import (
    "echo_api_prac/infra"
    "echo_api_prac/models"
)

func CreatePost(post *models.Post) error {
    return infra.DB.Create(post).Error
}

func GetPostsByTag(tag string) ([]models.Post, error) {
    var posts []models.Post
    err := infra.DB.
		Preload("User").
		Preload("Comments.User").
		Where("tag = ?", tag).Find(&posts).Error
    return posts, err
}