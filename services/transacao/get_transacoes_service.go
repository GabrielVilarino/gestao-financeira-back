package transacao

import (
	"fmt"
	"time"

	"github.com/GabrielVilarino/gestao-financeira-back.git/exceptions"
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetTransacoesService(idUsuario int, dataInicio, dataFim *string, tipo *string, idGrupo *int) ([]schemas.GetTransacaoResponse, error) {
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

	if tipo != nil && *tipo != "RECEITA" && *tipo != "DESPESA" {
		return nil, &exceptions.ValidationError{Message: "tipo deve ser RECEITA ou DESPESA"}
	}

	transacoes, err := repository.GetTransacoesRepository(idUsuario, dataInicio, dataFim, tipo, idGrupo)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar transações: %w", err)
	}

	return transacoes, nil
}

func GetTransacaoByIDService(id string) (*schemas.GetTransacaoByIDResponse, error) {
	transacao, err := repository.GetTransacaoByIDRepository(id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar transação: %w", err)
	}

	response := transacaoEntityToGetByIDSchema(*transacao)
	return &response, nil
}
