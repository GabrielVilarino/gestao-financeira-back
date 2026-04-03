package ganho

import (
	"github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"
)

func DeleteGanhoService(id int) error {
	return repository.DeleteGanhoRepository(id)
}
