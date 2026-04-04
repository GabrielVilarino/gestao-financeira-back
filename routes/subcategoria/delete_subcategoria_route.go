package subcategoria

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/subcategoria"
	"github.com/gin-gonic/gin"
)

func DeleteSubCategoriaRoute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	idUsuario := c.GetInt("user_id")

	if err := subcategoria.DeleteSubCategoriaController(idUsuario, id); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao excluir subcategoria",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Subcategoria excluída com sucesso",
	})
}
