package services

import (
    "echo_api_prac/models"
    "echo_api_prac/repositories"
)

func CreatePost(user_id uint, title, content string, tag string) (*models.Post, error) {
    post := &models.Post{ Content: content, Title: title, Tag: tag, UserID: user_id,  }
    err := repositories.CreatePost(post)
    return post, err
}

func GetPostsByTag(tag string) ([]models.Post, error) {
    return repositories.GetPostsByTag(tag)
}
