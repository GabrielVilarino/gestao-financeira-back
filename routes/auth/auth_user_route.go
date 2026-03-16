package auth

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/auth"
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

	response, token, err := auth.AuthUserController(request)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credenciais inválidas",
		})
		return
	}

	c.SetCookie(
		"auth_token",
		*token,
		3600*24,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, response)
}
