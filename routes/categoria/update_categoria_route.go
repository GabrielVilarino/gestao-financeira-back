package categoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/categoria"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateCategoriaRoute(c *gin.Context) {
	var request schemas.UpdateCategoriaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	idUsuario := c.GetInt("user_id")

	categoria, err := categoria.UpdateCategoriaController(idUsuario, request)
	if err != nil {
		configs.Log.Errorf("Failed to update categoria: %v", err)
		c.JSON(500, gin.H{"error": "Failed to update categoria"})
		return
	}

	c.JSON(200, categoria)
}
