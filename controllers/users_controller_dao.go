package controllers

import (
	"net/http"

	"github.com/danial2026/golang-proj/domain"
	"github.com/gin-gonic/gin"
)

var (
	UsersController = usersController{}
)

type UsersController struct{}

func (controller usersController) StoreNewUser(c *gin.Context) {

	value, ok := c.Request.URL.Query()["email"]
	if !ok || len(value[0]) < 1 {
		UsersController.BadRequestResponse(c)
		return
	}

	c.Header("Content-Type", "application/json")

	var newUser domain.User

	newUser.Email = string(value[0])

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": newUser,
	})
}
