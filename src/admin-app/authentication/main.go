package main

import (
	"InterviewPractice/src/utils/postgres"
	"admin-app/authentiction/router"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

func main() {

	if err := postgres.InitPostgresClient("../../configs"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("postgres client initialized")
	startRouter()
}

func startRouter() {
	logger := logrus.New()
	router := router.GetRouter()
	logger.Info(fmt.Sprint("Running on port ", 8080))
	router.Run(fmt.Sprintf(":%d", 8080))
}
