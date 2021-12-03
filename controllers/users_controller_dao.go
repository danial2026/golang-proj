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
	valueEmail, ok := c.Request.URL.Query()["email"]
	if !ok || len(valueEmail[0]) < 1 {
		UsersController.BadRequestResponse(c)
		return
	}
	newEmail := string(valueEmail[0])

	valueName, ok := c.Request.URL.Query()["name"]
	if !ok || len(valueName[0]) < 1 {
		// UsersController.BadRequestResponse(c)
		return
	}
	newName := string(valueName[0])

	valueDescription, ok := c.Request.URL.Query()["description"]
	if !ok || len(valueDescription[0]) < 1 {
		UsersController.BadRequestResponse(c)
		return
	}
	newDescription := string(valueDescription[0])

	c.Header("Content-Type", "application/json")

	users := services.UsersService.GetByEmail(newEmail)

	statusCode := http.StatusOK

	if len(users) >= 1 {

		UsersController.BadRequestResponse(c)
		return
	}

	newUser := &domain.User{
		ID:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Email:       newEmail,
		Name:        newName,
		Description: newDescription,
	}

	services.UsersService.Save(newUser)

	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": newUser,
	})
}
