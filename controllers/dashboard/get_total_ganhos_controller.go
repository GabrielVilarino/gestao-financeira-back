package dashboard

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	dashboardService "github.com/GabrielVilarino/gestao-financeira-back.git/services/dashboard"
)

func GetDashboardTotalGanhosController(idUsuario int, idGrupoJWT *int, dataInicio, dataFim *string, idGrupo *int) (*schemas.GetDashboardTotalGanhosResponse, error) {
	return dashboardService.GetDashboardTotalGanhosService(idUsuario, idGrupoJWT, dataInicio, dataFim, idGrupo)
}
