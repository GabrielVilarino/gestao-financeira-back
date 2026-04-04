package despesa

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services"
)

func UpdateDespesaService(idUsuario int, idGrupoJWT *int, request schemas.UpdateDespesaRequest) (*schemas.UpdateDespesaResponse, error) {
	if err := services.ValidarIDGrupo(request.IDGrupo, idGrupoJWT); err != nil {
		return nil, err
	}

	despesa := DespesaUpdateSchemaToEntity(request)
	despesa.IDUsuario = idUsuario
	despesa.IDGrupo = idGrupoJWT

	despesaUpdated, err := repository.UpdateDespesaRepository(despesa)
	if err != nil {
		return nil, err
	}

	response := DespesaEntityToUpdateSchema(*despesaUpdated)
	return &response, nil
}
