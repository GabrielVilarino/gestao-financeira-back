package subcategoria

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/subcategoria"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateSubCategoriaRoute(c *gin.Context) {
	var request schemas.UpdateSubCategoriaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	idUsuario := c.GetInt("user_id")

	subcategoria, err := subcategoria.UpdateSubCategoriaController(idUsuario, request)
	if err != nil {
		configs.Log.Errorf("Failed to update subcategoria: %v", err)
		c.JSON(500, gin.H{"error": "Failed to update subcategoria"})
		return
	}

	c.JSON(200, subcategoria)
}
