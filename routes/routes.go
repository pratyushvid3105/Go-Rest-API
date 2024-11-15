package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/middlewares"
)

// Here, we want to get our server, we want to get a pointer to the server to be precise, so at the gin.Engine like this. Because we can then use this server to register those routes just as we did it in the main.go file with server.GET and server.POST and so on.
func RegisterRoutes(server *gin.Engine){
	server.GET("/events/:id", getEvent)
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	// The first argument after the path, before createEvent, we can use the middlewares package And then refer to that Authenticate function like this because with Gin, we can register multiple request handlers and they will be executed from left to right.
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users", getUsers)
	// So that's how we now register routes in this function. And since we're always operating on exactly the same server value, since we're using a pointer here, we don't have to return anything here or do anything like that. Instead we are manipulating the original server when this function (RegisterRoutes) here is executed.
}