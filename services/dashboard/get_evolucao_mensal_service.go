package dashboard

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetDashboardEvolucaoMensalService(idUsuario int, idGrupoJWT *int, dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetDashboardEvolucaoMensalItemResponse, error) {
	if err := validateDashboardPeriod(dataInicio, dataFim); err != nil {
		return nil, err
	}

	idGrupoResolved, err := resolveDashboardGroup(idGrupo, idGrupoJWT)
	if err != nil {
		return nil, err
	}

	evolucaoMensal, err := repository.GetDashboardEvolucaoMensalRepository(idUsuario, idGrupoResolved, *dataInicio, *dataFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar evolução mensal do dashboard: %w", err)
	}

	return evolucaoMensal, nil
}
