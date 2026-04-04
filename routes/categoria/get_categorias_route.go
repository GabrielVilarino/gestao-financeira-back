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

	response, err := categoriaController.GetCategoriasController(idUsuario, tipo)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar categorias",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
