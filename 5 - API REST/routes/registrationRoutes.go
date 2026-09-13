package routes

import (
	"gabriel/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	// Obter o id indicado no requisição
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	// Obter o event indicado pelo ID e verifiar se existe
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find ID."})
		return
	}

	// Obter id do usuário
	userId := context.GetInt64("userId")

	// Criar o registro no banco
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	// Mensagem de sucesso
	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})

}

func cancelRegistration(context *gin.Context) {

	// Obter o id indicado no requisição
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	var event models.Event
	event.ID = eventId

	// Obter id do usuário
	userId := context.GetInt64("userId")

	err = event.DeleteRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete user registration for event."})
		return
	}

	// Mensagem de sucesso
	context.JSON(http.StatusOK, gin.H{"message": "Registration deleted!"})

}
