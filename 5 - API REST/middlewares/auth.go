package middlewares

import (
	"gabriel/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authenticate é o middleware que valida o token JWT da requisição e libera o acesso às rotas protegidas.
func Authenticate(context *gin.Context) {
	// Ler o token enviado no header
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Validar o token e extrair o ID do usuário
	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Disponibilizar o ID do usuário para os próximos handlers
	context.Set("userID", userID)
	context.Next()
}
