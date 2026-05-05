package dashboard

import (
	"net/http"

	"github.com/GabrielVilarino/gestao-financeira-back.git/configs"
	"github.com/GabrielVilarino/gestao-financeira-back.git/controllers"
	dashboardController "github.com/GabrielVilarino/gestao-financeira-back.git/controllers/dashboard"
	"github.com/gin-gonic/gin"
)

func GetDashboardSaldoLiquidoRoute(c *gin.Context) {
	filters, err := getDashboardFilters(c)
	if err != nil {
		configs.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idUsuario := c.GetInt("user_id")
	response, err := dashboardController.GetDashboardSaldoLiquidoController(
		idUsuario,
		filters.idGrupoJWT,
		filters.dataInicio,
		filters.dataFim,
		filters.idGrupo,
	)
	if err != nil {
		configs.Log.Error(err)
		if controllers.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar saldo líquido do dashboard",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
