package dashboard

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetDashboardTotalDespesasService(idUsuario int, idGrupoJWT *int, dataInicio, dataFim *string, idGrupo *int) (*schemas.GetDashboardTotalDespesasResponse, error) {
	if err := validateDashboardPeriod(dataInicio, dataFim); err != nil {
		return nil, err
	}

	idGrupoResolved, err := resolveDashboardGroup(idGrupo, idGrupoJWT)
	if err != nil {
		return nil, err
	}

	totalDespesas, err := repository.GetDashboardTotalDespesasRepository(idUsuario, idGrupoResolved, *dataInicio, *dataFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar total de despesas do dashboard: %w", err)
	}

	return &schemas.GetDashboardTotalDespesasResponse{TotalDespesas: totalDespesas}, nil
}
