package main

import (
	"example/rest_api/db"
	"example/rest_api/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

