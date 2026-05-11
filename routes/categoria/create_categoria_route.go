package categoria

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/categoria"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateCategoriaRoute(c *gin.Context) {
	var request schemas.CreateCategoriaRequest

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

	categoriaResp, err := categoria.CreateCategoriaController(idUsuario, idGrupo, request)
	if err != nil {
		configs.Log.Errorf("Failed to create categoria: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create categoria"})
		return
	}

	c.JSON(http.StatusCreated, categoriaResp)
}
