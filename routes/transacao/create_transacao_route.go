package transacao

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	transacaoController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/transacao"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTransacaoRoute(c *gin.Context) {
	var request schemas.CreateTransacaoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao registrar transação",
		})
		return
	}

	idUsuario := c.GetInt("user_id")
	idGroupRaw := c.GetInt("id_group")
	var idGrupoJWT *int
	if idGroupRaw != 0 {
		idGrupoJWT = &idGroupRaw
	}

	response, err := transacaoController.CreateTransacaoController(idUsuario, idGrupoJWT, request)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao registrar transação",
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}
