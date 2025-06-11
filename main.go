package main

import (
	"example.com/m/v2/database"
	"example.com/m/v2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
