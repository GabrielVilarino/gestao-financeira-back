package despesa

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/despesa"
)

func GetDespesasController(idUsuario int, dataInicio, dataFim *string, idGrupo *int) ([]schemas.GetDespesaResponse, error) {
	response, err := despesa.GetDespesasService(idUsuario, dataInicio, dataFim, idGrupo)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetDespesaByIDController(id string) (*schemas.GetDespesaByIDResponse, error) {
	return despesa.GetDespesaByIDService(id)
}
