package dashboard

import (
	"fmt"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetDashboardSaldoLiquidoService(idUsuario int, idGrupoJWT *int, dataInicio, dataFim *string, idGrupo *int) (*schemas.GetDashboardSaldoLiquidoResponse, error) {
	if err := validateDashboardPeriod(dataInicio, dataFim); err != nil {
		return nil, err
	}

	idGrupoResolved, err := resolveDashboardGroup(idGrupo, idGrupoJWT)
	if err != nil {
		return nil, err
	}

	saldoLiquido, err := repository.GetDashboardSaldoLiquidoRepository(idUsuario, idGrupoResolved, *dataInicio, *dataFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar saldo líquido do dashboard: %w", err)
	}

	return &schemas.GetDashboardSaldoLiquidoResponse{SaldoLiquido: saldoLiquido}, nil
}
