package ganho

import (
	"fmt"
	"time"

	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func GetGanhosService(dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetGanhoResponse, error) {
	if dataInicio != nil {
		if _, err := time.Parse("2006-01-02", *dataInicio); err != nil {
			return nil, &ValidationError{Message: "data_inicio inválida, use o formato AAAA-MM-DD"}
		}
	}
	if dataFim != nil {
		if _, err := time.Parse("2006-01-02", *dataFim); err != nil {
			return nil, &ValidationError{Message: "data_fim inválida, use o formato AAAA-MM-DD"}
		}
	}
	if dataInicio != nil && dataFim != nil {
		inicio, _ := time.Parse("2006-01-02", *dataInicio)
		fim, _ := time.Parse("2006-01-02", *dataFim)
		if inicio.After(fim) {
			return nil, &ValidationError{Message: "data_inicio não pode ser maior que data_fim"}
		}
	}

	ganhos, err := repository.GetGanhosRepository(dataInicio, dataFim, idGrupo)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar ganhos: %w", err)
	}

	var response []schemas.GetGanhoResponse
	for _, g := range ganhos {
		response = append(response, schemas.GetGanhoResponse{
			ID:              g.ID,
			IDUsuario:       g.IDUsuario,
			IDGrupo:         g.IDGrupo,
			IDCategoria:     g.IDCategoria,
			IDSubcategoria:  g.IDSubcategoria,
			TipoTransacao:   g.TipoTransacao,
			Valor:           g.Valor,
			DataRecebimento: g.DataRecebimento,
			DataCriacao:     g.DataCriacao,
		})
	}

	return response, nil
}
