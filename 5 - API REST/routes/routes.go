package routes

import (
	"gabriel/api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)    // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // events/1

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // Garante que a autenticação será calculado para todos do grupo
	authenticated.POST("/events", createEvent)  // Executa da esquerda para direita
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
