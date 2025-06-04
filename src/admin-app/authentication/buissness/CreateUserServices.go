package buissness

import (
	genericModels "InterviewPractice/src/models"
	"InterviewPractice/src/utils/postgres"
	"admin-app/authentiction/models"
	"admin-app/authentiction/repositories"
	"context"
)

type CreateUserService struct {
	CreateUserRepository repositories.CreateUserRepository
}

func NewCreateUserSService(createUserRepository repositories.CreateUserRepository) *CreateUserService {
	return &CreateUserService{
		CreateUserRepository: createUserRepository,
	}
}

func (service *CreateUserService) CreateNewUser(ctx context.Context, spanCtx context.Context, bffCreateUserRequest models.BFFCreateUserRequest) error {
	postgresClient := postgres.GetPostgresClient()
	db := postgresClient.GormDB

	user := &genericModels.User{
		Username:    bffCreateUserRequest.UserName,
		Age:         bffCreateUserRequest.Age,
		Email:       bffCreateUserRequest.Email,
		Password:    bffCreateUserRequest.Password,
		PhoneNumber: bffCreateUserRequest.PhoneNumber,
	}

	err := service.CreateUserRepository.CreateUser(ctx, db, user)
	if err != nil {
		return err
	}
	return nil
}
