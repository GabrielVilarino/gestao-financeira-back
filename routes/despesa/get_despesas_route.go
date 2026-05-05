package despesa

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/despesa"
	"github.com/gin-gonic/gin"
)

func GetDespesasRoute(c *gin.Context) {
	idUsuario := c.GetInt("user_id")

	var dataInicio, dataFim *string
	var idGrupo *int

	if val := c.Query("data_inicio"); val != "" {
		dataInicio = &val
	}

	if val := c.Query("data_fim"); val != "" {
		dataFim = &val
	}

	if c.Query("group") == "true" {
		groupID := c.GetInt("id_group")
		idGrupo = &groupID
	}
	response, err := despesa.GetDespesasController(idUsuario, dataInicio, dataFim, idGrupo)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar despesas",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetDespesaByIDRoute(c *gin.Context) {
	id := c.Param("id")

	response, err := despesa.GetDespesaByIDController(id)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar despesas",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
