package dashboard

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	dashboardService "github.com/GabrielVilarino/gestao-financeira-back.git/services/dashboard"
)

func GetDashboardTotalDespesasController(idUsuario int, idGrupoJWT *int, dataInicio, dataFim *string, idGrupo *int) (*schemas.GetDashboardTotalDespesasResponse, error) {
	return dashboardService.GetDashboardTotalDespesasService(idUsuario, idGrupoJWT, dataInicio, dataFim, idGrupo)
}
