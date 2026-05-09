package transacao

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	transacaoController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/transacao"
	"github.com/gin-gonic/gin"
)

func GetTransacoesRoute(c *gin.Context) {
	idUsuario := c.GetInt("user_id")

	var dataInicio, dataFim, tipo *string
	var idGrupo *int

	if val := c.Query("data_inicio"); val != "" {
		dataInicio = &val
	}
	if val := c.Query("data_fim"); val != "" {
		dataFim = &val
	}
	if val := c.Query("tipo"); val != "" {
		tipo = &val
	}
	if c.Query("group") == "true" {
		groupID := c.GetInt("id_group")
		idGrupo = &groupID
	}

	response, err := transacaoController.GetTransacoesController(idUsuario, dataInicio, dataFim, tipo, idGrupo)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar transações",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetTransacaoByIDRoute(c *gin.Context) {
	id := c.Param("id")

	response, err := transacaoController.GetTransacaoByIDController(id)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar transação",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
