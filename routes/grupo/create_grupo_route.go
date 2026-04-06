package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers/grupo"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateGrupoRoute(c *gin.Context) {
	var request schemas.CreateGrupoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idUsuario := c.GetInt("user_id")

	// Validar se o usuario possui grupo, se possui nao pode criar
	idGrupo := c.GetInt("id_group")
	if idGrupo != 0 {
		c.JSON(400, gin.H{"error": "É necessário sair do grupo atual para criar um novo grupo"})
		return
	}

	grupo, err := grupo.CreateGrupoController(idUsuario, request)
	if err != nil {
		configs.Log.Errorf("Erro ao criar grupo: %v", err)
		c.JSON(500, gin.H{"error": "Erro ao criar grupo"})
		return
	}

	c.JSON(201, grupo)
}
