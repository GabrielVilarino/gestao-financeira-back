package despesa

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/despesa"
)

func UpdateDespesaController(idUsuario int, idGrupo *int, request schemas.UpdateDespesaRequest) (*schemas.UpdateDespesaResponse, error) {
	response, err := despesa.UpdateDespesaService(idUsuario, idGrupo, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
