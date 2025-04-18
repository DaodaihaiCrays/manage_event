package routes

import (
	"example/rest_api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Cách 1 dùng middlewares
	// server.POST("/events", middlewares.Authenticate, createEvent)

	// Cách 2
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("events/:id/register", RegisterForEvent)
	authenticated.DELETE("events/:id/register", CancelRegistration)
	
	server.POST("/signup", Signup)
	server.POST("/login", Login)
}