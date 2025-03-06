package repositories

import (
	"fmt"
	"time"

	"echo_api_prac/infra"
	"echo_api_prac/models"
)

func GetUserByID(id uint) (*models.User, error) {
	cacheKey := fmt.Sprintf("user:%d", id)
	var user models.User

	// キャッシュから取得
	err := infra.GetCache(cacheKey, &user)
	if err == nil {
		return &user, nil // キャッシュヒット時
	}

	// キャッシュがない場合、DBから取得
	err = infra.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	// キャッシュに保存（TTL: 10分）
	_ = infra.SetCache(cacheKey, user, 10*time.Minute)

	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	cacheKey := "users"
	var users []models.User

	// キャッシュから取得
	err := infra.GetCache(cacheKey, &users)
	if err == nil {
		return users, nil
	}

	// キャッシュがない場合、DBから取得
	err = infra.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	// キャッシュに保存（TTL: 5分）
	_ = infra.SetCache(cacheKey, users, 5*time.Minute)

	return users, nil
}

func CreateUser(user *models.User) error {
	err := infra.DB.Create(user).Error
	if err != nil {
		return err
	}

	// キャッシュ削除（一覧を更新する必要がある）
	_ = infra.DeleteCache("users")

	return nil
}