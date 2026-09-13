package routes

import (
	"gabriel/api/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes declara todas as rotas da API, separando as públicas das que exigem autenticação.
func RegisterRoutes(server *gin.Engine) {

	// Rotas públicas de eventos (leitura)
	server.GET("/events", getEvents)    // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // events/1

	// Rotas que exigem usuário autenticado
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // Garante que a autenticação será calculado para todos do grupo
	authenticated.POST("/events", createEvent)  // Executa da esquerda para direita
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// Rotas públicas de autenticação
	server.POST("/signup", signup)
	server.POST("/login", login)
}
