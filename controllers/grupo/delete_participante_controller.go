package grupo

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/services/grupo"
)

func DeleteParticipanteController(idParticipante int, idGroup int) error {
	err := grupo.DeleteParticipanteService(idParticipante, idGroup)

	if err != nil {
		return err
	}

	return nil
}
