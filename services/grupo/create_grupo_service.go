package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func CreateGrupoService(idUsuario int, request schemas.CreateGrupoRequest) (*schemas.GrupoResponse, error) {
	grupoEntity := CreateGrupoSchemaToEntity(request)

	grupo, err := repository.CreateGrupoRepository(idUsuario, grupoEntity)

	if err != nil {
		return nil, err
	}

	response := EntityToGrupoResponse(*grupo)
	return &response, nil
}
