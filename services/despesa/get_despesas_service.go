package despesa

import (
	"fmt"
	"time"

	"github.com/GabrielVilarino/gestao-financeira-back.git/exceptions"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetDespesasService(idUsuario int, dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetDespesaResponse, error) {
	if dataInicio != nil {
		if _, err := time.Parse("2006-01-02", *dataInicio); err != nil {
			return nil, &exceptions.ValidationError{Message: "data_inicio inválida, use o formato AAAA-MM-DD"}
		}
	}
	if dataFim != nil {
		if _, err := time.Parse("2006-01-02", *dataFim); err != nil {
			return nil, &exceptions.ValidationError{Message: "data_fim inválida, use o formato AAAA-MM-DD"}
		}
	}
	if dataInicio != nil && dataFim != nil {
		inicio, _ := time.Parse("2006-01-02", *dataInicio)
		fim, _ := time.Parse("2006-01-02", *dataFim)
		if inicio.After(fim) {
			return nil, &exceptions.ValidationError{Message: "data_inicio não pode ser maior que data_fim"}
		}
	}

	despesas, err := repository.GetDespesasRepository(idUsuario, dataInicio, dataFim, idGrupo)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar despesas: %w", err)
	}

	return despesas, nil
}

func GetDespesaByIDService(id string) (*schemas.GetDespesaByIDResponse, error) {
	despesa, err := repository.GetDespesasByIDRepository(id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar despesas: %w", err)
	}

	response := DespesaEntityToGetSchema(*despesa)

	return &response, nil
}
