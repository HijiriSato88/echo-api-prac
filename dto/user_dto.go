package dto

type CreateUserInput struct {
	Firstname   string `json:"firstname" binding:"required,min=2"`
	Lastname    string `json:"lastname" binding:"required,min=2"`
	Email       string `json:"email" binding:"required,email"`
	Phonenumber string `json:"phonenumber" binding:"required"`
	Age         uint   `json:"age" binding:"required,min=0"`
	Is_adult    bool   `json:"is_adult"`
}
