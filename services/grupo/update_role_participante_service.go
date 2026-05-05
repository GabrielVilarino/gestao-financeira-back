package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
)

func UpdateRoleParticipanteService(isAdmin bool, idGroup int, idParticipante int) error {
	err := repository.UpdateRoleParticipanteRepository(isAdmin, idGroup, idParticipante)
	if err != nil {
		return err
	}

	return nil
}
