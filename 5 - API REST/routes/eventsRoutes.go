package routes

import (
	"fmt"
	"gabriel/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getEvents responde com a lista de todos os eventos cadastrados.
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}
	// context.JSON(http.StatusOK, gin.H{"e": "Hello!"})
	context.JSON(http.StatusOK, events)
}

// createEvent lê o evento do corpo da requisição e salva no banco, associado ao usuário logado.
func createEvent(context *gin.Context) {

	// Ler e validar o corpo da requisição
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Associar o evento ao usuário autenticado
	userID := context.GetInt64("userID")
	event.UserID = userID

	// Salvar no banco
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event}) // Codigo 201

}

// getEvent responde com um único evento, identificado pelo ID na URL.
func getEvent(context *gin.Context) {
	// Obter o id indicado na requisição
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	// Buscar o evento no banco
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)

}

// updateEvent atualiza um evento existente, apenas se o usuário logado for o criador.
func updateEvent(context *gin.Context) {

	// Obter o id indicado no requisição
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	// Obter o event indicado pelo ID
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	// Verificar se o usuário logado é o mesmo que criou o evento
	userID := context.GetInt64("userID")
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	// Ler o body da requisição e criar o evento atualizado
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	updatedEvent.ID = eventId

	// Atualizar evento
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

// deleteEvent remove um evento existente, apenas se o usuário logado for o criador.
func deleteEvent(context *gin.Context) {

	// Obter o id indicado no requisição
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	// Obter o event indicado pelo ID
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find ID."})
		return
	}

	// Verificar se o usuário logado é o mesmo que criou o evento
	userID := context.GetInt64("userID")
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	// Deletar evento do banco
	err = event.DeleteEvent()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event. "})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
