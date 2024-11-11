package routes

import (
	"github.com/gin-gonic/gin"
	"go-events-booking-api/models"
	"net/http"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not retrieve event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, event)
	return
}

func updateEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	event, err := updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
	return
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
	return
}
