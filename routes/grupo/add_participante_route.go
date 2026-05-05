package grupo

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/grupo"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func AddParticipanteRoute(c *gin.Context) {
	idGrupo, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	idGrupoJWT := c.GetInt("id_group")

	if idGrupo != idGrupoJWT {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Não autorizado",
		})
		return
	}

	request := schemas.AddParticipanteRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados de participante inválidos",
		})
		return
	}

	err = grupo.AddParticipanteController(request, idGrupo)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao adicionar participante",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Participante adicionado com sucesso",
	})
}
