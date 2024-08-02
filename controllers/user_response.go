package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"millanuka.com/personal-website-backend/models"
)

func AddUserResponse(context *gin.Context) {
	var userResponse models.UserResponse

	if err := context.BindJSON(&userResponse); err != nil {
		return
	}

	// TODO: Add saving to CSV file here

	context.JSON(http.StatusCreated, userResponse)

}
