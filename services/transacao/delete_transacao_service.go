package transacao

import "github.com/GabrielVilarino/gestao-financeira-back.git/models/repository"

func DeleteTransacaoService(id int) error {
	return repository.DeleteTransacaoRepository(id)
}
