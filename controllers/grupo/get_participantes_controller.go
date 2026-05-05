package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/grupo"
)

func GetParticipantesByIDController(idGrupo int) (*schemas.ParticipanteResponse, error) {
	participantes, err := grupo.GetParticipantesByIDService(idGrupo)

	if err != nil {
		return nil, err
	}

	return participantes, nil
}
