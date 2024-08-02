package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"millanuka.com/personal-website-backend/routes"
)

func rootUrl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This works"})
}

func main() {
	router := routes.InitRouter()

	router.Run("localhost:8000")
}
