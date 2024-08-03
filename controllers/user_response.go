package controllers

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"millanuka.com/personal-website-backend/models"
)

func AddUserResponse(context *gin.Context) {
	var userResponse models.UserResponse

	if err := context.BindJSON(&userResponse); err != nil {
		return
	}

	file, err := os.OpenFile("./data/user_responses.csv", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Print("There was an error reading in the file\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error while saving response"})
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{userResponse.Email, userResponse.Response})

	context.JSON(http.StatusCreated, userResponse)

}
