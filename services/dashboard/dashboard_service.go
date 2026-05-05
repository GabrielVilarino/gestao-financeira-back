package dashboard

import (
	"time"

	"github.com/GabrielVilarino/gestao-financeira-back.git/exceptions"
	baseServices "github.com/GabrielVilarino/gestao-financeira-back.git/services"
)

func validateDashboardPeriod(dataInicio, dataFim *string) error {
	if dataInicio == nil || dataFim == nil {
		return &exceptions.ValidationError{Message: "data_inicio e data_fim são obrigatórias para o dashboard"}
	}

	inicio, err := time.Parse("2006-01-02", *dataInicio)
	if err != nil {
		return &exceptions.ValidationError{Message: "data_inicio inválida, use o formato AAAA-MM-DD"}
	}

	fim, err := time.Parse("2006-01-02", *dataFim)
	if err != nil {
		return &exceptions.ValidationError{Message: "data_fim inválida, use o formato AAAA-MM-DD"}
	}

	if inicio.After(fim) {
		return &exceptions.ValidationError{Message: "data_inicio não pode ser maior que data_fim"}
	}

	return nil
}

func resolveDashboardGroup(idGrupoRequest, idGrupoJWT *int) (*int, error) {
	if err := baseServices.ValidarIDGrupo(idGrupoRequest, idGrupoJWT); err != nil {
		return nil, err
	}

	return idGrupoRequest, nil
}
