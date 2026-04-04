package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services"
)

func UpdateGanhoService(idUsuario int, idGrupoJWT *int, request schemas.UpdateGanhoRequest) (*schemas.UpdateGanhoResponse, error) {
	if err := services.ValidarIDGrupo(request.IDGrupo, idGrupoJWT); err != nil {
		return nil, err
	}

	ganho := GanhoUpdateSchemaToEntity(request)
	ganho.IDUsuario = idUsuario
	ganho.IDGrupo = idGrupoJWT

	ganhoUpdated, err := repository.UpdateGanhoRepository(ganho)
	if err != nil {
		return nil, err
	}

	response := GanhoEntityToUpdateSchema(*ganhoUpdated)
	return &response, nil
}
