package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func CreateGanhoService(idUsuario int, idGrupoJWT *int, request schemas.CreateGanhoRequest) (*schemas.CreateGanhoResponse, error) {
	if err := validarIDGrupo(request.IDGrupo, idGrupoJWT); err != nil {
		return nil, err
	}

	ganho := GanhoSchemaToEntity(request)
	ganho.IDUsuario = idUsuario
	if request.IDGrupo != nil {
		ganho.IDGrupo = request.IDGrupo
	} else {
		ganho.IDGrupo = idGrupoJWT
	}

	ganhoCreated, err := repository.CreateGanhoRepository(ganho)
	if err != nil {
		return nil, err
	}

	response := GanhoEntityToCreateSchema(*ganhoCreated)
	return &response, nil
}
