package despesa

import "github.com/GabrielVilarino/gestao-financeira-back.git/services/despesa"

func DeleteDespesaController(id int) error {
	return despesa.DeleteDespesaService(id)
}
