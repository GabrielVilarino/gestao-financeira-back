package categoria

import (
	"net/http"
	"strconv"

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

	response, err := categoriaController.GetCategoriasController(tipo)
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
