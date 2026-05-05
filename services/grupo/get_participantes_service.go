package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
)

func GetParticipantesByIDService(idGrupo int) (*schemas.ParticipanteResponse, error) {
	grupo, err := repository.GetParticipantesByIDRepository(idGrupo)
	if err != nil {
		return nil, err
	}

	response := EntityToParticipanteResponse(*grupo)

	return &response, nil
}
