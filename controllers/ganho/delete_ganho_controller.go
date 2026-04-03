package ganho

import (
	ganhoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/ganho"
)

func DeleteGanhoController(id int) error {
	return ganhoService.DeleteGanhoService(id)
}
