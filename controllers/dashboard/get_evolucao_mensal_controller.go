package dashboard

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	dashboardService "github.com/GabrielVilarino/gestao-financeira-back.git/services/dashboard"
)

func GetDashboardEvolucaoMensalController(idUsuario int, idGrupoJWT *int, dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetDashboardEvolucaoMensalItemResponse, error) {
	return dashboardService.GetDashboardEvolucaoMensalService(idUsuario, idGrupoJWT, dataInicio, dataFim, idGrupo)
}
