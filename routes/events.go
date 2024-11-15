package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/models"
)

func getEvent(context *gin.Context){
	eventId, err1 := strconv.ParseInt(context.Param("id"), 10, 64)
	if err1 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later.", "error": err1})
		return
	}
	event, err2 := models.GetEventById(eventId)
	if err2 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later.", "error": err2})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events. Try again later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later.", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later.", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later.", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later.", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event. Try again later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}