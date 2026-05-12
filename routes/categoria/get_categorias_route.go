package categoria

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	categoriaController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/categoria"
	"github.com/gin-gonic/gin"
)

func GetCategoriasRoute(c *gin.Context) {
	var tipo *string

	if val := c.Query("tipo"); val != "" {
		if val != "receita" && val != "despesa" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "tipo inválido",
			})
			return
		}
		tipo = &val
	}

	idUsuario := c.GetInt("user_id")

	var idGrupo *int
	if c.Query("group") == "true" {
		idGroupRaw := c.GetInt("id_group")
		if idGroupRaw != 0 {
			idGrupo = &idGroupRaw
		}
	}

	response, err := categoriaController.GetCategoriasController(idUsuario, tipo, idGrupo)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar categorias",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
