package models

type BFFCreateUserRequest struct {
	UserName    string `json:"username" validate:"required" example:"xyz"`
	Email       string `json:"email" validate:"required" example:"xyz@gmail.com"`
	Password    string `json:"password" validate:"min=8 max=12" example:"xyz123456"`
	PhoneNumber string `json:"phonenumber" validate:"min=10 max=12" example:"9999999999"`
	Age         int    `json:"age" validate:"min=18 max 40" example:"29"`
}

type BFFCreateUserResponse struct {
	Message string `json:"message" example:"User is successfully created"`
}
