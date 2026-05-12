package subcategoria

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/subcategoria"
	"github.com/gin-gonic/gin"
)

func GetSubcategoriasRoute(c *gin.Context) {
	var idCategoria *int

	if val := c.Query("id_categoria"); val != "" {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id_categoria inválido",
			})
			return
		}
		idCategoria = &parsed
	}

	var idGrupo *int
	if c.Query("group") == "true" {
		idGroupRaw := c.GetInt("id_group")
		if idGroupRaw != 0 {
			idGrupo = &idGroupRaw
		}
	}

	response, err := subcategoria.GetSubcategoriasController(idCategoria, idGrupo)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar subcategorias",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
