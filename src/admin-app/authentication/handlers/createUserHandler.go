package handlers

import (
	"admin-app/authentiction/buissness"
	"admin-app/authentiction/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	service *buissness.CreateUserService
}

func NewCreateUserController(service *buissness.CreateUserService) *CreateUserController {
	return &CreateUserController{
		service: service,
	}
}

func (controller *CreateUserController) CreateUserHandler(c *gin.Context) {
	var bffCreateUserRequest models.BFFCreateUserRequest
	if err := c.ShouldBindJSON(&bffCreateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.service.CreateNewUser(c, c, bffCreateUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
