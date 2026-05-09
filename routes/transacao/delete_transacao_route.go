package transacao

import (
	"net/http"
	"strconv"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	transacaoController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/transacao"
	"github.com/gin-gonic/gin"
)

func DeleteTransacaoRoute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	if err := transacaoController.DeleteTransacaoController(id); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao excluir transação",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transação excluída com sucesso",
	})
}
