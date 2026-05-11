package subcategoria

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/subcategoria"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateSubCategoriaRoute(c *gin.Context) {
	var request schemas.CreateSubCategoriaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	idUsuario := c.GetInt("user_id")

	var idGrupo *int
	if c.Query("group") == "true" {
		idGroupRaw := c.GetInt("id_group")
		if idGroupRaw != 0 {
			idGrupo = &idGroupRaw
		}
	}

	subcategoriaResp, err := subcategoria.CreateSubCategoriaController(idUsuario, idGrupo, request)
	if err != nil {
		configs.Log.Errorf("Failed to create subcategoria: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subcategoria"})
		return
	}

	c.JSON(http.StatusCreated, subcategoriaResp)
}
