package routes

import (
	"github.com/gin-gonic/gin"
	"go-events-booking-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)

	authenticatedGroup := server.Group("/")
	authenticatedGroup.Use(middlewares.Authenticate)

	authenticatedGroup.POST("/events", middlewares.Authenticate, createEvent)
	authenticatedGroup.PUT("/events/:id", updateEvent)
	authenticatedGroup.DELETE("/events/:id", deleteEvent)
	authenticatedGroup.POST("events/:id/register", registerForEvent)
	authenticatedGroup.DELETE("events/:id/register", cancelRegistration)
}
