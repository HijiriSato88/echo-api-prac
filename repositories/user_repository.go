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

// NewUserRepository は UserRepository 構造体のポインタを、UUserRepositoryインターフェース型として返す
func NewUserRepository(db *gorm.DB) UUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(newUser models.User) (*models.User, error) {
	result := r.db.Create(&newUser) // GORMはポインタで渡すことで、新しく生成されたID（主キ）、created_at, updated_atなどを構造体に自動で設定してくれるみたい
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil // 保存されたユーザー情報（自動で設定されたIDなど）
}

func (c *UserRepository) FindAll() (*[]models.User, error) {
	var users []models.User
	result := c.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}