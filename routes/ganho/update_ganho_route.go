package ganho

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	ganhoController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/ganho"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateGanhoRoute(c *gin.Context) {
	var request schemas.UpdateGanhoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao editar ganho",
		})
		return
	}

	response, err := ganhoController.UpdateGanhoController(request)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao editar ganho",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
