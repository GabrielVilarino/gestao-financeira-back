package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/schemas"
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/grupo"
)

func UpdateRoleParticipanteController(request schemas.UpdateRoleParticipanteRequest, idGroup int, idParticipante int) error {
	err := grupo.UpdateRoleParticipanteService(request.IsAdmin, idGroup, idParticipante)

	if err != nil {
		return err
	}

	return nil
}
