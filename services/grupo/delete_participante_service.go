package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
)

func DeleteParticipanteService(idParticipante int, idGroup int) error {
	err := repository.DeleteParticipanteRepository(idParticipante, idGroup)
	if err != nil {
		return err
	}

	return nil
}
