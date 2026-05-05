package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
)

func AddParticipanteService(email string, idGroup int) error {
	err := repository.AddParticipanteRepository(email, idGroup)
	if err != nil {
		return err
	}

	return nil
}
