package categoria

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	categoriaController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/categoria"
	"github.com/gin-gonic/gin"
)

func GetCategoriasRoute(c *gin.Context) {
	response, err := categoriaController.GetCategoriasController()
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar categorias",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

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

	response, err := categoriaController.GetSubcategoriasController(idCategoria)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar subcategorias",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
