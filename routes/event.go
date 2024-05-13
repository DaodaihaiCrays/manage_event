package routes

import (
	"example/rest_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int
	Message string
}

var globalResponseSuccess = Response{
	Code:    200,
	Message: "Success",
}

var globalResponseFail = Response{
	Code:    400,
	Message: "Fail",
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get event data"})
		return
	}
	// context.JSON(http.StatusOK, gin.H{"message": "Hello"})
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get an event data by id"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": globalResponseSuccess, "data": event})
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event) //Convert JSON to object

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Fail to convert JSON to object"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't insert event data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": globalResponseSuccess, "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if event.UserId != userId{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not authorized to update event"})
		return
	}
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get an event data by id"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Fail to convert JSON to object"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't update an event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": globalResponseSuccess})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get event id"})
		
		return
	}

	// var event models.Event
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get an event data by id"})
		return
	}

	if event.UserId != userId{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not authorized to delete event"})
		return
	}
	// event.ID = eventId
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't delete an event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": globalResponseSuccess})
}