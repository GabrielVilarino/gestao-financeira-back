package transacao

import transacaoService "github.com/GabrielVilarino/gestao-financeira-back.git/services/transacao"

func DeleteTransacaoController(id int) error {
	return transacaoService.DeleteTransacaoService(id)
}
