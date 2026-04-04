package despesa

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/despesa"
)

func CreateDespesaController(idUsuario int, idGrupoJWT *int, request schemas.CreateDespesaRequest) (*schemas.CreateDespesaResponse, error) {
	response, err := despesa.CreateDespesaService(idUsuario, idGrupoJWT, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
