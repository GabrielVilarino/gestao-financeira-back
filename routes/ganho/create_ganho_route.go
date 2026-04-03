package ganho

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	ganhoController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/ganho"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateGanhoRoute(c *gin.Context) {
	var request schemas.CreateGanhoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao registrar ganho",
		})
		return
	}

	response, err := ganhoController.CreateGanhoController(request)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao registrar ganho",
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}
