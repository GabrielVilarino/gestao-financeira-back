package user

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/user"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func AuthUserRoute(c *gin.Context) {
	var request schemas.AuthUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados de autenticação inválidos",
		})
		return
	}

	response, err := user.AuthUserController(request)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credenciais inválidas",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
