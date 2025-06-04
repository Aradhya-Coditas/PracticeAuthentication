package postgres

import (
	"InterviewPractice/src/constants"
	"InterviewPractice/src/models"
	config "InterviewPractice/src/utils/configs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dataBaseConfiguration *models.DatabaseConfiguration
var postgresConfiguration *models.PostgresConfiguration

func InitPostgresClient(configPath string) error {
	if err := initPostgresConfig(configPath); err != nil {
		return err
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", postgresConfiguration.Host, postgresConfiguration.Port, postgresConfiguration.Username, postgresConfiguration.Password, postgresConfiguration.DBName, postgresConfiguration.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	setPostgresClient(db)
	return nil
}

func setPostgresClient(db *gorm.DB) {
	dataBaseConfiguration = &models.DatabaseConfiguration{
		GormDB: db,
	}
}

func GetPostgresClient() *models.DatabaseConfiguration {
	return dataBaseConfiguration
}

func initPostgresConfig(configPath string) error {
	var err error
	postgresConfiguration, err = config.LoadConfig[models.PostgresConfiguration](configPath, constants.ConfigName, constants.ConfigType)
	return err
}
