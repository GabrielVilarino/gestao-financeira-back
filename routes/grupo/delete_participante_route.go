package grupo

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/grupo"
	"github.com/gin-gonic/gin"
)

func DeleteParticipanteRoute(c *gin.Context) {
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

	idParticipante, err := strconv.Atoi(c.Param("id_participante"))
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do participante inválido",
		})
		return
	}

	if idParticipante == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID do participante é obrigatório",
		})
		return
	}

	err = grupo.DeleteParticipanteController(idParticipante, idGrupo)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao remover participante",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Participante removido com sucesso",
	})
}
