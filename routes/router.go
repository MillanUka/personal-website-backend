package routes

import (
	"github.com/gin-gonic/gin"
	"millanuka.com/personal-website-backend/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/send-user-response", controllers.AddUserResponse)

	return router
}
