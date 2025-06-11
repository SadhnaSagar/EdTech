package routes

import (
	"example.com/m/v2/controllers"
	"example.com/m/v2/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)

	protected := router.Group("/api")
	protected.Use(middleware.Authorize("read:dashboard"))
	{
		protected.GET("/dashboard", controllers.Protected)
	}
}
