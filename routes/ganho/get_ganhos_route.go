package ganho

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	ganhoController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/ganho"
	"github.com/gin-gonic/gin"
)

func GetGanhosRoute(c *gin.Context) {
	idUsuario := c.GetInt("user_id")

	var dataInicio, dataFim *string
	var idGrupo *int

	if val := c.Query("data_inicio"); val != "" {
		dataInicio = &val
	}
	if val := c.Query("data_fim"); val != "" {
		dataFim = &val
	}
	if val := c.Query("id_grupo"); val != "" {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id_grupo inválido",
			})
			return
		}
		idGrupo = &parsed
	}

	response, err := ganhoController.GetGanhosController(idUsuario, dataInicio, dataFim, idGrupo)
	if err != nil {
		configs.Log.Error(err)
		if ganhoController.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar ganhos",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetGanhosByIDRoute(c *gin.Context) {
	id := c.Param("id")

	response, err := ganhoController.GetGanhosByIDController(id)
	if err != nil {
		configs.Log.Error(err)
		if ganhoController.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar ganhos",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
