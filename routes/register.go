package routes

import (
	"github.com/gin-gonic/gin"
	"go-events-booking-api/models"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to register event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered"})
	return
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
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

	if userId != event.UserID {
		context.JSON(http.StatusForbidden, gin.H{"message": "userId and eventId do not match"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to cancel registration"})
		return
	}
}
