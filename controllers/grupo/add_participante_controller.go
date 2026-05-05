package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/grupo"
)

func AddParticipanteController(request schemas.AddParticipanteRequest, idGroup int) error {
	err := grupo.AddParticipanteService(request.Email, idGroup)

	if err != nil {
		return err
	}

	return nil
}
