package controllers

import (
	"net/http"
	"time"

	"github.com/danial2026/golang-proj/domain"
	"github.com/danial2026/golang-proj/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	UsersController = usersController{}
)

type usersController struct{}

func (controller usersController) StoreNewUser(c *gin.Context) {
	value, ok := c.Request.URL.Query()["email"]
	if !ok || len(value[0]) < 1 {
		UsersController.BadRequestResponse(c)
		return
	}

	c.Header("Content-Type", "application/json")

	newEmail := string(value[0])

	users := services.UsersService.GetByEmail(newEmail)
	
	userResponse := &domain.User{}
	statusCode := http.StatusOK

	if len(users) < 1 {
		newUser := &domain.User{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     newEmail,
		}

		services.UsersService.Save(newUser)

		userResponse = newUser
	} else {

		userResponse = users[len(users)-1]
		statusCode = http.StatusBadRequest
	}

	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": userResponse,
	})
}
