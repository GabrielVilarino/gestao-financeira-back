package despesa

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services"
)

func CreateDespesaService(idUsuario int, idGrupoJWT *int, request schemas.CreateDespesaRequest) (*schemas.CreateDespesaResponse, error) {
	if err := services.ValidarIDGrupo(request.IDGrupo, idGrupoJWT); err != nil {
		return nil, err
	}

	despesa := DespesaSchemaToEntity(request)

	despesa.IDUsuario = idUsuario
	if request.IDGrupo != nil {
		despesa.IDGrupo = request.IDGrupo
	} else {
		despesa.IDGrupo = idGrupoJWT
	}

	despesaCreated, err := repository.CreateDespesaRepository(despesa)
	if err != nil {
		return nil, err
	}

	response := DespesaEntityToCreateSchema(*despesaCreated)

	return &response, nil
}
