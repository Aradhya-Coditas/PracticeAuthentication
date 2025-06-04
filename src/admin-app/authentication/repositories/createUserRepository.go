package repositories

import (
	"InterviewPractice/src/models"
	"context"

	"gorm.io/gorm"
)

type createUserRepository struct{}

type CreateUserRepository interface {
	CreateUser(ctx context.Context, db *gorm.DB, User *models.User) error
}

func NewCreateUserWatchlist() *createUserRepository {
	return &createUserRepository{}
}

func GetCreateUserWatchlist(useDBMocks bool) CreateUserRepository {
	return NewCreateUserWatchlist()
}

func (c *createUserRepository) CreateUser(ctx context.Context, db *gorm.DB, User *models.User) error {
	// user := User
	result := db.WithContext(ctx).Model(&models.User{}).Create(&User)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}