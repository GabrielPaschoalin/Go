package routes

import (
	"gabriel/api/models"
	"gabriel/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup lê os dados do corpo da requisição e cria um novo usuário.
func signup(context *gin.Context) {

	// Ler e validar o corpo da requisição
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Salvar o usuário no banco
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

// login valida as credenciais informadas e retorna um token JWT em caso de sucesso.
func login(context *gin.Context) {
	// Ler e validar o corpo da requisição
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Verificar email e senha
	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	// Gerar token de autenticação
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!.", "token": token})

}
