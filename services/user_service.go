package services

import (
	"echo_api_prac/dto"
	"echo_api_prac/models"
	"echo_api_prac/repositories"
)

type UUserService interface {
	Create(createUserInput dto.CreateUserInput) (*models.User, error)
	FindAll() (*[]models.User, error)
}

type UserService struct {
	repository repositories.UUserRepository
}

func NewUserService(repository repositories.UUserRepository) UUserService {
	return &UserService{repository: repository}
}

// DTO（Data Transfer Object）から、実際に保存するための ドメインモデル（models.User）を生成
// dto.CreateUserInput : 外部からの入力専用（リクエスト用データ）, models.User : アプリケーション内部で使用されるデータ構造（DB保存用など）
// 変換の意味：リクエスト仕様の変更が内部ロジックに影響しにくくなる。
func (s *UserService) Create(createUserInput dto.CreateUserInput) (*models.User, error) {

	isAdult := createUserInput.Age >= 20
	
	newUser := models.User{
		Firstname:        	createUserInput.Firstname,
		Lastname:        	createUserInput.Lastname,
		Email:       		createUserInput.Email,
		Phonenumber: 		createUserInput.Phonenumber,
		Age: 				createUserInput.Age,
		Is_adult:    		isAdult,
	}
	return s.repository.Create(newUser)
}

func (s *UserService) FindAll() (*[]models.User, error) {
	return s.repository.FindAll()
}
