package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/db"
	"github.com/pratyushvid3105/Go-Rest-API/models"
)

func main(){
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	server.POST("/events", createEvent)

	server.Run(":8083") // localhost:8080
}

func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}