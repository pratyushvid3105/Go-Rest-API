package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/middlewares"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events/:id", getEvent)
	server.GET("/events", getEvents)
	server.GET("/users", getUsers)
	server.GET("/registrations", getRegistrations)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}