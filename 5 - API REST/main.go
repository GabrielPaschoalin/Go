package main

import (
	"gabriel/api/db"
	"gabriel/api/routes"

	"github.com/gin-gonic/gin"
)

// main inicializa o banco de dados, registra as rotas e sobe o servidor HTTP.
func main() {
	// Conectar ao banco e garantir que as tabelas existam
	db.InitDB()

	server := gin.Default()

	// Registrar todas as rotas da API
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
