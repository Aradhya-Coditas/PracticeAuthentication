package main

import (
	"InterviewPractice/src/models"
	"InterviewPractice/src/utils/postgres"
	"log"
)

func main() {
	err := postgres.InitPostgresClient("./src/configs")
	if err != nil {
		log.Fatalf("failed to initialize db %s", err)
		return
	}
	dbClient := postgres.GetPostgresClient()
	if err := dbClient.GormDB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed error %s", err)
		return
	}
	log.Println("Migration successfully done")
}
