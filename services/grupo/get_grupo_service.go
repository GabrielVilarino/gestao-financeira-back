package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetGrupoServiceByID(idGrupo int) (*schemas.GrupoResponse, error) {
	grupo, err := repository.GetGrupoByIDRepository(idGrupo)
	if err != nil {
		return nil, err
	}

	response := EntityToGrupoResponse(*grupo)

	return &response, nil
}
