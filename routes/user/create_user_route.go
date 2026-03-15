package user

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/user"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateUserRoute(c *gin.Context) {
	var request schemas.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao criar usuário",
		})
		return
	}

	response, err := user.CreateUserController(request)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar usuário",
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}
