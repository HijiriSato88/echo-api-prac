package repositories

import (
	"gorm.io/gorm"
	"echo_api_prac/models"
)

type UUserRepository interface {
	Create(newUser models.User) (*models.User, error)
	FindAll() (*[]models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(newUser models.User) (*models.User, error) {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}

func (c *UserRepository) FindAll() (*[]models.User, error) {
	var users []models.User
	result := c.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}