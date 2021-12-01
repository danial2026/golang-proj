package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (controller usersController) BadRequestResponse(c *gin.Context) {

    UsersController.JsonResponse(c, http.StatusBadRequest, "Bad Request")
}

func (controller usersController) NotFoundResponse(c *gin.Context) {

    UsersController.JsonResponse(c, http.StatusNotFound, "Not Found")
}

func (controller usersController) JsonResponse(c *gin.Context, http_status int, messege string) {

    c.Header("Content-Type", "application/json")

    var response_json response

    response_json.Status=http_status
    response_json.Message=messege

    time_stamp , _ := time.Now().UTC().MarshalText()
    response_json.Timestamp=string(time_stamp)

    response_json.Path=c.Request.URL.Path
    c.JSON(response_json.Status, response_json)
}