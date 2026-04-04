package despesa

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/despesa"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateDespesaRoute(c *gin.Context) {
	var request schemas.UpdateDespesaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao editar despesa",
		})
		return
	}

	idUsuario := c.GetInt("user_id")
	idGroupRaw := c.GetInt("id_group")
	var idGrupo *int
	if idGroupRaw != 0 {
		idGrupo = &idGroupRaw
	}

	response, err := despesa.UpdateDespesaController(idUsuario, idGrupo, request)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao editar despesa",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
