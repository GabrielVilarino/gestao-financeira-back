package despesa

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/despesa"
	"github.com/gin-gonic/gin"
)

func DeleteDespesaRoute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	if err := despesa.DeleteDespesaController(id); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao excluir despesa",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Despesa excluída com sucesso",
	})
}
