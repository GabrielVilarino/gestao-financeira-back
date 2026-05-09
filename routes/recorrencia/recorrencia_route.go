package recorrencia

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	recorrenciaController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/recorrencia"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateRecorrenciaRoute(c *gin.Context) {
	var req schemas.CreateRecorrenciaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao registrar recorrência",
		})
		return
	}

	idUsuario := c.GetInt("user_id")
	idGroupRaw := c.GetInt("id_group")
	var idGrupoJWT *int
	if idGroupRaw != 0 {
		idGrupoJWT = &idGroupRaw
	}

	response, err := recorrenciaController.CreateRecorrenciaController(idUsuario, idGrupoJWT, req)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao registrar recorrência",
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func GetRecorrenciasRoute(c *gin.Context) {
	idUsuario := c.GetInt("user_id")
	idGroupRaw := c.GetInt("id_group")
	var idGrupoJWT *int
	if idGroupRaw != 0 {
		idGrupoJWT = &idGroupRaw
	}

	var idGrupo *int
	if c.Query("group") == "true" && idGrupoJWT != nil {
		idGrupo = idGrupoJWT
	}

	response, err := recorrenciaController.GetRecorrenciasController(idUsuario, idGrupoJWT, idGrupo)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar recorrências",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func EncerrarRecorrenciaRoute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	if err := recorrenciaController.EncerrarRecorrenciaController(id); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao encerrar recorrência",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Recorrência encerrada com sucesso",
	})
}
