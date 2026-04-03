package ganho

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	ganhoController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/ganho"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateGanhoRoute(c *gin.Context) {
	var request schemas.UpdateGanhoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao editar ganho",
		})
		return
	}

	idUsuario := c.GetInt("user_id")
	idGroupRaw := c.GetInt("id_group")
	var idGrupo *int
	if idGroupRaw != 0 {
		idGrupo = &idGroupRaw
	}

	response, err := ganhoController.UpdateGanhoController(idUsuario, idGrupo, request)
	if err != nil {
		configs.Log.Error(err)
		if ganhoController.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao editar ganho",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
