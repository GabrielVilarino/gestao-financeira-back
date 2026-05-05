package dashboard

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetDashboardTotalGanhosService(idUsuario int, idGrupoJWT *int, dataInicio, dataFim *string, idGrupo *int) (*schemas.GetDashboardTotalGanhosResponse, error) {
	if err := validateDashboardPeriod(dataInicio, dataFim); err != nil {
		return nil, err
	}

	idGrupoResolved, err := resolveDashboardGroup(idGrupo, idGrupoJWT)
	if err != nil {
		return nil, err
	}

	totalGanhos, err := repository.GetDashboardTotalGanhosRepository(idUsuario, idGrupoResolved, *dataInicio, *dataFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar total de ganhos do dashboard: %w", err)
	}

	return &schemas.GetDashboardTotalGanhosResponse{TotalGanhos: totalGanhos}, nil
}
