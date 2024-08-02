package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootUrl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This works"})
}

func main() {
	router := gin.Default()
	router.GET("/", rootUrl)

	router.Run("localhost:8000")
}
