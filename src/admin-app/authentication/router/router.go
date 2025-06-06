package router

import (
	"admin-app/authentiction/business"
	"admin-app/authentiction/handlers"
	"admin-app/authentiction/repositories"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.New()

	//db instance

	// router.POST("Create Login", "Handelers")

	createUserRepository := repositories.NewCreateUserWatchlist()
	createUserService := business.NewCreateUserSService(createUserRepository)
	createUserController := handlers.NewCreateUserController(createUserService)

	router.POST("/createUser", createUserController.CreateUserHandler)

	return router
}
