package handlers

import (
	"admin-app/authentiction/business"
	"admin-app/authentiction/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type createUserController struct {
	service *business.CreateUserService
}

// @Summary Create new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.BFFCreateUserRequest true "User object"
// @Success 200 {object} models.BFFCreateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /createUser [post]
func NewCreateUserController(service *business.CreateUserService) *createUserController {
	return &createUserController{
		service: service,
	}
}

func (controller *createUserController) CreateUserHandler(c *gin.Context) {
	var bffCreateUserRequest models.BFFCreateUserRequest
	if err := c.ShouldBindJSON(&bffCreateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&bffCreateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.service.CreateNewUser(c, c, bffCreateUserRequest)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			errorResponse := models.ErrorAPIResponse{
				Error:   http.StatusNotFound,
				Message: "Record not found",
			}
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}
		if strings.Contains(err.Error(), "record not found") {
			errorResponse := models.ErrorMessage{
				Key:         http.StatusNotFound,
				ErrorMessage: "Record not found",
			}
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
