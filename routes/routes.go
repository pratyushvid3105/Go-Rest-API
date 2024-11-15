package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/middlewares"
)

// Here, we want to get our server, we want to get a pointer to the server to be precise, so at the gin.Engine like this. Because we can then use this server to register those routes just as we did it in the main.go file with server.GET and server.POST and so on.
func RegisterRoutes(server *gin.Engine){
	server.GET("/events/:id", getEvent)
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.GET("/users", getUsers)
	server.GET("/registrations", getRegistrations)

	// We'll use authentication on multiple routes, also on PUT and DELETE, for example, there is an easier approach we can use. We can use that server here and call the Group method to create a group of routes. Now, Group then wants a path with which all grouped routes will start and here we'll just use "/" because that's the one thing all grouped routes will have in common. Then as a result, we'll get back a router group, which we should store in a variable (here "authenticated"). And now we can use that router group here to add routes to it.
	authenticated := server.Group("/")

	// To protect the routes, we should also use that Group to set up a middleware that will always be executed, and we do that with the "Use" method.
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// We will add more routes that belong to this authenticated group, because both registering and canceling should only be possible for locked in users.
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
	// So that's how we now register routes in this function. And since we're always operating on exactly the same server value, since we're using a pointer here, we don't have to return anything here or do anything like that. Instead we are manipulating the original server when this function (RegisterRoutes) here is executed.
}