package dashboard

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	dashboardService "github.com/GabrielVilarino/gestao-financeira-back.git/services/dashboard"
)

func GetDashboardSaldoLiquidoController(idUsuario int, idGrupoJWT *int, dataInicio, dataFim *string, idGrupo *int) (*schemas.GetDashboardSaldoLiquidoResponse, error) {
	return dashboardService.GetDashboardSaldoLiquidoService(idUsuario, idGrupoJWT, dataInicio, dataFim, idGrupo)
}
