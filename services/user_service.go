package services

import (
	"gin_api_prac/dto"
	"gin_api_prac/models"
	"gin_api_prac/repositories"
)

type UUserService interface {
	Create(createUserInput dto.CreateUserInput) (*models.User, error)
}

type UserService struct {
	repository repositories.UUserRepository
}

func NewUserService(repository repositories.UUserRepository) UUserService {
	return &UserService{repository: repository}
}

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
