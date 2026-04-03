package ganho

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	ganhoController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/ganho"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateGanhoRoute(c *gin.Context) {
	var request schemas.CreateGanhoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao registrar ganho",
		})
		return
	}

	idUsuario := c.GetInt("user_id")
	idGroupRaw := c.GetInt("id_group")
	var idGrupoJWT *int
	if idGroupRaw != 0 {
		idGrupoJWT = &idGroupRaw
	}

	response, err := ganhoController.CreateGanhoController(idUsuario, idGrupoJWT, request)
	if err != nil {
		configs.Log.Error(err)
		if ganhoController.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao registrar ganho",
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}
