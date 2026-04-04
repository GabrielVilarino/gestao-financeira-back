package despesa

import "github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"

func DeleteDespesaService(id int) error {
	return repository.DeleteDespesaRepository(id)
}
