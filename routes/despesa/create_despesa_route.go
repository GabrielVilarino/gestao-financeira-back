package despesa

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/despesa"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateDespesaRoute(c *gin.Context) {
	var request schemas.CreateDespesaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Erro ao registrar despesa",
		})
		return
	}

	idUsuario := c.GetInt("user_id")
	idGroupRaw := c.GetInt("id_group")
	var idGrupoJWT *int
	if idGroupRaw != 0 {
		idGrupoJWT = &idGroupRaw
	}

	response, err := despesa.CreateDespesaController(idUsuario, idGrupoJWT, request)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"error": "Erro ao registrar despesa",
		})
		return
	}

	c.JSON(201, response)
}
